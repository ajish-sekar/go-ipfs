package middleware

import (
	"net/http"

	"github.com/ipfs/go-ipfs/core/middleware/contracts"
	"github.com/ipfs/go-ipfs/core/middleware/auth"
	"github.com/ipfs/go-ipfs/core/middleware/logger"
)

// register your middlewares here. example of
// middlewares is stored inside "./middlewares/"
var middlewares = []contracts.Middleware{
	logger.Handler(),
	auth.Handler(),
}


func Handler(w http.ResponseWriter, r *http.Request, headers map[string][]string) bool {
	for i := 0; i < len(middlewares); i++ {
		if ok := middlewares[i].Handle(w, r, headers); !ok {
			return false
		}
	}
	return true
}