package me

import (
	"net/http"
)

func GetTemplateParamFunc() func(http.ResponseWriter, *http.Request, map[string]interface{}) {
	return nil
}
