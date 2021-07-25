package dbrepo

import (
	"context"
	"errors"
	"time"

	"github.com/nchukkaio/goweblearning/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func (m *postgresDBRepo) AllUsers() bool {
	return true
}

// InsertReservation inserts a reservation into the database
func (m *postgresDBRepo) InsertReservation(res models.Reservation) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var newId int
	stmt := `insert into reservations (first_name,last_name,email,phone,start_date,
				end_date,room_id,created_at,updated_at) values($1,$2,$3,$4,$5,$6,$7,$8,$9)
				returning id`
	err := m.DB.QueryRowContext(ctx, stmt,
		res.FirstName,
		res.LastName,
		res.Email,
		res.Phone,
		res.StartDate,
		res.EndDate,
		res.RoomId,
		time.Now(),
		time.Now(),
	).Scan(&newId)
	if err != nil {
		return 0, err
	}
	return newId, nil
}

// InsertRoomRestriction inserts row into room_restrictions table
func (m *postgresDBRepo) InsertRoomRestriction(rr models.RoomRestriction) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	stmt := `insert into room_restrictions (start_date,end_date,room_id,reservation_id,restriction_id,
				created_at,updated_at) values($1,$2,$3,$4,$5,$6,$7)
			`
	_, err := m.DB.ExecContext(ctx, stmt,
		rr.StartDate,
		rr.EndDate,
		rr.RoomId,
		rr.ReservationId,
		rr.RestrictionId,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		return err
	}
	return nil
}

// SearchAvailabilityByDatesByRoomId returns true if room with id available between the given dates
func (m *postgresDBRepo) SearchAvailabilityByDatesByRoomId(start, end time.Time, roomId int) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `
				select count(id)
				from
				room_restrictions
				where
					room_id=$1 and $2 < end_date and $3 > start_date 
			`
	var numRows int
	row := m.DB.QueryRowContext(ctx, stmt, roomId, start, end)
	err := row.Scan(&numRows)
	if err != nil {
		return false, err
	}
	if numRows == 0 {
		return true, nil
	}

	return false, nil
}

// SearchAvailabilityForAllRooms gives all available rooms if any, for a given date range
func (m *postgresDBRepo) SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var rooms []models.Room
	stmt := `
			
				select r.id,r.room_name
				from rooms as r
				where 
				r.id not in (select room_id from room_restrictions as rr 
				where $1 < end_date and $2 > start_date)

			`
	rows, err := m.DB.QueryContext(ctx, stmt, start, end)
	if err != nil {
		return rooms, err
	}
	for rows.Next() {
		var room models.Room
		err := rows.Scan(&room.ID, &room.RoomName)
		if err != nil {
			return rooms, err
		}
		rooms = append(rooms, room)
	}

	return rooms, nil
}

// GetRoomById returns a room by id
func (m *postgresDBRepo) GetRoomById(id int) (models.Room, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var room models.Room
	stmt := `
			
			select id,room_name,created_at,updated_at from rooms where id = $1

		`
	row := m.DB.QueryRowContext(ctx, stmt, id)
	err := row.Scan(&room.ID, &room.RoomName, &room.CreatedAt, &room.UpdatedAt)

	if err != nil {
		return room, err
	}

	return room, err
}

// GetRoomById returns a User by ID
func (m *postgresDBRepo) GetUserById(id int) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var u models.User
	stmt := `

			select id,first_name,last_name,password,email,access_level,created_at,updated_at from users where id = $1

		`
	row := m.DB.QueryRowContext(ctx, stmt, id)
	err := row.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Password, &u.Email, &u.AccessLevel, &u.CreatedAt, &u.UpdatedAt)

	if err != nil {
		return u, err
	}

	return u, err
}

// UpdateUser updates user in the database
func (m *postgresDBRepo) UpdateUser(u models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	query := `
		update users set first_name=$1,last_name=$2,email=$3,access_level=$4,updated_at=$5
	`

	_, err := m.DB.ExecContext(ctx, query, u.FirstName, u.LastName, u.AccessLevel, time.Now())

	if err != nil {
		return err
	}
	return nil
}

// Authenticate tells identifies the user
func (m *postgresDBRepo) Authenticate(email, testPassword string) (int, string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var id int
	var hashedPassword string

	row := m.DB.QueryRowContext(ctx, "select id,password from users where email=$1", email)

	err := row.Scan(&id, &hashedPassword)

	if err != nil {
		return id, "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(testPassword))

	if err == bcrypt.ErrMismatchedHashAndPassword {
		return 0, "", errors.New("incorrect password")
	} else if err != nil {
		return 0, "", err
	}

	return id, hashedPassword, nil
}

// // Authenticate tells identifies the user
// func (m *postgresDBRepo) RegisterUser(email, testPassword string) (int, string, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
// 	defer cancel()

// 	var newId int
// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(testPassword), 12)
// 	if err != nil {
// 		return newId, "", err
// 	}
// 	stmt := `
// 		insert into users (first_name,last_name,email,password,access_level,
// 		created_at,updated_at) values($1,$2,$3,$4,$5,$6,$7)
// 		returning id
// 	`
// 	err = m.DB.QueryRowContext(ctx, stmt, "test", "test", email, hashedPassword, 1, time.Now(), time.Now()).Scan(&newId)

// 	if err != nil {
// 		return newId, "", err
// 	}

// 	return newId, string(hashedPassword), nil
// }
