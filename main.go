package main

import (
    "os"
    "log"
    "net/http"

    "github.com/plankiton/ServMan/api"

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
func main() {
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

    router := mux.NewRouter()

    // SockIoAPI(router, db)
    HttpAPI(router, db)
    p := getEnv("HTTP_PORT", "8000")
    log.Fatal(http.ListenAndServe(p, router))
}
