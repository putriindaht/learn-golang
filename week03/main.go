package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const baseUrl = "https://jsonplaceholder.typicode.com/posts/"

type ResponseData struct {
	Status       int `json:"status"`
	ResponseBody any `json:"responseBody"`
}

func getPostsHandler(w http.ResponseWriter, r *http.Request) {
	client := &http.Client{}

	// Create GET request
	req, err := http.NewRequest("GET", baseUrl, nil)
	if err != nil {
		fmt.Println("Error creating GET request:", err)
		http.Error(w, "Error creating GET request", http.StatusInternalServerError)
		return
	}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending GET request:", err)
		http.Error(w, "Error sending GET request", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		http.Error(w, "Failed to read response body", http.StatusInternalServerError)
		return
	}

	// Decode the response JSON into a slice of maps
	var responseBody []map[string]interface{}
	err = json.Unmarshal(body, &responseBody)
	if err != nil {
		fmt.Println("Failed to parse JSON response:", err)
		http.Error(w, "Failed to parse JSON response", http.StatusInternalServerError)
		return
	}

	// Wrap the data in your response struct
	response := ResponseData{
		Status:       resp.StatusCode,
		ResponseBody: responseBody,
	}

	// Optional: Pretty print the response for logging
	prettyJSON, err := json.MarshalIndent(responseBody, "", "  ")
	if err != nil {
		fmt.Println("Failed to format JSON for logging:", err)
	} else {
		fmt.Println("Response status:", resp.StatusCode)
		fmt.Println("Response body:", string(prettyJSON))
	}

	// Send the final response to the client
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	json.NewEncoder(w).Encode(response)
}

func postPostsHandler(w http.ResponseWriter, r *http.Request) {
	client := &http.Client{}

	// Prepare JSON data to create a new post
	jsonData := `{
		"title": "Post with Authorization",
		"body": "This post includes an authorization header.",
		"userId": 1
	}`

	// Create a new POST request
	req, err := http.NewRequest("POST", baseUrl, bytes.NewBuffer([]byte(jsonData)))
	if err != nil {
		fmt.Println("Error creating POST request:", err)
		http.Error(w, "Error creating POST request", http.StatusInternalServerError)
		return
	}

	// Set request headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer 123456")

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending POST request:", err)
		http.Error(w, "Error sending POST request", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to read POST response body:", err)
		http.Error(w, "Failed to read response body", http.StatusInternalServerError)
		return
	}

	// Decode the JSON response into a map
	var responseBody map[string]interface{}
	err = json.Unmarshal(bodyBytes, &responseBody)
	if err != nil {
		fmt.Println("Failed to parse JSON response:", err)
		http.Error(w, "Failed to parse JSON response", http.StatusInternalServerError)
		return
	}

	// Wrap the response for standardized output
	response := ResponseData{
		Status:       resp.StatusCode,
		ResponseBody: responseBody,
	}

	// Pretty print the response for logging
	prettyJSON, err := json.MarshalIndent(responseBody, "", "  ")
	if err != nil {
		fmt.Println("Failed to format JSON:", err)
	} else {
		fmt.Println("Response status:", resp.StatusCode)
		fmt.Println("Response body:", string(prettyJSON))
	}

	// Send the final response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	json.NewEncoder(w).Encode(response)
}

func putPostsHandler(w http.ResponseWriter, r *http.Request) {
	client := &http.Client{}

	// Prepare JSON data for update
	dataJson := `{
		"title": "Updated Post",
		"body": "This post has been updated.",
		"userId": 1
	}`

	// Build PUT request
	req, err := http.NewRequest("PUT", baseUrl+"/1", bytes.NewBuffer([]byte(dataJson)))
	if err != nil {
		fmt.Println("Error creating PUT request:", err)
		http.Error(w, "Error creating PUT request", http.StatusInternalServerError)
		return
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer 123456")

	// Send request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending PUT request:", err)
		http.Error(w, "Error sending PUT request", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Read response body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to read response body:", err)
		http.Error(w, "Failed to read response body", http.StatusInternalServerError)
		return
	}

	// Decode JSON body to map
	var responseBody map[string]interface{}
	err = json.Unmarshal(bodyBytes, &responseBody)
	if err != nil {
		fmt.Println("Failed to parse JSON:", err)
		http.Error(w, "Failed to parse JSON", http.StatusInternalServerError)
		return
	}

	// Wrap into standardized response
	response := ResponseData{
		Status:       resp.StatusCode,
		ResponseBody: responseBody,
	}

	// Pretty-print for logging
	prettyJSON, err := json.MarshalIndent(responseBody, "", "  ")
	if err != nil {
		fmt.Println("Failed to format JSON:", err)
	} else {
		fmt.Println("Response status:", resp.StatusCode)
		fmt.Println("Response body:", string(prettyJSON))
	}

	// Send final response to client
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	json.NewEncoder(w).Encode(response)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getPostsHandler(w, r)
	case http.MethodPost:
		postPostsHandler(w, r)
	case http.MethodPut:
		putPostsHandler(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

}

func main() {

	// handle routes '/posts'
	http.HandleFunc("/posts", postHandler)

	// start server
	fmt.Println("Server run in PORT: 8080")
	http.ListenAndServe(":8080", nil)

}
