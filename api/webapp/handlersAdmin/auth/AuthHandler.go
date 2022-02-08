package auth

import (
	"microseviceAdmin/domain/model"
	"microseviceAdmin/domain/store"
	"microseviceAdmin/webapp/session"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// AuthAdmin ...
func AuthAdmin(s *store.Store) httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		EmailForm := r.FormValue("email")
		Password := r.FormValue("password")
		s.Open()
		user, err := s.User().FindByEmail(EmailForm)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Eror during checking users email or password. Err msg: %s", err.Error())
			http.Redirect(w, r, "/admin/login", http.StatusFound)
			return
		}

		userID := user.UserID
		hashPassword := user.Password

		isConfirmed := model.CheckPasswordHash(hashPassword, Password)
		if isConfirmed != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Eror during checking users email or password. Err msg: %s", err.Error())
			http.Redirect(w, r, "/admin/login", http.StatusFound)
			return
		}

		if user.Role != "employee" {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("You are not employee")
			http.Redirect(w, r, "/admin/login", http.StatusFound)
			return
		}

		employee, err := s.Employee().FindByUserID(userID)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Eror during getting employee. Err msg: %s", err.Error())
			http.Redirect(w, r, "/admin/login", http.StatusFound)
			return
		}
		permissions := []model.Permission{}
		if employee.Position == "employee" {
			permissions = model.DefaultPermissoins
		}

		session.AuthSession(w, r, employee, permissions)

		http.Redirect(w, r, "/admin/home", http.StatusFound)

	}
}
