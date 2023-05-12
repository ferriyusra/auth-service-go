package handler

import (
	"net/http"
)

type CrmAuthHandler interface {
	Register(w http.ResponseWriter, r *http.Request)
}
