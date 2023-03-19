package main

import (
	"fmt"
	"github.com/mw/go-course/pkg/handlers"
	"net/http"
)

const portNumber = ":8080"

// main is the main function
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))
	httpError := http.ListenAndServe(portNumber, nil)
	if httpError != nil {
		fmt.Println(httpError)
	}
}
