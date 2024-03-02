package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Animals struct {
	ID             string `json:"id"`
	Kind_Of_Animal string `json:"kind_of_animal"`
	Kind_Of_Breed  string `json:"kind_of_breed"`
	Name           string `json:"name"`
	Age            string `json:"age"`
	Description    string `json:"description"`
}

type Users struct {
	UserID               string `json:"userid"`
	User_Email           string `json:"user_email"`
	Username             string `json:"username"`
	Password             string `json:"password"`
	Number_of_phone_user string `json:"number_of_phone_user"`
}

var animal []Animals

func getAnimalList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(animal)
}

func getAnimal(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range animal {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Animals{})
}

func addAnimal(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var fluf Animals
	_ = json.NewDecoder(r.Body).Decode(&fluf)
	fluf.ID = strconv.Itoa(rand.Intn(1000000))
	animal = append(animal, fluf)
	json.NewEncoder(w).Encode(fluf)
}

func updateAnimal(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range animal {
		if item.ID == params["id"] {
			animal = append(animal[:index], animal[index+1:]...)
			var fluf Animals
			_ = json.NewDecoder(r.Body).Decode(&fluf)
			fluf.ID = params["id"]
			animal = append(animal, fluf)
			json.NewEncoder(w).Encode(fluf)
			return
		}
	}
	json.NewEncoder(w).Encode(animal)
}

func deleteAnimal(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range animal {
		if item.ID == params["id"] {
			animal = append(animal[:index], animal[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(animal)
}

func main() {
	r := mux.NewRouter()
	animal = append(animal, Animals{ID: "1", Kind_Of_Animal: "Cat", Kind_Of_Breed: "Siamese", Name: "Mittens", Age: "3", Description: "Playful and affectionate, enjoys chasing toys."})
	animal = append(animal, Animals{ID: "2", Kind_Of_Animal: "Dog", Kind_Of_Breed: "Golden Retriever", Name: "Max", Age: "2", Description: "Friendly and loyal, loves being around people."})
	r.HandleFunc("/animal", getAnimalList).Methods("GET")
	r.HandleFunc("/animal/{id}", getAnimal).Methods("GET")
	r.HandleFunc("/animal", addAnimal).Methods("POST")
	r.HandleFunc("/animal/{id}", updateAnimal).Methods("PUT")
	r.HandleFunc("/animal/{id}", deleteAnimal).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}
