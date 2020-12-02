package main

import (
    "net/http"
    "project/view"
)


func main() {
    http.HandleFunc("/users", view.UserView)

    http.ListenAndServe("127.0.0.1:8080", nil)
}
