package middleware

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"goapi/api"
	"goapi/internal/tools"
	"net/http"
)

var UnAnthorizedError = errors.New("Un Anthorized Error")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username := r.URL.Query().Get("username")
		token := r.Header.Get("Authorization")
		var err error

		if username == "" || token == "" {
			log.Error(UnAnthorizedError)
			api.RequestErrorHandler(w, UnAnthorizedError)
			return
		}

		var database *tools.DatabaseInterface
		database, err = tools.NewDatabase()
		if err != nil {
			api.InternalErrorHandler(w)
			return
		}

		var loginDetails *tools.LoginDetails
		loginDetails = (*database).GetUserLoginDetails(username)

		if loginDetails == nil || token != (*loginDetails).AuthToken {
			log.Error(UnAnthorizedError)
			api.RequestErrorHandler(w, UnAnthorizedError)
			return
		}
		next.ServeHTTP(w, r)
	})
}
