package me

import (
	"gopre"
	"net/http"
)

func templateHandler(w http.ResponseWriter, r *http.Request) {
	gopre.Display(w, r,
		gopre.St2Sl("me/menu.tmpl", "me/template.tmpl"), nil)
}
