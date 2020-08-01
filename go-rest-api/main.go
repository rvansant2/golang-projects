package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	. "go-rest-api/config"
	. "go-rest-api/dataObjects"
	. "go-rest-api/models"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

var config = Config{}
var rDO = RecipesDataObject{}

func validateAPIKey(next http.HandlerFunc) http.HandlerFunc {
	config.Read()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		keys := r.Header.Get("x-api-key")
		if len(keys) < 1 {
			respondWithError(w, http.StatusUnauthorized, "API key is missing.")
			return
		}

		if string(keys) != config.APIKey {
			respondWithError(w, http.StatusBadRequest, "API key is not authorized.")
			return
		}
		next.ServeHTTP(w, r)
	})
}

func GetAllRecipes(w http.ResponseWriter, r *http.Request) {
	recipes, err := rDO.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, recipes)
}

func CreateRecipe(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var recipe Recipe
	if err := json.NewDecoder(r.Body).Decode(&recipe); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload.")
		return
	}
	recipe.UniqueID = bson.NewObjectId()
	if err := rDO.Insert(recipe); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, recipe)
}

func FindRecipeById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	recipe, err := rDO.FindById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid recipe id.")
		return
	}
	respondWithJson(w, http.StatusOK, recipe)
}

func UpdateRecipeById(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	var recipe Recipe
	if err := json.NewDecoder(r.Body).Decode(&recipe); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload.")
		return
	}
	if err := rDO.Update(params["id"], recipe); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func DeleteRecipeById(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	if err := rDO.Delete(params["id"]); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func RateRecipeById(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	var rating Rating
	if err := json.NewDecoder(r.Body).Decode(&rating); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload.")
		return
	}

	if err := rDO.RateRecipeById(params["id"], rating); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})

}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func init() {
	config.Read()
	rDO.Server = config.Server
	rDO.Database = config.Database
	rDO.Connect()
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/recipes", GetAllRecipes).Methods("GET")
	r.HandleFunc("/recipes", validateAPIKey(CreateRecipe)).Methods("POST")
	r.HandleFunc("/recipes/{id}", FindRecipeById).Methods("GET")
	r.HandleFunc("/recipes/{id}", validateAPIKey(UpdateRecipeById)).Methods("PUT")
	r.HandleFunc("/recipes/{id}", validateAPIKey(DeleteRecipeById)).Methods("DELETE")
	r.HandleFunc("/recipes/{id}/rating", RateRecipeById).Methods("POST")
	fmt.Println("Starting up on 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
