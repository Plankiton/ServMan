package util

import (
    "crypto/sha1"
    "fmt"
)

type Response struct {
    Message   string       `json:"message,omitempty"`
    Code      string       `json:"code,omitempty"`
    Type      string       `json:"type,omitempty"`
    Data      interface{}  `json:"data,omitempty"`
}

type Request struct {
    Token   string             `json:"message,omitempty"`
    Data    map[string]string  `json:"data,omitempty"`
}

func ToHash(s string) string {
    h := sha1.New()
    return fmt.Sprintf("%x", h.Sum([]byte(s)))
}
