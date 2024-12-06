package middleware

import (
	"errors"
	"net/http"

	"github.com/gabriel-q7/go-rest-api/api"
	log "github.com/sirupsen/logrus"
)

var UnauthorizedError = errors.New("Invalida username or token.")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var username string = r.URL.Query().Get("username")
		var token string = r.Header.Get("authorization")
		var err error

		if username == "" || token == "" {
			log.Error(UnauthorizedError)
			api.RequestErrorHandler(w, UnauthorizedError)
			return
		}

		var database *tools.DatabaseInterface
		database, err = tools.NewDatabase()
		if err != nil {
			api.InternalErrorHandler(w)
			return
		}

		var loginDetails = (*database).GetUserLoginDetails(username)
	})
}
