package farm

import (
    "encoding/json"
    "github.com/plankiton/mux"
    "net/http"
    "time"
    "fmt"
    "errors"

    "gorm.io/gorm"
    "../util"
)

type Addr struct {
    ID           string `json:"id,omitempty" gorm:"index"`
    Street       string `json:"street,omitempty"`
    State        string `json:"state,omitempty"`
    Number       string `json:"number,omitempty"`
    Code         string `json:"cep,omitempty"`
    City         string `json:"city,omitempty"`
    Neightbourn  string `json:"neightborhood,omitempty"`
}

type Farm struct {
    ID        string `json:"id,omitempty" gorm:"primaryKey"`
    PersonId  string `json:"person_id,omitempty"`
    AddressId string `json:"address_id,omitempty" gorm:"uniqueIndex"`
    Name      string `json:"name,omitempty" gorm:"index"`

    CreateTime time.Time `json:"created_at,omitempty" gorm:"index"`
    UpdateTime time.Time `json:"updated_at,omitempty" gorm:"index"`
}


var database *gorm.DB
func PopulateDB(db *gorm.DB) {
    database = db
    database.AutoMigrate(&Farm{})
    database.AutoMigrate(&Addr{})
}

// GetPeople mostra todos os contatos da variável farm
func GetAllFarms(w http.ResponseWriter, r *http.Request) {
    farm := []Farm{}
    database.Find(&farm)

    // TODO: sentence for validate logged user

    json.NewEncoder(w).Encode(util.Response{
            Code: "GetFarms",
            Type: "sucess",
            Data: farm,
        })
}
func GetFarms(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)

    farm := []Farm{}
    res := database.Where("person_id = ?", params["id"]).Find(&farm)
    if errors.Is(res.Error, gorm.ErrRecordNotFound) {
        w.WriteHeader(404)

        json.NewEncoder(w).Encode(util.Response{
            Message: "The farm not found!",
            Code: "NotFound",
            Type: "error",
            Data: nil,
        })

        return
    }

    // TODO: sentence for validate logged user
    json.NewEncoder(w).Encode(util.Response{
            Code: "GetFarms",
            Type: "sucess",
            Data: farm,
        })
}

// GetFarm mostra apenas um contato
func GetFarm(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    farm := Farm{}

    res := database.First(&farm, params["id"])
    if errors.Is(res.Error, gorm.ErrRecordNotFound) {
        w.WriteHeader(404)

        json.NewEncoder(w).Encode(util.Response{
            Message: "The farm not found!",
            Code: "NotFound",
            Type: "error",
            Data: nil,
        })

        return
    }

    // TODO: sentence for validate logged user
    json.NewEncoder(w).Encode(util.Response{
            Code: "GetFarm",
            Type: "sucess",
            Data: farm,
        })
}

// CreateFarm cria um novo contato
func CreateFarm(w http.ResponseWriter, r *http.Request) {
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

    farm := Farm {
        PersonId: params["id"],
    }

    address := Addr{}
    r_addr, err := http.Get(fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", body.Data["cep"]))
    if err == nil {
        json.NewDecoder(r_addr.Body).Decode(&address)
    }

    for i, prop := range(body.Data) {
        switch i {
        case "name":
            farm.Name = prop
        case "cep":
            if address.Code == "" {
                address.Code = prop
            }
        case "street":
            if address.Street == "" {
                address.Street = prop
            }
        case "number":
            if address.Number == "" {
                address.Number = prop
            }
        case "neighborhood":
            if address.Neightbourn == "" {
                address.Neightbourn = prop
            }
        case "state":
            if address.State == "" {
                address.State = prop
            }
        case "city":
            if address.City == "" {
                address.City = prop
            }
        }
    }

    farm.CreateTime = time.Now()
    farm.UpdateTime = time.Now()

    farm.ID = util.ToHash(farm.Name+
        farm.CreateTime.Format("%Y%m%d%H%M%S"))

    address.ID = util.ToHash(address.Code+address.Street+address.Number)
    {
        addr := Addr{}
        res := database.Where("id = ?", address.ID).First(&addr)
        if !errors.Is(res.Error, gorm.ErrRecordNotFound) {
            address = addr
        } else {
            database.Create(&address)
            database.Commit()
        }
    }
    farm.AddressId = address.ID

    res := database.Where("address_id = ?", farm.AddressId).First(&farm)
    if !errors.Is(res.Error, gorm.ErrRecordNotFound) {
        w.WriteHeader(403)

        json.NewEncoder(w).Encode(util.Response{
            Message: "The farm already exists!",
            Code: "AlreadyExists",
            Type: "error",
            Data: nil,
        })

        return
    }

    // TODO: sentence for validate logged user
    // Sending all to database
    database.Create(&farm)
    database.Commit()
    json.NewEncoder(w).Encode(util.Response {
        Type:    "sucess",
        Code:    "CreatedFarm",
        Data:    farm,
    })
}

func UpdateFarm(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)

    var body util.Request
    json.NewDecoder(r.Body).Decode(&body)

    farm, address := Farm{}, Addr{}

    res := database.First(&farm, params["id"])
    if errors.Is(res.Error, gorm.ErrRecordNotFound) {
        w.WriteHeader(404)

        json.NewEncoder(w).Encode(util.Response{
            Message: "The farm not found!",
            Code: "NotFound",
            Type: "error",
        })

        return
    }
    res = database.First(&address, farm.AddressId)

    for i, prop := range(body.Data) {
        switch i {
        case "name":
            farm.Name = prop
        case "cep":
            address.Code = prop

            r_addr, err := http.Get(fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", body.Data["cep"]))
            if err == nil {
                json.NewDecoder(r_addr.Body).Decode(&address)
            }
        case "street":
            address.Street = prop
        case "number":
            address.Number = prop
        case "neighborhood":
            address.Neightbourn = prop
        case "state":
            address.State = prop
        case "city":
            address.City = prop
        }
    }

    address.ID = util.ToHash(address.Code+address.Street+address.Number)
    {
        addr := Addr{}
        res := database.Where("id = ?", address.ID).First(&addr)
        if !errors.Is(res.Error, gorm.ErrRecordNotFound) {
            address = addr
        } else {
            database.Create(&address)
            database.Commit()
        }
    }

    farm.UpdateTime = time.Now()

    farm.AddressId = address.ID

    // TODO: sentence for validate logged user

    database.Save(&farm)
    database.Commit()
    json.NewEncoder(w).Encode(util.Response{
        Message: fmt.Sprintf("Farm %s did updated!", farm.Name),
        Code: "UpdatedFarm",
        Type: "sucess",
        Data: farm,
    })


}

// DeleteFarm deleta um contato
func DeleteFarm(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)

    farm := Farm{}
    res := database.First(&farm, params["id"])
    if errors.Is(res.Error, gorm.ErrRecordNotFound) {
        w.WriteHeader(404)

        json.NewEncoder(w).Encode(util.Response{
            Message: "The farm not found!",
            Code: "NotFound",
            Type: "error",
            Data: nil,
        })

        return
    }

    // TODO: sentence for validate logged user

    database.Delete(&farm)
    database.Commit()
    json.NewEncoder(w).Encode(util.Response{
        Type: "sucess",
        Code: "DeleteFarm",
        Data: farm,
    })
}
