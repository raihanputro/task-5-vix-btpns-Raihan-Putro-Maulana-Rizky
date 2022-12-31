package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
)

type Photo struct {
	ID int `json:"ID"`
	Title string `json:"Title"`
	Caption string `json:"Caption"`
	PhotoUrl string `json:"PhotoUrl"`
	UserID int `json:UserID`
}

type Photos []Photo

func allPhotos(w http.ResponseWriter, r*http.Request) {
	photos := Photos{
		Photo{
			ID: 1, 
			Title: "Monas", 
			Caption: "Landmark Jakarta", 
			PhotoUrl: "https://commons.wikimedia.org/wiki/File:National_Monument_(Monas)_Jakarta.jpg", 
			UserID: 1,
		},
	}

	fmt.Println("Endpoint Hit: All Photos Endpoint")
	json.NewEncoder(w).Encode(photos)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage endpoint Hit")
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/photos", allPhotos)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequest()
}