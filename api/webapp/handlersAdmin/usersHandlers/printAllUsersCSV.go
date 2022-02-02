package usershandlers

import (
	"fmt"
	"microseviceAdmin/domain/store"
	"microseviceAdmin/pkg/csv"
	"microseviceAdmin/webapp/session"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// PrintAllUsersCSV in csv file
func PrintAllUsersCSV(s *store.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		session.CheckSession(w, r)
		err := session.CheckRigths(w, r)
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

		users, err := s.User().GetAll()
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			s.Logger.Errorf("Can't find users. Err msg: %v", err)
			return
		}
		name := "allusers"
		
		path, err := csv.MakeCSV(users, name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			s.Logger.Errorf("error writing record to csv:", err)
			return
		}

		s.Logger.Info("Csv is created")		

		url := fmt.Sprint("http://localhost:8080/admin/users/download/", name, ".csv")
		s.Logger.Info("url: ", url)
		

		http.ServeFile(w, r, path)

		//http.Redirect(w, r, "/admin/homeusers", http.StatusFound)	

	}
}
