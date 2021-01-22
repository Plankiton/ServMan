package handler

import (
    "os"
    "log"
    "net/http"

    "strings"
    re "regexp"
    "github.com/gorilla/mux"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

func getEnv(key string, def string) string {
    val, ok := os.LookupEnv(key)
    if !ok {
        return def
    }
    return val
}

// função principal para executar a api
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

    routes := map[string] func(*http.Request)util.Response {
        "/auth/login": auth.CreateToken,
        "/auth/logout": auth.DeleteToken,
        "/auth/check": auth.TestToken,
        "/user": user.GetPeople,
        "/user": user.CreatePerson,
        "/user/{id}": user.GetPerson,
        "/user/{id}": user.DeletePerson,
        "/user/{id}": user.UpdatePerson,
        "/farm": farm.GetAllFarms,
        "/user/{id}/farm": farm.GetFarms,
        "/user/{id}/farm": farm.CreateFarm,
        "/farm/{id}": farm.GetFarm,
        "/farm/{id}": farm.DeleteFarm,
        "/farm/{id}": farm.UpdateFarm,
        "/farm/{id}/addr": farm.GetAddr,
        "/addr/{cep}": farm.GetAddrFromCep,
        "/serv": serv.GetAllServs,
        "/user/{id}/serv": serv.GetServs,
        "/farm/{id}/serv": serv.GetServs,
        "/user/{id}/serv": serv.CreateServ,
        "/user/farm/{id}/serv": serv.CreateServ,
        "/serv/{id}/mark": serv.MarkServTime,
        "/serv/{id}": serv.GetServ,
        "/serv/{id}": serv.DeleteServ,
        "/serv/{id}": serv.UpdateServ,
    };

    for i, route := range(routes) {
        if r.path == i {
            res := route(r)
            if res.Status != 0 {
                w.WriteHeader(res.Status)
            }

            json.NewEncoder(w).Write(res)
            return
        }

        novars := re.MustCompile("\\{*\\}").Split(i, -1)
        values := []string{};

        for i :=0 ; i < len(novars) ; i += len(match) {
        }

    }
}
