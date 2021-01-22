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
    Stoped      bool      `json:"stoped,omitempty"`
}

var database *gorm.DB
func PopulateDB(db *gorm.DB) {
    database = db
    database.AutoMigrate(&Serv{})
}

// GetPeople mostra todos os contatos da variável serv
func GetAllServs(r *http.Request) util.Response {
    serv := []Serv{}
    database.Find(&serv)

    // TODO: sentence for validate logged user

    return (util.Response{
        Code: "GetServices",
        Type: "sucess",
        Data: serv,
    })
}

func GetServs(r *http.Request) util.Response {
    params := mux.Vars(r)

    serv := []Serv{}
    res := database.Where("employee_id = ? ", params["id"]).Find(&serv)
    if errors.Is(res.Error, gorm.ErrRecordNotFound) {
        res = database.Where("farm_id = ? ", params["id"]).Find(&serv)
        if errors.Is(res.Error, gorm.ErrRecordNotFound) {

            return (util.Response{
                Message: "The person not found!",
                Code: "NotFound",
                Status: 404,
                Type: "error",
                Data: nil,
            })

        }
    }

    // TODO: sentence for validate logged user

    return (util.Response{
        Code: "GetServices",
        Type: "sucess",
        Data: serv,
    })

}

// GetServ mostra apenas um contato
func GetServ(r *http.Request) util.Response {
    params := mux.Vars(r)

    serv := Serv{}
    res := database.Where("id = ?", params["id"]).First(&serv, params["id"])
    if errors.Is(res.Error, gorm.ErrRecordNotFound) {

        return (util.Response{
            Status: 404,
            Message: "The person not found!",
            Code: "NotFound",
            Type: "error",
            Data: nil,
        })

    }

    // TODO: sentence for validate logged user

    return (util.Response{
        Code: "GetService",
        Type: "sucess",
        Data: serv,
    })
}

// CreateServ cria um novo contato
func CreateServ(r *http.Request) util.Response {
    params := mux.Vars(r)

    var body util.Request
    json.NewDecoder(r.Body).Decode(&body)

    if len(body.Data) == 0 {

        return (util.Response{
            Status: 400,
            Message: "The data sent is invalid!"+
                     `(must be {"data": "..."})`,
            Code: "BadRequest",
            Type: "error",
            Data: nil,
        })

    }


    serv := Serv {}
    if r.URL.Path[:6+4] == "/user/farm" {

        farm := farm.Farm{}
        res := database.Where("id = ?", params["id"]).First(&farm)
        if errors.Is(res.Error, gorm.ErrRecordNotFound) {

            return (util.Response{
                Status: 404,
                Message: "The farm not found!",
                Code: "NotFound",
                Type: "error",
            })

        }

        serv.FarmId = farm.ID
        serv.EmployeeId = body.Data["employee"]

    } else {

        person := user.Person{}
        res := database.Where("doc_value = ? OR id = ?",
        params["id"], params["id"]).First(&person)
        if errors.Is(res.Error, gorm.ErrRecordNotFound) {

            return (util.Response{
                Status: 404,
                Message: "The person not found!",
                Code: "NotFound",
                Type: "error",
                Data: nil,
            })

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

        return (util.Response{
            Status: 400,
            Message: "The data sent is invalid!"+
                     `(must be {"data": "..."})`,
            Code: "BadRequest",
            Type: "error",
            Data: nil,
        })

    }

    serv.BeginTime = time.Now()
    serv.EndTime = time.Now()

    // TODO: sentence for validate logged user

    database.Create(serv)
    database.Commit()
    return (util.Response {
        Type:    "sucess",
        Code:    "CreatedServ",
        Data:    serv,
    })
}

func MarkServTime(r *http.Request) util.Response {
    params := mux.Vars(r)

    var body util.Request
    json.NewDecoder(r.Body).Decode(&body)

    serv := Serv{}
    res := database.Where("id = ?", params["id"]).First(&serv)
    if errors.Is(res.Error, gorm.ErrRecordNotFound) {

        return (util.Response{
            Status: 404,
            Message: "The service not found!",
            Code: "NotFound",
            Type: "error",
            Data: nil,
        })

    }

    set := false

    if body.Data["value"] == "" {
        if body.Data["type"] == "begin" {
            serv.BeginTime = time.Now()
            serv.Stoped = false
        } else {
            serv.EndTime = time.Now()
            serv.Stoped = true
        }

        set = true
    }

    if !set {
        layout := "2006-01-02T15:04:05.000Z"
        if body.Data["type"] == "begin" {
            serv.BeginTime, _ = time.Parse(layout, body.Data["value"])
            serv.Stoped = false
        } else {
            serv.EndTime, _ = time.Parse(layout, body.Data["value"])
            serv.Stoped = true
        }
    }

    database.Save(&serv)
    database.Commit()

    return (util.Response{
        Code: "UpdatedServ",
        Type: "sucess",
        Data: serv,
    })
}

func UpdateServ(r *http.Request) util.Response {
    params := mux.Vars(r)

    var body util.Request
    json.NewDecoder(r.Body).Decode(&body)

    serv := Serv{}
    res := database.Where("id = ?", params["id"]).First(&serv)
    if errors.Is(res.Error, gorm.ErrRecordNotFound) {

        return (util.Response{
            Status: 404,
            Message: "The service not found!",
            Code: "NotFound",
            Type: "error",
            Data: nil,
        })

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
    return (util.Response{
        Code: "UpdatedServ",
        Type: "sucess",
        Data: serv,
    })
}

// DeleteServ deleta um contato
func DeleteServ(r *http.Request) util.Response {
    params := mux.Vars(r)

    serv := Serv{}
    res := database.Where("id = ?", params["id"]).First(&serv)
    if errors.Is(res.Error, gorm.ErrRecordNotFound) {

        return (util.Response{
            Status: 404,
            Message: "The person not found!",
            Code: "NotFound",
            Type: "error",
            Data: nil,
        })

    }

    // TODO: sentence for validate logged user

    database.Delete(&serv)
    database.Commit()
    return (util.Response{
        Code: "DeleteService",
        Type: "sucess",
        Data: serv,
    })
}
