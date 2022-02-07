package model

//Permissons like creat_user / delete_hotel
type Permission struct {
	PermissionID int
	Name         string
	Descriptoin  string
}

var DefaultPermissoins []Permission = []Permission{{0, "read_user ", "Read user"},
	{0, "read_hotel", "Read hotel"}}
