package router

import (
	"http3tcc/infra/entities"
	"net/http"
)

var mux = http.NewServeMux()

func AddRoute(routes []entities.Route) {

	for _, route := range routes {
		mux.HandleFunc(route.Name, route.Callback)
	}

}

func GetMuxer() *http.ServeMux {
	return mux
}
