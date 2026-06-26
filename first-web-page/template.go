package main

import "html/template"

var Tmpl = template.Must(
	template.ParseFiles("templates/index.html"),
)
var Suc = template.Must(
	template.ParseFiles("templates/success.html"),
)