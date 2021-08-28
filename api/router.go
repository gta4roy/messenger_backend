package api

import (
	"gta4roy/messenger/log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

const (
	BaseURL = "/api/v1/message"

	HealthChecURL = "/health"

	AddMessageURL = BaseURL + "/add"

	ModifyMessageURL = BaseURL + "/modify"

	SearchMessageURL = BaseURL + "/search"

	PrintAllMessageURL = BaseURL + "/getall"

	DeleteMessageURL = BaseURL + "/remove"
)

var routes = Routes{
	Route{
		"HealthCheck", "GET", HealthChecURL, handleGetHealth,
	},
	Route{
		"AddMessenger", "POST", AddMessageURL, handleAddMessage,
	},
	// Route{
	// 	"ModifyAddress", "POST", ModifyMessageURL, handleModifyMessage,
	// },
	// Route{
	// 	"SearchAddress", "GET", SearchMessageURL, handleSearchMessage,
	// },
	Route{
		"PrintAllAddress", "GET", PrintAllMessageURL, handlePrintAllMessage,
	},

	// Route{
	// 	"DeleteAddress", "GET", DeleteMessageURL, handleDeleteMessage,
	// },
}

func logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		inner.ServeHTTP(w, r)
		log.Trace.Println("%s %s 5s %s", r.Method, r.RequestURI, name, time.Since(start))
	})
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = logger(handler, route.Name)
		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(handler)
	}
	return router
}
