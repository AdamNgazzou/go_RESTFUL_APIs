package middleware

import (
	"errors"
	"net/http"

	"github.com/AdamNgazzou/go_RESTFUL_APIs/api"
	"github.com/AdamNgazzou/go_RESTFUL_APIs/internal/tools"
	log "github.com/sirupsen/logrus"
)

var unAuthorizedError = errors.New("Invalid username or token.")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var username string = r.URL.Query().Get("username")
		var token = r.Header.Get("Authorization")
		var err error

		if username == "" || token == "" {
			log.Error(unAuthorizedError)
			api.RequestErrorHandler(w, r)
			return
		}

		var database *tools.DatabaseInterface
		database, err = tools.NewDatabase()
		if err != nil {
			api.InternalErrorHandler(w, r)
			return
		}
		var loginDetails *tools.LoginDetails
		loginDetails = (*database).GetUserLoginDetails(username)

		if loginDetails == nil || (token != (*loginDetails).AuthToken) {
			log.Error(unAuthorizedError)
			api.RequestErrorHandler(w, r)
			return
		}

		next.ServeHTTP(w, r)
	})
}
