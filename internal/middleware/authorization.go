package middleware

import (
	"errors"
	"net/http"

	"github.com/dexterhere04/Learn_go/api"
	"github.com/dexterhere04/Learn_go/internal/tools"
	log "github.com/sirupsen/logrus"
)

var UnAuthorizedError = errors.New("Invalid username or token")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var username string = r.URL.Query().Get("username")
		var token = r.Header.Get("Authorization")
		var err error
		if username == "" || token == "" {
			log.Error(UnAuthorizedError)
			api.RequestInvalidError(w, UnAuthorizedError)
			return
		}
		var database tools.DatabaseInterface
		database, err = tools.GetDatabaseInstance()
		if err != nil {
			api.InternalErrorHandler(w, err)
			return
		}

		var loginDetails *tools.LoginDetails
		loginDetails = database.GetUserLoginDetails(username)
		if loginDetails == nil || loginDetails.AuthToken != token {
			log.Error(UnAuthorizedError)
			api.RequestInvalidError(w, UnAuthorizedError)
			return
		}
		next.ServeHTTP(w, r)
	})
}
