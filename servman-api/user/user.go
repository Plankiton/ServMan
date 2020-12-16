package user

import (
    "encoding/json"
    "github.com/plankiton/mux"
    "net/http"
    "time"
    "errors"
    "fmt"

    "gorm.io/gorm"
    "../util"
)

type Role struct {
    ID         string    `json:"id,omitempty" gorm:"primaryKey"`
    Name       string    `json:"name"`
}

type RoleShip struct {
    RoleId     string    `json:"id,omitempty" gorm:"primaryKey"`
    PersonId   string    `json:"person_id,omitempty" gorm:"index"`
}

type Person struct {
    ID         string    `json:"id,omitempty" gorm:"primaryKey"`
    DocValue   string    `json:"document,omitempty" gorm:"uniqueIndex"`
    DocType    string    `json:"doc_type,omitempty"`
    Telephone  string    `json:"telephone,omitempty" gorm:"index,default:null"`
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
        roleships, types := []RoleShip{}, []string{}
        database.Where("person_id = ? ", p.ID).Find(&roleships)

        for _, r := range(roleships) {
            role := Role{}
            database.First(&role, r.RoleId)

            types = append(types, role.Name)
        }

        people = append(people, Person {
            ID: p.ID,
            Name: p.Name,
            DocValue: p.DocValue,
            Telephone: p.Telephone,
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

    roleships, types := []RoleShip{}, []string{}
    database.Where("person_id = ? ", person.ID).Find(&roleships)

    for _, r := range(roleships) {
        role := Role{}
        database.First(&role, r.RoleId)

        types = append(types, role.Name)
    }


    // TODO: sentence for validate logged user
    json.NewEncoder(w).Encode(util.Response{
        Code: "GetPerson",
        Type: "sucess",
        Data: Person {
            ID: person.ID,
            Name: person.Name,
            DocValue: person.DocValue,
            Telephone: person.Telephone,
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

    person.Name      = body.Data["name"]
    person.DocValue  = body.Data["document"]
    person.Telephone = body.Data["telephone"]
    person.DocType   = body.Data["doc_type"]

    role, ship := Role{}, RoleShip{}

    // TODO: sentence for validate logged user
    if body.Data["type"] == "root" {
        res := database.Find(&[]Person{})
        if res.RowsAffected > 0 {

            database.First(&role, "employee")

            ship.PersonId = person.ID
            ship.RoleId = role.ID

        } else {

            database.First(&role, "root")

            ship.PersonId = person.ID
            ship.RoleId = role.ID

        }
    }

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

    person.ID = util.ToHash(person.DocValue)
    person.CreateTime = time.Now()
    person.UpdateTime = time.Now()

    database.Create(person)
    database.Create(ship)
    database.Commit()

    roleships, types := []RoleShip{}, []string{}
    database.Where("person_id = ? ", person.ID).Find(&roleships)

    for _, r := range(roleships) {
        role := Role{}
        database.First(&role, r.RoleId)

        types = append(types, role.Name)
    }

    json.NewEncoder(w).Encode(util.Response {
        Type: "sucess",
        Code: "CreatedPerson",
        Data: Person {
            ID: person.ID,
            Name: person.Name,
            DocValue: person.DocValue,
            Telephone: person.Telephone,
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
        database.First(&role, r.RoleId)

        types = append(types, role.Name)
    }

    json.NewEncoder(w).Encode(util.Response{
        Code: "UpdatedUser",
        Type: "sucess",
        Data: Person {
            ID: person.ID,
            Name: person.Name,
            DocValue: person.DocValue,
            Telephone: person.Telephone,
            CreateTime: person.CreateTime,
            UpdateTime: person.UpdateTime,
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
        Data: Person {
            ID: person.ID,
            Name: person.Name,
            DocValue: person.DocValue,
            Telephone: person.Telephone,
            CreateTime: person.CreateTime,
            UpdateTime: person.UpdateTime,
        },
    })
}
