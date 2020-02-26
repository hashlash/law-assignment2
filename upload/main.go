package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Uploading File...\n")

    r.ParseMultipartForm(5 << 20)

    file, handler, err := r.FormFile("myFile")
    if err != nil {
        fmt.Println("Error retrieving the file")
        fmt.Println(err)
        return
    }
    defer file.Close()

    fmt.Printf("Uploaded File: %+v\n", handler.Filename)
    fmt.Printf("File Size: %+v\n", handler.Size)
    fmt.Printf("MIME Header: %+v\n", handler.Header)

    tempFile, err := ioutil.TempFile("files", "upload-*.png")
    if err != nil {
        fmt.Println(err)
    }
    defer tempFile.Close()

    fileBytes, err := ioutil.ReadAll(file)
    if err != nil {
        fmt.Println(err)
    }

    tempFile.Write(fileBytes)
    fmt.Printf("Uploaded to: %v\n", tempFile.Name())

    fmt.Fprintf(w, "Successfully Uploaded File\n")
}

func downloadFile(w http.ResponseWriter, r *http.Request) {
    filename := r.URL.Query().Get("filename")

    file, err := ioutil.ReadFile("files/" + filename)
    if err != nil {
        fmt.Println(err)
    }

    w.Write(file)
}

func setupRoutes() {
    http.HandleFunc("/upload/", uploadFile)
    http.HandleFunc("/download/", downloadFile)
    //http.Handle("/download/", http.FileServer(http.Dir("./files")))
    fmt.Println(http.ListenAndServe(":8888", nil))
}

func main() {
    fmt.Println("Hello World")
    setupRoutes()
}
