package roomhandlers

import (
	"log"
	"microseviceAdmin/domain/store"
	"microseviceAdmin/webapp/session"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// DeleteUser ...
func DeleteRooms(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		session.CheckSession(w, r)
		err := session.CheckRigths(w, r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v. ", err)
			return
		}

		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("id"))
			http.Redirect(w, r, "/admin/homerooms", http.StatusFound)
			return
		}
		err = s.Open()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			s.Logger.Errorf("Can't open DB. Err msg:%v.", err)
			return
		}
		err = s.Room().Delete(id)
		if err != nil {
			log.Print(err)
			s.Logger.Errorf("Can't delete room. Err msg:%v.", err)
			return
		}
		s.Logger.Info("Delete room with id = %d", id)
		http.Redirect(w, r, "/admin/homerooms", http.StatusFound)

	}
}
