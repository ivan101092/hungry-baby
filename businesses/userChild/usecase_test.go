package userChild_test

import (
	"context"
	"errors"
	"hungry-baby/businesses"
	userChild "hungry-baby/businesses/userChild"
	userChildMock "hungry-baby/mocks/userChild"
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
			userChildRepository userChildMock.Repository
		)
		userChildUsecase := userChild.NewUserChildUsecase(2, &userChildRepository)

		domain := []userChild.Domain{
			{
				ID:     1,
				UserID: 1,
				Name:   "name",
			},
		}
		userChildRepository.On("FindAll", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("int")).Return(domain, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		result, err := userChildUsecase.FindAll(ctx, "1", 1)

		assert.Nil(t, err)
		assert.Equal(t, domain, result)
	})

	t.Run("test case 2, error", func(t *testing.T) {
		var (
			userChildRepository userChildMock.Repository
		)
		userChildUsecase := userChild.NewUserChildUsecase(2, &userChildRepository)

		errError := errors.New("Error Repo")
		userChildRepository.On("FindAll", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("int")).Return(nil, errError).Once()
		ctx := context.WithValue(context.Background(), "userID", 1)
		_, err := userChildUsecase.FindAll(ctx, "1", 1)

		assert.Equal(t, errError, err)
	})
}

func TestFind(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		var (
			userChildRepository userChildMock.Repository
		)
		userChildUsecase := userChild.NewUserChildUsecase(2, &userChildRepository)

		domain := []userChild.Domain{
			{
				ID:     1,
				UserID: 1,
				Name:   "name",
			},
		}
		userChildRepository.On("Find", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("int"), mock.AnythingOfType("int"),
			mock.AnythingOfType("int")).Return(domain, 1, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		result, total, err := userChildUsecase.Find(ctx, "1", 1, 0, 0)

		assert.Nil(t, err)
		assert.Equal(t, domain, result)
		assert.Equal(t, 1, total)
	})

	t.Run("test case 2, error", func(t *testing.T) {
		var (
			userChildRepository userChildMock.Repository
		)
		userChildUsecase := userChild.NewUserChildUsecase(2, &userChildRepository)

		errError := errors.New("Error Repo")
		userChildRepository.On("Find", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("int"), mock.AnythingOfType("int"),
			mock.AnythingOfType("int")).Return(nil, 0, errError).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		_, _, err := userChildUsecase.Find(ctx, "1", 1, 0, 0)

		assert.Equal(t, errError, err)
	})
}

func TestFindByID(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		var (
			userChildRepository userChildMock.Repository
		)
		userChildUsecase := userChild.NewUserChildUsecase(2, &userChildRepository)

		domain := userChild.Domain{
			ID:     1,
			UserID: 1,
			Name:   "name",
		}
		userChildRepository.On("FindByID", mock.Anything, mock.AnythingOfType("int")).Return(domain, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		result, err := userChildUsecase.FindByID(ctx, 1)

		assert.Nil(t, err)
		assert.Equal(t, domain, result)
	})

	t.Run("test case 2, invalid id", func(t *testing.T) {
		var (
			userChildRepository userChildMock.Repository
		)
		userChildUsecase := userChild.NewUserChildUsecase(2, &userChildRepository)

		domain := userChild.Domain{
			ID:     1,
			UserID: 1,
			Name:   "name",
		}
		userChildRepository.On("FindByID", mock.Anything, mock.AnythingOfType("int")).Return(domain, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		_, err := userChildUsecase.FindByID(ctx, -1)

		assert.Equal(t, businesses.ErrIDNotFound, err)
	})

	t.Run("test case 3, error", func(t *testing.T) {
		var (
			userChildRepository userChildMock.Repository
		)
		userChildUsecase := userChild.NewUserChildUsecase(2, &userChildRepository)

		errError := errors.New("Error Repo")
		userChildRepository.On("FindByID", mock.Anything, mock.AnythingOfType("int")).Return(userChild.Domain{}, errError).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		_, err := userChildUsecase.FindByID(ctx, 1)

		assert.Equal(t, errError, err)
	})
}

func TestStore(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		var (
			userChildRepository userChildMock.Repository
		)
		userChildUsecase := userChild.NewUserChildUsecase(2, &userChildRepository)

		domain := userChild.Domain{
			ID:     1,
			UserID: 1,
			Name:   "name",
		}
		errError := errors.New("record not found")
		userChildRepository.On("FindByCode", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(userChild.Domain{}, errError).Once()
		userChildRepository.On("Store", mock.Anything, mock.AnythingOfType("*userChild.Domain")).Return(domain, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		result, err := userChildUsecase.Store(ctx, &domain)

		assert.Nil(t, err)
		assert.Equal(t, domain, result)
	})

	t.Run("test case 4, store error", func(t *testing.T) {
		var (
			userChildRepository userChildMock.Repository
		)
		userChildUsecase := userChild.NewUserChildUsecase(2, &userChildRepository)

		domain := userChild.Domain{
			ID:     1,
			UserID: 1,
			Name:   "name",
		}
		errRepo := errors.New("repo error")
		userChildRepository.On("Store", mock.Anything, mock.AnythingOfType("*userChild.Domain")).Return(userChild.Domain{}, errRepo).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		_, err := userChildUsecase.Store(ctx, &domain)

		assert.Equal(t, errRepo, err)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		var (
			userChildRepository userChildMock.Repository
		)
		userChildUsecase := userChild.NewUserChildUsecase(2, &userChildRepository)

		domain := userChild.Domain{
			ID:     1,
			UserID: 1,
			Name:   "name",
		}
		userChildRepository.On("Update", mock.Anything, mock.AnythingOfType("*userChild.Domain")).Return(domain, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		result, err := userChildUsecase.Update(ctx, &domain)

		assert.Nil(t, err)
		assert.Equal(t, domain, result)
	})

	t.Run("test case 2, store error", func(t *testing.T) {
		var (
			userChildRepository userChildMock.Repository
		)
		userChildUsecase := userChild.NewUserChildUsecase(2, &userChildRepository)

		domain := userChild.Domain{
			ID:     1,
			UserID: 1,
			Name:   "name",
		}
		errRepo := errors.New("repo error")
		userChildRepository.On("Update", mock.Anything, mock.AnythingOfType("*userChild.Domain")).Return(userChild.Domain{}, errRepo).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		_, err := userChildUsecase.Update(ctx, &domain)

		assert.Equal(t, errRepo, err)
	})
}

func TestDelete(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		var (
			userChildRepository userChildMock.Repository
		)
		userChildUsecase := userChild.NewUserChildUsecase(2, &userChildRepository)

		domain := userChild.Domain{
			ID:     1,
			UserID: 1,
			Name:   "name",
		}
		userChildRepository.On("FindByID", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("string")).Return(domain, nil).Once()
		userChildRepository.On("Delete", mock.Anything, mock.AnythingOfType("*userChild.Domain")).Return(domain, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		result, err := userChildUsecase.Delete(ctx, &domain)

		assert.Nil(t, err)
		assert.Equal(t, domain, result)
	})

	t.Run("test case 2, find error", func(t *testing.T) {
		var (
			userChildRepository userChildMock.Repository
		)
		userChildUsecase := userChild.NewUserChildUsecase(2, &userChildRepository)

		domain := userChild.Domain{
			ID:     1,
			UserID: 1,
			Name:   "name",
		}
		errError := errors.New("record not found")
		userChildRepository.On("FindByID", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("string")).Return(userChild.Domain{}, errError).Once()
		userChildRepository.On("Delete", mock.Anything, mock.AnythingOfType("*userChild.Domain")).Return(userChild.Domain{}, errError).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		_, err := userChildUsecase.Delete(ctx, &domain)

		assert.Equal(t, errError, err)
	})

	t.Run("test case 3, delete error", func(t *testing.T) {
		var (
			userChildRepository userChildMock.Repository
		)
		userChildUsecase := userChild.NewUserChildUsecase(2, &userChildRepository)

		domain := userChild.Domain{
			ID:     1,
			UserID: 1,
			Name:   "name",
		}
		errError := errors.New("repo error")
		userChildRepository.On("FindByID", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("string")).Return(domain, nil).Once()
		userChildRepository.On("Delete", mock.Anything, mock.AnythingOfType("*userChild.Domain")).Return(userChild.Domain{}, errError).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		_, err := userChildUsecase.Delete(ctx, &domain)

		assert.Equal(t, errError, err)
	})
}
