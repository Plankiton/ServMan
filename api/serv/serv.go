package serv

import (
    "encoding/json"
    "github.com/gorilla/mux"
    "net/http"
    "strconv"
    "time"
    "errors"

    "gorm.io/gorm"

    "github.com/plankiton/ServMan/api/util"
    "github.com/plankiton/ServMan/api/user"
    "github.com/plankiton/ServMan/api/farm"
)

type Serv struct {
    ID          string    `json:"id,omitempty" gorm:"primaryKey"`
    Description string    `json:"description,omitempty"`
    Price       float64   `json:"price,omitempty" gorm:"index"`
    Begin_photo string    `json:"begin,omitempty"`
    End_photo   string    `json:"end,omitempty"`

    EmployeeId  string    `json:"employee,omitempty" gorm:"index"`
    FarmId      string    `json:"farm,omitempty" gorm:"index"`

    BeginTime   time.Time `json:"started_at,omitempty" gorm:"index"`
    EndTime     time.Time `json:"finished_at,omitempty" gorm:"index"`
}

var database *gorm.DB
func PopulateDB(db *gorm.DB) {
    database = db
    database.AutoMigrate(&Serv{})
}

// GetPeople mostra todos os contatos da variável serv
func GetAllServs(w http.ResponseWriter, r *http.Request) {
    serv := []Serv{}
    database.Find(&serv)

    // TODO: sentence for validate logged user

    json.NewEncoder(w).Encode(util.Response{
        Code: "GetServices",
        Type: "sucess",
        Data: serv,
    })
}

func GetServs(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)

    serv := []Serv{}
    res := database.Where("employee_id = ? ", params["id"]).Find(&serv)
    if errors.Is(res.Error, gorm.ErrRecordNotFound) {
        res = database.Where("farm_id = ? ", params["id"]).Find(&serv)
        if errors.Is(res.Error, gorm.ErrRecordNotFound) {
            w.WriteHeader(404)

            json.NewEncoder(w).Encode(util.Response{
                Message: "The person not found!",
                Code: "NotFound",
                Type: "error",
                Data: nil,
            })

            return
        }
    }

    // TODO: sentence for validate logged user

    json.NewEncoder(w).Encode(util.Response{
        Code: "GetServices",
        Type: "sucess",
        Data: serv,
    })

}

// GetServ mostra apenas um contato
func GetServ(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)

    serv := Serv{}
    res := database.Where("id = ?", params["id"]).First(&serv, params["id"])
    if errors.Is(res.Error, gorm.ErrRecordNotFound) {
        w.WriteHeader(404)

        json.NewEncoder(w).Encode(util.Response{
            Message: "The person not found!",
            Code: "NotFound",
            Type: "error",
            Data: nil,
        })

        return
    }

    // TODO: sentence for validate logged user

    json.NewEncoder(w).Encode(util.Response{
        Code: "GetService",
        Type: "sucess",
        Data: serv,
    })
}

// CreateServ cria um novo contato
func CreateServ(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)

    var body util.Request
    json.NewDecoder(r.Body).Decode(&body)

    if len(body.Data) == 0 {
        w.WriteHeader(400)

        json.NewEncoder(w).Encode(util.Response{
            Message: "The data sent is invalid!"+
                     `(must be {"data": "..."})`,
            Code: "BadRequest",
            Type: "error",
            Data: nil,
        })

        return
    }


    serv := Serv {}
    if r.URL.Path[:6+4] == "/user/farm" {

        farm := farm.Farm{}
        res := database.Where("id = ?", params["id"]).First(&farm)
        if errors.Is(res.Error, gorm.ErrRecordNotFound) {
            w.WriteHeader(404)

            json.NewEncoder(w).Encode(util.Response{
                Message: "The farm not found!",
                Code: "NotFound",
                Type: "error",
            })

            return
        }

        serv.FarmId = farm.ID
        serv.EmployeeId = body.Data["employee"]

    } else {

        person := user.Person{}
        res := database.Where("doc_value = ? OR id = ?",
        params["id"], params["id"]).First(&person)
        if errors.Is(res.Error, gorm.ErrRecordNotFound) {
            w.WriteHeader(404)

            json.NewEncoder(w).Encode(util.Response{
                Message: "The person not found!",
                Code: "NotFound",
                Type: "error",
                Data: nil,
            })

            return
        }

        serv.EmployeeId = person.ID
        serv.FarmId = body.Data["farm"]

    }

    price, err := strconv.ParseFloat(body.Data["price"], 64)
    if err != nil {
        price = 0.0
    }
    serv.Price = price

    serv.Description = body.Data["description"]
    serv.ID = util.ToHash(serv.Description+serv.EmployeeId+serv.FarmId)
    if serv.ID == "" {
        w.WriteHeader(400)

        json.NewEncoder(w).Encode(util.Response{
            Message: "The data sent is invalid!"+
                     `(must be {"data": "..."})`,
            Code: "BadRequest",
            Type: "error",
            Data: nil,
        })

        return
    }

    serv.BeginTime = time.Now()
    serv.EndTime = time.Now()

    // TODO: sentence for validate logged user

    database.Create(serv)
    database.Commit()
    json.NewEncoder(w).Encode(util.Response {
        Type:    "sucess",
        Code:    "CreatedServ",
        Data:    serv,
    })
}

func MarkServTime(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)

    var body util.Request
    json.NewDecoder(r.Body).Decode(&body)

    serv := Serv{}
    res := database.Where("id = ?", params["id"]).First(&serv)
    if errors.Is(res.Error, gorm.ErrRecordNotFound) {
        w.WriteHeader(404)

        json.NewEncoder(w).Encode(util.Response{
            Message: "The service not found!",
            Code: "NotFound",
            Type: "error",
            Data: nil,
        })

        return
    }

    set := false
    timestamp, err := strconv.ParseInt(body.Data["value"], 10, 64)

    if body.Data["value"] == "" {
        if body.Data["type"] == "begin" {
            serv.BeginTime = time.Now()
        } else {
            serv.EndTime = time.Now()
        }

        set = true
    }

    if !set {
        if err != nil {
            w.WriteHeader(400)

            json.NewEncoder(w).Encode(util.Response{
                Message: "The data sent is invalid!"+
                `(must be {"data": "..."})`,
                Code: "BadRequest",
                Type: "error",
                Data: nil,
            })

            return
        }

        if body.Data["type"] == "begin" {
            serv.BeginTime = time.Unix(timestamp, 0)
        } else {
            serv.EndTime = time.Unix(timestamp, 0)
        }
    }

    database.Save(&serv)
    database.Commit()

    json.NewEncoder(w).Encode(util.Response{
        Code: "UpdatedServ",
        Type: "sucess",
        Data: serv,
    })
}

func UpdateServ(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)

    var body util.Request
    json.NewDecoder(r.Body).Decode(&body)

    serv := Serv{}
    res := database.Where("id = ?", params["id"]).First(&serv)
    if errors.Is(res.Error, gorm.ErrRecordNotFound) {
        w.WriteHeader(404)

        json.NewEncoder(w).Encode(util.Response{
            Message: "The service not found!",
            Code: "NotFound",
            Type: "error",
            Data: nil,
        })

        return
    }

    for i, prop := range(body.Data) {
        switch i {
        case "description":
            serv.Description = prop
        case "price":
            var err error
            serv.Price, err = strconv.ParseFloat(prop, 64)
            if err != nil {
                serv.Price = 0.0
            }
        }
    }


    // TODO: sentence for validate logged user

    database.Save(&serv)
    database.Commit()
    json.NewEncoder(w).Encode(util.Response{
        Code: "UpdatedServ",
        Type: "sucess",
        Data: serv,
    })
}

// DeleteServ deleta um contato
func DeleteServ(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)

    serv := Serv{}
    res := database.Where("id = ?", params["id"]).First(&serv)
    if errors.Is(res.Error, gorm.ErrRecordNotFound) {
        w.WriteHeader(404)

        json.NewEncoder(w).Encode(util.Response{
            Message: "The person not found!",
            Code: "NotFound",
            Type: "error",
            Data: nil,
        })

        return
    }

    // TODO: sentence for validate logged user

    database.Delete(&serv)
    database.Commit()
    json.NewEncoder(w).Encode(util.Response{
        Code: "DeleteService",
        Type: "sucess",
        Data: serv,
    })
}
