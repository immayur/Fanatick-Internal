package postgres

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/segmentio/ksuid"
	"github.com/swensonhe/fanatick-backend/fanatick"
)

// EventDB is a database for events.
type SeatDB struct {
	*DB
}

// EventTx is an event transaction.
type SeatTx struct {
	*Tx
}

// Get returns an Seat.
func (db *DB) GetSeatByID(id string) (*fanatick.Seat, error) {
	var seat fanatick.Seat

	query := `
		SELECT 
			seats.id,
			seats.no,
			seats.section,
			seats.row,
			seats.image_url,
			seats.created_at,
			seats.updated_at,
			seats.deleted_at
		FROM seats
		WHERE seats.id = $1`

	err := db.QueryRow(query, id).Scan(
		&seat.ID,
		&seat.No,
		&seat.Section,
		&seat.Row,
		&seat.SeatURL,
		&seat.CreatedAt,
		&seat.UpdatedAt,
		&seat.DeletedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fanatick.ErrorNotFound
		}

		return nil, err
	}

	return &seat, nil
}

// Query returns a list of events.
func (db *DB) QuerySeatsbyEventID(id string, params map[fanatick.SeatQueryParam]interface{}) ([]*fanatick.Seat, error) {
	query := `
		SELECT
		seats.id,
		seats.no,
		seats.section,
		seats.row,
		seats.image_url,
		seats.created_at,
		seats.updated_at,
		seats.deleted_at
		FROM seats
		WHERE %s
		LIMIT %s
	`

	args := []interface{}{}
	wheres := []string{}
	limit := `NULL`

	fmt.Println(params)

	wheres = append(wheres, fmt.Sprintf(`seats.id = '%s'`, id))

	if params != nil {
		if param, ok := params[fanatick.SeatQueryParamBefore]; ok {
			args = append(args, param)
			wheres = append(wheres, fmt.Sprintf(`seats.id < $%d`, len(args)))
		}

		if param, ok := params[fanatick.SeatQueryParamLimit]; ok {
			args = append(args, param)
			limit = fmt.Sprintf(`$%d`, len(args))
		}
	}

	fmt.Println(args)
	fmt.Println(limit)

	query = fmt.Sprintf(query, strings.Join(wheres, " AND "), limit)

	fmt.Println(query)

	rows, err := db.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}

	seats := []*fanatick.Seat{}
	for rows.Next() {
		var seat fanatick.Seat
		err := rows.Scan(
			&seat.ID,
			&seat.No,
			&seat.Section,
			&seat.Row,
			&seat.SeatURL,
			&seat.CreatedAt,
			&seat.UpdatedAt,
			&seat.DeletedAt,
		)
		if err != nil && err != sql.ErrNoRows {
			return nil, err
		}

		seats = append(seats, &seat)
	}

	return seats, nil
}

// BeginTx begins a transaction.
func (db *DB) BeginSeatTx() fanatick.SeatTx {
	return &SeatTx{Tx: db.Begin()}
}

func (tx *SeatTx) CommitSeatTx() error {

	err := tx.Tx.Commit()
	if err != nil {
		err := tx.Tx.Rollback()
		if err != nil {
			return fmt.Errorf("falied to rollback EventTx after commit EventTx failed")
		}
		return fmt.Errorf("falied to commit EventTx")
	}

	return nil
}

// // Create creates an event.
func (tx *SeatTx) CreateSeat(id string, seat *fanatick.Seat) error {
	query := `
		INSERT INTO seats(id,no,row,section,image_url,event_id)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	seat.ID = ksuid.New().String()
	seat.CreatedAt = time.Now()
	fmt.Print(seat)
	_, err := tx.Exec(query, seat.ID, seat.No, seat.Row, seat.Section, seat.SeatURL, id)
	if err != nil {
		return err
	}

	return nil
}

// Update updates an event.
func (tx *SeatTx) UpdateSeat(event *fanatick.Seat) error {
	// query := `
	// 	UPDATE seats
	// 	SET name=$2, start_at=$3, updated_at=$4
	// 	WHERE id = $1
	// 	RETURNING updated_at
	// `
	// err := tx.QueryRow(query, event.ID, event.Name, event.StartAt, time.Now()).Scan(&event.UpdatedAt)
	// if err != nil {
	// 	return err
	// }

	return nil
}

// Delete deletes an event.
func (tx *SeatTx) DeleteSeat(id string) error {
	query := `
		UPDATE events
		SET deleted_at=$2
		WHERE id = $1
	`

	_, err := tx.Exec(query, id, time.Now())
	if err != nil {
		return err
	}

	return nil
}
