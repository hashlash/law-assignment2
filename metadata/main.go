package main

import (
    "os"
    "fmt"
    "encoding/json"
    "net/http"
    "./models"
)

func uploadData(w http.ResponseWriter, r *http.Request) {
    var fileData models.FileData
    err := json.NewDecoder(r.Body).Decode(&fileData)
    if err != nil {
        fmt.Println(err)
    }

    fmt.Println(fileData)

    fileData.Insert()
}

func setupRoutes() {
    http.HandleFunc("/upload/", uploadData)
    //http.HandleFunc("/files/", files)
    fmt.Println(http.ListenAndServe(":8888", nil))
}

func main() {
    fmt.Println("Hello World")
    models.InitDB(os.Getenv("DATABASE_URL"))
    setupRoutes()
}
