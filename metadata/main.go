package main

import (
    "os"
    "fmt"
    "strconv"
    "database/sql"
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

func files(w http.ResponseWriter, r *http.Request) {
    dataId := r.URL.Query().Get("id")

    var fileData interface{}
    var err error

    if dataId == "" {
         fileData, err = models.AllFileData("admin")
    } else {
         dataId, _ := strconv.Atoi(dataId)
         fileData, err = models.GetFileData(dataId, "admin")
    }

    switch err {
         case sql.ErrNoRows:
              fmt.Println("No rows")
         case nil:
              jsonData, _ := json.Marshal(fileData)
              w.Header().Set("Content-Type", "application/json")
              w.Write(jsonData)
         default:
              fmt.Println(err)
    }
}

func setupRoutes() {
    http.HandleFunc("/upload/", uploadData)
    http.HandleFunc("/files/", files)
    fmt.Println(http.ListenAndServe(":8888", nil))
}

func main() {
    fmt.Println("Hello World")
    models.InitDB(os.Getenv("DATABASE_URL"))
    setupRoutes()
}
