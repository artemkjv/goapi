package middleware

import (
	"errors"
	"net/http"
)

var UnauthorizedError = errors.New("Invalid username or token.")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var username string = r.URL.Query().Get("username")
		var token string = r.URL.Query().Get("token")
		var err error
		if username == "" || token == "" {
			err = UnauthorizedError
		}
	})
}
