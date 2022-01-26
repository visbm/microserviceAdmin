package auth

import (
	"html/template"
	"microseviceAdmin/domain/store"
	"microseviceAdmin/webapp/session"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// HomeAdmin ...
func HomeAdmin(s *store.Store) httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		exist := session.IsExist(w, r)
		if exist {
			HomePage(w, s)

			return
		}
		s.Logger.Errorf("Unauthorized")
		http.Redirect(w, r, "/admin/login", http.StatusFound)
	}
}

// HomePage ...
func HomePage(w http.ResponseWriter, s *store.Store) {
	files := []string{
		"/api/webapp/tamplates/homeAdmin.html",
		"/api/webapp/tamplates/base.html",
	}
	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		http.Error(w, err.Error(), 400)
		s.Logger.Errorf("Can not parse template: %v", err)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), 400)
		s.Logger.Errorf("Can not parse template: %v", err)
		return
	}
}
