package usershandlers

import (
	"log"
	"microseviceAdmin/domain/model"
	"microseviceAdmin/domain/store"
	"microseviceAdmin/webapp/session"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

var permission_delete model.Permission = model.Permission{
	PermissionID: 0,
	Name:         "read_uers",
	Descriptoin:  "ability to read a user"}

// DeleteUser ...
func DeleteUser(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		session.CheckSession(w, r)
		err := session.CheckRigths(w, r, permission_delete.Name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			s.Logger.Errorf("Bad request. Err msg:%v. ", err)
			return
		}

		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			s.Logger.Errorf("Bad request. Err msg:%v. Requests body: %v", err, r.FormValue("id"))
			http.Redirect(w, r, "/admin/homeusers", http.StatusFound)
			return
		}
		err = s.Open()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			s.Logger.Errorf("Can't open DB. Err msg:%v.", err)
			return
		}
		err = s.User().Delete(id)
		if err != nil {
			log.Print(err)
			s.Logger.Errorf("Can't delete user. Err msg:%v.", err)
			return
		}
		s.Logger.Info("Delete user with id = %d", id)
		http.Redirect(w, r, "/admin/homeusers", http.StatusFound)

	}
}
