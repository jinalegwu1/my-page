package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", homeHandler)
	//http.HandleFunc("/students")
	fmt.Println("Server: Running Successuffully!")
	http.ListenAndServe(":8080", nil)
}