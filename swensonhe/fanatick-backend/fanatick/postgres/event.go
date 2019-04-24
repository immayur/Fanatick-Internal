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
type EventDB struct {
	*DB
}

// EventTx is an event transaction.
type EventTx struct {
	*Tx
}

// Get returns an event.
func (db *EventDB) GetEvent(id string) (*fanatick.Event, error) {
	var event fanatick.Event

	query := `
		SELECT 
			events.id,
			events.name,
			events.start_at,
			events.created_at,
			events.updated_at,
			events.deleted_at
		FROM events
		WHERE events.id = $1 AND events.deleted_at IS NULL`

	err := db.QueryRow(query, id).Scan(
		&event.ID,
		&event.Name,
		&event.StartAt,
		&event.CreatedAt,
		&event.UpdatedAt,
		&event.DeletedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fanatick.ErrorNotFound
		}

		return nil, err
	}

	return &event, nil
}

// Query returns a list of events.
func (db *EventDB) QueryEvent(params map[fanatick.EventQueryParam]interface{}) ([]*fanatick.Event, error) {
	query := `
		SELECT
			events.id,
			events.name,
			events.start_at,
			events.created_at,
			events.updated_at,
			events.deleted_at
		FROM events
		WHERE %s
		ORDER BY events.id collate "C" DESC
		LIMIT %s
	`

	args := []interface{}{}
	wheres := []string{`events.deleted_at IS NULL`}
	limit := `NULL`

	fmt.Println(params)

	if params != nil {
		if param, ok := params[fanatick.EventQueryParamBefore]; ok {
			args = append(args, param)
			wheres = append(wheres, fmt.Sprintf(`events.id < $%d`, len(args)))
		}

		if param, ok := params[fanatick.EventQueryParamLimit]; ok {
			args = append(args, param)
			limit = fmt.Sprintf(`$%d`, len(args))
		}
	}

	query = fmt.Sprintf(query, strings.Join(wheres, " AND "), limit)

	rows, err := db.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}

	events := []*fanatick.Event{}
	for rows.Next() {
		var event fanatick.Event
		err := rows.Scan(
			&event.ID,
			&event.Name,
			&event.StartAt,
			&event.CreatedAt,
			&event.UpdatedAt,
			&event.DeletedAt,
		)
		if err != nil && err != sql.ErrNoRows {
			return nil, err
		}

		events = append(events, &event)
	}

	return events, nil
}

// BeginTx begins a transaction.
func (db *EventDB) BeginEventTx() fanatick.EventTx {
	return &EventTx{Tx: db.Begin()}
}

func (tx *EventTx) CommitEventTx() error {

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

// Create creates an event.
func (tx *EventTx) CreateEvent(event *fanatick.Event) error {
	query := `
		INSERT INTO events (id, name, start_at)
		VALUES ($1, $2, $3)
	`

	event.ID = ksuid.New().String()
	event.CreatedAt = time.Now()

	_, err := tx.Exec(query, event.ID, event.Name, event.StartAt)
	if err != nil {
		return err
	}

	return nil
}

// Update updates an event.
func (tx *EventTx) UpdateEvent(event *fanatick.Event) error {
	query := `
		UPDATE events
		SET name=$2, start_at=$3, updated_at=$4
		WHERE id = $1
		RETURNING updated_at
	`
	err := tx.QueryRow(query, event.ID, event.Name, event.StartAt, time.Now()).Scan(&event.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

// Delete deletes an event.
func (tx *EventTx) DeleteEvent(id string) error {
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
