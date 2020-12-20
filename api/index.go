package api

import (
    "log"
    "net/http"

    "github.com/plankiton/ServMan/api"

    "github.com/gorilla/mux"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)


// função principal para executar a api
func main() {
    dsn := "host=localhost user=plankiton password=joaojoao dbname=servman port=5432 sslmode=disable TimeZone=America/Araguaina"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return
    }

    db.Migrator().CurrentDatabase()

    router := mux.NewRouter()

    // SockIoAPI(router, db)
    HttpAPI(router, db)

    log.Fatal(http.ListenAndServe(":8000", router))
}
