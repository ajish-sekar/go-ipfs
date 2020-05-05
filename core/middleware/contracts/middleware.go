package contracts

import "net/http"

type Middleware interface {
	Handle(w http.ResponseWriter, r *http.Request, headers map[string][]string) bool
}