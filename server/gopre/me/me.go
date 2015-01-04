package me

import (
	"net/http"
)

func init() {
	router()
}

func router() {
	http.HandleFunc("/me/", dashboardHandler)
	http.HandleFunc("/me/slides", slideHandler)
	http.HandleFunc("/me/templates", templateHandler)
}
