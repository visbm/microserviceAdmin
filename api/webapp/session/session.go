package session

import (
	"fmt"
	"microseviceAdmin/domain/model"
	"net/http"
	"os"
	"time"

	"github.com/antonlindstrom/pgstore"
)

//var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
//var sessionRepository = store.New(s.config)
//var config = &webapp.Config{}
//config.NewConfig()

//var db, _ = webapp.ConnectDb()
//var PGStore, _ = pgstore.NewPGStoreFromPool(db, []byte(os.Getenv("SESSION_KEY")))
//var STORE = store.New(s.config)

var PGStore, _ = pgstore.NewPGStore("postgres://user:userpass@postgresql_database:5432/adminDB?sslmode=disable", []byte(os.Getenv("SESSION_KEY")))

// CheckSession ...
func CheckSession(w http.ResponseWriter, r *http.Request) {
	defer PGStore.StopCleanup(PGStore.Cleanup(time.Minute * 5))

	session, err := PGStore.Get(r, "session-key")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, ok := session.Values["EmployeeID"]
	if !ok {
		http.Redirect(w, r, "/admin/login", http.StatusUnauthorized)
		return
	}
}

// AuthSession ...
func AuthSession(w http.ResponseWriter, r *http.Request, employee *model.Employee) {

	session, err := PGStore.Get(r, "session-key")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session.Values["EmployeeID"] = employee.EmployeeID
	position := string(employee.Position)
	session.Values["Position"] = position
	session.Values["Employee_HotelID"] = employee.Hotel.HotelID
	//	session.Values["Permissions"] = permissions

	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

// Logout ...
func Logout(w http.ResponseWriter, r *http.Request) {
	session, err := PGStore.Get(r, "session-key")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session.Options.MaxAge = -1
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/admin/login", http.StatusFound)
}

// IsExist ...
func IsExist(w http.ResponseWriter, r *http.Request) bool {
	session, _ := PGStore.Get(r, "session")
	_, ok := session.Values["EmployeeID"]
	return ok
}

//CheckRigths of employee and return err if not enough rights
func CheckRigths(w http.ResponseWriter, r *http.Request) error {
	method := r.Method

	session, err := PGStore.Get(r, "session-key")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	position := session.Values["Position"]

	if method == "POST" && position == "employee" {
		http.Error(w, "You don't have enough rights", http.StatusForbidden)
		return fmt.Errorf("you don't have enough rights")
	}

	return nil
}
