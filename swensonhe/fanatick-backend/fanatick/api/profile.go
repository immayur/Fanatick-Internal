package api

import (
	"github.com/swensonhe/fanatick-backend/fanatick"
)

// UserProfileService performs operations on userProfiles.
type UserProfileService struct {
	UserProfileStore fanatick.UserProfileStore
	UserProfileTx    fanatick.UserProfileTx
	Logger           fanatick.Logger
}

// GetUserProfile returns an userProfile.
func (svc *UserProfileService) GetUserProfile(id string) (*fanatick.UserProfile, error) {
	userProfile, err := svc.UserProfileStore.GetUserProfile(id)
	if err != nil {
		if err == fanatick.ErrorNotFound {
			return nil, ErrorNotFound(`UserProfile not found.`)
		}

		go svc.Logger.Error(err)
		return nil, ErrorInternal(err.Error())
	}

	return userProfile, nil
}

//GetUserProfileByUserId  returns an userProfile by user id
func (svc *UserProfileService) GetUserProfileByUserId(userID string) (*fanatick.UserProfile, error) {
	userProfile, err := svc.UserProfileStore.GetUserProfileByUserId(userID)
	if err != nil {
		if err == fanatick.ErrorNotFound {
			return nil, ErrorNotFound(`UserProfile not found.`)
		}

		go svc.Logger.Error(err)
		return nil, ErrorInternal(err.Error())
	}

	return userProfile, nil
}

// QueryUserProfile returns userProfiles.
func (svc *UserProfileService) QueryUserProfile(params map[fanatick.UserProfileQueryParam]interface{}) ([]*fanatick.UserProfile, error) {
	userProfiles, err := svc.UserProfileStore.QueryUserProfile(params)
	if err != nil {
		go svc.Logger.Error(err)
		return nil, ErrorInternal(err.Error())
	}

	return userProfiles, nil
}

//CreateUserProfile creates a user profile
func (svc *UserProfileService) CreateUserProfile(userProfile *fanatick.UserProfile) error {
	tx := svc.UserProfileStore.BeginUserProfileTx()

	err := tx.CreateUserProfile(userProfile)
	if err != nil {
		if err == fanatick.ErrorInternal {
			return ErrorNotFound(`Unable to Create UserProfile.`)
		}

		go svc.Logger.Error(err)
		return ErrorInternal(err.Error())
	}

	err = tx.CommitUserProfileTx()
	if err != nil {
		return ErrorInternal(err.Error())
	}
	return nil
}

//UpdateUserProfile  Updates User Profile
func (svc *UserProfileService) UpdateUserProfile(userProfile *fanatick.UserProfile) error {

	tx := svc.UserProfileStore.BeginUserProfileTx()

	// TODO: if we want to fetch details from DB
	//  and then update in case all attributes are not passed in request body

	//existingUserProfile, err := svc.UserProfileStore.GetUserProfileByUserId(userProfile.UserID)
	//if err != nil {
	//	if err == fanatick.ErrorNotFound {
	//		return ErrorNotFound(`UserProfile not found.`)
	//	}
	//
	//	return ErrorInternal(err.Error())
	//}
	//
	//userProfile = existingUserProfile

	err := tx.UpdateUserProfile(userProfile)
	if err != nil {
		if err == fanatick.ErrorInternal {
			return ErrorNotFound(`Unable to Update UserProfile.`)
		}

		go svc.Logger.Error(err)
		return ErrorInternal(err.Error())
	}

	err = tx.CommitUserProfileTx()
	if err != nil {
		return ErrorInternal(err.Error())
	}

	return nil
}
