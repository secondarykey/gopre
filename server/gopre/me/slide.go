package me

import (
	"net/http"
)

func GetSlideParamFunc() func(http.ResponseWriter, *http.Request, map[string]interface{}) {
	return nil
}
