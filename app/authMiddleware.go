package app

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/gtaylor314/Banking-Lib/errs"
	"github.com/gtaylor314/Banking-MS/domain"
)

type AuthMiddleware struct {
	repo domain.AuthRepository
}

func (authMid AuthMiddleware) authorizationHandler() func(http.Handler) http.Handler {
	return func(nextMidware http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// grab current route from the http.Request object
			currentRoute := mux.CurrentRoute(r)
			// grab the route variables from the http.Request object
			routeVars := mux.Vars(r)
			// Get() returns the first value at the specified key - the header is essentially a map[string][]string object
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				writeResponse(w, http.StatusUnauthorized, "missing token")
				return
			}
			// grab token using getTokenFromHeader()
			token := getTokenFromHeader(authHeader)
			// verify if the token, routeName, and route variables are authorized for the individual
			if !authMid.repo.IsAuthorized(token, currentRoute.GetName(), routeVars) {
				// if the verification fails
				appErr := errs.AppError{Code: http.StatusForbidden, Message: "Unauthorized"}
				writeResponse(w, appErr.Code, appErr.MessageOnly())
				return
			}
			// if the verification passes
			// ServeHTTP() will write reply headers and data to the response writer and then return - upon return, the request
			// is completed
			nextMidware.ServeHTTP(w, r)
		})
	}
}

func getTokenFromHeader(h string) string {
	// the token value is preceded by its type - in this case "Bearer" e.g. Bearer aaaa.bbbb.cccc
	splitToken := strings.Split(h, "Bearer")
	if len(splitToken) != 2 {
		return ""
	}
	// TrimSpace() removes all leading and trailing white space from the string passed in
	return strings.TrimSpace(splitToken[1])
}
