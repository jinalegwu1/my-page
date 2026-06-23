package main

import (
	"net/http"
)

// Tmpl represents your parsed template. It is declared here to avoid compilation errors.
//var Tmpl *template.Template

type Welcome struct {
	SchoolName string
	RegNumber  string
	Students   []Student
}

type Address struct {
	Country string
	State   string
	City    string
	LGA     string
	Street  string
}

type Student struct {
	Fullname    string
	Age         int
	Height      float64
	Email       string
	PhoneNumber string
	Language    string
	Id          string
	Address     Address
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	welcome := Welcome{
		SchoolName: "Glorious Nursery & Primary School", // Added quotes
		RegNumber:  "BN/ot/3245",                        // Added quotes
		Students: []Student{ // Fixed struct name from []students
			{
				Fullname:    "Adah Adanu",
				Age:         32,
				Height:      3.6,
				Email:       "adahadanu0@gmail.com",
				PhoneNumber: "09012345678", 
				Language:    "Idoma",
				Id:          "001",
				Address: Address{
					Country: "Nigeria",
					State:   "Benue State",
					City:    "Otukpa",
					LGA:     "Ogbadibo Local Government Area", // Fixed comma spacing
					Street:  "No 32, Orokam Street Otukpa",
				},
			},
			{
				Fullname:    "Audu Adamu",
				Age:         25,
				Height:      5.6,
				//Email:       "auduhadamu0@gmail.com",
				PhoneNumber: "08087654321", 
				Language:    "Idoma",
				Id:          "002",
				Address: Address{
					Country: "Nigeria",
					State:   "Benue State",
					City:    "Otukpo",
					LGA:     "Otukpo Local Government Area", 
					Street:  "No 32, Abel Ujah Street Otukpo",
				},
			},
			{
				Fullname:    "Gabriel Ogah",
				Age:         37,
				Height:      4.6,
				Email:       "Gabrelogah@gmail.com",
				//PhoneNumber: "", 
				Language:    "Idoma, Igala, Jukun",
				Id:          "003",
				Address: Address{
					Country: "Nigeria",
					State:   "Benue State",
					City:    "Ohimini",
					LGA:     "Ohimini Local Government Area", 
					Street:  "No 32, Orokam Street Otukpa",
				},
			},
		},
	}

	if err := Tmpl.Execute(w, welcome); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
