package models

import (
	"errors"
	"time"

	"example.com/go-api/db"
)

type Event struct {
	Id       int64
	Name     string    `binding: "required"`
	Desc     string    `binding: "required"`
	Location string    `binding: "required"`
	DateTime time.Time `binding: "required"`
	UserId   int64
}

func (event *Event) Save() error {
	query := `INSERT INTO events(name, desc, location, datetime, user_id) values (?,?,?,?,?)`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(event.Name, event.Desc, event.Location, event.DateTime, event.UserId)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	event.Id = id
	return err
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.Id, &event.Name, &event.Desc, &event.Location, &event.DateTime, &event.UserId)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	query := "SELECT * FROM events where id = ?"
	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.Id, &event.Name, &event.Desc, &event.Location, &event.DateTime, &event.UserId)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (e Event) Update() error {
	query := `UPDATE events
	SET name = ?, desc = ?, location = ?, datetime = ?
	WHERE id = ?`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.Name, e.Desc, e.Location, e.DateTime, e.Id)
	return err
}

func (e Event) Delete() error {
	query := `DELETE FROM events WHERE id = ?`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.Id)
	return err
}

func (event *Event) Register(userId int64) error {
	query := `INSERT INTO registrations(event_id, user_id) values (?,?)`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.Id, userId)

	if err != nil {
		return err
	}

	return err
}

func (event *Event) CancelRegistration(userId int64) error {
	query := `DELETE FROM registrations WHERE event_id = ? AND user_id = ?`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(event.Id, userId)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("no entry found")
	}

	return err
}
