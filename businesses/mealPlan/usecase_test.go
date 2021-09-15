package mealPlan_test

import (
	"context"
	"errors"
	"hungry-baby/businesses"
	mealPlan "hungry-baby/businesses/mealPlan"
	mealPlanMock "hungry-baby/mocks/mealPlan"
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
			mealPlanRepository mealPlanMock.Repository
		)
		mealPlanUsecase := mealPlan.NewMealPlanUsecase(2, &mealPlanRepository)

		domain := []mealPlan.Domain{
			{
				ID:                 1,
				UserID:             1,
				Name:               "name",
				MinAge:             0,
				MaxAge:             0,
				Interval:           0,
				SuggestionQuantity: 0,
				Unit:               "",
				Status:             true,
			},
		}
		mealPlanRepository.On("FindAll", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("int"), mock.AnythingOfType("string")).Return(domain, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		result, err := mealPlanUsecase.FindAll(ctx, "1", 1, "true")

		assert.Nil(t, err)
		assert.Equal(t, domain, result)
	})

	t.Run("test case 2, error", func(t *testing.T) {
		var (
			mealPlanRepository mealPlanMock.Repository
		)
		mealPlanUsecase := mealPlan.NewMealPlanUsecase(2, &mealPlanRepository)

		errError := errors.New("Error Repo")
		mealPlanRepository.On("FindAll", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("int"), mock.AnythingOfType("string")).Return(nil, errError).Once()
		ctx := context.WithValue(context.Background(), "userID", 1)
		_, err := mealPlanUsecase.FindAll(ctx, "1", 1, "true")

		assert.Equal(t, errError, err)
	})
}

func TestFind(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		var (
			mealPlanRepository mealPlanMock.Repository
		)
		mealPlanUsecase := mealPlan.NewMealPlanUsecase(2, &mealPlanRepository)

		domain := []mealPlan.Domain{
			{
				ID:                 1,
				UserID:             1,
				Name:               "name",
				MinAge:             0,
				MaxAge:             0,
				Interval:           0,
				SuggestionQuantity: 0,
				Unit:               "",
				Status:             true,
			},
		}
		mealPlanRepository.On("Find", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("int"),
			mock.AnythingOfType("int")).Return(domain, 1, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		result, total, err := mealPlanUsecase.Find(ctx, "1", 1, "true", 0, 0)

		assert.Nil(t, err)
		assert.Equal(t, domain, result)
		assert.Equal(t, 1, total)
	})

	t.Run("test case 2, error", func(t *testing.T) {
		var (
			mealPlanRepository mealPlanMock.Repository
		)
		mealPlanUsecase := mealPlan.NewMealPlanUsecase(2, &mealPlanRepository)

		errError := errors.New("Error Repo")
		mealPlanRepository.On("Find", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("int"),
			mock.AnythingOfType("int")).Return(nil, 0, errError).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		_, _, err := mealPlanUsecase.Find(ctx, "1", 1, "true", 0, 0)

		assert.Equal(t, errError, err)
	})
}

func TestFindByID(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		var (
			mealPlanRepository mealPlanMock.Repository
		)
		mealPlanUsecase := mealPlan.NewMealPlanUsecase(2, &mealPlanRepository)

		domain := mealPlan.Domain{
			ID:                 1,
			UserID:             1,
			Name:               "name",
			MinAge:             0,
			MaxAge:             0,
			Interval:           0,
			SuggestionQuantity: 0,
			Unit:               "",
			Status:             true,
		}
		mealPlanRepository.On("FindByID", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("string")).Return(domain, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		result, err := mealPlanUsecase.FindByID(ctx, 1, "true")

		assert.Nil(t, err)
		assert.Equal(t, domain, result)
	})

	t.Run("test case 2, invalid id", func(t *testing.T) {
		var (
			mealPlanRepository mealPlanMock.Repository
		)
		mealPlanUsecase := mealPlan.NewMealPlanUsecase(2, &mealPlanRepository)

		domain := mealPlan.Domain{
			ID:                 1,
			UserID:             1,
			Name:               "name",
			MinAge:             0,
			MaxAge:             0,
			Interval:           0,
			SuggestionQuantity: 0,
			Unit:               "",
			Status:             true,
		}
		mealPlanRepository.On("FindByID", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("string")).Return(domain, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		_, err := mealPlanUsecase.FindByID(ctx, -1, "true")

		assert.Equal(t, businesses.ErrIDNotFound, err)
	})

	t.Run("test case 3, error", func(t *testing.T) {
		var (
			mealPlanRepository mealPlanMock.Repository
		)
		mealPlanUsecase := mealPlan.NewMealPlanUsecase(2, &mealPlanRepository)

		errError := errors.New("Error Repo")
		mealPlanRepository.On("FindByID", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("string")).Return(mealPlan.Domain{}, errError).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		_, err := mealPlanUsecase.FindByID(ctx, 1, "true")

		assert.Equal(t, errError, err)
	})
}

func TestStore(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		var (
			mealPlanRepository mealPlanMock.Repository
		)
		mealPlanUsecase := mealPlan.NewMealPlanUsecase(2, &mealPlanRepository)

		domain := mealPlan.Domain{
			ID:                 1,
			UserID:             1,
			Name:               "name",
			MinAge:             0,
			MaxAge:             0,
			Interval:           0,
			SuggestionQuantity: 0,
			Unit:               "",
			Status:             true,
		}
		mealPlanRepository.On("Store", mock.Anything, mock.AnythingOfType("*mealPlan.Domain")).Return(domain, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		result, err := mealPlanUsecase.Store(ctx, &domain)

		assert.Nil(t, err)
		assert.Equal(t, domain, result)
	})

	t.Run("test case 2, store error", func(t *testing.T) {
		var (
			mealPlanRepository mealPlanMock.Repository
		)
		mealPlanUsecase := mealPlan.NewMealPlanUsecase(2, &mealPlanRepository)

		domain := mealPlan.Domain{
			ID:                 1,
			UserID:             1,
			Name:               "name",
			MinAge:             0,
			MaxAge:             0,
			Interval:           0,
			SuggestionQuantity: 0,
			Unit:               "",
			Status:             true,
		}
		errRepo := errors.New("repo error")
		mealPlanRepository.On("Store", mock.Anything, mock.AnythingOfType("*mealPlan.Domain")).Return(mealPlan.Domain{}, errRepo).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		_, err := mealPlanUsecase.Store(ctx, &domain)

		assert.Equal(t, errRepo, err)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		var (
			mealPlanRepository mealPlanMock.Repository
		)
		mealPlanUsecase := mealPlan.NewMealPlanUsecase(2, &mealPlanRepository)

		domain := mealPlan.Domain{
			ID:                 1,
			UserID:             1,
			Name:               "name",
			MinAge:             0,
			MaxAge:             0,
			Interval:           0,
			SuggestionQuantity: 0,
			Unit:               "",
			Status:             true,
		}
		mealPlanRepository.On("Update", mock.Anything, mock.AnythingOfType("*mealPlan.Domain")).Return(domain, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		result, err := mealPlanUsecase.Update(ctx, &domain)

		assert.Nil(t, err)
		assert.Equal(t, domain, result)
	})

	t.Run("test case 2, store error", func(t *testing.T) {
		var (
			mealPlanRepository mealPlanMock.Repository
		)
		mealPlanUsecase := mealPlan.NewMealPlanUsecase(2, &mealPlanRepository)

		domain := mealPlan.Domain{
			ID:                 1,
			UserID:             1,
			Name:               "name",
			MinAge:             0,
			MaxAge:             0,
			Interval:           0,
			SuggestionQuantity: 0,
			Unit:               "",
			Status:             true,
		}
		errRepo := errors.New("repo error")
		mealPlanRepository.On("Update", mock.Anything, mock.AnythingOfType("*mealPlan.Domain")).Return(mealPlan.Domain{}, errRepo).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		_, err := mealPlanUsecase.Update(ctx, &domain)

		assert.Equal(t, errRepo, err)
	})
}

func TestDelete(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		var (
			mealPlanRepository mealPlanMock.Repository
		)
		mealPlanUsecase := mealPlan.NewMealPlanUsecase(2, &mealPlanRepository)

		domain := mealPlan.Domain{
			ID:                 1,
			UserID:             1,
			Name:               "name",
			MinAge:             0,
			MaxAge:             0,
			Interval:           0,
			SuggestionQuantity: 0,
			Unit:               "",
			Status:             true,
		}
		mealPlanRepository.On("FindByID", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("string")).Return(domain, nil).Once()
		mealPlanRepository.On("Delete", mock.Anything, mock.AnythingOfType("*mealPlan.Domain")).Return(domain, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		result, err := mealPlanUsecase.Delete(ctx, &domain)

		assert.Nil(t, err)
		assert.Equal(t, domain, result)
	})

	t.Run("test case 2, find error", func(t *testing.T) {
		var (
			mealPlanRepository mealPlanMock.Repository
		)
		mealPlanUsecase := mealPlan.NewMealPlanUsecase(2, &mealPlanRepository)

		domain := mealPlan.Domain{
			ID:                 1,
			UserID:             1,
			Name:               "name",
			MinAge:             0,
			MaxAge:             0,
			Interval:           0,
			SuggestionQuantity: 0,
			Unit:               "",
			Status:             true,
		}
		errError := errors.New("record not found")
		mealPlanRepository.On("FindByID", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("string")).Return(mealPlan.Domain{}, errError).Once()
		mealPlanRepository.On("Delete", mock.Anything, mock.AnythingOfType("*mealPlan.Domain")).Return(mealPlan.Domain{}, errError).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		_, err := mealPlanUsecase.Delete(ctx, &domain)

		assert.Equal(t, errError, err)
	})

	t.Run("test case 3, delete error", func(t *testing.T) {
		var (
			mealPlanRepository mealPlanMock.Repository
		)
		mealPlanUsecase := mealPlan.NewMealPlanUsecase(2, &mealPlanRepository)

		domain := mealPlan.Domain{
			ID:                 1,
			UserID:             1,
			Name:               "name",
			MinAge:             0,
			MaxAge:             0,
			Interval:           0,
			SuggestionQuantity: 0,
			Unit:               "",
			Status:             true,
		}
		errError := errors.New("repo error")
		mealPlanRepository.On("FindByID", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("string")).Return(domain, nil).Once()
		mealPlanRepository.On("Delete", mock.Anything, mock.AnythingOfType("*mealPlan.Domain")).Return(mealPlan.Domain{}, errError).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		_, err := mealPlanUsecase.Delete(ctx, &domain)

		assert.Equal(t, errError, err)
	})
}
