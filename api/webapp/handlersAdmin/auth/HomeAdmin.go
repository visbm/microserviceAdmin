package auth

import (
	"html/template"
	"microseviceAdmin/domain/store"
	viewdata "microseviceAdmin/pkg/csv/viewData"
	"microseviceAdmin/webapp/session"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// HomeAdmin ...
func HomeAdmin(s *store.Store) httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		session.CheckSession(w, r)

		vd := viewdata.ViewData{
			ResponseWriter: w,
			Request: r,
		}

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

		err = tmpl.Execute(w, vd)
		if err != nil {
			http.Error(w, err.Error(), 400)
			s.Logger.Errorf("Can not parse template: %v", err)
			return
		}
	}
}
