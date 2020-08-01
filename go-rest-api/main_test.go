package main_test

import (
	"bytes"
	"encoding/json"
	. "go-rest-api/models"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestGetRecipesAPIURL(t *testing.T) {
	t.Log("Testing GET URL for recipes.")
	response, err := http.Get("http://localhost:8080/recipes")
	if err != nil {
		t.Errorf("Error thrown: %s", err)
	}
	body, err := ioutil.ReadAll(response.Body)
	t.Log("Testing GET URL status code for recipes.")
	if response.StatusCode != 200 && response.StatusCode != 201 {
		t.Errorf("Expected response code %d or %d, but got %d\n", http.StatusOK, response.StatusCode, response.StatusCode)
	}
	t.Log("Testing GET URL for JSON.")
	if err != nil {
		t.Errorf("Expected to read the response body.")
	}

	t.Log("Testing GET URL for a non-empty body.")
	if len(body) == 0 {
		t.Errorf("Expected the response body was not empty.")
	}
}

func TestPostRecipesNoAuthAPIURL(t *testing.T) {
	t.Log("Testing POST URL for recipes.")
	url := "http://localhost:8080/recipes"

	var jsonString = []byte(`{"name": "TestSteak2","prep_time": "1. Cook it. 2. Bath it in sage, rosemary and thyme.  3. Then eat it.","difficulty": "2","vegetarian": "false"}`)
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonString))
	request.Header.Set("X-api-key", "")
	request.Header.Set("Content-Type", "application/json")

	t.Log("Calling URL: ", url)
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		t.Error(err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	t.Log("Response Body: ", string(body))
	if string(body) != "{\"error\":\"API key is missing.\"}" {
		t.Errorf("Expected the response body: API key is missing.")
	}
}

func TestPostRecipesAuthorizedAPIURL(t *testing.T) {
	t.Log("Testing POST URL for recipes.")
	url := "http://localhost:8080/recipes"

	var jsonString = []byte(`{"name": "TestSteak2","prep_time": "1. Cook it. 2. Bath it in sage, rosemary and thyme.  3. Then eat it.","difficulty": "2","vegetarian": "false"}`)
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonString))
	request.Header.Set("X-api-key", "apiKey12345")
	request.Header.Set("Content-Type", "application/json")

	t.Log("Calling URL: ", url)
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		t.Error(err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	t.Log("Response Body: ", string(body))
	var recipe Recipe
	jsonErr := json.Unmarshal(body, &recipe)
	if jsonErr != nil {
		t.Error(jsonErr)
	}
	if string(body) == "{\"error\":\"API key is not authorized.\"}" {
		t.Errorf("Expected the response body not to be: API key is not authorized.")
	}
	if recipe.Name == "" {
		t.Errorf("Expected the response body JSON to have property of Name to be not empty.")
	}
	t.Log(recipe.Name)
}
