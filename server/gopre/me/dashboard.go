package me

import (
	"net/http"
)

func GetDashboardParamFunc() func(http.ResponseWriter, *http.Request, map[string]interface{}) {
	return nil
}