package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"simple-crud-golang/api"
	"simple-crud-golang/api/language"
	"strconv"
)

func getPalindromeCheck(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	text := params["text"]
	isPalindrome := api.IsPalindrome(text)

	if (!isPalindrome) {
		w.WriteHeader(http.StatusMethodNotAllowed)
		io.WriteString(w, "Not Palindrome\n")
		return
	}

	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "Palindrome\n")
}

func getLanguagesEndpoint(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(language.GetLanguages())
}

func getLanguageEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	paramsId, err := strconv.Atoi(params["id"])
	if err != nil {
		// handle error
		fmt.Println(err)
		json.NewEncoder(w).Encode(language.Language{})
	}

	// the best way to check for an empty Person
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(language.GetLanguage(paramsId))
}

func postLanguageEndpoint(w http.ResponseWriter, req *http.Request) {
	var model language.Language
	_ = json.NewDecoder(req.Body).Decode(&model)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(language.CreateLanguage(model))
}

func patchLanguageEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	paramsId, err := strconv.Atoi(params["id"])
	if err != nil {
		// handle error
		fmt.Println(err)
		json.NewEncoder(w).Encode(language.Language{})
	}

	var model language.Language
	_ = json.NewDecoder(req.Body).Decode(&model)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(language.UpdateLanguage(paramsId, model))
}

func deleteLanguageEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	paramsId, err := strconv.Atoi(params["id"])
	if err != nil {
		// handle error
		fmt.Println(err)
		json.NewEncoder(w).Encode(language.Language{})
	}
	// the best way to check for an empty Person
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(language.DeleteLanguage(paramsId))
}

//Init data
func init() {
	language.CreateLanguage(language.Language{
		Language: "C",
		Appeared: 1972,
		Functional: true,
		ObjectOriented: false,
		Created: []string{"Dennis Ritchie"},
		Relation: &language.Relation{
			InfluencedBy: []string{"B", "ALGOL 68", "Assembly", "FORTRAN"},
			Influences: []string{"C++", "Objective-C", "C#", "Java", "JavaScript", "PHP", "Go"},
		},
	})
}

func main() {
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello Go developers\n")
	}

	router := mux.NewRouter()
	router.HandleFunc("/", helloHandler).Methods("GET")
	router.HandleFunc("/palindrome/{text}", getPalindromeCheck).Methods("GET")
	router.HandleFunc("/languages", getLanguagesEndpoint).Methods("GET")
	router.HandleFunc("/language/{id}", getLanguageEndpoint).Methods("GET")
	router.HandleFunc("/language", postLanguageEndpoint).Methods("POST")
	router.HandleFunc("/language/{id}", patchLanguageEndpoint).Methods("PATCH")
	router.HandleFunc("/language/{id}", deleteLanguageEndpoint).Methods("DELETE")

	log.Println("Listing for requests at http://localhost:8000/")
	log.Fatal(http.ListenAndServe(":8000", router))
}
