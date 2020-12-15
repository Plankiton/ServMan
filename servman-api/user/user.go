package user

import (
    "encoding/json"
    "github.com/plankiton/mux"
    "net/http"
    "time"
    "errors"

    "gorm.io/gorm"
    "../util"
)

type Person struct {
    ID         string    `json:"id,omitempty" gorm:"index"`
    DocValue   string    `json:"document,omitempty" gorm:"uniqueIndex"`
    DocType    string    `json:"doc_type,omitempty" gorm:"uniqueIndex"`
    Telephone  string    `json:"telephone,omitempty" gorm:"uniqueIndex"`
    Name       string    `json:"name,omitempty" gorm:"index"`
    Type       string    `json:"type,omitempty" gorm:"index"`

    PassHash   string    `json:"password,omitempty" gorm:"embendded,embenddedPrefix=Pass"`

    CreateTime time.Time `json:"created_at,omitempty"`
    UpdateTime time.Time `json:"updated_at,omitempty"`
}

func (self *Person) CheckPass(s string) bool {
    byteHash := []byte(self.PassHash)
    err := util.CheckPass(byteHash, s)
    if err != nil {
        return false
    }
    return true
}

func (self *Person) SetPass(s string) (string, error) {
    hash, err := util.PassHash(s)
    if err != nil {
        return "", nil
    }

    self.PassHash = hash
    return self.PassHash, nil
}

var database *gorm.DB
func PopulateDB(db *gorm.DB) {
    database = db
    database.AutoMigrate(&Person{})
}


// GetPeople mostra todos os contatos da variável person
func GetPeople(w http.ResponseWriter, r *http.Request) {

    person := []Person{}
    res := database.Find(&person)
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
    people := []Person{}
    for _, p := range(person) {
        people = append(people, Person {
            Name: p.Name,
            DocValue: p.DocValue,
            Type: p.Type,
            CreateTime: p.CreateTime,
            UpdateTime: p.UpdateTime,
        })
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
    res := database.Where("doc_value = ?", params["id"]).First(&person)
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
        Code: "GetPerson",
        Type: "sucess",
        Data: Person {
            ID: person.ID,
            Name: person.Name,
            DocValue: person.DocValue,
            Type: person.Type,
            CreateTime: person.CreateTime,
            UpdateTime: person.UpdateTime,
        },
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


    person := Person {}

    person.Name     = body.Data["name"]
    person.DocValue = body.Data["document"]
    person.DocType  = body.Data["doc_type"]
    person.Type     = body.Data["type"]
    _, err := person.SetPass(body.Data["password"])
    if person.Type == "root" {
        res := database.Where("doc_value = ?", person.DocValue).First(&person)
        if res.RowAffected > 0 {
            person.Type = "employee"
            // TODO: sentence for validate logged user
        }
    }

    if err != nil {
        w.WriteHeader(500)

        json.NewEncoder(w).Encode(util.Response{
            Message: "Error on creation of password hash on database!",
            Code: "DbError",
            Type: "error",
        })

        return
    }

    {
        person := Person{DocValue: person.DocValue,}
        res := database.Where("doc_value = ?", person.DocValue).First(&person)
        if !errors.Is(res.Error, gorm.ErrRecordNotFound) {
            w.WriteHeader(403)

            json.NewEncoder(w).Encode(util.Response{
                Message: "The user already exist!",
                Code: "AlreadyExist",
                Type: "error",
            })

            return
        }
    }

    person.ID = util.ToHash(person.DocValue)
    person.CreateTime = time.Now()
    person.UpdateTime = time.Now()

    database.Create(person)
    database.Commit()

    json.NewEncoder(w).Encode(util.Response {
        Type: "sucess",
        Code: "CreatedPerson",
        Data: Person {
            Name: person.Name,
            DocValue: person.DocValue,
            Type: person.Type,
            CreateTime: person.CreateTime,
            UpdateTime: person.UpdateTime,
        },
    })
}

func UpdatePerson(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)

    var body util.Request
    json.NewDecoder(r.Body).Decode(&body)

    person := Person{}
    res := database.Where("id = ?", params["id"]).First(&person)
    if errors.Is(res.Error, gorm.ErrRecordNotFound) {
        res = database.Where("doc_value = ?", params["id"]).First(&person)
        if errors.Is(res.Error, gorm.ErrRecordNotFound) {
            w.WriteHeader(404)

            json.NewEncoder(w).Encode(util.Response{
                Message: "The user not found!",
                Code: "NotFound",
                Type: "error",
                Data: nil,
            })

            return
        }
    }

    for i, prop := range(body.Data) {
        switch i {
        case "name":
            person.Name = prop
        case "document":
            person.DocValue = prop
        case "doc_type":
            person.DocType = prop
        case "type":
            person.Type = prop
            // TODO: sentence for validate logged user
        case "password":
            _, err := person.SetPass(prop)
            if err != nil {
                w.WriteHeader(500)

                json.NewEncoder(w).Encode(util.Response{
                    Message: "Error on creation of password hash on database!",
                    Code: "DbError",
                    Type: "error",
                })

                return
            }

        }
    }


    person.UpdateTime = time.Now()

    // TODO: sentence for validate logged user
    database.Save(&person)
    database.Commit()
    json.NewEncoder(w).Encode(util.Response{
        Code: "UpdatedUser",
        Type: "sucess",
        Data: Person {
            Name: person.Name,
            DocValue: person.DocValue,
            Type: person.Type,
            CreateTime: person.CreateTime,
            UpdateTime: person.UpdateTime,
        },
    })
}

// DeletePerson deleta um contato
func DeletePerson(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    person := Person{}
    res := database.Where("id = ?", params["id"]).First(&person)
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
    database.Delete(&person)
    database.Commit()
    json.NewEncoder(w).Encode(util.Response{
        Code: "DeleteUser",
        Type: "sucess",
        Data: Person {
            Name: person.Name,
            DocValue: person.DocValue,
            Type: person.Type,
            CreateTime: person.CreateTime,
            UpdateTime: person.UpdateTime,
        },
    })
}
