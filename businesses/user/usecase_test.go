package user_test

import (
	"context"
	"errors"
	"hungry-baby/businesses"
	user "hungry-baby/businesses/user"
	minioMock "hungry-baby/mocks/minio"
	userMock "hungry-baby/mocks/user"
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
			userRepository  userMock.Repository
			minioRepository minioMock.Repository
		)
		userUsecase := user.NewUserUsecase(2, &userRepository, &minioRepository)

		domain := []user.Domain{
			{
				ID:     1,
				Code:   "code",
				Name:   "name",
				Status: true,
			},
		}
		userRepository.On("FindAll", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(domain, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		result, err := userUsecase.FindAll(ctx, "1", "true")

		assert.Nil(t, err)
		assert.Equal(t, domain, result)
	})

	t.Run("test case 2, error", func(t *testing.T) {
		var (
			userRepository  userMock.Repository
			minioRepository minioMock.Repository
		)
		userUsecase := user.NewUserUsecase(2, &userRepository, &minioRepository)

		errError := errors.New("Error Repo")
		userRepository.On("FindAll", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(nil, errError).Once()
		ctx := context.WithValue(context.Background(), "userID", 1)
		_, err := userUsecase.FindAll(ctx, "1", "true")

		assert.Equal(t, errError, err)
	})
}

func TestFind(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		var (
			userRepository  userMock.Repository
			minioRepository minioMock.Repository
		)
		userUsecase := user.NewUserUsecase(2, &userRepository, &minioRepository)

		domain := []user.Domain{
			{
				ID:     1,
				Code:   "code",
				Name:   "name",
				Status: true,
			},
		}
		userRepository.On("Find", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("string"), mock.AnythingOfType("int"),
			mock.AnythingOfType("int")).Return(domain, 1, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		result, total, err := userUsecase.Find(ctx, "1", "true", 0, 0)

		assert.Nil(t, err)
		assert.Equal(t, domain, result)
		assert.Equal(t, 1, total)
	})

	t.Run("test case 2, error", func(t *testing.T) {
		var (
			userRepository  userMock.Repository
			minioRepository minioMock.Repository
		)
		userUsecase := user.NewUserUsecase(2, &userRepository, &minioRepository)

		errError := errors.New("Error Repo")
		userRepository.On("Find", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("string"), mock.AnythingOfType("int"),
			mock.AnythingOfType("int")).Return(nil, 0, errError).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		_, _, err := userUsecase.Find(ctx, "1", "true", 0, 0)

		assert.Equal(t, errError, err)
	})
}

func TestFindByID(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		var (
			userRepository  userMock.Repository
			minioRepository minioMock.Repository
		)
		userUsecase := user.NewUserUsecase(2, &userRepository, &minioRepository)

		domain := user.Domain{
			ID:              1,
			Code:            "code",
			Name:            "name",
			ProfileImageURL: "url",
			Status:          true,
		}
		userRepository.On("FindByID", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("string")).Return(domain, nil).Once()
		minioRepository.On("GetFile", mock.AnythingOfType("string")).Return("url", nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		result, err := userUsecase.FindByID(ctx, 1, "true")

		assert.Nil(t, err)
		assert.Equal(t, domain, result)
	})

	t.Run("test case 2, invalid id", func(t *testing.T) {
		var (
			userRepository  userMock.Repository
			minioRepository minioMock.Repository
		)
		userUsecase := user.NewUserUsecase(2, &userRepository, &minioRepository)

		domain := user.Domain{
			ID:     1,
			Code:   "code",
			Name:   "name",
			Status: true,
		}
		userRepository.On("FindByID", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("string")).Return(domain, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		_, err := userUsecase.FindByID(ctx, -1, "true")

		assert.Equal(t, businesses.ErrIDNotFound, err)
	})

	t.Run("test case 3, error", func(t *testing.T) {
		var (
			userRepository  userMock.Repository
			minioRepository minioMock.Repository
		)
		userUsecase := user.NewUserUsecase(2, &userRepository, &minioRepository)

		errError := errors.New("Error Repo")
		userRepository.On("FindByID", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("string")).Return(user.Domain{}, errError).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		_, err := userUsecase.FindByID(ctx, 1, "true")

		assert.Equal(t, errError, err)
	})

	t.Run("test case 4, minio error", func(t *testing.T) {
		var (
			userRepository  userMock.Repository
			minioRepository minioMock.Repository
		)
		userUsecase := user.NewUserUsecase(2, &userRepository, &minioRepository)

		domain := user.Domain{
			ID:              1,
			Code:            "code",
			Name:            "name",
			ProfileImageURL: "url",
			Status:          true,
		}
		errError := errors.New("Error Minio")
		userRepository.On("FindByID", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("string")).Return(domain, nil).Once()
		minioRepository.On("GetFile", mock.AnythingOfType("string")).Return("", errError).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		_, err := userUsecase.FindByID(ctx, 1, "true")

		assert.Equal(t, errError, err)
	})
}

func TestFindByCode(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		var (
			userRepository  userMock.Repository
			minioRepository minioMock.Repository
		)
		userUsecase := user.NewUserUsecase(2, &userRepository, &minioRepository)

		domain := user.Domain{
			ID:     1,
			Code:   "code",
			Name:   "name",
			Status: true,
		}
		userRepository.On("FindByCode", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(domain, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		result, err := userUsecase.FindByCode(ctx, "code", "true")

		assert.Nil(t, err)
		assert.Equal(t, domain, result)
	})

	t.Run("test case 2, error", func(t *testing.T) {
		var (
			userRepository  userMock.Repository
			minioRepository minioMock.Repository
		)
		userUsecase := user.NewUserUsecase(2, &userRepository, &minioRepository)

		errError := errors.New("Error Repo")
		userRepository.On("FindByCode", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(user.Domain{}, errError).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		_, err := userUsecase.FindByCode(ctx, "code", "true")

		assert.Equal(t, errError, err)
	})
}

func TestFindByEmail(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		var (
			userRepository  userMock.Repository
			minioRepository minioMock.Repository
		)
		userUsecase := user.NewUserUsecase(2, &userRepository, &minioRepository)

		domain := user.Domain{
			ID:     1,
			Code:   "code",
			Name:   "name",
			Status: true,
		}
		userRepository.On("FindByEmail", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(domain, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		result, err := userUsecase.FindByEmail(ctx, "code", "true")

		assert.Nil(t, err)
		assert.Equal(t, domain, result)
	})

	t.Run("test case 2, error", func(t *testing.T) {
		var (
			userRepository  userMock.Repository
			minioRepository minioMock.Repository
		)
		userUsecase := user.NewUserUsecase(2, &userRepository, &minioRepository)

		errError := errors.New("Error Repo")
		userRepository.On("FindByEmail", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(user.Domain{}, errError).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		_, err := userUsecase.FindByEmail(ctx, "code", "true")

		assert.Equal(t, errError, err)
	})
}

func TestStore(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		var (
			userRepository  userMock.Repository
			minioRepository minioMock.Repository
		)
		userUsecase := user.NewUserUsecase(2, &userRepository, &minioRepository)

		domain := user.Domain{
			ID:     1,
			Code:   "code",
			Name:   "name",
			Status: true,
		}
		errError := errors.New("record not found")
		userRepository.On("FindByEmail", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(user.Domain{}, errError).Once()
		userRepository.On("Store", mock.Anything, mock.AnythingOfType("*user.Domain")).Return(domain, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		result, err := userUsecase.Store(ctx, &domain)

		assert.Nil(t, err)
		assert.Equal(t, domain, result)
	})

	t.Run("test case 2, find error", func(t *testing.T) {
		var (
			userRepository  userMock.Repository
			minioRepository minioMock.Repository
		)
		userUsecase := user.NewUserUsecase(2, &userRepository, &minioRepository)

		domain := user.Domain{
			ID:     1,
			Code:   "code",
			Name:   "name",
			Status: true,
		}
		errError := errors.New("internal server")
		userRepository.On("FindByEmail", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(user.Domain{}, errError).Once()
		userRepository.On("Store", mock.Anything, mock.AnythingOfType("*user.Domain")).Return(domain, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		_, err := userUsecase.Store(ctx, &domain)

		assert.Equal(t, errError, err)
	})

	t.Run("test case 3, find exist", func(t *testing.T) {
		var (
			userRepository  userMock.Repository
			minioRepository minioMock.Repository
		)
		userUsecase := user.NewUserUsecase(2, &userRepository, &minioRepository)

		domain := user.Domain{
			ID:     1,
			Code:   "code",
			Name:   "name",
			Status: true,
		}
		userRepository.On("FindByEmail", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(domain, nil).Once()
		userRepository.On("Store", mock.Anything, mock.AnythingOfType("*user.Domain")).Return(domain, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		_, err := userUsecase.Store(ctx, &domain)

		assert.Equal(t, businesses.ErrDuplicateData, err)
	})

	t.Run("test case 4, store error", func(t *testing.T) {
		var (
			userRepository  userMock.Repository
			minioRepository minioMock.Repository
		)
		userUsecase := user.NewUserUsecase(2, &userRepository, &minioRepository)

		domain := user.Domain{
			ID:     1,
			Code:   "code",
			Name:   "name",
			Status: true,
		}
		errError := errors.New("record not found")
		errRepo := errors.New("repo error")
		userRepository.On("FindByEmail", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(user.Domain{}, errError).Once()
		userRepository.On("Store", mock.Anything, mock.AnythingOfType("*user.Domain")).Return(user.Domain{}, errRepo).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		_, err := userUsecase.Store(ctx, &domain)

		assert.Equal(t, errRepo, err)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		var (
			userRepository  userMock.Repository
			minioRepository minioMock.Repository
		)
		userUsecase := user.NewUserUsecase(2, &userRepository, &minioRepository)

		domain := user.Domain{
			ID:     1,
			Code:   "code",
			Name:   "name",
			Status: true,
		}
		errError := errors.New("record not found")
		userRepository.On("FindByID", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("string")).Return(domain, nil).Once()
		userRepository.On("FindByEmail", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(user.Domain{}, errError).Once()
		userRepository.On("Update", mock.Anything, mock.AnythingOfType("*user.Domain")).Return(domain, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		result, err := userUsecase.Update(ctx, &domain)

		assert.Nil(t, err)
		assert.Equal(t, domain, result)
	})

	t.Run("test case 2, find error", func(t *testing.T) {
		var (
			userRepository  userMock.Repository
			minioRepository minioMock.Repository
		)
		userUsecase := user.NewUserUsecase(2, &userRepository, &minioRepository)

		domain := user.Domain{
			ID:     1,
			Code:   "code",
			Name:   "name",
			Status: true,
		}
		errError := errors.New("internal server")
		userRepository.On("FindByID", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("string")).Return(domain, nil).Once()
		userRepository.On("FindByEmail", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(user.Domain{}, errError).Once()
		userRepository.On("Update", mock.Anything, mock.AnythingOfType("*user.Domain")).Return(domain, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		_, err := userUsecase.Update(ctx, &domain)

		assert.Equal(t, errError, err)
	})

	t.Run("test case 3, find exist", func(t *testing.T) {
		var (
			userRepository  userMock.Repository
			minioRepository minioMock.Repository
		)
		userUsecase := user.NewUserUsecase(2, &userRepository, &minioRepository)

		domain := user.Domain{
			ID:     1,
			Code:   "code",
			Name:   "name",
			Status: true,
		}
		userRepository.On("FindByID", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("string")).Return(domain, nil).Once()
		userRepository.On("FindByEmail", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(domain, nil).Once()
		userRepository.On("Update", mock.Anything, mock.AnythingOfType("*user.Domain")).Return(domain, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		domain.ID = 2
		_, err := userUsecase.Update(ctx, &domain)

		assert.Equal(t, businesses.ErrDuplicateData, err)
	})

	t.Run("test case 4, store error", func(t *testing.T) {
		var (
			userRepository  userMock.Repository
			minioRepository minioMock.Repository
		)
		userUsecase := user.NewUserUsecase(2, &userRepository, &minioRepository)

		domain := user.Domain{
			ID:     1,
			Code:   "code",
			Name:   "name",
			Status: true,
		}
		errError := errors.New("record not found")
		errRepo := errors.New("repo error")
		userRepository.On("FindByID", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("string")).Return(domain, nil).Once()
		userRepository.On("FindByEmail", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(user.Domain{}, errError).Once()
		userRepository.On("Update", mock.Anything, mock.AnythingOfType("*user.Domain")).Return(user.Domain{}, errRepo).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		_, err := userUsecase.Update(ctx, &domain)

		assert.Equal(t, errRepo, err)
	})

	t.Run("test case 5, not found error", func(t *testing.T) {
		var (
			userRepository  userMock.Repository
			minioRepository minioMock.Repository
		)
		userUsecase := user.NewUserUsecase(2, &userRepository, &minioRepository)

		domain := user.Domain{
			ID:     1,
			Code:   "code",
			Name:   "name",
			Status: true,
		}
		errError := errors.New("record not found")
		userRepository.On("FindByID", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("string")).Return(domain, errError).Once()
		userRepository.On("FindByEmail", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(user.Domain{}, errError).Once()
		userRepository.On("Update", mock.Anything, mock.AnythingOfType("*user.Domain")).Return(user.Domain{}, errError).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		_, err := userUsecase.Update(ctx, &domain)

		assert.Equal(t, errError, err)
	})
}

func TestDelete(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		var (
			userRepository  userMock.Repository
			minioRepository minioMock.Repository
		)
		userUsecase := user.NewUserUsecase(2, &userRepository, &minioRepository)

		domain := user.Domain{
			ID:     1,
			Code:   "code",
			Name:   "name",
			Status: true,
		}
		userRepository.On("FindByID", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("string")).Return(domain, nil).Once()
		userRepository.On("Delete", mock.Anything, mock.AnythingOfType("*user.Domain")).Return(domain, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		result, err := userUsecase.Delete(ctx, &domain)

		assert.Nil(t, err)
		assert.Equal(t, domain, result)
	})

	t.Run("test case 2, find error", func(t *testing.T) {
		var (
			userRepository  userMock.Repository
			minioRepository minioMock.Repository
		)
		userUsecase := user.NewUserUsecase(2, &userRepository, &minioRepository)

		domain := user.Domain{
			ID:     1,
			Code:   "code",
			Name:   "name",
			Status: true,
		}
		errError := errors.New("record not found")
		userRepository.On("FindByID", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("string")).Return(user.Domain{}, errError).Once()
		userRepository.On("Delete", mock.Anything, mock.AnythingOfType("*user.Domain")).Return(user.Domain{}, errError).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		_, err := userUsecase.Delete(ctx, &domain)

		assert.Equal(t, errError, err)
	})

	t.Run("test case 3, delete error", func(t *testing.T) {
		var (
			userRepository  userMock.Repository
			minioRepository minioMock.Repository
		)
		userUsecase := user.NewUserUsecase(2, &userRepository, &minioRepository)

		domain := user.Domain{
			ID:     1,
			Code:   "code",
			Name:   "name",
			Status: true,
		}
		errError := errors.New("repo error")
		userRepository.On("FindByID", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("string")).Return(domain, nil).Once()
		userRepository.On("Delete", mock.Anything, mock.AnythingOfType("*user.Domain")).Return(user.Domain{}, errError).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		_, err := userUsecase.Delete(ctx, &domain)

		assert.Equal(t, errError, err)
	})
}
