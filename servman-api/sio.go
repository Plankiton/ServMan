package main

import (
    "fmt"
    "net/http"

    "gorm.io/gorm"
    "github.com/plankiton/mux"
    sio "github.com/googollee/go-socket.io"
)

func SockIoAPI(router *mux.Router, db *gorm.DB) {
    server := sio.NewServer(nil)

    server.OnConnect("/", func(s sio.Conn) error {
        s.SetContext("")
        fmt.Println("connected:", s.ID())
        return nil
    })

    server.OnEvent("/", "notice", func(s sio.Conn, msg string) {
        fmt.Println("notice:", msg)
        s.Emit("reply", "have "+msg)
    })

    server.OnEvent("/chat", "msg", func(s sio.Conn, msg string) string {
        s.SetContext(msg)
        return "recv " + msg
    })

    server.OnEvent("/", "bye", func(s sio.Conn) string {
        last := s.Context().(string)
        s.Emit("bye", last)
        s.Close()
        return last
    })

    server.OnError("/", func(s sio.Conn, e error) {
        fmt.Println("meet error:", e)
    })

    server.OnDisconnect("/", func(s sio.Conn, reason string) {
        fmt.Println("closed", reason)
    })

    go server.Serve()
    defer server.Close()

    router.Handle("/sio/", server)
    router.Handle("/", http.FileServer(http.Dir("./upload")))
}
