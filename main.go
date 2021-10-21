package main

import (
    "bracket-backend/routers"
    "log"
    "net/http"
)

// our main function
func main() {
    router := routers.InitRouter()
    log.Fatal(http.ListenAndServe(":8000", router))
}
