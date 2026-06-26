package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", HomeHandler)
	//http.HandleFunc("/student", StudentLoadHandler)
	fmt.Println("Server: Running Successfully!")
	http.ListenAndServe(":8080", nil)
}
