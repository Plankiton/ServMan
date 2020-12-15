package user

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

type Doc struct {
    Type      string `json:"type,omitempty"`
           // cpf | cnpj
    Value     string `json:"value,omitempty"`
}

type Pass struct {
    Hash     string     `json:"hash,omitempty"`
}

type Person struct {
    ID         string `json:"id,omitempty" gorm:"index"`
    Document   *Doc      `json:"document,omitempty" gorm:"embendded,embenddedPrefix=Doc"`
    Telephone  string    `json:"telephone,omitempty" gorm:"uniqueIndex"`
    Name       string    `json:"name,omitempty" gorm:"index"`
    Type       string    `json:"type,omitempty" gorm:"index"`

    Password   *Pass     `json:"password,omitempty" gorm:"embendded,embenddedPrefix=Pass"`

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

var database *gorm.DB
func PopulateDB(db *gorm.DB) {
    database = db
}


// GetPeople mostra todos os contatos da variável people
func GetPeople(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)

    people := []Person{}
    res := database.First(&people, params["id"])
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

    json.NewEncoder(w).Encode(util.Response{
        Code: "GetPeople",
        Type: "sucess",
        Data: people,
    })

}

// GetPerson mostra apenas um contato
func GetPerson(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)

    person := Person{}
    res := database.First(&person, params["id"])
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

    json.NewEncoder(w).Encode(util.Response{
        Code: "GetPerson",
        Type: "sucess",
        Data: person,
    })
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

    {
        person := Person{Document: person.Document,}
        res := database.Where("DocValue = ?", person.Document.Value).Find(&person)
        if !errors.Is(res.Error, gorm.ErrRecordNotFound) {
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

    database.Create(person)

    json.NewEncoder(w).Encode(util.Response {
        Type:    "sucess",
        Code:    "CreatedPerson",
        Data:    person,
    })
}

func UpdatePerson(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)

    var body util.Request
    json.NewDecoder(r.Body).Decode(&body)

    person := Person{}
    res := database.First(&person, params["id"])
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

        database.Save(&person)
        json.NewEncoder(w).Encode(util.Response{
            Message: fmt.Sprintf("User %s did updated!", person.Name),
            Code: "UpdatedUser",
            Type: "sucess",
            Data: person,
        })
        return
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
    person := Person{}
    res := database.First(&person, params["id"])
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

    database.Delete(&person)
    json.NewEncoder(w).Encode(util.Response{
        Code: "DeleteUser",
        Type: "sucess",
        Data: person,
    })
}
