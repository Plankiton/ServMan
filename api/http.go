package handler
import (
    "log"

    "github.com/plankiton/ServMan/api/user"
    "github.com/plankiton/ServMan/api/farm"
    "github.com/plankiton/ServMan/api/serv"
    "github.com/plankiton/ServMan/api/auth"

    "github.com/gorilla/mux"
    "gorm.io/gorm"
)

func HttpAPI(router *mux.Router, db *gorm.DB) {
    // Auth
    auth.PopulateDB(db)
    router.HandleFunc("/auth/login", auth.CreateToken).Methods("POST")
    router.HandleFunc("/auth/logout", auth.DeleteToken).Methods("POST")
    router.HandleFunc("/auth/check", auth.TestToken).Methods("POST")

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

    router.HandleFunc("/serv/{id}/mark", serv.MarkServTime).Methods("POST")
    router.HandleFunc("/serv/{id}", serv.GetServ).Methods("GET")
    router.HandleFunc("/serv/{id}", serv.DeleteServ).Methods("DELETE")
    router.HandleFunc("/serv/{id}", serv.UpdateServ).Methods("POST")
    log.Output(2, "Routing Service operations")
}
