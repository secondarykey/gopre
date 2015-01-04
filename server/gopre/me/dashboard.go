package me

import (
	"gopre"
	"net/http"
)

func dashboardHandler(w http.ResponseWriter, r *http.Request) {
	gopre.Display(w, r,
		gopre.St2Sl("me/menu.tmpl", "me/dashboard.tmpl"), nil)
}
