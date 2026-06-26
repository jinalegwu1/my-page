package main

import "html/template"

var SCH = template.Must(
	template.ParseFiles("template/search.html"),
)
var RST = template.Must(
	template.ParseFiles("template/result.html"),
)