package user

import (
    "encoding/json"
    "github.com/gorilla/mux"
    "net/http"
    "time"
    "errors"
    "fmt"
    "strings"

    "gorm.io/gorm"
    "github.com/plankiton/ServMan/api/util"
)

type Role struct {
    ID         int    `json:"id,omitempty" gorm:"primaryKey"`
    Name       string `json:"name,omitempty" gorm:"index"`
}

type RoleShip struct {
    RoleId     int       `json:"id,omitempty" gorm:"primaryKey"`
    PersonId   string    `json:"person_id,omitempty" gorm:"index"`
}

type Person struct {
    ID         string    `json:"id,omitempty" gorm:"primaryKey"`
    DocValue   string    `json:"document,omitempty" gorm:"uniqueIndex"`
    DocType    string    `json:"doc_type,omitempty" gorm:"default:'cpf'"`
    Phone      string    `json:"phone,omitempty" gorm:"index,default:null"`
    Name       string    `json:"name,omitempty" gorm:"index"`

    PassHash   string    `json:"password,omitempty"`

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
    database.AutoMigrate(&Person{}, &Role{}, &RoleShip{})

    roles := []Role{}
    database.Find(&roles)
    if len(roles) == 0 {
        database.Create(&Role {
            Name:   "root",
        })
        database.Create(&Role {
            Name:   "employee",
        })
        database.Create(&Role {
            Name:   "client",
        })
        database.Create(&Role {
            Name:   "admin",
        })
    }

    database.Commit()
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
    people := [](map[string]interface{}){}
    for _, p := range(person) {
        roleships, types := []RoleShip{}, []string{}
        database.Where("person_id = ? ", p.ID).Find(&roleships)

        for _, r := range(roleships) {
            role := Role{}
            database.First(&role, r.RoleId)

            types = append(types, role.Name)
        }

        people = append(people, map[string]interface{} {
            "id": p.ID,
            "name": p.Name,
            "document": p.DocValue,
            "phone": p.Phone,
            "roles": types,
            "created_at": p.CreateTime,
            "updated_at": p.UpdateTime,
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

    roleships, types := []RoleShip{}, []string{}
    database.Where("person_id = ? ", person.ID).Find(&roleships)

    for _, r := range(roleships) {
        role := Role{}
        res := database.First(&role, r.RoleId)
        if !errors.Is(res.Error, gorm.ErrRecordNotFound) {
            types = append(types, role.Name)
        }
    }


    // TODO: sentence for validate logged user
    json.NewEncoder(w).Encode(util.Response{
        Code: "GetPerson",
        Type: "sucess",
        Data: map[string]interface{} {
            "id": person.ID,
            "name": person.Name,
            "document": person.DocValue,
            "phone": person.Phone,
            "roles": types,
            "created_at": person.CreateTime,
            "updated_at": person.UpdateTime,
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
    } else {
        needed := []string{"password", "name", "document", "phone"}
        for _, prop := range(needed) {
            if body.Data[prop] == "" {
                w.WriteHeader(400)

                json.NewEncoder(w).Encode(util.Response{
                    Message: fmt.Sprintf(`"%s" is required!`, prop),
                    Code: "BadRequest",
                    Type: "error",
                    Data: nil,
                })

                return
            }
        }
    }

    person := Person {}

    person.Name      = body.Data["name"]
    person.DocValue  = body.Data["document"]
    person.Phone = body.Data["phone"]
    person.DocType   = body.Data["doc_type"]
    if person.DocType == "" {
        person.DocType = "cpf"
    }

    person.ID = util.ToHash(person.DocValue)

    person.CreateTime = time.Now()
    person.UpdateTime = time.Now()

    _, err := person.SetPass(body.Data["password"])

    if err != nil {
        w.WriteHeader(500)

        json.NewEncoder(w).Encode(util.Response{
            Message: "Error on creation of password hash on database!",
            Code: "DbError",
            Type: "error",
        })

        return
    }

    res := database.Where("doc_value = ?", person.DocValue).First(&Person{})
    if !errors.Is(res.Error, gorm.ErrRecordNotFound) {
        if res.Error != nil {
            w.WriteHeader(400)

            json.NewEncoder(w).Encode(util.Response{
                Message: fmt.Sprintf("Database error: %s", res.Error),
                Code: "BadRequest",
                Type: "error",
                Data: nil,
            })

            return
        }

        w.WriteHeader(403)

        json.NewEncoder(w).Encode(util.Response{
            Message: "The user already exist!",
            Code: "AlreadyExist",
            Type: "error",
        })

        return
    }


    if body.Data["type"] != "" {
        prop := body.Data["type"];
        roleships := []RoleShip{}
        database.Where("person_id = ? ", person.ID).Find(&roleships)
        for _, r := range(roleships) {
            database.Delete(r)
        }

        types := strings.Split(prop, ",")

        for _, r := range(types) {
            role, ship := Role{}, RoleShip{}

            if r == "root" {

                // TODO: Substituir por checagem de auth

                res := database.Find(&[]Person{})
                if res.RowsAffected > 0 {

                    database.Where("name = ?", "root").First(&role)
                    database.Where("role_id = ? AND person_id = ?",
                    role.ID, person.ID).First(&ship)

                    ship.PersonId = person.ID
                    ship.RoleId = role.ID

                } else {

                    database.Where("name = ?", "admin").First(&role)

                    ship.PersonId = person.ID
                    ship.RoleId = role.ID

                }
            } else {
                database.Where("name = ?",  r).First(&role)

                ship.PersonId = person.ID
                ship.RoleId = role.ID
                if role.Name == "" {
                    database.First(&role, "client")

                    ship.PersonId = person.ID
                    ship.RoleId = role.ID
                }
            }

            if ship.PersonId != "" && ship.RoleId != 0 {
                database.Create(&ship)
            }
        }
    }

    roleships, types := []RoleShip{}, []string{}
    database.Where("person_id = ? ", person.ID).Find(&roleships)

    for _, r := range(roleships) {
        role := Role{}
        database.Where("id = ?", r.RoleId).First(&role)

        types = append(types, role.Name)
    }

    database.Create(person)
    database.Commit()

    json.NewEncoder(w).Encode(util.Response {
        Type: "sucess",
        Code: "CreatedPerson",
        Data: map[string]interface{} {
            "id": person.ID,
            "name": person.Name,
            "document": person.DocValue,
            "phone": person.Phone,
            "roles": types,
            "created_at": person.CreateTime,
            "updated_at": person.UpdateTime,
        },
    })
}

func UpdatePerson(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)

    var body util.Request
    json.NewDecoder(r.Body).Decode(&body)

    person := Person{}
    res := database.Where("id = ? OR doc_value = ?", params["id"], params["id"]).First(&person)
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

        for i, prop := range(body.Data) {
            switch i {
            case "name":
                person.Name = prop
            case "document":
                person.DocValue = prop
            case "doc_type":
                person.DocType = prop
            case "type":

                // TODO: sentence for validate logged user

                {
                    roleships := []RoleShip{}
                    database.Where("person_id = ? ", person.ID).Find(&roleships)
                    for _, r := range(roleships) {
                        database.Delete(r)
                    }
                }

                types := strings.Split(prop, ",")

                for _, r := range(types) {
                    role, ship := Role{}, RoleShip{}
                    ship_exists := false

                    if r == "root" {
                        // TODO: Substituir por checagem de auth
                        res := database.Find(&[]Person{})
                        if res.RowsAffected > 0 {

                            database.Where("name = ?", "root").First(&role)
                            if database.Where("role_id = ? AND person_id = ?",
                            role.ID, person.ID).First(&ship) == nil {
                                ship_exists = true
                            }

                            ship.PersonId = person.ID
                            ship.RoleId = role.ID

                        } else {

                            database.Where("name = ?", "admin").First(&role)

                            ship.PersonId = person.ID
                            ship.RoleId = role.ID

                        }
                    } else {
                        database.Where("name = ?",  r).First(&role)

                        ship.PersonId = person.ID
                        ship.RoleId = role.ID
                        if role.Name == "" {
                            database.Where("id = ?", "client").First(&role)

                            ship.PersonId = person.ID
                            ship.RoleId = role.ID
                        }
                    }

                    if ship.PersonId != "" && ship.RoleId != 0 {
                        if ship_exists {
                            database.Save(&ship)
                        } else {
                            database.Create(&ship)
                        }
                    }
                }

                database.Commit()

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

        roleships, types := []RoleShip{}, []string{}
        database.Where("person_id = ? ", person.ID).Find(&roleships)

        for _, r := range(roleships) {
            role := Role{}
            database.Where("id = ?", r.RoleId).First(&role)

            types = append(types, role.Name)
        }

        json.NewEncoder(w).Encode(util.Response{
            Code: "UpdatedUser",
            Type: "sucess",
            Data: map[string]interface{} {
                "id": person.ID,
                "name": person.Name,
                "document": person.DocValue,
                "phone": person.Phone,
                "roles": types,
                "createtime": person.CreateTime,
                "updatetime": person.UpdateTime,
            },
        })
    }

    // DeletePerson deleta um contato
    func DeletePerson(w http.ResponseWriter, r *http.Request) {
        params := mux.Vars(r)
        person := Person{}
        res := database.Where("id = ? OR doc_value = ?", params["id"], params["id"]).First(&person)
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

        roleships, types := []RoleShip{}, []string{}
        database.Where("person_id = ? ", person.ID).Find(&roleships)

        for _, r := range(roleships) {
            role := Role{}
            database.First(&role, r.RoleId)

            types = append(types, role.Name)
        }

        json.NewEncoder(w).Encode(util.Response{
            Code: "DeleteUser",
            Type: "sucess",
            Data: map[string]interface{} {
                "id": person.ID,
                "name": person.Name,
                "document": person.DocValue,
                "phone": person.Phone,
                "roles": types,
                "createtime": person.CreateTime,
                "updatetime": person.UpdateTime,
            },
        })
    }
