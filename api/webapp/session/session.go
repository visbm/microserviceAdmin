package session

import (
	"encoding/gob"
	"fmt"
	"microseviceAdmin/domain/model"
	"net/http"
	"strings"
)

// CheckSession ...
func CheckSession(w http.ResponseWriter, r *http.Request) {

	session, err := sstore.PGStore.Get(r, "session-key")

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
func AuthSession(w http.ResponseWriter, r *http.Request, employee *model.Employee, permissions *[]model.Permission) {

	session, err := sstore.PGStore.Get(r, "session-key")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	gob.Register(model.Employee{})
	session.Values["Employee"] = employee
	session.Values["EmployeeID"] = employee.EmployeeID

	gob.Register([]model.Permission{})
	session.Values["Permissions"] = permissions

	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

// Logout ...
func Logout(w http.ResponseWriter, r *http.Request) {
	session, err := sstore.PGStore.Get(r, "session-key")
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

	session, err := sstore.PGStore.Get(r, "session-key")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return false
	}
	_, ok := session.Values["EmployeeID"]

	return ok
}

//CheckRigths of employee and return err if not enough rights
func CheckRigths(w http.ResponseWriter, r *http.Request) error {

	session, err := sstore.PGStore.Get(r, "session-key")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	employee, ok := session.Values["Employee"]
	if !ok {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return fmt.Errorf("no employee in session")
	}
	fmt.Println("employee: ", employee)

	permissions, ok := session.Values["Permissions"]
	if !ok {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return fmt.Errorf("no permissions in session")
	}

	str := fmt.Sprintf("%v", permissions)

	fmt.Println("permissions: ", str)
	lookFor := "delete_user"
	contain := strings.Contains(str, lookFor)
	

	fmt.Println("contain: ", contain)
	

	return nil
}


func CheckRigths2(w http.ResponseWriter, r *http.Request , name string) error {

	session, err := sstore.PGStore.Get(r, "session-key")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	employee, ok := session.Values["Employee"]
	if !ok {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return fmt.Errorf("no employee in session")
	}
	fmt.Println("employee: ", employee)

	permissions, ok := session.Values["Permissions"]
	if !ok {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return fmt.Errorf("no permissions in session")
	}

	str := fmt.Sprintf("%v", permissions)

	fmt.Println("permissions: ", str)	
	contain := strings.Contains(str, name)
	if !contain {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return fmt.Errorf("not enough rights")
	}
	fmt.Println("contain: ", contain)
	return nil
}