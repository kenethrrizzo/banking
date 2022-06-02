package domain

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/kenethrrizzo/banking/logger"
)

type AuthRepository interface {
	IsAuthorized(token, routeName string, vars map[string]string) bool
}

type RemoteAuthRepository struct{}

func (r RemoteAuthRepository) IsAuthorized(token, routeName string, vars map[string]string) bool {
	verifyURL := buildVerifyURL(token, routeName, vars)

	resp, err := http.Get(verifyURL)
	if err != nil {
		logger.Error(err.Error())
		return false
	}
	body := map[string]bool{}
	err = json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		logger.Error(err.Error())
		return false
	}
	return body["is_authorized"]
}

func buildVerifyURL(token, routeName string, vars map[string]string) string {
	url := url.URL{Host: "localhost:8181", Path: "/auth/verify", Scheme: "http"}
	query := url.Query()

	query.Add("token", token)
	query.Add("route_name", routeName)

	for k, v := range vars {
		query.Add(k, v)
	}

	url.RawQuery = query.Encode()
	return url.String()
}

func NewAuthRepository() RemoteAuthRepository {
	return RemoteAuthRepository{}
}
