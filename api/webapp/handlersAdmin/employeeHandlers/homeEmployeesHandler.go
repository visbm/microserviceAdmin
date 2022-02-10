package employeehandlers

import (
	"microseviceAdmin/domain/store"
	"microseviceAdmin/webapp/session"
	"net/http"
	"text/template"

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

// HomeEmployeesHandler ...
func HomeEmployeesHandler(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		session.CheckSession(w, r)
		err := session.CheckRigths(w, r, permission_read.Name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			s.Logger.Errorf("Bad request. Err msg:%v. ", err)
			return
		}

		vd := ViewData{
			w: w,
			r: r}

		files := []string{
			"/api/webapp/tamplates/employeeHome.html",
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
