package model

//Permissons like creat_user / delete_hotel
type Permission struct {
	PermissionID int
	Name         string
	Descriptoin  string
}

var DefaultPermissoins []Permission = []Permission{{0, "read_user ", "Read user"},
	{0, "read_hotel", "Read hotel"}}

func  find(per *[]Permission ,name string) bool {

	for i := 0; i < len(per); i++ {
		
	}

	for _ , i := range per{
		if i.Name == name{
			return true
		}
	}
	
}