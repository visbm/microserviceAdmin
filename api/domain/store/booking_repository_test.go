package store_test

import (
	"microseviceAdmin/domain/model"
	"microseviceAdmin/domain/store"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBookingRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, host, dbName, user, password, port, sslMode)
	t.Cleanup(teardown)
	t.Run("valid", func(t *testing.T) {
		b := model.TestBooking()

		p, _ := s.Pet().Create(model.TestPet())
		se, _ := s.Seat().Create(model.TestSeat())
		e, _ := s.Employee().Create(model.TestEmployee())
		u , _:= s.User().Create(model.TestUser())

		b.Employee = *e
		b.Seat = *se
		b.Pet = *p
		p.Owner = *u

		b, err := s.Booking().Create(b)

		assert.NoError(t, err)
		assert.NotNil(t, b)
	})
}

/*
func TestBookingRepository_Delete(t *testing.T) {
	s, teardown := store.TestStore(t, host, dbName, user, password, port, sslMode)
	t.Cleanup(teardown)

	t.Run("invalid id", func(t *testing.T) {
		id := -1
		err := s.Employee().Delete(id)
		assert.Error(t, err)
	})
	t.Run("valid", func(t *testing.T) {
		b := model.TestBooking()
		u, _ := s.User().Create(model.TestUser())
		h, _ := s.Hotel().Create(model.TestHotel())
		e.User = *u
		e.Hotel = *h
		e, _ = s.Employee().Create(e)
		err := s.Employee().Delete(e.EmployeeID)
		assert.NoError(t, err)
	})
}

func TestBookingRepository_FindByID(t *testing.T) {
	s, teardown := store.TestStore(t, host, dbName, user, password, port, sslMode)
	t.Cleanup(teardown)
	t.Run("invalid id", func(t *testing.T) {
		id := -1
		_, err := s.Employee().FindByID(id)
		assert.Error(t, err)
	})
	t.Run("valid", func(t *testing.T) {
		b := model.TestBooking()
		u, _ := s.User().Create(model.TestUser())
		h, _ := s.Hotel().Create(model.TestHotel())
		e.User = *u
		e.Hotel = *h
		e, _ = s.Employee().Create(e)
		e, err := s.Employee().FindByID(e.EmployeeID)
		assert.NoError(t, err)
		assert.NotNil(t, b)
	})
}

func TestBookingRepository_GetAll(t *testing.T) {
	s, teardown := store.TestStore(t, host, dbName, user, password, port, sslMode)
	t.Cleanup(teardown)
	t.Run("valid ", func(t *testing.T) {
		e, err := s.Employee().GetAll()
		assert.NoError(t, err)
		assert.NotNil(t, b)
	})
}

func TestBookingRepository_FindByUserID(t *testing.T) {
	s, teardown := store.TestStore(t, host, dbName, user, password, port, sslMode)
	t.Cleanup(teardown)
	t.Run("invalid id", func(t *testing.T) {
		id := -1
		_, err := s.Employee().FindByID(id)
		assert.Error(t, err)
	})
	t.Run("valid", func(t *testing.T) {
		b := model.TestBooking()
		u, _ := s.User().Create(model.TestUser())
		h, _ := s.Hotel().Create(model.TestHotel())
		e.User = *u
		e.Hotel = *h
		e, _ = s.Employee().Create(e)
		e, err := s.Employee().FindByUserID(e.UserID)
		assert.NoError(t, err)
		assert.NotNil(t, b)
	})
}
*/
