package fanatick

// User is a Fanatick user.
type User struct {
	ID string `json:"id"`
}

// UserCreator defines the operation to create a user.
type UserCreator interface {
	Create(user *User) error
}
