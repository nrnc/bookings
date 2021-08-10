package dbrepo

import (
	"errors"
	"time"

	"github.com/nchukkaio/goweblearning/internal/models"
)

func (m *testDBRepo) AllUsers() bool {
	return true
}

// InsertReservation inserts a reservation into the database
func (m *testDBRepo) InsertReservation(res models.Reservation) (int, error) {
	if res.RoomId == 2 {
		return 0, errors.New("error created for testing")
	}
	return 1, nil
}

// InsertRoomRestriction inserts row into room_restrictions table
func (m *testDBRepo) InsertRoomRestriction(rr models.RoomRestriction) error {
	if rr.RoomId == 1000 {
		return errors.New("error created for testing")
	}
	return nil
}

// SearchAvailabilityByDatesByRoomId returns true if room with id available between the given dates
func (m *testDBRepo) SearchAvailabilityByDatesByRoomId(start, end time.Time, roomId int) (bool, error) {
	if roomId == 3 {
		return false, errors.New("this room can't be chosen")
	}
	return true, nil
}

// SearchAvailabilityForAllRooms gives all available rooms if any, for a given date range
func (m *testDBRepo) SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error) {
	var rooms []models.Room
	if start == end {
		return rooms, nil
	}
	layout := "2006-01-02"
	testDateToFail, _ := time.Parse(layout, "2060-01-01")
	if testDateToFail == start {
		return rooms, errors.New("no rooms")
	}
	rooms = []models.Room{
		{
			ID:       1,
			RoomName: "Genaral's Quarters",
		},
	}
	return rooms, nil
}

func (m *testDBRepo) GetRoomById(id int) (models.Room, error) {
	var room models.Room
	if id > 2 {
		return room, errors.New("room with id's more than 2 doesn't exists")
	}
	return room, nil
}

func (m *testDBRepo) GetUserById(id int) (models.User, error) {
	var user models.User

	return user, nil
}

func (m *testDBRepo) UpdateUser(u models.User) error {
	return nil
}

func (m *testDBRepo) Authenticate(email, testPassword string) (int, string, error) {
	return 0, "", nil
}

// func (m *testDBRepo) RegisterUser(email, testPassword string) (int, string, error) {
// 	return 0, "", nil
// }
func (m *testDBRepo) AllReservations() ([]models.Reservation, error) {
	return nil, nil
}
func (m *testDBRepo) NewReservations() ([]models.Reservation, error) {
	return nil, nil
}

func (m *testDBRepo) GetReservationById(id int) (models.Reservation, error) {
	var reservation models.Reservation
	return reservation, nil
}

// UpdateReservation updates reservation in the database
func (m *testDBRepo) UpdateReservation(r models.Reservation) error {
	return nil
}

func (m *testDBRepo) DeleteReservation(id int) error {
	return nil
}

// UpdateProcessedForReservation updates processed status of reservation in the database
func (m *testDBRepo) UpdateProcessedForReservation(id, processed int) error {
	return nil
}

func (m *testDBRepo) AllRooms() ([]models.Room, error) {
	var rooms []models.Room
	return rooms, nil
}

func (m *testDBRepo) GetRestrictionsRoomByDate(roomID int, start, end time.Time) ([]models.RoomRestriction, error) {
	return nil, nil
}

// InsertBlockForRoom inserts a room restriction
func (m *testDBRepo) InsertBlockForRoom(id int, startDate time.Time) error {

	return nil
}

// DeleteBlockForRoom inserts a room restriction
func (m *testDBRepo) DeleteBlockById(id int) error {

	return nil
}
