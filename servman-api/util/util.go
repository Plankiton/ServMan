package util

import (
    "golang.org/x/crypto/bcrypt"
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
func PassHash(s string) (string, error) {
    hash, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
    return string(hash), err
}

func CheckPass(p []byte, s string) (error) {
    err := bcrypt.CompareHashAndPassword(p, []byte(s))
    return err
}
