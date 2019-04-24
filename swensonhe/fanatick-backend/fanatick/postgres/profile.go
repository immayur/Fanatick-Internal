package postgres

import (
	"database/sql"
	"fmt"
	"github.com/segmentio/ksuid"
	"github.com/swensonhe/fanatick-backend/fanatick"
	"time"
)

// UserProfileDB is a database for user profiles.
type UserProfileDB struct {
	*DB
}

// UserProfileTx is a user profile transaction.
type UserProfileTx struct {
	*Tx
}

// GetUserProfile returns a user profile.
func (db *UserProfileDB) GetUserProfile(id string) (*fanatick.UserProfile, error) {
	var profile fanatick.UserProfile
	//todo
	return &profile, nil
}

//GetUserProfileByUserId returns an user profile by user id.
func (db *UserProfileDB) GetUserProfileByUserId(userId string) (*fanatick.UserProfile, error) {
	var profile fanatick.UserProfile

	query := `
		SELECT 
			user_profile.id,
			user_profile.user_id,
			user_profile.first_name,
			user_profile.last_name,
			user_profile.profile_pic_url,
			user_profile.created_at
		FROM user_profile
		WHERE user_profile.user_id = $1`

	err := db.QueryRow(query, userId).Scan(
		&profile.ID,
		&profile.UserID,
		&profile.FirstName,
		&profile.LastName,
		&profile.ProfilePicURL,
		&profile.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fanatick.ErrorNotFound
		}

		return nil, err
	}
	return &profile, nil
}

// Query returns a list of user profiles.
func (db *UserProfileDB) QueryUserProfile(params map[fanatick.UserProfileQueryParam]interface{}) ([]*fanatick.UserProfile, error) {

	profiles := []*fanatick.UserProfile{}
	//todo
	return profiles, nil
}

// BeginTx begins a transaction.
func (db *UserProfileDB) BeginUserProfileTx() fanatick.UserProfileTx {
	return &UserProfileTx{Tx: db.Begin()}
}

// CommitUserProfileTx commits user profile transactions
func (tx *UserProfileTx) CommitUserProfileTx() error {
	err := tx.Tx.Commit()
	if err != nil {
		err := tx.Tx.Rollback()
		if err != nil {
			return fmt.Errorf("falied to rollback UserProfileTx after commit UserProfileTx failed")
		}
		return fmt.Errorf("falied to commit UserProfileTx")
	}
	return nil
}

// Create creates a user profile.
func (tx *UserProfileTx) CreateUserProfile(profile *fanatick.UserProfile) error {

	profile.ID = ksuid.New().String()
	profile.CreatedAt = time.Now()

	query := `
		INSERT INTO 
		    user_profile (
		            id,
					user_id,
					first_name,
					last_name,
					profile_pic_url,
		            created_at
		        )
		VALUES ($1, $2, $3, $4, $5, $6)	`

	parameters := make([]interface{}, 6)
	parameters[0] = profile.ID
	parameters[1] = profile.UserID
	parameters[2] = profile.FirstName
	parameters[3] = profile.LastName
	parameters[4] = profile.ProfilePicURL
	parameters[5] = profile.CreatedAt

	_, err := tx.Exec(query, parameters...)
	if err != nil {
		return err
	}

	return nil
}

// Update updates a user profile.
func (tx *UserProfileTx) UpdateUserProfile(profile *fanatick.UserProfile) error {
	t := time.Now()
	profile.UpdatedAt = &t
	query := `
		UPDATE user_profile
		SET first_name =$2, last_name =$3,profile_pic_url=$4 ,updated_at=$5
		WHERE user_id = $1
	`
	parameters := make([]interface{}, 5)
	parameters[0] = profile.UserID
	parameters[1] = profile.FirstName
	parameters[2] = profile.LastName
	parameters[3] = profile.ProfilePicURL
	parameters[4] = profile.UpdatedAt

	_, err := tx.Exec(query, parameters...)
	if err != nil {
		return err
	}

	return nil
}

// Delete deletes a user profile.
func (tx *UserProfileTx) DeleteUserProfile(id string) error {
	//todo
	return nil
}
