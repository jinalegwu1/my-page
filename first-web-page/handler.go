package main

import "net/http"

type Address struct{
	State string
	City string
	LGA string
	Street string
}
type Student struct{
	FullName string
	Age int
	Phone string
	Address Address
	

}
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if err := Tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
}

func StudentLoadHandler(w http.ResponseWriter, r *http.Request){
	student := []Student{
		{
			FullName: "Gabriel Odeh",
			Phone: "09060908844",
			Age: 23,
			Address: Address{
				State: "Benue State",
				City: "Otukpo",
				LGA: "Otukpo, Local Government Area.",
				Street: "No_ 45, David Otafu, Street",
			},

		},
		{
			FullName: "John Okoh",
			Phone: "07012345665",
			Age: 20,
			Address: Address{
				State: "Benue State",
				City: "Awume",
				LGA: "Ohimini, Local Government Area.",
				Street: "No_12, LeefCenter Road",
			},
		},
		{
			FullName: "Nnena Chioma",
			Age: 20,
			Phone: "08120204040",
			Address: Address{
				State: "Enugu",
				City: "Ubolo",
				LGA: "Ebanor, Local Government Area.",
				Street: "No_23, igwe Iyala, Street",
			},
		},
		
	}
	if err := Tmpl.Execute(w, student); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
}