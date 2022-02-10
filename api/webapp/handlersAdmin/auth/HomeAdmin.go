package auth

import (
	"html/template"
	"microseviceAdmin/domain/store"
	"microseviceAdmin/webapp/session"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type ViewData struct {
	w http.ResponseWriter
	r *http.Request
}

func (vd ViewData) HasPermission(name string) bool {
	err := session.CheckRigths(vd.w, vd.r, name)
	return err == nil
}

// HomeAdmin ...
func HomeAdmin(s *store.Store) httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		session.CheckSession(w, r)

		vd := ViewData{
			w: w,
			r: r}

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
