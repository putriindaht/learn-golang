package main

import (
	"fmt"
	"net/http"

	"github.com/putriindah/learn-golang-week03/handlers"
)

func main() {

	// handle routes '/posts'
	http.HandleFunc("/posts", handlers.PostHandler)

	// start server
	fmt.Println("Server run in PORT: 8080")
	http.ListenAndServe(":8080", nil)

}
