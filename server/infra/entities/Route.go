package entities

import "net/http"

type Route struct {
	Name     string
	Callback func(w http.ResponseWriter, r *http.Request)
}
