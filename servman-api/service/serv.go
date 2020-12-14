package farm

import (
    "encoding/json"
    "github.com/plankiton/mux"
    "net/http"
    "time"
    "fmt"

    "../util"
)

type Serv struct {
    Description string  `json:"description,omitempty"`
    Price       float64 `json:"price,omitempty"`
    Begin_photo string  `json:"begin,omitempty"`
    End_photo   string  `json:"end,omitempty"`

    BeginTime time.Time `json:"started_at,omitempty"`
    EndTime time.Time `json:"finished_at,omitempty"`
}

