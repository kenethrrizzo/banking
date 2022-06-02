package app

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/kenethrrizzo/banking/domain"
)

type AuthMiddleware struct {
	repo domain.AuthRepository
}

func (a AuthMiddleware) authorizationHandler() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			currentRoute := mux.CurrentRoute(r)
			currentRouteVars := mux.Vars(r)
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				writeResponse(rw, http.StatusUnauthorized, "Missing token")
				return
			}
			token := getTokenFromHeader(authHeader)
			isAuthorized := a.repo.IsAuthorized(token, currentRoute.GetName(), currentRouteVars)

			if !isAuthorized {
				writeResponse(rw, http.StatusUnauthorized, "Unauthorized")
				return
			}

			next.ServeHTTP(rw, r)
		})
	}
}

func getTokenFromHeader(header string) string {
	tokenSplited := strings.Split(header, "Bearer")
	if len(tokenSplited) == 2 {
		return tokenSplited[1]
	}
	return ""
}
