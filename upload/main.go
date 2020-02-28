package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
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

	tempFile, err := ioutil.TempFile("files", "*-"+handler.Filename)
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	client := &http.Client{}
	reqBody, _ := json.Marshal(map[string]string{
		"FileName": handler.Filename,
		"Path":     tempFile.Name(),
	})
	u, _ := url.Parse(os.Getenv("METADATA_HOST"))
	u.Path = path.Join(u.Path, "upload") + "/"
	req, _ := http.NewRequest("POST", u.String(), bytes.NewBuffer(reqBody))
	req.Header.Add("Authorization", r.Header.Get("Authorization"))
	resp, _ := client.Do(req)

	if resp.StatusCode == http.StatusUnauthorized {
		os.Remove(tempFile.Name())
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

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
	fmt.Println(http.ListenAndServe(":8889", nil))
}

func main() {
	fmt.Println("Hello World")
	setupRoutes()
}
