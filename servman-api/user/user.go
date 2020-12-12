package user

import (
    "encoding/json"
    "github.com/plankiton/mux"
    "net/http"
    "time"
    "fmt"

    "../util"
)

type Doc struct {
    Type      string `json:"type,omitempty"`
           // cpf | cnpj
    Value     string `json:"value,omitempty"`
}

type Pass struct {
    Hash     string     `json:"hash,omitempty"`
}

type Person struct {
    ID         string    `json:"id,omitempty"`
    Document   *Doc      `json:"document,omitempty"`
    Telephone  string    `json:"telephone,omitempty"`
    Name       string    `json:"name,omitempty"`
    Type       string    `json:"type,omitempty"`
            // employee | root | client
    Password   *Pass     `json:"password,omitempty"`

    CreateTime time.Time `json:"created_at,omitempty"`
    UpdateTime time.Time `json:"updated_at,omitempty"`
}

func (self Person) CheckPass(s string) bool {
    created_at := self.CreateTime.Format("%Y%m%d%H%M%S")
    pass_hash := fmt.Sprintf("plankhash$%x$%x$%x",
                        util.ToHash(s),
                        util.ToHash(self.Document.Value),
                        util.ToHash(created_at))
    return self.Password.Hash == pass_hash
}

func (self Person) SetPass(s string) string {
    created_at := self.CreateTime.Format("%Y%m%d%H%M%S")

    hash := fmt.Sprintf("plankhash$%x$%x$%x",
                         util.ToHash(s),
                         util.ToHash(self.Document.Value),
                         util.ToHash(created_at))

    self.Password.Hash = hash
    return self.Password.Hash
}

var people []Person

// GetPeople mostra todos os contatos da variável people
func GetPeople(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(people)
}

// GetPerson mostra apenas um contato
func GetPerson(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for _, item := range people {
        if item.ID == params["id"] || item.Document.Value == params["id"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    json.NewEncoder(w).Encode(&Person{})
}

// CreatePerson cria um novo contato
func CreatePerson(w http.ResponseWriter, r *http.Request) {
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


    person := Person {
        Document: &Doc {},
        Password: &Pass {},
    }

    for i, prop := range(body.Data) {
        switch i {
        case "name":
            person.Name = prop
        case "document":
            person.Document.Value = prop
        case "doc_type":
            person.Document.Type = prop
        case "password":
            person.SetPass(prop)
        }
    }

    for _, item := range people {
        if item.Document.Value == person.Document.Value {
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

    person.CreateTime = time.Now()
    person.UpdateTime = time.Now()

    people = append(people, person)
    res := util.Response {
        Type:    "sucess",
        Message: "Person created!",
        Code:    "CreatedPerson",
        Data:    person,
    }

    json.NewEncoder(w).Encode(res)
}

func UpdatePerson(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)

    var body util.Request
    json.NewDecoder(r.Body).Decode(&body)

    for index, item := range people {
        if item.ID == params["id"] || item.Document.Value == params["id"] {
            person := people[index]
            not_set := false

            for i, prop := range(body.Data) {
                switch i {
                case "name":
                    person.Name = prop
                case "document":
                    person.Document.Value = prop
                case "doc_type":
                    person.Document.Type = prop
                case "password":
                    person.SetPass(prop)
                default:
                    not_set = true
                }
            }


            if !not_set {
                person.UpdateTime = time.Now()
            }

            people[index] = person
            json.NewEncoder(w).Encode(util.Response{
                Message: fmt.Sprintf("User %s did updated!", person.Name),
                Code: "UpdatedUser",
                Type: "sucess",
                Data: person,
            })
        }
    }

    w.WriteHeader(404)

    json.NewEncoder(w).Encode(util.Response{
        Message: "The person not found!",
        Code: "NotFound",
        Type: "error",
    })
}

// DeletePerson deleta um contato
func DeletePerson(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for index, item := range people {
        if item.ID == params["id"] || item.Document.Value == params["id"] {
            people = append(people[:index], people[index+1:]...)
            break
        }
        json.NewEncoder(w).Encode(people)
    }
}

func AppendPerson(p Person) {
    people = append(people, p)
}

func People() []Person {
    return people
}
