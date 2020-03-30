package auth

import (
	"fmt"
	"net/http"
	"strings"
)

type authMiddleware struct{}

func Handler() *authMiddleware {
	return &authMiddleware{}
}

func unauthorized(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
	fmt.Fprintf(w, "Unauthorized User\n")
}

func (_ *authMiddleware) Handle(w http.ResponseWriter, r *http.Request) bool {

	if(r.Method != http.MethodGet){
		return true
	}

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		unauthorized(w)
		return false
	}

	authToken := strings.SplitN(authHeader, " ", 2)
	if len(authToken) < 2 {
		unauthorized(w)
		return false
	}

	// validate token here. token can be an access token,
	// bearer, or jwt. commonly, the value of `authType`
	// would be "Bearer", or "JWT".
	// authType, token := authToken[0], authToken[1]

	return true
}