package serv

import (
    "encoding/json"
    "github.com/plankiton/mux"
    "net/http"
    "strconv"
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

    founds := []Serv{}
    for _, f := range(services) {
        if r.URL.Path[:6+4] == "/user/farm" {
            if (f.FarmId == params["id"]){
                founds = append(founds, f)
            }
        } else {
            if (f.EmployeeId == params["id"]){
                founds = append(founds, f)
            }
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


    service := Serv {}
    if r.URL.Path[:6+4] == "/user/farm" {
        service.FarmId = params["id"]
    } else {
        service.EmployeeId = params["id"]
    }

    for i, prop := range(body.Data) {
        switch i {
        case "description":
            service.Description = prop
        case "price":
            var err error
            service.Price, err = strconv.ParseFloat(prop, 64)
            if err != nil {
                service.Price = 0.0
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

    service.BeginTime = time.Now()

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
                case "description":
                    service.Description = prop
                case "price":
                    var err error
                    service.Price, err = strconv.ParseFloat(prop, 64)
                    if err != nil {
                        service.Price = 0.0
                    }
                default:
                    not_set = true
                }
            }


            if !not_set {
                services[index] = service
                json.NewEncoder(w).Encode(util.Response{
                    Message: fmt.Sprintf("Serv \"%s\" did updated!", service.Description),
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
