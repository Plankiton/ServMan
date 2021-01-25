package api

import (
    "os"
    "net/http"
    "strings"
    "encoding/json"

    re "regexp"

    "github.com/gorilla/mux"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"

    "github.com/plankiton/ServMan/api/user"
    "github.com/plankiton/ServMan/api/serv"
    "github.com/plankiton/ServMan/api/util"
    "github.com/plankiton/ServMan/api/farm"
    "github.com/plankiton/ServMan/api/auth"
)

func getEnv(key string, def string) string {
    val, ok := os.LookupEnv(key)
    if !ok {
        return def
    }
    return val
}

func GetRouter() *mux.Router {
    r := mux.NewRouter()
    r.HandleFunc("/", Handler).Methods(http.MethodGet)
    return r
}

func Handler(w http.ResponseWriter, r *http.Request) {
    dsn := getEnv("DB_URI",
    "host=localhost "+
        "user=plankiton "+
        "password=joaojoao "+
        "dbname=servman "+
        "port=5432 "+
        "sslmode=disable "+
        "TimeZone=America/Araguaina")
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return
    }
    db.Migrator().CurrentDatabase()

    routes := map[string] (map[string]func(*http.Request)util.Response) {
        "/auth/login": {"POST": auth.CreateToken},
        "/auth/logout": {"POST": auth.DeleteToken},
        "/auth/check": {"POST": auth.TestToken},
        "/user": {
            "GET": user.GetPeople,
            "POST": user.CreatePerson,
        },
        "/user/{id}": {
            "GET": user.GetPerson,
            "DELETE": user.DeletePerson,
            "POST": user.UpdatePerson,
        },
        "/farm": {"GET": farm.GetAllFarms},
        "/user/{id}/farm": {
            "GET": farm.GetFarms,
            "POST": farm.CreateFarm,
        },
        "/farm/{id}": {
            "GET": farm.GetFarm,
            "DELETE": farm.DeleteFarm,
            "POST": farm.UpdateFarm,
        },
        "/farm/{id}/addr": {"GET": farm.GetAddr},
        "/addr/{cep}": {"GET": farm.GetAddrFromCep},
        "/serv": {"GET": serv.GetAllServs},
        "/user/{id}/serv": {
            "GET": serv.GetServs,
            "POST": serv.CreateServ,
        },
        "/farm/{id}/serv": {
            "GET": serv.GetServs, 
            "POST": serv.CreateServ,
        },
        "/serv/{id}/mark": {"POST": serv.MarkServTime},
        "/serv/{id}": {
            "GET": serv.GetServ,
            "DELETE": serv.DeleteServ,
            "POST": serv.UpdateServ,
        },
    };

    returned := false
    for i, route := range(routes) {
        action := route[r.Method];
        if action == nil {
            continue
        }

        get_res := func () {
            returned = true
            res := action(r)
            if res.Status != 0 {
                w.WriteHeader(res.Status)
            }

            json.NewEncoder(w).Encode(res)
        }

        if r.URL.Path == i {
            get_res()
            break
        }

        novars := re.MustCompile("\\{*\\}").Split(i, -1)

        pattern := strings.Join(novars, "*")
        match, _ := re.MatchString(pattern, r.URL.Path)
        if match {
            get_res()
            break
        }
    }

    if !returned {
        w.WriteHeader(404)
        json.NewEncoder(w).Encode(util.Response {
            Code: "NotFound",
            Status: 404,
            Message: "Route not found",
            Type: "Error",
        })
    }
}
