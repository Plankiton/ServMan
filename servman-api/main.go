package main

import (
    "github.com/plankiton/mux"
    "log"
    "net/http"

    "./user"
)

// função principal para executar a api
func main() {
    router := mux.NewRouter()
    router.Name("ServMan API")

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

    p = user.Person{
        ID: "2",
        Password: &user.Pass{},
        Document: &user.Doc{
            Type: "cpf",
            Value: "123456789",
        },
        Name: "Joao Da Silva",
        Type: "root",
    }
    p.SetPass("joao")
    user.AppendPerson(p)

    router.HandleFunc("/user", user.GetPeople).Methods("GET")

    router.HandleFunc("/user", user.CreatePerson).Methods("POST")
    router.HandleFunc("/user/{id}", user.GetPerson).Methods("GET")
    router.HandleFunc("/user/{id}", user.DeletePerson).Methods("DELETE")
    router.HandleFunc("/user/{id}", user.UpdatePerson).Methods("POST")
    log.Output(2, "Routing /user - User operations")

    log.Fatal(http.ListenAndServe(":8000", router))
}
