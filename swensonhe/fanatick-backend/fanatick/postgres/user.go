package postgres

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/segmentio/ksuid"
	"github.com/swensonhe/fanatick-backend/fanatick"
)

// UserDB is a database for users.
type UserDB struct {
	*DB
}

// UserTx is an user transaction.
type UserTx struct {
	*Tx
}

// Get returns an user.
func (db *UserDB) GetUser(id string) (*fanatick.User, error) {
	var user fanatick.User

	query := `
		SELECT 
			users.id,
			users.firebase_uuid,
			users.phone_number,
			users.is_active,
			users.is_seller,
			users.last_login_time,
			users.created_at,
		    users.updated_at
		FROM users
		WHERE users.id = $1`

	err := db.QueryRow(query, id).Scan(
		&user.ID,
		&user.FireBaseUUID,
		&user.PhoneNumber,
		&user.IsActive,
		&user.IsSeller,
		&user.LastLoginTime,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fanatick.ErrorNotFound
		}
		return nil, err
	}

	return &user, nil
}


//GetUserByFireBaseUUID returns an user profile by user id.
func (db *UserDB) GetUserByFireBaseUUID(firebaseUUID string) (*fanatick.User, error) {
	var user fanatick.User

	query := `
		SELECT 
			users.id,
			users.firebase_uuid,
			users.phone_number,
			users.is_active,
			users.is_seller,
			users.last_login_time,
			users.created_at,
		    users.updated_at
		FROM users
		WHERE users.firebase_uuid = $1`

	err := db.QueryRow(query, firebaseUUID).Scan(
		&user.ID,
		&user.FireBaseUUID,
		&user.PhoneNumber,
		&user.IsActive,
		&user.IsSeller,
		&user.LastLoginTime,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fanatick.ErrorNotFound
		}
		return nil, err
	}

	return &user, nil
}

//GetUserByFireBaseUUID returns an user profile by user id.
func (tx *UserTx) GetUserByFireBaseUUID(firebaseUUID string) (*fanatick.User, error) {
	var user fanatick.User

	query := `
		SELECT 
			users.id,
			users.firebase_uuid,
			users.phone_number,
			users.is_active,
			users.is_seller,
			users.last_login_time,
			users.created_at,
		    users.updated_at
		FROM users
		WHERE users.firebase_uuid = $1`

	err := tx.QueryRow(query, firebaseUUID).Scan(
		&user.ID,
		&user.FireBaseUUID,
		&user.PhoneNumber,
		&user.IsActive,
		&user.IsSeller,
		&user.LastLoginTime,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fanatick.ErrorNotFound
		}
		return nil, err
	}

	return &user, nil
}

// Query returns a list of users.
func (db *UserDB) QueryUser(params map[fanatick.UserQueryParam]interface{}) ([]*fanatick.User, error) {

	users := []*fanatick.User{}
	//todo
	return users, nil
}

// BeginTx begins a transaction.
func (db *UserDB) BeginUserTx() fanatick.UserTx {
	return &UserTx{Tx: db.Begin()}
}

// CommitUserTx commits user transactions
func (tx *UserTx) CommitUserTx() error {
	err := tx.Tx.Commit()
	if err != nil {
		err := tx.Tx.Rollback()
		if err != nil {
			return fmt.Errorf("falied to rollback UserTx after commit UserTx failed")
		}
		return fmt.Errorf("falied to commit UserTx")
	}

	return nil
}

// Create creates an user.
func (tx *UserTx) CreateUser(user *fanatick.User) error {
	user.ID = ksuid.New().String()
	user.CreatedAt = time.Now()
	user.LastLoginTime = time.Now()
	user.IsActive = true

	query := `
		INSERT INTO 
		    users (
		            id, 
		            firebase_uuid, 
		            phone_number, 
		            is_seller, 
		            is_active,
		            created_at,
		           	last_login_time
		        )
		VALUES ($1, $2, $3, $4, $5, $6, $7)	`

	parameters := make([]interface{}, 7)
	parameters[0] = user.ID
	parameters[1] = user.FireBaseUUID
	parameters[2] = user.PhoneNumber
	parameters[3] = user.IsSeller
	parameters[4] = user.IsActive
	parameters[5] = user.CreatedAt
	parameters[6] = user.LastLoginTime

	_, err := tx.Exec(query, parameters...)
	if err != nil {
		return err
	}

	return nil
}

// Update updates an user.
func (tx *UserTx) UpdateUser(user *fanatick.User) error {
	return nil
}

// UpdateUserLastLoginTime updates an user last login time.
func (tx *UserTx) UpdateUserLastLoginTime(user *fanatick.User) error {
	user.LastLoginTime = time.Now()
	query := `
		UPDATE users
		SET last_login_time=$2
		WHERE firebase_uuid = $1
		RETURNING last_login_time
	`
	err := tx.QueryRow(query, user.FireBaseUUID, user.LastLoginTime).Scan(&user.LastLoginTime)
	if err != nil {
		return err
	}

	return nil
}

// Delete deletes an user.
func (tx *UserTx) DeleteUser(id string) error {
	return nil
}
