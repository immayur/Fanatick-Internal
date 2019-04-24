package api

import (
	"github.com/swensonhe/fanatick-backend/fanatick"
	"github.com/swensonhe/fanatick-backend/fanatick/firebase"
)

// UserService performs operations on users.
type UserService struct {
	UserStore      fanatick.UserStore
	Logger         fanatick.Logger
	FirebaseClient *firebase.Client
}

// GetUser returns an user.
func (svc *UserService) GetUser(id string) (*fanatick.User, error) {
	user, err := svc.UserStore.GetUser(id)
	if err != nil {
		if err == fanatick.ErrorNotFound {
			return nil, ErrorNotFound(`User not found.`)
		}

		go svc.Logger.Error(err)
		return nil, ErrorInternal(err.Error())
	}

	return user, nil
}

// GetUserByFireBaseUUID returns an user.
func (svc *UserService) GetUserByFireBaseUUID(firebaseUUID string) (*fanatick.User, error) {
	user, err := svc.UserStore.GetUserByFireBaseUUID(firebaseUUID)
	if err != nil {
		if err == fanatick.ErrorNotFound {
			return nil, ErrorNotFound(`User not found.`)
		}

		go svc.Logger.Error(err)
		return nil, ErrorInternal(err.Error())
	}

	return user, nil
}

// QueryUser returns users.
func (svc *UserService) QueryUser(params map[fanatick.UserQueryParam]interface{}) ([]*fanatick.User, error) {
	users, err := svc.UserStore.QueryUser(params)
	if err != nil {
		go svc.Logger.Error(err)
		return nil, ErrorInternal(err.Error())
	}

	return users, nil
}

//CreateUser creates a new user if not exists
func (svc *UserService) CreateUser(user *fanatick.User) error {

	//if !svc.FirebaseClient.UserExists(*user.FireBaseUUID){
	//	return ErrorUnauthorized(`unable to validate user in firebase`)
	//}

	tx := svc.UserStore.BeginUserTx()
	existingUser, err := tx.GetUserByFireBaseUUID(user.FireBaseUUID)
	if err != nil {
		if err != fanatick.ErrorNotFound {
			return ErrorInternal(err.Error())
		}
		existingUser = user
		err = tx.CreateUser(existingUser)
		if err != nil {
			if err == fanatick.ErrorInternal {
				return ErrorNotFound(`Unable to Create User.`)
			}

			go svc.Logger.Error(err)
			return ErrorInternal(err.Error())
		}
	}

	//// copy to passed user instance
	//if user.FireBaseUUID == existingUser.FireBaseUUID {
	//	fmt.Println(existingUser, user)
	//	user = existingUser
	//}

	err = tx.UpdateUserLastLoginTime(existingUser)
	if err != nil {
		return ErrorInternal(err.Error())
	}

	*user = *existingUser

	err = tx.CommitUserTx()
	if err != nil {
		return ErrorInternal(err.Error())
	}

	return nil
}
