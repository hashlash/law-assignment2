package main

import (
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"net/http"
)

func compressHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(5 << 20)
	file, fileHandler, _ := r.FormFile("myFile")
	fileBytes, _ := ioutil.ReadAll(file)

	w.Header().Add(
		"Content-Disposition",
		fmt.Sprintf("attachment; filename=\"%v.gz\"", fileHandler.Filename),
	)

	zw, _ := gzip.NewWriterLevel(w, gzip.BestCompression)
	zw.Write(fileBytes)
	zw.Close()
}

func setupRoutes() {
	http.HandleFunc("/compress/", compressHandler)
	fmt.Println(http.ListenAndServe(":8887", nil))
}

func main() {
	fmt.Println("Hai")
	setupRoutes()
}
