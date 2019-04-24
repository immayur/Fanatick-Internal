package fanatick

import (
	"time"
)

// User is a Fanatick user.
type User struct {
	ID            string    `json:"id"`
	FireBaseUUID  string    `json:"firebase_uuid"`
	PhoneNumber   string    `json:"phone_number"`
	IsActive      bool      `json:"is_active"`
	IsSeller      bool      `json:"is_seller"` // false – buyer / true - seller
	LastLoginTime time.Time `json:"last_login_time"`
	CreateUpdateDetails
}

type UserCreateResponse struct {
	User
	Token string `json:"token"`
}

type CreateUpdateDetails struct {
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

// UserQueryParam is an user query parameter.
type UserQueryParam string

// UserGetter is the interface that wraps an user get request.
type UserGetter interface {
	GetUser(id string) (*User, error)
	GetUserByFireBaseUUID(firebaseUUID string) (*User, error)
}

// UserGetterByFireBaseUUID is the interface that wraps an user get request by firebaseUUID.
type UserGetterByFireBaseUUID interface {
	GetUserByFireBaseUUID(firebaseUUID string) (*User, error)
}

// UserQueryer is the interface that wraps an user query request.
type UserQueryer interface {
	QueryUser(params map[UserQueryParam]interface{}) ([]*User, error)
}

// UserCreator is the interface that wraps an user creation request.
type UserCreator interface {
	CreateUser(user *User) error
}

// UserUpdater is the interface that wraps an user update request.
type UserUpdater interface {
	UpdateUser(user *User) error
}

//UserLastLoginUpdater is the interface that wraps UpdateUserLastLoginTime.
type UserLastLoginUpdater interface {
	UpdateUserLastLoginTime(user *User) error
}

// UserDeleter is the interface that wraps an user delete request.
type UserDeleter interface {
	DeleteUser(id string) error
}

// UserTxBeginner is the interface that wraps an user transaction starter.
type UserTxBeginner interface {
	BeginUserTx() UserTx
}

// UserTxCommitter is the interface that wraps an User transaction Committer.
type UserTxCommitter interface {
	CommitUserTx() error
}

// UserStore defines the operations of an user store.
type UserStore interface {
	UserGetter
	UserQueryer
	UserTxBeginner

}

// UserTx defines the operations that may be performed on an user update transaction.
type UserTx interface {
	UserCreator
	UserUpdater
	UserDeleter
	UserTxCommitter
	UserGetterByFireBaseUUID
	UserLastLoginUpdater
}

// The user query params.
const (
	// UserQueryParamLimit indicates the maximum number of users to return.
	UserQueryParamLimit = UserQueryParam("limit")

	// UserQueryParamBefore indicates the last user of the previously queried results.
	UserQueryParamBefore = UserQueryParam("before")
)

// UserGetAndCreate defines the operations that are done on user login
// ie : create if user doesn't exist.
//type UserGetAndCreate interface {
//	UserCreator
//	UserGetterByFireBaseUUID
//}
