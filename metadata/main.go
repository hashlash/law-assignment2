package main

import (
	"./models"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func getOauthUser(token string) *string {
	client := &http.Client{}

	req, err := http.NewRequest("GET", os.Getenv("OAUTH_USER_RESOURCE"), nil)
	req.Header.Add("Authorization", token)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	if resp.StatusCode == http.StatusUnauthorized {
		return nil
	}

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	userId := result["user_id"].(string)

	return &userId
}

func uploadData(w http.ResponseWriter, r *http.Request) {
	userPtr := getOauthUser(r.Header.Get("Authorization"))
	if userPtr == nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var fileData models.FileData
	err := json.NewDecoder(r.Body).Decode(&fileData)
	if err != nil {
		fmt.Println(err)
	}

	fileData.Owner = *userPtr

	fmt.Println(fileData)

	fileData.Insert()
}

func files(w http.ResponseWriter, r *http.Request) {
	userPtr := getOauthUser(r.Header.Get("Authorization"))
	if userPtr == nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	dataId := r.URL.Query().Get("id")

	var fileData interface{}
	var err error

	if dataId == "" {
		fileData, err = models.AllFileData(*userPtr)
	} else {
		dataId, _ := strconv.Atoi(dataId)
		fileData, err = models.GetFileData(dataId, *userPtr)
	}

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No matching data")
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
