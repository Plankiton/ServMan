package farm

import (
    "encoding/json"
    "github.com/plankiton/mux"
    "net/http"
    "time"
    "fmt"

    "../util"
)

type Addr struct {
    Street       string `json:"street,omitempty"`
    State        string `json:"state,omitempty"`
    Number       string `json:"number,omitempty"`
    Code         string `json:"cep,omitempty"`
    City         string `json:"city,omitempty"`
    Neightbourn  string `json:"neightborhood,omitempty"`
}

type Farm struct {
    PersonId  string `json:"client_id,omitempty"`
    ID        string `json:"id,omitempty"`
    Address   *Addr  `json:"address,omitempty"`
    Name      string `json:"name,omitempty"`

    CreateTime time.Time `json:"created_at,omitempty"`
    UpdateTime time.Time `json:"updated_at,omitempty"`
}


var farms []Farm

// GetPeople mostra todos os contatos da variável farms
func GetFarms(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)

    var founds []Farm
    for _, f := range(farms) {
        if (f.PersonId == params["id"]){
            founds = append(founds, f)
        }
    }

    json.NewEncoder(w).Encode(founds)
}

// GetFarm mostra apenas um contato
func GetFarm(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for _, item := range farms {
        if item.ID == params["id"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    json.NewEncoder(w).Encode(&Farm{})
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
        Address: &Addr{},
        PersonId: params["id"],
    }

    r_addr, err := http.Get(fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", body.Data["cep"]))
    if err == nil {
        json.NewDecoder(r_addr.Body).Decode(&farm.Address)
    }

    for i, prop := range(body.Data) {
        switch i {
        case "name":
            farm.Name = prop
        case "cep":
            if farm.Address.Code == "" {
                farm.Address.Code = prop
            }
        case "street":
            if farm.Address.Street == "" {
                farm.Address.Street = prop
            }
        case "number":
            if farm.Address.Number == "" {
                farm.Address.Number = prop
            }
        case "neighborhood":
            if farm.Address.Neightbourn == "" {
                farm.Address.Neightbourn = prop
            }
        case "state":
            if farm.Address.State == "" {
                farm.Address.State = prop
            }
        case "city":
            if farm.Address.City == "" {
                farm.Address.City = prop
            }
        }
    }


    for _, item := range farms {
        if item.ID == farm.ID {
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

    farm.CreateTime = time.Now()
    farm.UpdateTime = time.Now()

    farms = append(farms, farm)
    res := util.Response {
        Type:    "sucess",
        Message: "Farm created!",
        Code:    "CreatedFarm",
        Data:    farm,
    }

    json.NewEncoder(w).Encode(res)
}

func UpdateFarm(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)

    var body util.Request
    json.NewDecoder(r.Body).Decode(&body)

    for index, item := range farms {
        if item.ID == params["id"] {
            farm := farms[index]
            not_set := false

            for i, prop := range(body.Data) {
                switch i {
                case "name":
                    farm.Name = prop
                case "cep":
                    farm.Address.Code = prop

                    r_addr, err := http.Get(fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", body.Data["cep"]))
                    if err == nil {
                        json.NewDecoder(r_addr.Body).Decode(&farm.Address)
                    }

                case "street":
                    farm.Address.Street = prop
                case "number":
                    farm.Address.Number = prop
                case "neighborhood":
                    farm.Address.Neightbourn = prop
                case "state":
                    farm.Address.State = prop
                case "city":
                    farm.Address.City = prop
                default:
                    not_set = true
                }
            }


            if !not_set {
                farm.UpdateTime = time.Now()

                farms[index] = farm
                json.NewEncoder(w).Encode(util.Response{
                    Message: fmt.Sprintf("Farm %s did updated!", farm.Name),
                    Code: "UpdatedFarm",
                    Type: "sucess",
                    Data: farm,
                })

                return
            }
        }
    }

    w.WriteHeader(404)

    json.NewEncoder(w).Encode(util.Response{
        Message: "The farm not found!",
        Code: "NotFound",
        Type: "error",
    })
}

// DeleteFarm deleta um contato
func DeleteFarm(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for index, item := range farms {
        if item.ID == params["id"] {
            farms = append(farms[:index], farms[index+1:]...)
            break
        }
        json.NewEncoder(w).Encode(farms)
    }
}

func AppendFarm(p Farm) {
    farms = append(farms, p)
}

func Farms() []Farm {
    return farms
}
