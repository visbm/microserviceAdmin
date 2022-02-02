package server

import (
	"microseviceAdmin/domain/store"
	"microseviceAdmin/webapp/handlersAdmin/auth"
	bookinghandlers "microseviceAdmin/webapp/handlersAdmin/bookingHandlers"
	employeehandlers "microseviceAdmin/webapp/handlersAdmin/employeeHandlers"
	hotelhandlers "microseviceAdmin/webapp/handlersAdmin/hotelHandlers"
	pethandlers "microseviceAdmin/webapp/handlersAdmin/petHandlers"
	roomhandlers "microseviceAdmin/webapp/handlersAdmin/roomHandlers"
	seathandlers "microseviceAdmin/webapp/handlersAdmin/seatHandlers"
	usershandlers "microseviceAdmin/webapp/handlersAdmin/usersHandlers"
	"microseviceAdmin/webapp/middlewear/download"
	"microseviceAdmin/webapp/middlewear/upload"
	"net/http"
)

func (s *Server) configureRoutes() {
	s.router.Handle("GET", "/admin/login", auth.LoginAdmin(store.New(s.config)))
	s.router.Handle("POST", "/admin/auth", auth.AuthAdmin(store.New(s.config)))
	s.router.Handle("GET", "/admin/logout", auth.LogoutAdmin(store.New(s.config)))

	s.router.Handle("GET", "/admin/home", auth.HomeAdmin(store.New(s.config)))

	s.router.Handle("GET", "/admin/homeusers", usershandlers.HomeUsersHandler(store.New(s.config)))
	s.router.Handle("GET", "/admin/users", usershandlers.AllUsersHandler(store.New(s.config)))
	s.router.Handle("POST", "/admin/user/new", usershandlers.NewUser(store.New(s.config)))
	s.router.Handle("GET", "/admin/users/id/", usershandlers.GetUserByID(store.New(s.config)))
	s.router.Handle("POST", "/admin/users/delete", usershandlers.DeleteUser(store.New(s.config)))

	s.router.Handle("GET", "/admin/users/csv/", usershandlers.PrintAllUsersCSV(store.New(s.config), download.DownloadFileHandler(store.New(s.config))))

	//s.router.Handle("GET", "/admin/users/csv/", usershandlers.PrintAllUsersCSV(store.New(s.config)))
	//s.router.Handler("GET", "/admin/users/download/",  download.DownloadFileHandler(store.New(s.config)))

	s.router.Handle("GET", "/admin/choose", upload.Choose(store.New(s.config)))
	s.router.Handle("POST", "/admin/upload", upload.UploadFileHandler(store.New(s.config)))

	s.router.Handle("GET", "/admin/homehotels", hotelhandlers.HomeHotelHandler(store.New(s.config)))
	s.router.Handle("GET", "/admin/hotels", hotelhandlers.AllHotelsHandler(store.New(s.config)))
	s.router.Handle("GET", "/admin/hotels/id", hotelhandlers.GetHotelByID(store.New(s.config)))
	s.router.Handle("POST", "/admin/hotels/delete", hotelhandlers.DeleteHotels(store.New(s.config)))

	s.router.Handle("GET", "/admin/homepets", pethandlers.HomePetsHandler(store.New(s.config)))
	s.router.Handle("GET", "/admin/pets", pethandlers.AllPetsHandler(store.New(s.config)))
	s.router.Handle("GET", "/admin/pets/id", pethandlers.GetPetByID(store.New(s.config)))
	s.router.Handle("POST", "/admin/pets/delete", pethandlers.DeletePets(store.New(s.config)))

	s.router.Handle("GET", "/admin/homerooms", roomhandlers.HomeRoomHandler(store.New(s.config)))
	s.router.Handle("GET", "/admin/rooms", roomhandlers.AllRoomsHandler(store.New(s.config)))
	s.router.Handle("GET", "/admin/rooms/id", roomhandlers.GetRoomByID(store.New(s.config)))
	s.router.Handle("POST", "/admin/rooms/delete", roomhandlers.DeleteRooms(store.New(s.config)))

	s.router.Handle("GET", "/admin/homeseats", seathandlers.HomeSeatsHandler(store.New(s.config)))
	s.router.Handle("GET", "/admin/seats", seathandlers.AllSeatsHandler(store.New(s.config)))
	s.router.Handle("GET", "/admin/seats/id", seathandlers.GetSeatByID(store.New(s.config)))
	s.router.Handle("POST", "/admin/seats/delete", seathandlers.DeleteSeats(store.New(s.config)))

	s.router.Handle("GET", "/admin/homebookings", bookinghandlers.HomeBookingHandler(store.New(s.config)))
	s.router.Handle("GET", "/admin/bookings", bookinghandlers.AllBookingsHandler(store.New(s.config)))
	s.router.Handle("GET", "/admin/bookings/id", bookinghandlers.GetBookingByID(store.New(s.config)))
	s.router.Handle("POST", "/admin/bookings/delete", bookinghandlers.DeleteBooking(store.New(s.config)))

	s.router.Handle("GET", "/admin/homeemployees", employeehandlers.HomeEmployeesHandler(store.New(s.config)))
	s.router.Handle("GET", "/admin/employees", employeehandlers.AllEmployeeHandler(store.New(s.config)))
	s.router.Handle("GET", "/admin/employees/id", employeehandlers.GetEmployeeByID(store.New(s.config)))
	s.router.Handle("POST", "/admin/employees/delete", employeehandlers.DeleteEmployee(store.New(s.config)))

	http.Handle("/files/", http.StripPrefix("/files/", http.FileServer(http.Dir("/api/pkg/csv/"))))

	s.router.ServeFiles("/admin/templates/*filepath", http.Dir("templates"))
}
