package main

// 
import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

// our main function
func main() {
    router := mux.NewRouter()
    router.HandleFunc("/test", GetTest).Methods("GET")
    log.Fatal(http.ListenAndServe(":8000", router))
}

func GetTest(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode("Hello")
}

