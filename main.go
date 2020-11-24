package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type am struct {
	Title string
	Qty string
}

type Spacecraft struct {
	ID int8 `json:"id"`
	Name string `json:"name"`
	Class string `json:"class"`
	Armament []am `json:"armament"`
	Crew int8 `json:"crew"`
	Image string `json:"image"`
	Value float32 `json:"value"`
	Status string `json:"status"`
}

var fleet []Spacecraft

// Get entire fleet
func getFleet(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(fleet)
}

// Get Spacecraft
func getSpacecraft(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range fleet {
		if item.Name == params["name"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Spacecraft{})
}

// Create Spacecraft
func createSpacecraft(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	var spacecraft Spacecraft
	_ = json.NewDecoder(r.Body).Decode(&spacecraft)
	fleet = append(fleet, spacecraft)
    json.NewEncoder(w).Encode(spacecraft)
}

// Edit/Update existing Spacecraft
func updateSpacecraft(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, item := range fleet{
		if item.Name == params["name"] {
			fleet = append(fleet[:i], fleet[i+1:]...)
			var spacecraft Spacecraft
			_ = json.NewDecoder(r.Body).Decode(&spacecraft)
			spacecraft.Name = params["name"]
			fleet = append(fleet, spacecraft)
			json.NewEncoder(w).Encode(spacecraft)
			return
		}
	}
}

// Delete Spacecraft
func deleteSpacecraft(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, item := range fleet {
		if item.Name == params["name"] {
			fleet = append(fleet[:1], fleet[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(fleet)
}

func main()  {
	router := mux.NewRouter()


	// Handlers
	router.HandleFunc("/fleet", getFleet).Methods("GET")
	router.HandleFunc("/fleet/{name}", getSpacecraft).Methods("GET")
	router.HandleFunc("/fleet", createSpacecraft).Methods("POST")
	router.HandleFunc("/fleet/{name}", updateSpacecraft).Methods("PUT")
	router.HandleFunc("/fleet/{name}", deleteSpacecraft).Methods("DELETE")

	// Start server
	log.Fatal(http.ListenAndServe(":3000", router))
}
