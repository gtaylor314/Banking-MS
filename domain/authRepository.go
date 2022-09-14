package domain

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/gtaylor314/Banking-Lib/logger"
)

type AuthRepository interface {
	IsAuthorized(token string, routeName string, vars map[string]string) bool
}

type RemoteAuthRepository struct{}

func (r RemoteAuthRepository) IsAuthorized(token string, routeName string, vars map[string]string) bool {
	url := buildVerifyURL(token, routeName, vars)
	// Get() issues a http GET to the passed in URL and returns a response and an error
	resp, err := http.Get(url)
	if err != nil {
		logger.Error("error while issuing a GET to the URL " + err.Error())
		return false
	}
	responseMap := make(map[string]bool)
	// create a decoder that reads from the received response body and decode the data into the response map
	err = json.NewDecoder(resp.Body).Decode(&responseMap)
	if err != nil {
		logger.Error("error while decoding response body " + err.Error())
		return false
	}
	// Banking-Auth microservice returns a response with an "isAuthorized" key set to true or false
	return responseMap["isAuthorized"]
}

func buildVerifyURL(token string, routeName string, vars map[string]string) string {
	// url.URL is a struct with which we initialize host, path, and scheme
	url := url.URL{Host: "localhost:8181", Path: "/auth/verify", Scheme: "http"}
	// Query() parses RawQuery and returns a map[string][]string with the values - we will use this map to fully build the URL
	buildURL := url.Query()
	// add token under key "token"
	buildURL.Add("token", token)
	// add routeName under key "routeName"
	buildURL.Add("routeName", routeName)
	// for-range across the passed in vars to add each key-value pair in vars to buildURL
	for key, value := range vars {
		buildURL.Add(key, value)
	}
	// Encode encodes the buildURL in URL encoded form - meaning key1=value1&key2=value2 for the URL
	// (e.g. ...routeName=GetCustomer&customer_id=1...)
	url.RawQuery = buildURL.Encode()
	return url.String()
}

func NewAuthRepository() RemoteAuthRepository {
	return RemoteAuthRepository{}
}
