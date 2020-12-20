package serv

import (
    "encoding/json"
    "github.com/gorilla/mux"
    "net/http"
    "strconv"
    "time"
    "errors"

    "gorm.io/gorm"

    "../util"
    "../user"
    "../farm"
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
    res := database.First(&serv, params["id"])
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
        serv.FarmId = params["id"]
        serv.EmployeeId = body.Data["employee"]
    } else {
        serv.EmployeeId = params["id"]
        serv.FarmId = body.Data["farm"]
    }

    empl, farm := user.Person{}, farm.Farm{}
    res := database.Where("id = ? AND type = 'employee'", serv.EmployeeId).First(&empl)
    if errors.Is(res.Error, gorm.ErrRecordNotFound) {
        w.WriteHeader(404)

        json.NewEncoder(w).Encode(util.Response{
            Message: "The employee not found!",
            Code: "NotFound",
            Type: "error",
        })

        return
    }
    res = database.Where("id = ? AND person_id = ?", serv.FarmId, empl.ID).First(&farm)
    if errors.Is(res.Error, gorm.ErrRecordNotFound) {
        w.WriteHeader(404)

        json.NewEncoder(w).Encode(util.Response{
            Message: "The farm not found or employee is not then owner!",
            Code: "NotFound",
            Type: "error",
        })

        return
    }

    serv.Description = body.Data["description"]
    serv.ID = util.ToHash(serv.Description+serv.EmployeeId+serv.FarmId)

    var err error
    serv.Price, err = strconv.ParseFloat(body.Data["price"], 64)
    if err != nil {
        serv.Price = 0.0
    }


    {
        serv := Serv{ID:serv.ID,}
        res := database.First(&serv, serv.ID)
        if errors.Is(res.Error, gorm.ErrRecordNotFound) {
            w.WriteHeader(403)

            json.NewEncoder(w).Encode(util.Response{
                Message: "The user already exists!",
                Code: "AlreadyExists",
                Type: "error",
                Data: nil,
            })

            return
        }
    }

    serv.BeginTime = time.Now()

    // TODO: sentence for validate logged user

    database.Create(serv)
    database.Commit()
    json.NewEncoder(w).Encode(util.Response {
        Type:    "sucess",
        Code:    "CreatedServ",
        Data:    serv,
    })
}

func UpdateServ(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)

    var body util.Request
    json.NewDecoder(r.Body).Decode(&body)

    serv := Serv{}
    res := database.First(&serv, params["id"])
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
    res := database.First(&serv, params["id"])
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
