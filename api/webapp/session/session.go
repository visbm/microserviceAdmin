package session

import (
	"fmt"
	"microseviceAdmin/domain/model"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

// CheckSession ...
func CheckSession(w http.ResponseWriter, r *http.Request) {

	session, err := store.Get(r, "session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, ok := session.Values["EmployeeID"]
	if !ok {
		http.Redirect(w, r, "/admin/login", http.StatusFound)
		return
	}
}

// AuthSession ...
func AuthSession(w http.ResponseWriter, r *http.Request, employee *model.Employee) {
	session, err := store.Get(r, "session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session.Values["EmployeeID"] = employee.EmployeeID
	position := employee.PositionString()
	session.Values["Position"] = position

	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Logout ...
func Logout(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	delete(session.Values, "EmployeeID")
	session.Save(r, w)
	http.Redirect(w, r, "/admin/login", http.StatusFound)
}

// IsExist ...
func IsExist(w http.ResponseWriter, r *http.Request) bool {
	session, _ := store.Get(r, "session")
	_, ok := session.Values["EmployeeID"]
	return ok
}

//CheckRigths of employee and return err if not enough rights
func CheckRigths(w http.ResponseWriter, r *http.Request) error {
	method := r.Method

	session, err := store.Get(r, "session")
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
