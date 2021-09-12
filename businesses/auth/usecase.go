package auth

import (
	"context"
	"errors"
	_middleware "hungry-baby/app/middleware"
	googleBusiness "hungry-baby/businesses/google"
	userBusiness "hungry-baby/businesses/user"
	userCredentialBusiness "hungry-baby/businesses/userCredential"
	"hungry-baby/drivers/thirdparties/google"
	googleHelpers "hungry-baby/helpers/google"
	"hungry-baby/helpers/interfacepkg"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/rs/xid"
)

type authUsecase struct {
	googleRepository      googleBusiness.Repository
	userUsecase           userBusiness.Usecase
	userCredentialUsecase userCredentialBusiness.Usecase
	configJWT             _middleware.ConfigJWT
	contextTimeout        time.Duration
}

func NewAuthUsecase(timeout time.Duration, googleRepo googleBusiness.Repository,
	userUsecase userBusiness.Usecase, userCredentialUsecase userCredentialBusiness.Usecase,
	configJWT _middleware.ConfigJWT) Usecase {
	return &authUsecase{
		googleRepository:      googleRepo,
		userUsecase:           userUsecase,
		userCredentialUsecase: userCredentialUsecase,
		configJWT:             configJWT,
		contextTimeout:        timeout,
	}
}

func (uc *authUsecase) GetGoogleLoginURL(ctx context.Context) (string, error) {
	res, err := uc.googleRepository.GetTokenFromWeb("", google.CalendarScopes)
	if err != nil {
		return res, err
	}

	return res, nil
}

// VerifyGoogleCode ...
func (uc *authUsecase) VerifyGoogleCode(ctx context.Context, code string) (Domain, error) {
	scopes := google.CalendarScopes
	tokenPath := "../files/" + xid.New().String() + ".json"
	err := uc.googleRepository.SaveTokenFromWeb("", scopes, code, tokenPath)
	if err != nil {
		return Domain{}, err
	}

	// Get Token
	tokenByte, err := ioutil.ReadFile(tokenPath)
	if err != nil {
		return Domain{}, err
	}
	var token userCredentialBusiness.DomainRegistrationDetails
	interfacepkg.UnmarshalCb(string(tokenByte), &token)
	defer os.Remove(tokenPath)

	// Verify token scopes
	err = googleHelpers.VerifyTokenScope(token.AccessToken, scopes)
	if err != nil {
		return Domain{}, err
	}

	// Verify token
	gmailUser, err := googleHelpers.GetGoogleProfile(token.AccessToken)
	if err != nil {
		return Domain{}, err
	}
	if gmailUser["email"] == nil {
		return Domain{}, errors.New("Invalid Email")
	}
	log.Print(gmailUser)

	// Find userCredential
	user, err := uc.userUsecase.FindByEmail(ctx, gmailUser["email"].(string), "")
	if user.ID == 0 {
		user, err = uc.userUsecase.Store(ctx, &userBusiness.Domain{
			Code:   xid.New().String(),
			Email:  gmailUser["email"].(string),
			Name:   gmailUser["email"].(string),
			Status: true,
		})
		if err != nil {
			return Domain{}, err
		}

		now := time.Now().UTC()
		_, err = uc.userCredentialUsecase.Store(ctx, &userCredentialBusiness.Domain{
			UserID:              user.ID,
			Type:                "gmail",
			Email:               user.Email,
			EmailValidAt:        now.Format(time.RFC3339),
			RegistrationDetails: token,
			Status:              true,
		})
		if err != nil {
			return Domain{}, err
		}
	} else {
		userCredential, err := uc.userCredentialUsecase.FindByEmail(ctx, user.Email, "")
		if err != nil {
			now := time.Now().UTC()
			_, err = uc.userCredentialUsecase.Store(ctx, &userCredentialBusiness.Domain{
				UserID:              user.ID,
				Type:                "gmail",
				Email:               user.Email,
				EmailValidAt:        now.Format(time.RFC3339),
				RegistrationDetails: token,
				Status:              true,
			})
			if err != nil {
				return Domain{}, err
			}
		} else {
			userCredential.RegistrationDetails = token
			_, err = uc.userCredentialUsecase.Update(ctx, &userCredential)
			if err != nil {
				return Domain{}, err
			}
		}
	}

	res := Domain{
		Token:     uc.configJWT.GenerateToken(user.ID),
		ExpiredAt: int(time.Now().Local().Add(time.Hour * time.Duration(int64(uc.configJWT.ExpiresDuration))).Unix()),
	}

	return res, err
}
