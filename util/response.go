package util

import (
	"encoding/json"
	"net/http"

	"skegsTech/auth-service-go/config"
)

func Success(w http.ResponseWriter, code int, data interface{}, message string) {
	response, err := json.Marshal(map[string]interface{}{
		"data":    data,
		"success": true,
		"message": message,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(code)
	w.Write(response)
}

func Error(w http.ResponseWriter, code int, data interface{}, message string) {
	response, err := json.Marshal(map[string]interface{}{
		"data":    data,
		"message": message,
		"success": false,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(code)
	w.Write(response)
}

func Errorf(w http.ResponseWriter, code int, data interface{}, err error) {
	message := err.Error()

	// error 500 & debug off
	var appDebug bool
	config, err := config.GetConfig()
	if err != nil {
		appDebug = false
	} else {
		appDebug = config.AppDebug
	}
	if code == http.StatusInternalServerError && !appDebug {
		message = "Internal Server Error"
	}

	Error(w, code, data, message)
}
