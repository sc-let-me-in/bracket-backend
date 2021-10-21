package main

// 
import (
    "github.com/gin-gonic/gin"
    "log"
    "net/http"
)

// our main function
func main() {
    router := gin.Default()
    router.GET("/", func(context *gin.Context) {
        context.PureJSON(http.StatusOK, "Hello")
    })
    log.Fatal(http.ListenAndServe(":8000", router))
}
