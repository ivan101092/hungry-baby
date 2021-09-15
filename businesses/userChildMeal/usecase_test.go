package userChildMeal_test

import (
	"context"
	"errors"
	"hungry-baby/businesses"
	"hungry-baby/businesses/calendar"
	mealPlan "hungry-baby/businesses/mealPlan"
	user "hungry-baby/businesses/user"
	userChild "hungry-baby/businesses/userChild"
	userChildMeal "hungry-baby/businesses/userChildMeal"
	calendarMock "hungry-baby/mocks/calendar"
	mealPlanMock "hungry-baby/mocks/mealPlan"
	userMock "hungry-baby/mocks/user"
	userChildMock "hungry-baby/mocks/userChild"
	userChildMealMock "hungry-baby/mocks/userChildMeal"
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
			userChildMealRepository userChildMealMock.Repository
			userChildUsecase        userChildMock.Usecase
			mealPlanUsecase         mealPlanMock.Usecase
			calendarUsecase         calendarMock.Usecase
			userUsecase             userMock.Usecase
		)
		userChildMealUsecase := userChildMeal.NewUserChildMealUsecase(2, &userChildMealRepository,
			&userChildUsecase, &mealPlanUsecase, &calendarUsecase, &userUsecase)

		domain := []userChildMeal.Domain{
			{
				ID:     1,
				Name:   "code",
				Unit:   "name",
				Status: "done",
			},
		}
		userChildMealRepository.On("FindAll", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("int")).Return(domain, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		result, err := userChildMealUsecase.FindAll(ctx, "1", 1)

		assert.Nil(t, err)
		assert.Equal(t, domain, result)
	})

	t.Run("test case 2, error", func(t *testing.T) {
		var (
			userChildMealRepository userChildMealMock.Repository
			userChildUsecase        userChildMock.Usecase
			mealPlanUsecase         mealPlanMock.Usecase
			calendarUsecase         calendarMock.Usecase
			userUsecase             userMock.Usecase
		)
		userChildMealUsecase := userChildMeal.NewUserChildMealUsecase(2, &userChildMealRepository,
			&userChildUsecase, &mealPlanUsecase, &calendarUsecase, &userUsecase)

		errError := errors.New("Error Repo")
		userChildMealRepository.On("FindAll", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("int")).Return(nil, errError).Once()
		ctx := context.WithValue(context.Background(), "userID", 1)
		_, err := userChildMealUsecase.FindAll(ctx, "1", 1)

		assert.Equal(t, errError, err)
	})
}

func TestFind(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		var (
			userChildMealRepository userChildMealMock.Repository
			userChildUsecase        userChildMock.Usecase
			mealPlanUsecase         mealPlanMock.Usecase
			calendarUsecase         calendarMock.Usecase
			userUsecase             userMock.Usecase
		)
		userChildMealUsecase := userChildMeal.NewUserChildMealUsecase(2, &userChildMealRepository,
			&userChildUsecase, &mealPlanUsecase, &calendarUsecase, &userUsecase)

		domain := []userChildMeal.Domain{
			{
				ID:     1,
				Name:   "code",
				Unit:   "name",
				Status: "done",
			},
		}
		userChildMealRepository.On("Find", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("int"), mock.AnythingOfType("int"),
			mock.AnythingOfType("int")).Return(domain, 1, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		result, total, err := userChildMealUsecase.Find(ctx, "1", 1, 0, 0)

		assert.Nil(t, err)
		assert.Equal(t, domain, result)
		assert.Equal(t, 1, total)
	})

	t.Run("test case 2, error", func(t *testing.T) {
		var (
			userChildMealRepository userChildMealMock.Repository
			userChildUsecase        userChildMock.Usecase
			mealPlanUsecase         mealPlanMock.Usecase
			calendarUsecase         calendarMock.Usecase
			userUsecase             userMock.Usecase
		)
		userChildMealUsecase := userChildMeal.NewUserChildMealUsecase(2, &userChildMealRepository,
			&userChildUsecase, &mealPlanUsecase, &calendarUsecase, &userUsecase)

		errError := errors.New("Error Repo")
		userChildMealRepository.On("Find", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("int"), mock.AnythingOfType("int"),
			mock.AnythingOfType("int")).Return(nil, 0, errError).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		_, _, err := userChildMealUsecase.Find(ctx, "1", 1, 0, 0)

		assert.Equal(t, errError, err)
	})
}

func TestFindByID(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		var (
			userChildMealRepository userChildMealMock.Repository
			userChildUsecase        userChildMock.Usecase
			mealPlanUsecase         mealPlanMock.Usecase
			calendarUsecase         calendarMock.Usecase
			userUsecase             userMock.Usecase
		)
		userChildMealUsecase := userChildMeal.NewUserChildMealUsecase(2, &userChildMealRepository,
			&userChildUsecase, &mealPlanUsecase, &calendarUsecase, &userUsecase)

		domain := userChildMeal.Domain{
			ID:     1,
			UserID: 1,
			Name:   "code",
			Unit:   "name",
			Status: "done",
		}
		userChildMealRepository.On("FindByID", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("string")).Return(domain, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		result, err := userChildMealUsecase.FindByID(ctx, 1)

		assert.Nil(t, err)
		assert.Equal(t, domain, result)
	})

	t.Run("test case 2, invalid id", func(t *testing.T) {
		var (
			userChildMealRepository userChildMealMock.Repository
			userChildUsecase        userChildMock.Usecase
			mealPlanUsecase         mealPlanMock.Usecase
			calendarUsecase         calendarMock.Usecase
			userUsecase             userMock.Usecase
		)
		userChildMealUsecase := userChildMeal.NewUserChildMealUsecase(2, &userChildMealRepository,
			&userChildUsecase, &mealPlanUsecase, &calendarUsecase, &userUsecase)

		domain := userChildMeal.Domain{
			ID:     1,
			UserID: 1,
			Name:   "code",
			Unit:   "name",
			Status: "done",
		}
		userChildMealRepository.On("FindByID", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("string")).Return(domain, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		_, err := userChildMealUsecase.FindByID(ctx, -1)

		assert.Equal(t, businesses.ErrIDNotFound, err)
	})

	t.Run("test case 3, error", func(t *testing.T) {
		var (
			userChildMealRepository userChildMealMock.Repository
			userChildUsecase        userChildMock.Usecase
			mealPlanUsecase         mealPlanMock.Usecase
			calendarUsecase         calendarMock.Usecase
			userUsecase             userMock.Usecase
		)
		userChildMealUsecase := userChildMeal.NewUserChildMealUsecase(2, &userChildMealRepository,
			&userChildUsecase, &mealPlanUsecase, &calendarUsecase, &userUsecase)

		errError := errors.New("Error Repo")
		userChildMealRepository.On("FindByID", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("string")).Return(userChildMeal.Domain{}, errError).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		_, err := userChildMealUsecase.FindByID(ctx, 1)

		assert.Equal(t, errError, err)
	})
}

func TestFindByChildMeal(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		var (
			userChildMealRepository userChildMealMock.Repository
			userChildUsecase        userChildMock.Usecase
			mealPlanUsecase         mealPlanMock.Usecase
			calendarUsecase         calendarMock.Usecase
			userUsecase             userMock.Usecase
		)
		userChildMealUsecase := userChildMeal.NewUserChildMealUsecase(2, &userChildMealRepository,
			&userChildUsecase, &mealPlanUsecase, &calendarUsecase, &userUsecase)

		domain := userChildMeal.Domain{
			ID:     1,
			UserID: 1,
			Name:   "code",
			Unit:   "name",
			Status: "done",
		}
		userChildMealRepository.On("FindByChildMeal", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("int")).Return(domain, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		result, err := userChildMealUsecase.FindByChildMeal(ctx, 1, 1)

		assert.Nil(t, err)
		assert.Equal(t, domain, result)
	})

	t.Run("test case 2, error", func(t *testing.T) {
		var (
			userChildMealRepository userChildMealMock.Repository
			userChildUsecase        userChildMock.Usecase
			mealPlanUsecase         mealPlanMock.Usecase
			calendarUsecase         calendarMock.Usecase
			userUsecase             userMock.Usecase
		)
		userChildMealUsecase := userChildMeal.NewUserChildMealUsecase(2, &userChildMealRepository,
			&userChildUsecase, &mealPlanUsecase, &calendarUsecase, &userUsecase)

		errError := errors.New("Error Repo")
		userChildMealRepository.On("FindByChildMeal", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("int")).Return(userChildMeal.Domain{}, errError).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		_, err := userChildMealUsecase.FindByChildMeal(ctx, 1, 1)

		assert.Equal(t, errError, err)
	})
}

func TestStore(t *testing.T) {
	t.Run("test case 1, pending", func(t *testing.T) {
		var (
			userChildMealRepository userChildMealMock.Repository
			userChildUsecase        userChildMock.Usecase
			mealPlanUsecase         mealPlanMock.Usecase
			calendarUsecase         calendarMock.Usecase
			userUsecase             userMock.Usecase
		)
		userChildMealUsecase := userChildMeal.NewUserChildMealUsecase(2, &userChildMealRepository,
			&userChildUsecase, &mealPlanUsecase, &calendarUsecase, &userUsecase)

		domain := userChildMeal.Domain{
			ID:          1,
			UserID:      1,
			Quantity:    1,
			Name:        "name",
			Unit:        "unit",
			ScheduledAt: "2021-10-10T00:00:00Z",
			Status:      "pending",
		}

		userChildDomain := userChild.Domain{
			ID:     1,
			UserID: 1,
			Name:   "name",
		}
		userChildUsecase.On("FindByID", mock.Anything, mock.AnythingOfType("int")).Return(userChildDomain, nil).Once()
		mealPlanDomain := mealPlan.Domain{
			ID:     1,
			UserID: 1,
			Name:   "name",
			Unit:   "unit",
		}
		mealPlanUsecase.On("FindByID", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("string")).Return(mealPlanDomain, nil).Once()
		calendarDomain := calendar.Domain{
			Title:   mealPlanDomain.Name,
			StartAt: domain.ScheduledAt,
			EndAt:   domain.ScheduledAt,
		}
		calendarUsecase.On("Store", mock.Anything, mock.AnythingOfType("*calendar.Domain")).Return(calendarDomain, nil).Once()
		userChildMealRepository.On("Store", mock.Anything, mock.AnythingOfType("*userChildMeal.Domain")).Return(domain, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		result, err := userChildMealUsecase.Store(ctx, &domain)

		assert.Nil(t, err)
		assert.Equal(t, domain, result)
	})

	t.Run("test case 2, done", func(t *testing.T) {
		var (
			userChildMealRepository userChildMealMock.Repository
			userChildUsecase        userChildMock.Usecase
			mealPlanUsecase         mealPlanMock.Usecase
			calendarUsecase         calendarMock.Usecase
			userUsecase             userMock.Usecase
		)
		userChildMealUsecase := userChildMeal.NewUserChildMealUsecase(2, &userChildMealRepository,
			&userChildUsecase, &mealPlanUsecase, &calendarUsecase, &userUsecase)

		domain := userChildMeal.Domain{
			ID:          1,
			UserID:      1,
			Quantity:    1,
			Name:        "name",
			Unit:        "unit",
			ScheduledAt: "2021-09-10T00:00:00Z",
			FinishAt:    "2021-09-11T00:00:00Z",
			Status:      "done",
		}

		userChildDomain := userChild.Domain{
			ID:     1,
			UserID: 1,
			Name:   "name",
		}
		userChildUsecase.On("FindByID", mock.Anything, mock.AnythingOfType("int")).Return(userChildDomain, nil).Once()
		mealPlanDomain := mealPlan.Domain{
			ID:     1,
			UserID: 1,
			Name:   "name",
			Unit:   "unit",
		}
		mealPlanUsecase.On("FindByID", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("string")).Return(mealPlanDomain, nil).Once()
		calendarDomain := calendar.Domain{
			Title:   mealPlanDomain.Name,
			StartAt: domain.ScheduledAt,
			EndAt:   domain.ScheduledAt,
		}
		calendarUsecase.On("Store", mock.Anything, mock.AnythingOfType("*calendar.Domain")).Return(calendarDomain, nil).Once()
		userChildMealRepository.On("Store", mock.Anything, mock.AnythingOfType("*userChildMeal.Domain")).Return(domain, nil).Once()
		userChildMealRepository.On("FindNextPending", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("int")).Return(userChildMeal.Domain{}, errors.New("data not found")).Once()
		userDomain := user.Domain{
			ID:       1,
			Settings: user.DomainSettings{AutoNotification: true},
		}
		userUsecase.On("FindByID", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("string")).Return(userDomain, nil).Once()
		userChildMealRepository.On("Store", mock.Anything, mock.AnythingOfType("*userChildMeal.Domain")).Return(userChildMeal.Domain{}, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		result, err := userChildMealUsecase.Store(ctx, &domain)

		assert.Nil(t, err)
		assert.Equal(t, domain, result)
	})

	t.Run("test case 3, done invalid finish at", func(t *testing.T) {
		var (
			userChildMealRepository userChildMealMock.Repository
			userChildUsecase        userChildMock.Usecase
			mealPlanUsecase         mealPlanMock.Usecase
			calendarUsecase         calendarMock.Usecase
			userUsecase             userMock.Usecase
		)
		userChildMealUsecase := userChildMeal.NewUserChildMealUsecase(2, &userChildMealRepository,
			&userChildUsecase, &mealPlanUsecase, &calendarUsecase, &userUsecase)

		domain := userChildMeal.Domain{
			ID:          1,
			UserID:      1,
			Quantity:    1,
			Name:        "name",
			Unit:        "unit",
			ScheduledAt: "2021-09-10T00:00:00Z",
			FinishAt:    "",
			Status:      "done",
		}
		errError := errors.New("Invalid finish at")
		ctx := context.WithValue(context.Background(), "userID", 1)
		_, err := userChildMealUsecase.Store(ctx, &domain)

		assert.Equal(t, errError, err)
	})
}
func TestUpdate(t *testing.T) {
	t.Run("test case 1, pending", func(t *testing.T) {
		var (
			userChildMealRepository userChildMealMock.Repository
			userChildUsecase        userChildMock.Usecase
			mealPlanUsecase         mealPlanMock.Usecase
			calendarUsecase         calendarMock.Usecase
			userUsecase             userMock.Usecase
		)
		userChildMealUsecase := userChildMeal.NewUserChildMealUsecase(2, &userChildMealRepository,
			&userChildUsecase, &mealPlanUsecase, &calendarUsecase, &userUsecase)

		domain := userChildMeal.Domain{
			ID:          1,
			UserID:      1,
			Quantity:    1,
			Name:        "name",
			Unit:        "unit",
			ScheduledAt: "2021-10-10T00:00:00Z",
			CalendarID:  "calendarid",
			Status:      "pending",
		}
		userChildMealRepository.On("FindByID", mock.Anything, mock.AnythingOfType("int")).Return(domain, nil).Once()
		userChildDomain := userChild.Domain{
			ID:     1,
			UserID: 1,
			Name:   "name",
		}
		userChildUsecase.On("FindByID", mock.Anything, mock.AnythingOfType("int")).Return(userChildDomain, nil).Once()
		mealPlanDomain := mealPlan.Domain{
			ID:     1,
			UserID: 1,
			Name:   "name",
			Unit:   "unit",
		}
		mealPlanUsecase.On("FindByID", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("string")).Return(mealPlanDomain, nil).Once()
		calendarUsecase.On("Delete", mock.Anything, mock.AnythingOfType("string")).Return(nil).Once()
		calendarDomain := calendar.Domain{
			Title:   mealPlanDomain.Name,
			StartAt: domain.ScheduledAt,
			EndAt:   domain.ScheduledAt,
		}
		calendarUsecase.On("Store", mock.Anything, mock.AnythingOfType("*calendar.Domain")).Return(calendarDomain, nil).Once()
		userChildMealRepository.On("Update", mock.Anything, mock.AnythingOfType("*userChildMeal.Domain")).Return(domain, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		domain.ScheduledAt = "2021-10-11T00:00:00Z"
		result, err := userChildMealUsecase.Update(ctx, &domain)
		domain.ScheduledAt = "2021-10-10T00:00:00Z"

		assert.Nil(t, err)
		assert.Equal(t, domain, result)
	})

	t.Run("test case 2, done", func(t *testing.T) {
		var (
			userChildMealRepository userChildMealMock.Repository
			userChildUsecase        userChildMock.Usecase
			mealPlanUsecase         mealPlanMock.Usecase
			calendarUsecase         calendarMock.Usecase
			userUsecase             userMock.Usecase
		)
		userChildMealUsecase := userChildMeal.NewUserChildMealUsecase(2, &userChildMealRepository,
			&userChildUsecase, &mealPlanUsecase, &calendarUsecase, &userUsecase)

		domain := userChildMeal.Domain{
			ID:          1,
			UserID:      1,
			Quantity:    1,
			Name:        "name",
			Unit:        "unit",
			ScheduledAt: "2021-09-10T00:00:00Z",
			FinishAt:    "2021-09-11T00:00:00Z",
			CalendarID:  "calendarid",
			Status:      "done",
		}
		userChildMealRepository.On("FindByID", mock.Anything, mock.AnythingOfType("int")).Return(domain, nil).Once()

		userChildDomain := userChild.Domain{
			ID:     1,
			UserID: 1,
			Name:   "name",
		}
		userChildUsecase.On("FindByID", mock.Anything, mock.AnythingOfType("int")).Return(userChildDomain, nil).Once()
		mealPlanDomain := mealPlan.Domain{
			ID:     1,
			UserID: 1,
			Name:   "name",
			Unit:   "unit",
		}
		mealPlanUsecase.On("FindByID", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("string")).Return(mealPlanDomain, nil).Once()
		calendarDomain := calendar.Domain{
			Title:   mealPlanDomain.Name,
			StartAt: domain.ScheduledAt,
			EndAt:   domain.ScheduledAt,
		}
		calendarUsecase.On("Store", mock.Anything, mock.AnythingOfType("*calendar.Domain")).Return(calendarDomain, nil).Once()
		calendarUsecase.On("Delete", mock.Anything, mock.AnythingOfType("string")).Return(nil).Once()
		userChildMealRepository.On("Update", mock.Anything, mock.AnythingOfType("*userChildMeal.Domain")).Return(domain, nil).Once()
		userChildMealRepository.On("FindNextPending", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("int")).Return(userChildMeal.Domain{}, errors.New("data not found")).Once()
		userDomain := user.Domain{
			ID:       1,
			Settings: user.DomainSettings{AutoNotification: true},
		}
		userUsecase.On("FindByID", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("string")).Return(userDomain, nil).Once()
		userChildMealRepository.On("Store", mock.Anything, mock.AnythingOfType("*userChildMeal.Domain")).Return(userChildMeal.Domain{}, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		result, err := userChildMealUsecase.Update(ctx, &domain)

		assert.Nil(t, err)
		assert.Equal(t, domain, result)
	})

	t.Run("test case 3, done invalid finish at", func(t *testing.T) {
		var (
			userChildMealRepository userChildMealMock.Repository
			userChildUsecase        userChildMock.Usecase
			mealPlanUsecase         mealPlanMock.Usecase
			calendarUsecase         calendarMock.Usecase
			userUsecase             userMock.Usecase
		)
		userChildMealUsecase := userChildMeal.NewUserChildMealUsecase(2, &userChildMealRepository,
			&userChildUsecase, &mealPlanUsecase, &calendarUsecase, &userUsecase)

		domain := userChildMeal.Domain{
			ID:          1,
			UserID:      1,
			Quantity:    1,
			Name:        "name",
			Unit:        "unit",
			ScheduledAt: "2021-09-10T00:00:00Z",
			FinishAt:    "",
			Status:      "done",
		}
		errError := errors.New("Invalid finish at")
		ctx := context.WithValue(context.Background(), "userID", 1)
		_, err := userChildMealUsecase.Update(ctx, &domain)

		assert.Equal(t, errError, err)
	})
}

func TestDelete(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		var (
			userChildMealRepository userChildMealMock.Repository
			userChildUsecase        userChildMock.Usecase
			mealPlanUsecase         mealPlanMock.Usecase
			calendarUsecase         calendarMock.Usecase
			userUsecase             userMock.Usecase
		)
		userChildMealUsecase := userChildMeal.NewUserChildMealUsecase(2, &userChildMealRepository,
			&userChildUsecase, &mealPlanUsecase, &calendarUsecase, &userUsecase)

		domain := userChildMeal.Domain{
			ID:     1,
			UserID: 1,
			Name:   "code",
			Unit:   "name",
			Status: "done",
		}
		userChildMealRepository.On("FindByID", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("string")).Return(domain, nil).Once()
		userChildDomain := userChild.Domain{
			ID:     1,
			UserID: 1,
			Name:   "name",
		}
		userChildUsecase.On("FindByID", mock.Anything, mock.AnythingOfType("int")).Return(userChildDomain, nil).Once()
		userChildMealRepository.On("Delete", mock.Anything, mock.AnythingOfType("*userChildMeal.Domain")).Return(domain, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		result, err := userChildMealUsecase.Delete(ctx, &domain)

		assert.Nil(t, err)
		assert.Equal(t, domain, result)
	})

	t.Run("test case 2, invalid access", func(t *testing.T) {
		var (
			userChildMealRepository userChildMealMock.Repository
			userChildUsecase        userChildMock.Usecase
			mealPlanUsecase         mealPlanMock.Usecase
			calendarUsecase         calendarMock.Usecase
			userUsecase             userMock.Usecase
		)
		userChildMealUsecase := userChildMeal.NewUserChildMealUsecase(2, &userChildMealRepository,
			&userChildUsecase, &mealPlanUsecase, &calendarUsecase, &userUsecase)

		errError := errors.New("Invalid access")
		domain := userChildMeal.Domain{
			ID:     1,
			UserID: 1,
			Name:   "code",
			Unit:   "name",
			Status: "done",
		}
		userChildMealRepository.On("FindByID", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("string")).Return(domain, nil).Once()
		userChildDomain := userChild.Domain{
			ID:     1,
			UserID: 2,
			Name:   "name",
		}
		userChildUsecase.On("FindByID", mock.Anything, mock.AnythingOfType("int")).Return(userChildDomain, nil).Once()
		userChildMealRepository.On("Delete", mock.Anything, mock.AnythingOfType("*userChildMeal.Domain")).Return(domain, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		_, err := userChildMealUsecase.Delete(ctx, &domain)

		assert.Equal(t, errError, err)
	})

	t.Run("test case 3, find error", func(t *testing.T) {
		var (
			userChildMealRepository userChildMealMock.Repository
			userChildUsecase        userChildMock.Usecase
			mealPlanUsecase         mealPlanMock.Usecase
			calendarUsecase         calendarMock.Usecase
			userUsecase             userMock.Usecase
		)
		userChildMealUsecase := userChildMeal.NewUserChildMealUsecase(2, &userChildMealRepository,
			&userChildUsecase, &mealPlanUsecase, &calendarUsecase, &userUsecase)

		domain := userChildMeal.Domain{
			ID:     1,
			UserID: 1,
			Name:   "code",
			Unit:   "name",
			Status: "done",
		}
		errError := errors.New("record not found")
		userChildMealRepository.On("FindByID", mock.Anything, mock.AnythingOfType("int")).Return(domain, errError).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		_, err := userChildMealUsecase.Delete(ctx, &domain)

		assert.Equal(t, errError, err)
	})

	t.Run("test case 4, user child error", func(t *testing.T) {
		var (
			userChildMealRepository userChildMealMock.Repository
			userChildUsecase        userChildMock.Usecase
			mealPlanUsecase         mealPlanMock.Usecase
			calendarUsecase         calendarMock.Usecase
			userUsecase             userMock.Usecase
		)
		userChildMealUsecase := userChildMeal.NewUserChildMealUsecase(2, &userChildMealRepository,
			&userChildUsecase, &mealPlanUsecase, &calendarUsecase, &userUsecase)

		errError := errors.New("user child error")
		domain := userChildMeal.Domain{
			ID:     1,
			UserID: 1,
			Name:   "code",
			Unit:   "name",
			Status: "done",
		}
		userChildMealRepository.On("FindByID", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("string")).Return(domain, nil).Once()
		userChildUsecase.On("FindByID", mock.Anything, mock.AnythingOfType("int")).Return(userChild.Domain{}, errError).Once()
		userChildMealRepository.On("Delete", mock.Anything, mock.AnythingOfType("*userChildMeal.Domain")).Return(domain, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		_, err := userChildMealUsecase.Delete(ctx, &domain)

		assert.Equal(t, errError, err)
	})

	t.Run("test case 5, delete error", func(t *testing.T) {
		var (
			userChildMealRepository userChildMealMock.Repository
			userChildUsecase        userChildMock.Usecase
			mealPlanUsecase         mealPlanMock.Usecase
			calendarUsecase         calendarMock.Usecase
			userUsecase             userMock.Usecase
		)
		userChildMealUsecase := userChildMeal.NewUserChildMealUsecase(2, &userChildMealRepository,
			&userChildUsecase, &mealPlanUsecase, &calendarUsecase, &userUsecase)

		errError := errors.New("repo error")
		domain := userChildMeal.Domain{
			ID:     1,
			UserID: 1,
			Name:   "code",
			Unit:   "name",
			Status: "done",
		}
		userChildMealRepository.On("FindByID", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("string")).Return(domain, nil).Once()
		userChildDomain := userChild.Domain{
			ID:     1,
			UserID: 1,
			Name:   "name",
		}
		userChildUsecase.On("FindByID", mock.Anything, mock.AnythingOfType("int")).Return(userChildDomain, nil).Once()
		userChildMealRepository.On("Delete", mock.Anything, mock.AnythingOfType("*userChildMeal.Domain")).Return(userChildMeal.Domain{}, errError).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		_, err := userChildMealUsecase.Delete(ctx, &domain)

		assert.Equal(t, errError, err)
	})
}
