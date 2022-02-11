package permission

import (
	"microseviceAdmin/domain/store"
	"microseviceAdmin/webapp/session"
	"net/http"
	"strconv"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

// Get all permissions that the employee has...
func GetPerByEmplID(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		session.CheckSession(w, r)
		err := session.IsAdmin(w, r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v. ", err)
			return
		}

		err = s.Open()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			s.Logger.Errorf("Can't open DB. Err msg:%v.", err)
			return
		}

		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("id"))
			return
		}

		per, err := s.Permossions().GetByEmployeeId(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			s.Logger.Errorf("Can't find permissions. Err msg: %v", err)
			return
		}

		files := []string{
			"/api/webapp/tamplates/allPermissions.html",
			"/api/webapp/tamplates/base.html",
		}

		tmpl, err := template.ParseFiles(files...)
		if err != nil {
			http.Error(w, err.Error(), 400)
			s.Logger.Errorf("Can not parse template: %v", err)
			return
		}

		err = tmpl.Execute(w, per)
		if err != nil {
			http.Error(w, err.Error(), 400)
			s.Logger.Errorf("Can not parse template: %v", err)
			return
		}

	}
}
