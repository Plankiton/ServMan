package main

import (
    "github.com/plankiton/mux"
    "log"
    "net/http"

    "./user"
    "./farm"
    "./serv"
)

// função principal para executar a api
func main() {
    router := mux.NewRouter()

    p := user.Person{
        ID: "1",
        Password: &user.Pass{},
        Document: &user.Doc{
            Type: "cpf",
            Value: "123456789",
        },
        Name: "Joao Da Silva",
    }
    p.SetPass("joao")
    user.AppendPerson(p)

    f := farm.Farm{
        ID: "1",
        PersonId: "1",
        Name: "Maria de cristo",
        Address: &farm.Addr{
            Street: "Joao de melo paia",
            City:   "Sao Paulo",
            State:  "SP",
            Code:   "123233445",
        },
    }
    farm.AppendFarm(f)

    // User
    router.HandleFunc("/user", user.GetPeople).Methods("GET")

    router.HandleFunc("/user", user.CreatePerson).Methods("POST")
    router.HandleFunc("/user/{id}", user.GetPerson).Methods("GET")
    router.HandleFunc("/user/{id}", user.DeletePerson).Methods("DELETE")
    router.HandleFunc("/user/{id}", user.UpdatePerson).Methods("POST")
    log.Output(2, "Routing User operations")

    // Farm
    router.HandleFunc("/user/{id}/farm", farm.GetFarms).Methods("GET")

    router.HandleFunc("/user/{id}/farm", farm.CreateFarm).Methods("POST")
    router.HandleFunc("/user/farm/{id}", farm.GetFarm).Methods("GET")
    router.HandleFunc("/user/farm/{id}", farm.DeleteFarm).Methods("DELETE")
    router.HandleFunc("/user/farm/{id}", farm.UpdateFarm).Methods("POST")
    log.Output(2, "Routing Farm operations")

    // Serv
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
