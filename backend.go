package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bxcodec/faker/v4"
)

// Create Response Structure

type Response struct {
	Message string `json:"message"`
}

type RandomData struct {
	Name     string `json:"name"`
	Company  string `json:"company"`
	JobTitle string `json:"job_title"`
	Sentence string `json:"sentence"`
}

// Create a handler function
func send_message(w http.ResponseWriter, r *http.Request) {
	// Set content type to JSON
	w.Header().Set("Content-Type", "application/json")
	// Get the name from the query parameter
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}

	response := Response{
		Message: fmt.Sprintf("Hello %s", name),
	}

	// Encode the response as JSON
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Write the JSON response
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func generate_random_data(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	randomData := RandomData{
		Name:     faker.Name(),
		Company:  faker.Word(),
		JobTitle: faker.Word(),
		Sentence: faker.Sentence(),
	}
	// json.Marshal converts it to:
	// []byte(`{"name":"John Doe","company":"TechCorp","job_title":"Developer","sentence":"Hello World"}`)
	jsonResponse, err := json.Marshal(randomData)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func main() {
	http.HandleFunc("/generate-random", generate_random_data)
	http.HandleFunc("/hello", send_message)
	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Server failed to start: %v\n", err)
	}

}
