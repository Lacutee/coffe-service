package routers

import (
	"coffe-service/database"
	"coffe-service/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAllCoffe(w http.ResponseWriter, r *http.Request) {
	var coffe []models.Coffe

	err := database.Connector.Find(&coffe).Error

	if err != nil {
		fmt.Println(err)
		json.NewEncoder(w).Encode(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(coffe)
}

func CreateNewCoffe(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	var coffe models.Coffe

	err := json.Unmarshal(reqBody, &coffe)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	err = database.Connector.Create(&coffe).Error

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Coffe has been created")
}

func UpdateCoffeByID(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	var CoffeUpdate models.Coffe
	var Coffe models.Coffe
	id := mux.Vars(r)
	key := id["id"]

	err := json.Unmarshal(reqBody, &CoffeUpdate)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err)
		return
	}

	err = database.Connector.First(&Coffe, key).Error

	if err != nil {
		fmt.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	err = database.Connector.Model(&Coffe).Updates(&CoffeUpdate).Error

	if err != nil {
		fmt.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Coffe has been updated")
}
