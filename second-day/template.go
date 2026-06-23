package main

import "html/template"

var Tmpl = template.Must(
	template.ParseFiles("template/index.html"),
)