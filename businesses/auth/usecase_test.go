package auth_test

import (
	"context"
	"hungry-baby/app/middleware"
	auth "hungry-baby/businesses/auth"
	user "hungry-baby/businesses/user"
	googleMock "hungry-baby/mocks/google"
	minioMock "hungry-baby/mocks/minio"
	userMock "hungry-baby/mocks/user"
	userCredentialMock "hungry-baby/mocks/userCredential"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	googleRepository      googleMock.Repository
	userUsecase           userMock.Usecase
	userCredentialUsecase userCredentialMock.Usecase
	minioRepository       minioMock.Repository
	configJWT             middleware.ConfigJWT
	authUsecase           auth.Usecase
)

func setup() {
	authUsecase = auth.NewAuthUsecase(2, &googleRepository, &userUsecase, &userCredentialUsecase, configJWT)
}

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestGetTokenFromWeb(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		res := "https://google.com/login"
		googleRepository.On("GetTokenFromWeb", mock.AnythingOfType("string"), mock.AnythingOfType("[]string")).Return(res, nil).Once()
		result, err := authUsecase.GetGoogleLoginURL(context.Background())

		assert.Nil(t, err)
		assert.Equal(t, res, result)
	})
}

func TestVerifyGoogleCode(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		domain := auth.Domain{
			Token:     "",
			ExpiredAt: 0,
		}
		ctx := context.WithValue(context.Background(), "userID", 1)
		googleRepository.On("SaveTokenFromWeb", mock.AnythingOfType("string"), mock.AnythingOfType("[]string"),
			mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(nil).Once()
		userDomain := user.Domain{
			ID:              1,
			Code:            "code",
			Name:            "name",
			ProfileImageID:  1,
			ProfileImageURL: "https://s3.hungrybaby.com/image.jpg",
		}
		userUsecase.On("FindByEmail", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(userDomain, nil).Once()
		result, err := authUsecase.VerifyGoogleCode(ctx, "code")

		assert.Equal(t, result, domain)
		assert.NotNil(t, err)
	})

	// t.Run("test case 2, repository error", func(t *testing.T) {
	// 	domain := auth.Domain{
	// 		ID:         1,
	// 		Type:       "auth",
	// 		URL:        "auth.jpg",
	// 		FullURL:    "https://s3.hungrybaby.com/auth.jpg",
	// 		UserUpload: "1",
	// 	}
	// 	ctx := context.WithValue(context.Background(), "userID", 1)
	// 	errNotFound := errors.New("(Repo) ID Not Found")
	// 	authRepository.On("VerifyGoogleCode", mock.Anything, &domain).Return(auth.Domain{}, errNotFound).Once()
	// 	minioRepository.On("GetAuth", mock.AnythingOfType("string")).Return("https://s3.hungrybaby.com/auth.jpg", nil).Once()
	// 	result, err := authUsecase.VerifyGoogleCode(ctx, &domain)

	// 	assert.Equal(t, result, auth.Domain{})
	// 	assert.Equal(t, err, errNotFound)
	// })
}
