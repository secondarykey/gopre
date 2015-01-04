package me

import (
	"gopre"
	"net/http"
)

func slideHandler(w http.ResponseWriter, r *http.Request) {
	gopre.Display(w, r,
		gopre.St2Sl("me/menu.tmpl", "me/slide.tmpl"), nil)
}
