package main

import (
    "github.com/plankiton/mux"
    "log"
    "net/http"

    "./user"
    "./farm"
    "./serv"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

// função principal para executar a api
func main() {
    dsn := "host=localhost user=plankiton password=joaojoao dbname=servman port=5432 sslmode=disable TimeZone=America/Araguaina"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil && db != nil {
    }
    db.Migrator().CurrentDatabase()

    router := mux.NewRouter()

    // User
    user.PopulateDB(db)
    router.HandleFunc("/user", user.GetPeople).Methods("GET")

    router.HandleFunc("/user", user.CreatePerson).Methods("POST")
    router.HandleFunc("/user/{id}", user.GetPerson).Methods("GET")
    router.HandleFunc("/user/{id}", user.DeletePerson).Methods("DELETE")
    router.HandleFunc("/user/{id}", user.UpdatePerson).Methods("POST")
    log.Output(2, "Routing User operations")

    // Farm
    farm.PopulateDB(db)
    router.HandleFunc("/farm", farm.GetAllFarms).Methods("GET")
    router.HandleFunc("/user/{id}/farm", farm.GetFarms).Methods("GET")

    router.HandleFunc("/user/{id}/farm", farm.CreateFarm).Methods("POST")
    router.HandleFunc("/user/farm/{id}", farm.GetFarm).Methods("GET")
    router.HandleFunc("/user/farm/{id}", farm.DeleteFarm).Methods("DELETE")
    router.HandleFunc("/user/farm/{id}", farm.UpdateFarm).Methods("POST")
    log.Output(2, "Routing Farm operations")

    // Serv
    serv.PopulateDB(db)
    router.HandleFunc("/serv",      serv.GetAllServs).Methods("GET")
    router.HandleFunc("/user/{id}/serv",      serv.GetServs).Methods("GET")
    router.HandleFunc("/user/farm/{id}/serv", serv.GetServs).Methods("GET")

    router.HandleFunc("/user/{id}/serv",      serv.CreateServ).Methods("POST")
    router.HandleFunc("/user/farm/{id}/serv", serv.CreateServ).Methods("POST")

    router.HandleFunc("/serv/{id}", serv.GetServ).Methods("GET")
    router.HandleFunc("/serv/{id}", serv.DeleteServ).Methods("DELETE")
    router.HandleFunc("/serv/{id}", serv.UpdateServ).Methods("POST")
    log.Output(2, "Routing Service operations")

    log.Fatal(http.ListenAndServe(":8000", router))
}
