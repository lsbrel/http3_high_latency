package typos

import "net/http"

type Endpoint struct {
	Name     string
	Callback func(w http.ResponseWriter, r *http.Request)
}
