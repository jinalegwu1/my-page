package main

import (
	"net/http"
)

type Address struct {
	State  string
	City   string
	LGA    string
	Street string
}
type Student struct {
	FullName string
	Age      string
	Phone    string
	Address  Address
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		Tmpl.ExecuteTemplate(w, "index.html", nil)
		return
	}
	members := []Student{
		{
			FullName: r.FormValue("fullname"), 
			Age:      r.FormValue("age"),
			Phone:    r.FormValue("phone"),
			Address: Address{
				State:  r.FormValue("state"),
				City:   r.FormValue("city"),
				LGA:    r.FormValue("lga"),
				Street: r.FormValue("street"),
			},
		},
	}
	if err := Suc.ExecuteTemplate(w, "success.html", members); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// func StudentLoadHandler(w http.ResponseWriter, r *http.Request) {
// 	members := []Student{
// 		{
// 			FullName: r.FormValue("fullname"),
// 			Age:      r.FormValue("age"),
// 			Phone:    r.FormValue("phone"),
// 			Address: Address{
// 				State:  r.FormValue("state"),
// 				City:   r.FormValue("city"),
// 				LGA:    r.FormValue("lga"),
// 				Street: r.FormValue("street"),
// 			},
// 		},
// 	}
// 	if err := Succ.Execute(w, members); err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// }
