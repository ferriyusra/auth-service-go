package route

import (
	"fmt"
	"net/http"
	"os"

	authHandler "skegsTech/auth-service-go/domain/auth/handler"

	"skegsTech/auth-service-go/middleware"

	"github.com/gorilla/mux"
)

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, os.Getenv("APP_NAME")+" "+os.Getenv("APP_ENV")+" "+os.Getenv("APP_VERSION"))
}

func NewRouter(crmAuthHandler authHandler.CrmAuthHandler) *mux.Router {

	r := mux.NewRouter()
	r.Use(middleware.GenerateRequestID) // generate request ID middleware
	r.Use(middleware.RequestLogger)     // log http request

	r.HandleFunc("/", root)

	// crm endpoints or web app endpoints
	crm := r.PathPrefix("/crm").Subrouter()

	account := r.PathPrefix("/account").Subrouter()
	account.Use(middleware.CrmAuthenticated) // user must be authenticated

	crm.HandleFunc("/v1/auth/register", crmAuthHandler.Register).Methods("POST")
	crm.HandleFunc("/v1/auth/login", crmAuthHandler.Login).Methods("POST")

	account.HandleFunc("/v1/profile", crmAuthHandler.Profile).Methods("GET")

	return r
}
