package serv

import (
    "encoding/json"
    "github.com/plankiton/mux"
    "net/http"
    "time"
    "fmt"

    "../util"
)

type Serv struct {
    ID          string    `json:"id,omitempty"`
    Description string    `json:"description,omitempty"`
    Price       float64   `json:"price,omitempty"`
    Begin_photo string    `json:"begin,omitempty"`
    End_photo   string    `json:"end,omitempty"`

    EmployeeId  string    `json:"employee,omitempty"`
    FarmId      string    `json:"farm,omitempty"`

    BeginTime   time.Time `json:"started_at,omitempty"`
    EndTime     time.Time `json:"finished_at,omitempty"`
}

var services []Serv

// GetPeople mostra todos os contatos da variável services
func GetServs(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)

    var founds []Serv
    for _, f := range(services) {
        if (f.EmployeeId == params["id"]){
            founds = append(founds, f)
        }
    }

    json.NewEncoder(w).Encode(founds)
}

// GetServ mostra apenas um contato
func GetServ(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for _, item := range services {
        if item.ID == params["id"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    json.NewEncoder(w).Encode(&Serv{})
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


    service := Serv {
        Address: &Addr{},
        PersonId: params["id"],
    }

    r_addr, err := http.Get(fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", body.Data["cep"]))
    if err == nil {
        json.NewDecoder(r_addr.Body).Decode(&service.Address)
    }

    for i, prop := range(body.Data) {
        switch i {
        case "name":
            service.Name = prop
        case "cep":
            if service.Address.Code == "" {
                service.Address.Code = prop
            }
        case "street":
            if service.Address.Street == "" {
                service.Address.Street = prop
            }
        case "number":
            if service.Address.Number == "" {
                service.Address.Number = prop
            }
        case "neighborhood":
            if service.Address.Neightbourn == "" {
                service.Address.Neightbourn = prop
            }
        case "state":
            if service.Address.State == "" {
                service.Address.State = prop
            }
        case "city":
            if service.Address.City == "" {
                service.Address.City = prop
            }
        }
    }


    for _, item := range services {
        if item.ID == service.ID {
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

    service.CreateTime = time.Now()
    service.UpdateTime = time.Now()

    services = append(services, service)
    res := util.Response {
        Type:    "sucess",
        Message: "Serv created!",
        Code:    "CreatedServ",
        Data:    service,
    }

    json.NewEncoder(w).Encode(res)
}

func UpdateServ(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)

    var body util.Request
    json.NewDecoder(r.Body).Decode(&body)

    for index, item := range services {
        if item.ID == params["id"] {
            service := services[index]
            not_set := false

            for i, prop := range(body.Data) {
                switch i {
                case "name":
                    service.Name = prop
                case "cep":
                    service.Address.Code = prop

                    r_addr, err := http.Get(fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", body.Data["cep"]))
                    if err == nil {
                        json.NewDecoder(r_addr.Body).Decode(&service.Address)
                    }

                case "street":
                    service.Address.Street = prop
                case "number":
                    service.Address.Number = prop
                case "neighborhood":
                    service.Address.Neightbourn = prop
                case "state":
                    service.Address.State = prop
                case "city":
                    service.Address.City = prop
                default:
                    not_set = true
                }
            }


            if !not_set {
                service.UpdateTime = time.Now()

                services[index] = service
                json.NewEncoder(w).Encode(util.Response{
                    Message: fmt.Sprintf("Serv %s did updated!", service.Name),
                    Code: "UpdatedServ",
                    Type: "sucess",
                    Data: service,
                })

                return
            }
        }
    }

    w.WriteHeader(404)

    json.NewEncoder(w).Encode(util.Response{
        Message: "The service not found!",
        Code: "NotFound",
        Type: "error",
    })
}

// DeleteServ deleta um contato
func DeleteServ(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for index, item := range services {
        if item.ID == params["id"] {
            services = append(services[:index], services[index+1:]...)
            break
        }
        json.NewEncoder(w).Encode(services)
    }
}

func AppendServ(p Serv) {
    services = append(services, p)
}

func Servs() []Serv {
    return services
}
