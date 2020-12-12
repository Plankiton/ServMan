package farm

import (
    "encoding/json"
    "github.com/plankiton/mux"
    "net/http"
    "time"
    "fmt"

    "../util"
)

type Addr struct {
    Street       string `json:"street,omitempty"`
    Country      string `json:"country,omitempty"`
    Number       string `json:"number,omitempty"`
    Code         string `json:"code,omitempty"`
    City         string `json:"city,omitempty"`
    Neightbourn  string `json:"neightbourn,omitempty"`
}

type Farm struct {
    ID      string `json:"id,omitempty"`
    Address  *Addr `json:"address,omitempty"`
    Name    string `json:"name,omitempty"`

    CreateTime time.Time `json:"created_at,omitempty"`
    UpdateTime time.Time `json:"updated_at,omitempty"`
}


