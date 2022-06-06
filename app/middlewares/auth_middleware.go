package middlewares

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	repo "github.com/kenethrrizzo/banking/domain/repositories"
	util "github.com/kenethrrizzo/banking/utils"
)

type AuthMiddleware struct {
	Repo repo.AuthRepository
}

func (a AuthMiddleware) AuthorizationHandler() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			currentRoute := mux.CurrentRoute(r)
			currentRouteVars := mux.Vars(r)
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				util.WriteResponse(rw, http.StatusUnauthorized, "Missing token")
				return
			}
			token := getTokenFromHeader(authHeader)
			isAuthorized := a.Repo.IsAuthorized(token, currentRoute.GetName(), currentRouteVars)

			if !isAuthorized {
				util.WriteResponse(rw, http.StatusUnauthorized, "Unauthorized")
				return
			}

			next.ServeHTTP(rw, r)
		})
	}
}

func getTokenFromHeader(header string) string {
	tokenSplited := strings.Split(header, "Bearer")
	if len(tokenSplited) == 2 {
		return strings.TrimSpace(tokenSplited[1])
	}
	return ""
}
