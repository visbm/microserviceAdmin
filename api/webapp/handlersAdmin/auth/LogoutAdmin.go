package auth

import (
	"microseviceAdmin/domain/store"
	"microseviceAdmin/webapp/session"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// LogoutAdmin ...
func LogoutAdmin(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		session.Logout(w, r)
	}
}
