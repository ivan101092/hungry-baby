package userCredential_test

import (
	"context"
	"errors"
	"hungry-baby/businesses"
	userCredential "hungry-baby/businesses/userCredential"
	userCredentialMock "hungry-baby/mocks/userCredential"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestFindAll(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		var (
			userCredentialRepository userCredentialMock.Repository
		)
		userCredentialUsecase := userCredential.NewUserCredentialUsecase(2, &userCredentialRepository)

		domain := []userCredential.Domain{
			{
				ID:     1,
				Type:   "code",
				Email:  "name",
				Status: true,
			},
		}
		userCredentialRepository.On("FindAll", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(domain, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		result, err := userCredentialUsecase.FindAll(ctx, "1", "true")

		assert.Nil(t, err)
		assert.Equal(t, domain, result)
	})

	t.Run("test case 2, error", func(t *testing.T) {
		var (
			userCredentialRepository userCredentialMock.Repository
		)
		userCredentialUsecase := userCredential.NewUserCredentialUsecase(2, &userCredentialRepository)

		errError := errors.New("Error Repo")
		userCredentialRepository.On("FindAll", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(nil, errError).Once()
		ctx := context.WithValue(context.Background(), "userID", 1)
		_, err := userCredentialUsecase.FindAll(ctx, "1", "true")

		assert.Equal(t, errError, err)
	})
}

func TestFind(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		var (
			userCredentialRepository userCredentialMock.Repository
		)
		userCredentialUsecase := userCredential.NewUserCredentialUsecase(2, &userCredentialRepository)

		domain := []userCredential.Domain{
			{
				ID:     1,
				Type:   "code",
				Email:  "name",
				Status: true,
			},
		}
		userCredentialRepository.On("Find", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("string"), mock.AnythingOfType("int"),
			mock.AnythingOfType("int")).Return(domain, 1, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		result, total, err := userCredentialUsecase.Find(ctx, "1", "true", 0, 0)

		assert.Nil(t, err)
		assert.Equal(t, domain, result)
		assert.Equal(t, 1, total)
	})

	t.Run("test case 2, error", func(t *testing.T) {
		var (
			userCredentialRepository userCredentialMock.Repository
		)
		userCredentialUsecase := userCredential.NewUserCredentialUsecase(2, &userCredentialRepository)

		errError := errors.New("Error Repo")
		userCredentialRepository.On("Find", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("string"), mock.AnythingOfType("int"),
			mock.AnythingOfType("int")).Return(nil, 0, errError).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		_, _, err := userCredentialUsecase.Find(ctx, "1", "true", 0, 0)

		assert.Equal(t, errError, err)
	})
}

func TestFindByID(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		var (
			userCredentialRepository userCredentialMock.Repository
		)
		userCredentialUsecase := userCredential.NewUserCredentialUsecase(2, &userCredentialRepository)

		domain := userCredential.Domain{
			ID:     1,
			Type:   "code",
			Email:  "name",
			Status: true,
		}
		userCredentialRepository.On("FindByID", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("string")).Return(domain, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		result, err := userCredentialUsecase.FindByID(ctx, 1, "true")

		assert.Nil(t, err)
		assert.Equal(t, domain, result)
	})

	t.Run("test case 2, invalid id", func(t *testing.T) {
		var (
			userCredentialRepository userCredentialMock.Repository
		)
		userCredentialUsecase := userCredential.NewUserCredentialUsecase(2, &userCredentialRepository)

		domain := userCredential.Domain{
			ID:     1,
			Type:   "code",
			Email:  "name",
			Status: true,
		}
		userCredentialRepository.On("FindByID", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("string")).Return(domain, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		_, err := userCredentialUsecase.FindByID(ctx, -1, "true")

		assert.Equal(t, businesses.ErrIDNotFound, err)
	})

	t.Run("test case 3, error", func(t *testing.T) {
		var (
			userCredentialRepository userCredentialMock.Repository
		)
		userCredentialUsecase := userCredential.NewUserCredentialUsecase(2, &userCredentialRepository)

		errError := errors.New("Error Repo")
		userCredentialRepository.On("FindByID", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("string")).Return(userCredential.Domain{}, errError).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		_, err := userCredentialUsecase.FindByID(ctx, 1, "true")

		assert.Equal(t, errError, err)
	})
}

func TestFindByEmail(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		var (
			userCredentialRepository userCredentialMock.Repository
		)
		userCredentialUsecase := userCredential.NewUserCredentialUsecase(2, &userCredentialRepository)

		domain := userCredential.Domain{
			ID:     1,
			Type:   "code",
			Email:  "name",
			Status: true,
		}
		userCredentialRepository.On("FindByEmail", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(domain, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		result, err := userCredentialUsecase.FindByEmail(ctx, "code", "true")

		assert.Nil(t, err)
		assert.Equal(t, domain, result)
	})

	t.Run("test case 2, error", func(t *testing.T) {
		var (
			userCredentialRepository userCredentialMock.Repository
		)
		userCredentialUsecase := userCredential.NewUserCredentialUsecase(2, &userCredentialRepository)

		errError := errors.New("Error Repo")
		userCredentialRepository.On("FindByEmail", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(userCredential.Domain{}, errError).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		_, err := userCredentialUsecase.FindByEmail(ctx, "code", "true")

		assert.Equal(t, errError, err)
	})
}

func TestFindByUserType(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		var (
			userCredentialRepository userCredentialMock.Repository
		)
		userCredentialUsecase := userCredential.NewUserCredentialUsecase(2, &userCredentialRepository)

		domain := userCredential.Domain{
			ID:     1,
			Type:   "code",
			Email:  "name",
			Status: true,
		}
		userCredentialRepository.On("FindByUserType", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(domain, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		result, err := userCredentialUsecase.FindByUserType(ctx, 1, "code", "true")

		assert.Nil(t, err)
		assert.Equal(t, domain, result)
	})

	t.Run("test case 2, error", func(t *testing.T) {
		var (
			userCredentialRepository userCredentialMock.Repository
		)
		userCredentialUsecase := userCredential.NewUserCredentialUsecase(2, &userCredentialRepository)

		errError := errors.New("Error Repo")
		userCredentialRepository.On("FindByUserType", mock.Anything, mock.AnythingOfType("int"), mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(userCredential.Domain{}, errError).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		_, err := userCredentialUsecase.FindByUserType(ctx, 1, "code", "true")

		assert.Equal(t, errError, err)
	})
}

func TestStore(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		var (
			userCredentialRepository userCredentialMock.Repository
		)
		userCredentialUsecase := userCredential.NewUserCredentialUsecase(2, &userCredentialRepository)

		domain := userCredential.Domain{
			ID:     1,
			Type:   "code",
			Email:  "name",
			Status: true,
		}
		errError := errors.New("record not found")
		userCredentialRepository.On("FindByUserType", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(userCredential.Domain{}, errError).Once()
		userCredentialRepository.On("Store", mock.Anything, mock.AnythingOfType("*userCredential.Domain")).Return(domain, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		result, err := userCredentialUsecase.Store(ctx, &domain)

		assert.Nil(t, err)
		assert.Equal(t, domain, result)
	})

	t.Run("test case 2, find error", func(t *testing.T) {
		var (
			userCredentialRepository userCredentialMock.Repository
		)
		userCredentialUsecase := userCredential.NewUserCredentialUsecase(2, &userCredentialRepository)

		domain := userCredential.Domain{
			ID:     1,
			Type:   "code",
			Email:  "name",
			Status: true,
		}
		errError := errors.New("internal server")
		userCredentialRepository.On("FindByUserType", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(userCredential.Domain{}, errError).Once()
		userCredentialRepository.On("Store", mock.Anything, mock.AnythingOfType("*userCredential.Domain")).Return(domain, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		_, err := userCredentialUsecase.Store(ctx, &domain)

		assert.Equal(t, errError, err)
	})

	t.Run("test case 3, find exist", func(t *testing.T) {
		var (
			userCredentialRepository userCredentialMock.Repository
		)
		userCredentialUsecase := userCredential.NewUserCredentialUsecase(2, &userCredentialRepository)

		domain := userCredential.Domain{
			ID:     1,
			Type:   "code",
			Email:  "name",
			Status: true,
		}
		userCredentialRepository.On("FindByUserType", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(domain, nil).Once()
		userCredentialRepository.On("Store", mock.Anything, mock.AnythingOfType("*userCredential.Domain")).Return(domain, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		_, err := userCredentialUsecase.Store(ctx, &domain)

		assert.Equal(t, businesses.ErrDuplicateData, err)
	})

	t.Run("test case 4, store error", func(t *testing.T) {
		var (
			userCredentialRepository userCredentialMock.Repository
		)
		userCredentialUsecase := userCredential.NewUserCredentialUsecase(2, &userCredentialRepository)

		domain := userCredential.Domain{
			ID:     1,
			Type:   "code",
			Email:  "name",
			Status: true,
		}
		errError := errors.New("record not found")
		errRepo := errors.New("repo error")
		userCredentialRepository.On("FindByUserType", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(userCredential.Domain{}, errError).Once()
		userCredentialRepository.On("Store", mock.Anything, mock.AnythingOfType("*userCredential.Domain")).Return(userCredential.Domain{}, errRepo).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		_, err := userCredentialUsecase.Store(ctx, &domain)

		assert.Equal(t, errRepo, err)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		var (
			userCredentialRepository userCredentialMock.Repository
		)
		userCredentialUsecase := userCredential.NewUserCredentialUsecase(2, &userCredentialRepository)

		domain := userCredential.Domain{
			ID:     1,
			Type:   "code",
			Email:  "name",
			Status: true,
		}
		errError := errors.New("record not found")
		userCredentialRepository.On("FindByUserType", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(userCredential.Domain{}, errError).Once()
		userCredentialRepository.On("Update", mock.Anything, mock.AnythingOfType("*userCredential.Domain")).Return(domain, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		result, err := userCredentialUsecase.Update(ctx, &domain)

		assert.Nil(t, err)
		assert.Equal(t, domain, result)
	})

	t.Run("test case 2, find error", func(t *testing.T) {
		var (
			userCredentialRepository userCredentialMock.Repository
		)
		userCredentialUsecase := userCredential.NewUserCredentialUsecase(2, &userCredentialRepository)

		domain := userCredential.Domain{
			ID:     1,
			Type:   "code",
			Email:  "name",
			Status: true,
		}
		errError := errors.New("internal server")
		userCredentialRepository.On("FindByUserType", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(userCredential.Domain{}, errError).Once()
		userCredentialRepository.On("Update", mock.Anything, mock.AnythingOfType("*userCredential.Domain")).Return(domain, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		_, err := userCredentialUsecase.Update(ctx, &domain)

		assert.Equal(t, errError, err)
	})

	t.Run("test case 3, find exist", func(t *testing.T) {
		var (
			userCredentialRepository userCredentialMock.Repository
		)
		userCredentialUsecase := userCredential.NewUserCredentialUsecase(2, &userCredentialRepository)

		domain := userCredential.Domain{
			ID:     1,
			Type:   "code",
			Email:  "name",
			Status: true,
		}
		userCredentialRepository.On("FindByUserType", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(domain, nil).Once()
		userCredentialRepository.On("Update", mock.Anything, mock.AnythingOfType("*userCredential.Domain")).Return(domain, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		domain.ID = 2
		_, err := userCredentialUsecase.Update(ctx, &domain)

		assert.Equal(t, businesses.ErrDuplicateData, err)
	})

	t.Run("test case 4, store error", func(t *testing.T) {
		var (
			userCredentialRepository userCredentialMock.Repository
		)
		userCredentialUsecase := userCredential.NewUserCredentialUsecase(2, &userCredentialRepository)

		domain := userCredential.Domain{
			ID:     1,
			Type:   "code",
			Email:  "name",
			Status: true,
		}
		errError := errors.New("record not found")
		errRepo := errors.New("repo error")
		userCredentialRepository.On("FindByUserType", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(userCredential.Domain{}, errError).Once()
		userCredentialRepository.On("Update", mock.Anything, mock.AnythingOfType("*userCredential.Domain")).Return(userCredential.Domain{}, errRepo).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		_, err := userCredentialUsecase.Update(ctx, &domain)

		assert.Equal(t, errRepo, err)
	})
}

func TestDelete(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		var (
			userCredentialRepository userCredentialMock.Repository
		)
		userCredentialUsecase := userCredential.NewUserCredentialUsecase(2, &userCredentialRepository)

		domain := userCredential.Domain{
			ID:     1,
			Type:   "code",
			Email:  "name",
			Status: true,
		}
		userCredentialRepository.On("FindByID", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("string")).Return(domain, nil).Once()
		userCredentialRepository.On("Delete", mock.Anything, mock.AnythingOfType("*userCredential.Domain")).Return(domain, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		result, err := userCredentialUsecase.Delete(ctx, &domain)

		assert.Nil(t, err)
		assert.Equal(t, domain, result)
	})

	t.Run("test case 2, find error", func(t *testing.T) {
		var (
			userCredentialRepository userCredentialMock.Repository
		)
		userCredentialUsecase := userCredential.NewUserCredentialUsecase(2, &userCredentialRepository)

		domain := userCredential.Domain{
			ID:     1,
			Type:   "code",
			Email:  "name",
			Status: true,
		}
		errError := errors.New("record not found")
		userCredentialRepository.On("FindByID", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("string")).Return(userCredential.Domain{}, errError).Once()
		userCredentialRepository.On("Delete", mock.Anything, mock.AnythingOfType("*userCredential.Domain")).Return(userCredential.Domain{}, errError).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		_, err := userCredentialUsecase.Delete(ctx, &domain)

		assert.Equal(t, errError, err)
	})

	t.Run("test case 3, delete error", func(t *testing.T) {
		var (
			userCredentialRepository userCredentialMock.Repository
		)
		userCredentialUsecase := userCredential.NewUserCredentialUsecase(2, &userCredentialRepository)

		domain := userCredential.Domain{
			ID:     1,
			Type:   "code",
			Email:  "name",
			Status: true,
		}
		errError := errors.New("repo error")
		userCredentialRepository.On("FindByID", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("string")).Return(domain, nil).Once()
		userCredentialRepository.On("Delete", mock.Anything, mock.AnythingOfType("*userCredential.Domain")).Return(userCredential.Domain{}, errError).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		_, err := userCredentialUsecase.Delete(ctx, &domain)

		assert.Equal(t, errError, err)
	})
}
