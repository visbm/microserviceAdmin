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
	"net/http"
)

func (s *Server) configureRoutes() {
	s.router.Handle("GET", "/admin/login", auth.LoginAdmin(store.New(s.config)))
	s.router.Handle("POST", "/admin/auth", auth.AuthAdmin(store.New(s.config)))
	s.router.Handle("GET", "/admin/logout", auth.LogoutAdmin(store.New(s.config)))

	s.router.Handle("GET", "/admin/home", auth.HomeAdmin(store.New(s.config)))

	s.router.Handle("GET", "/admin/users", usershandlers.AllUsersHandler(store.New(s.config)))
	s.router.Handle("POST", "/admin/user/new", usershandlers.NewUser(store.New(s.config)))
	s.router.Handle("GET", "/admin/users/id/", usershandlers.GetUserByID(store.New(s.config)))
	s.router.Handle("POST", "/admin/users/delete/", usershandlers.DeleteUser(store.New(s.config)))

	s.router.Handle("GET", "/admin/hotels", hotelhandlers.AllHotelsHandler(store.New(s.config)))
	s.router.Handle("GET", "/admin/hotels/id:id", hotelhandlers.GetHotelByID(store.New(s.config)))

	s.router.Handle("GET", "/admin/pets", pethandlers.AllPetsHandler(store.New(s.config)))
	s.router.Handle("GET", "/admin/pets/id:id", pethandlers.GetPetByID(store.New(s.config)))

	s.router.Handle("GET", "/admin/rooms", roomhandlers.AllRoomsHandler(store.New(s.config)))
	s.router.Handle("GET", "/admin/rooms/id:id", roomhandlers.GetRoomByID(store.New(s.config)))

	s.router.Handle("GET", "/admin/seats", seathandlers.AllSeatsHandler(store.New(s.config)))
	s.router.Handle("GET", "/admin/seats/id:id", seathandlers.GetSeatByID(store.New(s.config)))

	s.router.Handle("GET", "/admin/bookings", bookinghandlers.AllBookingsHandler(store.New(s.config)))
	s.router.Handle("GET", "/admin/bookings/id:id", bookinghandlers.GetBookingByID(store.New(s.config)))

	s.router.Handle("GET", "/admin/employees", employeehandlers.AllEmployeeHandler(store.New(s.config)))
	s.router.Handle("GET", "/admin/employees/id:id", employeehandlers.GetEmployeeByID(store.New(s.config)))

	s.router.ServeFiles("/admin/templates/*filepath", http.Dir("templates"))
}
