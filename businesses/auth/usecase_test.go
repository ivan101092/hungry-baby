package auth_test

import (
	"context"
	"encoding/json"
	"errors"
	"hungry-baby/app/middleware"
	auth "hungry-baby/businesses/auth"
	user "hungry-baby/businesses/user"
	userCredential "hungry-baby/businesses/userCredential"
	googleMock "hungry-baby/mocks/google"
	userMock "hungry-baby/mocks/user"
	userCredentialMock "hungry-baby/mocks/userCredential"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestGetTokenFromWeb(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		var (
			googleRepository      googleMock.Repository
			userUsecase           userMock.Usecase
			userCredentialUsecase userCredentialMock.Usecase
			configJWT             middleware.ConfigJWT
		)
		authUsecase := auth.NewAuthUsecase(2, &googleRepository, &userUsecase, &userCredentialUsecase, configJWT)

		res := "https://google.com/login"
		googleRepository.On("GetTokenFromWeb", mock.AnythingOfType("string"), mock.AnythingOfType("[]string")).Return(res, nil).Once()
		result, err := authUsecase.GetGoogleLoginURL(context.Background())

		assert.Nil(t, err)
		assert.Equal(t, res, result)
	})

	t.Run("test case 2, error", func(t *testing.T) {
		var (
			googleRepository      googleMock.Repository
			userUsecase           userMock.Usecase
			userCredentialUsecase userCredentialMock.Usecase
			configJWT             middleware.ConfigJWT
		)
		authUsecase := auth.NewAuthUsecase(2, &googleRepository, &userUsecase, &userCredentialUsecase, configJWT)

		errError := errors.New("Error")
		googleRepository.On("GetTokenFromWeb", mock.AnythingOfType("string"), mock.AnythingOfType("[]string")).Return("", errError).Once()
		_, err := authUsecase.GetGoogleLoginURL(context.Background())

		assert.Equal(t, errError, err)
	})
}

func TestVerifyGoogleCode(t *testing.T) {
	t.Run("test case 1, user exist, credential exist", func(t *testing.T) {
		var (
			googleRepository      googleMock.Repository
			userUsecase           userMock.Usecase
			userCredentialUsecase userCredentialMock.Usecase
			configJWT             middleware.ConfigJWT
		)
		authUsecase := auth.NewAuthUsecase(2, &googleRepository, &userUsecase, &userCredentialUsecase, configJWT)
		ctx := context.WithValue(context.Background(), "userID", 1)

		// Create temporary file
		b, _ := json.MarshalIndent(map[string]interface{}{"test": "test"}, "", " ")
		ioutil.WriteFile("../../files/test.json", b, 0644)

		googleRepository.On("SaveTokenFromWeb", mock.AnythingOfType("string"), mock.AnythingOfType("[]string"),
			mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return("../../files/test.json", nil).Once()
		googleRepository.On("VerifyTokenScope", mock.AnythingOfType("string"),
			mock.AnythingOfType("[]string")).Return(nil).Once()
		gmailUser := map[string]interface{}{"email": "mock@email.com"}
		googleRepository.On("GetGoogleProfile", mock.AnythingOfType("string")).Return(gmailUser, nil).Once()
		userDomain := user.Domain{
			ID:              1,
			Code:            "code",
			Name:            "name",
			ProfileImageID:  1,
			ProfileImageURL: "https://s3.hungrybaby.com/image.jpg",
		}
		userUsecase.On("FindByEmail", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(userDomain, nil).Once()
		userCredentialDomain := userCredential.Domain{
			ID:     1,
			UserID: 1,
			Type:   "gmail",
		}
		userCredentialUsecase.On("FindByUserType", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(userCredentialDomain, nil).Once()
		userCredentialUsecase.On("Update", mock.Anything,
			mock.AnythingOfType("*userCredential.Domain")).Return(userCredentialDomain, nil).Once()
		_, err := authUsecase.VerifyGoogleCode(ctx, "code")

		assert.Nil(t, err)
	})

	t.Run("test case 2, user exist, credential not exist", func(t *testing.T) {
		var (
			googleRepository      googleMock.Repository
			userUsecase           userMock.Usecase
			userCredentialUsecase userCredentialMock.Usecase
			configJWT             middleware.ConfigJWT
		)
		authUsecase := auth.NewAuthUsecase(2, &googleRepository, &userUsecase, &userCredentialUsecase, configJWT)
		ctx := context.WithValue(context.Background(), "userID", 1)

		// Create temporary file
		b, _ := json.MarshalIndent(map[string]interface{}{"test": "test"}, "", " ")
		ioutil.WriteFile("../../files/test.json", b, 0644)

		googleRepository.On("SaveTokenFromWeb", mock.AnythingOfType("string"), mock.AnythingOfType("[]string"),
			mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return("../../files/test.json", nil).Once()
		googleRepository.On("VerifyTokenScope", mock.AnythingOfType("string"),
			mock.AnythingOfType("[]string")).Return(nil).Once()
		gmailUser := map[string]interface{}{"email": "mock@email.com"}
		googleRepository.On("GetGoogleProfile", mock.AnythingOfType("string")).Return(gmailUser, nil).Once()
		userDomain := user.Domain{
			ID:              1,
			Code:            "code",
			Name:            "name",
			ProfileImageID:  1,
			ProfileImageURL: "https://s3.hungrybaby.com/image.jpg",
		}
		userUsecase.On("FindByEmail", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(userDomain, nil).Once()
		userCredentialDomain := userCredential.Domain{
			ID:     1,
			UserID: 1,
			Type:   "gmail",
		}
		userCredentialUsecase.On("FindByUserType", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(userCredential.Domain{}, errors.New("data not found")).Once()
		userCredentialUsecase.On("Store", mock.Anything,
			mock.AnythingOfType("*userCredential.Domain")).Return(userCredentialDomain, nil).Once()
		_, err := authUsecase.VerifyGoogleCode(ctx, "code")

		assert.Nil(t, err)
	})

	t.Run("test case 3, user not exist", func(t *testing.T) {
		var (
			googleRepository      googleMock.Repository
			userUsecase           userMock.Usecase
			userCredentialUsecase userCredentialMock.Usecase
			configJWT             middleware.ConfigJWT
		)
		authUsecase := auth.NewAuthUsecase(2, &googleRepository, &userUsecase, &userCredentialUsecase, configJWT)
		ctx := context.WithValue(context.Background(), "userID", 1)

		// Create temporary file
		b, _ := json.MarshalIndent(map[string]interface{}{"test": "test"}, "", " ")
		ioutil.WriteFile("../../files/test.json", b, 0644)

		googleRepository.On("SaveTokenFromWeb", mock.AnythingOfType("string"), mock.AnythingOfType("[]string"),
			mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return("../../files/test.json", nil).Once()
		googleRepository.On("VerifyTokenScope", mock.AnythingOfType("string"),
			mock.AnythingOfType("[]string")).Return(nil).Once()
		gmailUser := map[string]interface{}{"email": "mock@email.com"}
		googleRepository.On("GetGoogleProfile", mock.AnythingOfType("string")).Return(gmailUser, nil).Once()
		userUsecase.On("FindByEmail", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(user.Domain{}, errors.New("data not found")).Once()
		userDomain := user.Domain{
			ID:              1,
			Code:            "code",
			Name:            "name",
			ProfileImageID:  1,
			ProfileImageURL: "https://s3.hungrybaby.com/image.jpg",
		}
		userUsecase.On("Store", mock.Anything, mock.AnythingOfType("*user.Domain")).Return(userDomain, nil).Once()
		userCredentialDomain := userCredential.Domain{
			ID:     1,
			UserID: 1,
			Type:   "gmail",
		}
		userCredentialUsecase.On("Store", mock.Anything,
			mock.AnythingOfType("*userCredential.Domain")).Return(userCredentialDomain, nil).Once()
		_, err := authUsecase.VerifyGoogleCode(ctx, "code")

		assert.Nil(t, err)
	})
}
