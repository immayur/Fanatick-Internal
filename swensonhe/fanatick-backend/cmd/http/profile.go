package main

import (
	"github.com/go-chi/chi"
	"net/http"

	"github.com/swensonhe/fanatick-backend/fanatick"
)

// GetUserProfile godoc
// @Summary Show a userProfile
// @Description get userProfile by ID
// @ID int
// @Accept  json
// @Produce  json
// @Param id path int true "UserProfile ID"
// @Success 200
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /userProfiles/{id} [get]
func GetUserProfileByUserIdHandler(userProfileGetterByUserId fanatick.UserProfileGetterByUserId) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: move decode json and fetch vars into function
		userID := chi.URLParam(r, "user_id")
		userProfile, err := userProfileGetterByUserId.GetUserProfileByUserId(userID)
		if err != nil {
			NewErrorWriter(w).Write(err)
			return
		}
		res := newCommonSuccessResponse(userProfile, CommonStatus200GETCallResponseMessage)
		NewJSONWriter(w).Write(res, http.StatusOK)
	}
}

// GetUserProfiles godoc
// @Summary Show a userProfiles
// @Description get userProfiles
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /userProfiles [get]
func QueryUserProfilesHandler(userProfileQueryer fanatick.UserProfileQueryer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := map[fanatick.UserProfileQueryParam]interface{}{}

		if limit := r.URL.Query().Get("limit"); limit != "" {
			params[fanatick.UserProfileQueryParamLimit] = limit
		}

		if before := r.URL.Query().Get("before"); before != "" {
			params[fanatick.UserProfileQueryParamBefore] = before
		}

		userProfiles, err := userProfileQueryer.QueryUserProfile(params)
		if err != nil {
			NewErrorWriter(w).Write(err)
			return
		}

		NewJSONWriter(w).Write(userProfiles, http.StatusOK)
	}
}

// PostUserProfileHandler godoc
// @Summary create a userProfile
// @Description create a userProfile
// @ID int
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /userProfiles/ [post]
func PostUserProfileHandler(userProfileCreator fanatick.UserProfileCreator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userProfile := fanatick.UserProfile{}

		err := decodeJSON(r, &userProfile)
		if err != nil {
			NewErrorWriter(w).Write(err)
			return
		}

		userID := chi.URLParam(r, "user_id")
		userProfile.UserID = userID
		err = userProfileCreator.CreateUserProfile(&userProfile)
		if err != nil {
			NewErrorWriter(w).Write(err)
			return
		}

		res := newCommonSuccessResponse(userProfile, CommonStatus200POSTCallResponseMessage)
		NewJSONWriter(w).Write(res, http.StatusOK)
	}
}

func PatchUserProfileHandler(UserProfileUpdater fanatick.UserProfileUpdater) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userProfile := fanatick.UserProfile{}

		userID := chi.URLParam(r, "user_id")
		userProfile.UserID = userID

		err := decodeJSON(r, &userProfile)
		if err != nil {
			NewErrorWriter(w).Write(err)
			return
		}

		err = UserProfileUpdater.UpdateUserProfile(&userProfile)
		if err != nil {
			NewErrorWriter(w).Write(err)
			return
		}

		res := newCommonSuccessResponse(userProfile, CommonStatus200PATCHCallResponseMessage)
		NewJSONWriter(w).Write(res, http.StatusOK)
	}
}

