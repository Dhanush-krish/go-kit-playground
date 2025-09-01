package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	// generated docs package

	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/gorilla/mux"
)

type addResponse struct {
	Sum int `json:"sum"`
}

func AddTwoNumber(num1, num2 int) int {
	return num1 + num2
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Come back after a month :), site is under construction! \n*****AVAILABLE API's***** \n/hello/{name} - sends greeting \n/add - adds two integer"))

}

// HelloHandler godoc
// @Summary Greet a user
// @Description Greets a user by name
// @Param name path string true "Name of the user"
// @Success 200 {string} string "Greeting message"
// @Router /hello/{name} [get]
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	fmt.Fprintf(w, "Hello, %s!\n", name)
}

// Home Handler
// @Summary show home page
// @Description Get a welcome page
// @Router / [get]
func AddHandler(w http.ResponseWriter, r *http.Request) {
	aStr := r.URL.Query().Get("a")
	bStr := r.URL.Query().Get("b")

	a, aErr := strconv.Atoi(aStr)
	b, bErr := strconv.Atoi(bStr)

	if aErr != nil || bErr != nil {
		http.Error(w, "Invalid parameters. Use /add?a=5&b=10", http.StatusBadRequest)
		return
	}

	resp := addResponse{Sum: AddTwoNumber(a, b)}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)

}

// @title Add API Example
// @version 1.0
// @description This is a sample API using gorill Mux and Swgger
// @host localhost:8080
// @BasePath /
func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler).Methods("GET")
	r.HandleFunc("/hello/{name}", HelloHandler).Methods("GET")
	r.HandleFunc("/add", AddHandler).Methods("GET")

	// Swagger endpoint
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// Start server
	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", r)

}
