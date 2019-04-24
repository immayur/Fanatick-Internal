package fanatick

// UserProfile is a Fanatick user UserProfile with user details.
type UserProfile struct {
	ID            string  `json:"id"`
	UserID        string  `json:"user_id"`
	FirstName     string  `json:"first_name"`
	LastName      string  `json:"last_name"`
	ProfilePicURL *string `json:"profile_pic_url"`
	CreateUpdateDetails
}

// UserProfileQueryParam is an profile query parameter.
type UserProfileQueryParam string

// UserProfileGetter is the interface that wraps an profile get request.
type UserProfileGetter interface {
	GetUserProfile(id string) (*UserProfile, error)
}

// UserProfileGetterByUserId is the interface that wraps an profile get request by user id.
type UserProfileGetterByUserId interface {
	GetUserProfileByUserId(userId string) (*UserProfile, error)
}

// UserProfileQueryer is the interface that wraps an profile query request.
type UserProfileQueryer interface {
	QueryUserProfile(params map[UserProfileQueryParam]interface{}) ([]*UserProfile, error)
}

// UserProfileCreator is the interface that wraps an profile creation request.
type UserProfileCreator interface {
	CreateUserProfile(profile *UserProfile) error
}

// UserProfileUpdater is the interface that wraps an profile update request.
type UserProfileUpdater interface {
	UpdateUserProfile(profile *UserProfile) error
}

// UserProfileDeleter is the interface that wraps an profile delete request.
type UserProfileDeleter interface {
	DeleteUserProfile(id string) error
}

// UserProfileTxBeginner is the interface that wraps an profile transaction starter.
type UserProfileTxBeginner interface {
	BeginUserProfileTx() UserProfileTx
}

// UserProfileTxCommitter is the interface that wraps an User profile transaction Committer.
type UserProfileTxCommitter interface {
	CommitUserProfileTx() error
}

// UserProfileStore defines the operations of an profile store.
type UserProfileStore interface {
	UserProfileGetter
	UserProfileGetterByUserId
	UserProfileQueryer
	UserProfileTxBeginner
}

// UserProfileTx defines the operations that may be performed on an profile update transaction.
type UserProfileTx interface {
	UserProfileCreator
	UserProfileUpdater
	UserProfileDeleter
	UserProfileTxCommitter
}

// The profile query params.
const (
	// UserProfileQueryParamLimit indicates the maximum number of profiles to return.
	UserProfileQueryParamLimit = UserProfileQueryParam("limit")

	// UserProfileQueryParamBefore indicates the last profile of the previously queried results.
	UserProfileQueryParamBefore = UserProfileQueryParam("before")
)
