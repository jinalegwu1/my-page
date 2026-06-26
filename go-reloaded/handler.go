package main

import (
	"encoding/json"
	"net/http"
	"os"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if err := SCH.ExecuteTemplate(w, "search.html", nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	if r.Method == "POST" {
		
		server := convert(.FormValue("hex"))
			//Bin:  r.FormValue("dec"),
		}

		if err := RST.ExecuteTemplate(w, "result.html", server); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		result, err := json.Marshal(server)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	err = os.WriteFile("report.json", []byte(result), 0644)
	RST.Execute(w, result)
	}

