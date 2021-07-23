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

	return rooms, nil
}

func (m *testDBRepo) GetRoomById(id int) (models.Room, error) {
	var room models.Room
	if id > 2 {
		return room, errors.New("room with id's more than 2 doesn't exists")
	}
	return room, nil
}
