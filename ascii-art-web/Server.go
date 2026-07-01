package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", homeHandler)
	fmt.Println("Server: Running Successfully!")
	http.ListenAndServe(":8080", nil)
}
