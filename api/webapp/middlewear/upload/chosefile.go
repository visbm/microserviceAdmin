package upload

import (
	"html/template"
	"microseviceAdmin/domain/store"
	"microseviceAdmin/webapp/session"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Choose(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		session.CheckSession(w, r)
		err := session.CheckRigths(w, r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v. ", err)
			return
		}

		t, _ := template.ParseFiles("/api/webapp/tamplates/upload.html")
		t.Execute(w, nil)

		w.Write([]byte("SUCCESS"))
	}

}
