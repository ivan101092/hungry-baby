package city_test

import (
	"context"
	"errors"
	"hungry-baby/businesses"
	city "hungry-baby/businesses/city"
	cityMock "hungry-baby/mocks/city"
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
			cityRepository cityMock.Repository
		)
		cityUsecase := city.NewCityUsecase(2, &cityRepository)

		domain := []city.Domain{
			{
				ID:         1,
				ProvinceID: 1,
				Code:       "code",
				Name:       "name",
				Status:     true,
			},
		}
		cityRepository.On("FindAll", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("int"), mock.AnythingOfType("string")).Return(domain, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		result, err := cityUsecase.FindAll(ctx, "1", 1, "true")

		assert.Nil(t, err)
		assert.Equal(t, domain, result)
	})

	t.Run("test case 2, error", func(t *testing.T) {
		var (
			cityRepository cityMock.Repository
		)
		cityUsecase := city.NewCityUsecase(2, &cityRepository)

		errError := errors.New("Error Repo")
		cityRepository.On("FindAll", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("int"), mock.AnythingOfType("string")).Return(nil, errError).Once()
		ctx := context.WithValue(context.Background(), "userID", 1)
		_, err := cityUsecase.FindAll(ctx, "1", 1, "true")

		assert.Equal(t, errError, err)
	})
}

func TestFind(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		var (
			cityRepository cityMock.Repository
		)
		cityUsecase := city.NewCityUsecase(2, &cityRepository)

		domain := []city.Domain{
			{
				ID:         1,
				ProvinceID: 1,
				Code:       "code",
				Name:       "name",
				Status:     true,
			},
		}
		cityRepository.On("Find", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("int"),
			mock.AnythingOfType("int")).Return(domain, 1, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		result, total, err := cityUsecase.Find(ctx, "1", 1, "true", 0, 0)

		assert.Nil(t, err)
		assert.Equal(t, domain, result)
		assert.Equal(t, 1, total)
	})

	t.Run("test case 2, error", func(t *testing.T) {
		var (
			cityRepository cityMock.Repository
		)
		cityUsecase := city.NewCityUsecase(2, &cityRepository)

		errError := errors.New("Error Repo")
		cityRepository.On("Find", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("int"),
			mock.AnythingOfType("int")).Return(nil, 0, errError).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		_, _, err := cityUsecase.Find(ctx, "1", 1, "true", 0, 0)

		assert.Equal(t, errError, err)
	})
}

func TestFindByID(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		var (
			cityRepository cityMock.Repository
		)
		cityUsecase := city.NewCityUsecase(2, &cityRepository)

		domain := city.Domain{
			ID:         1,
			ProvinceID: 1,
			Code:       "code",
			Name:       "name",
			Status:     true,
		}
		cityRepository.On("FindByID", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("string")).Return(domain, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		result, err := cityUsecase.FindByID(ctx, 1, "true")

		assert.Nil(t, err)
		assert.Equal(t, domain, result)
	})

	t.Run("test case 2, invalid id", func(t *testing.T) {
		var (
			cityRepository cityMock.Repository
		)
		cityUsecase := city.NewCityUsecase(2, &cityRepository)

		domain := city.Domain{
			ID:         1,
			ProvinceID: 1,
			Code:       "code",
			Name:       "name",
			Status:     true,
		}
		cityRepository.On("FindByID", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("string")).Return(domain, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		_, err := cityUsecase.FindByID(ctx, -1, "true")

		assert.Equal(t, businesses.ErrIDNotFound, err)
	})

	t.Run("test case 3, error", func(t *testing.T) {
		var (
			cityRepository cityMock.Repository
		)
		cityUsecase := city.NewCityUsecase(2, &cityRepository)

		errError := errors.New("Error Repo")
		cityRepository.On("FindByID", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("string")).Return(city.Domain{}, errError).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		_, err := cityUsecase.FindByID(ctx, 1, "true")

		assert.Equal(t, errError, err)
	})
}

func TestFindByCode(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		var (
			cityRepository cityMock.Repository
		)
		cityUsecase := city.NewCityUsecase(2, &cityRepository)

		domain := city.Domain{
			ID:         1,
			ProvinceID: 1,
			Code:       "code",
			Name:       "name",
			Status:     true,
		}
		cityRepository.On("FindByCode", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(domain, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		result, err := cityUsecase.FindByCode(ctx, "code", "true")

		assert.Nil(t, err)
		assert.Equal(t, domain, result)
	})

	t.Run("test case 2, error", func(t *testing.T) {
		var (
			cityRepository cityMock.Repository
		)
		cityUsecase := city.NewCityUsecase(2, &cityRepository)

		errError := errors.New("Error Repo")
		cityRepository.On("FindByCode", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(city.Domain{}, errError).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		_, err := cityUsecase.FindByCode(ctx, "code", "true")

		assert.Equal(t, errError, err)
	})
}

func TestStore(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		var (
			cityRepository cityMock.Repository
		)
		cityUsecase := city.NewCityUsecase(2, &cityRepository)

		domain := city.Domain{
			ID:         1,
			ProvinceID: 1,
			Code:       "code",
			Name:       "name",
			Status:     true,
		}
		errError := errors.New("record not found")
		cityRepository.On("FindByCode", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(city.Domain{}, errError).Once()
		cityRepository.On("Store", mock.Anything, mock.AnythingOfType("*city.Domain")).Return(domain, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		result, err := cityUsecase.Store(ctx, &domain)

		assert.Nil(t, err)
		assert.Equal(t, domain, result)
	})

	t.Run("test case 2, find error", func(t *testing.T) {
		var (
			cityRepository cityMock.Repository
		)
		cityUsecase := city.NewCityUsecase(2, &cityRepository)

		domain := city.Domain{
			ID:         1,
			ProvinceID: 1,
			Code:       "code",
			Name:       "name",
			Status:     true,
		}
		errError := errors.New("internal server")
		cityRepository.On("FindByCode", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(city.Domain{}, errError).Once()
		cityRepository.On("Store", mock.Anything, mock.AnythingOfType("*city.Domain")).Return(domain, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		_, err := cityUsecase.Store(ctx, &domain)

		assert.Equal(t, errError, err)
	})

	t.Run("test case 3, find exist", func(t *testing.T) {
		var (
			cityRepository cityMock.Repository
		)
		cityUsecase := city.NewCityUsecase(2, &cityRepository)

		domain := city.Domain{
			ID:         1,
			ProvinceID: 1,
			Code:       "code",
			Name:       "name",
			Status:     true,
		}
		cityRepository.On("FindByCode", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(domain, nil).Once()
		cityRepository.On("Store", mock.Anything, mock.AnythingOfType("*city.Domain")).Return(domain, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		_, err := cityUsecase.Store(ctx, &domain)

		assert.Equal(t, businesses.ErrDuplicateData, err)
	})

	t.Run("test case 4, store error", func(t *testing.T) {
		var (
			cityRepository cityMock.Repository
		)
		cityUsecase := city.NewCityUsecase(2, &cityRepository)

		domain := city.Domain{
			ID:         1,
			ProvinceID: 1,
			Code:       "code",
			Name:       "name",
			Status:     true,
		}
		errError := errors.New("record not found")
		errRepo := errors.New("repo error")
		cityRepository.On("FindByCode", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(city.Domain{}, errError).Once()
		cityRepository.On("Store", mock.Anything, mock.AnythingOfType("*city.Domain")).Return(city.Domain{}, errRepo).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		_, err := cityUsecase.Store(ctx, &domain)

		assert.Equal(t, errRepo, err)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		var (
			cityRepository cityMock.Repository
		)
		cityUsecase := city.NewCityUsecase(2, &cityRepository)

		domain := city.Domain{
			ID:         1,
			ProvinceID: 1,
			Code:       "code",
			Name:       "name",
			Status:     true,
		}
		errError := errors.New("record not found")
		cityRepository.On("FindByCode", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(city.Domain{}, errError).Once()
		cityRepository.On("Update", mock.Anything, mock.AnythingOfType("*city.Domain")).Return(domain, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		result, err := cityUsecase.Update(ctx, &domain)

		assert.Nil(t, err)
		assert.Equal(t, domain, result)
	})

	t.Run("test case 2, find error", func(t *testing.T) {
		var (
			cityRepository cityMock.Repository
		)
		cityUsecase := city.NewCityUsecase(2, &cityRepository)

		domain := city.Domain{
			ID:         1,
			ProvinceID: 1,
			Code:       "code",
			Name:       "name",
			Status:     true,
		}
		errError := errors.New("internal server")
		cityRepository.On("FindByCode", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(city.Domain{}, errError).Once()
		cityRepository.On("Update", mock.Anything, mock.AnythingOfType("*city.Domain")).Return(domain, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		_, err := cityUsecase.Update(ctx, &domain)

		assert.Equal(t, errError, err)
	})

	t.Run("test case 3, find exist", func(t *testing.T) {
		var (
			cityRepository cityMock.Repository
		)
		cityUsecase := city.NewCityUsecase(2, &cityRepository)

		domain := city.Domain{
			ID:         1,
			ProvinceID: 1,
			Code:       "code",
			Name:       "name",
			Status:     true,
		}
		cityRepository.On("FindByCode", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(domain, nil).Once()
		cityRepository.On("Update", mock.Anything, mock.AnythingOfType("*city.Domain")).Return(domain, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		domain.ID = 2
		_, err := cityUsecase.Update(ctx, &domain)

		assert.Equal(t, businesses.ErrDuplicateData, err)
	})

	t.Run("test case 4, store error", func(t *testing.T) {
		var (
			cityRepository cityMock.Repository
		)
		cityUsecase := city.NewCityUsecase(2, &cityRepository)

		domain := city.Domain{
			ID:         1,
			ProvinceID: 1,
			Code:       "code",
			Name:       "name",
			Status:     true,
		}
		errError := errors.New("record not found")
		errRepo := errors.New("repo error")
		cityRepository.On("FindByCode", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(city.Domain{}, errError).Once()
		cityRepository.On("Update", mock.Anything, mock.AnythingOfType("*city.Domain")).Return(city.Domain{}, errRepo).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		_, err := cityUsecase.Update(ctx, &domain)

		assert.Equal(t, errRepo, err)
	})
}

func TestDelete(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		var (
			cityRepository cityMock.Repository
		)
		cityUsecase := city.NewCityUsecase(2, &cityRepository)

		domain := city.Domain{
			ID:         1,
			ProvinceID: 1,
			Code:       "code",
			Name:       "name",
			Status:     true,
		}
		cityRepository.On("FindByID", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("string")).Return(domain, nil).Once()
		cityRepository.On("Delete", mock.Anything, mock.AnythingOfType("*city.Domain")).Return(domain, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		result, err := cityUsecase.Delete(ctx, &domain)

		assert.Nil(t, err)
		assert.Equal(t, domain, result)
	})

	t.Run("test case 2, find error", func(t *testing.T) {
		var (
			cityRepository cityMock.Repository
		)
		cityUsecase := city.NewCityUsecase(2, &cityRepository)

		domain := city.Domain{
			ID:         1,
			ProvinceID: 1,
			Code:       "code",
			Name:       "name",
			Status:     true,
		}
		errError := errors.New("record not found")
		cityRepository.On("FindByID", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("string")).Return(city.Domain{}, errError).Once()
		cityRepository.On("Delete", mock.Anything, mock.AnythingOfType("*city.Domain")).Return(city.Domain{}, errError).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		_, err := cityUsecase.Delete(ctx, &domain)

		assert.Equal(t, errError, err)
	})

	t.Run("test case 3, delete error", func(t *testing.T) {
		var (
			cityRepository cityMock.Repository
		)
		cityUsecase := city.NewCityUsecase(2, &cityRepository)

		domain := city.Domain{
			ID:         1,
			ProvinceID: 1,
			Code:       "code",
			Name:       "name",
			Status:     true,
		}
		errError := errors.New("repo error")
		cityRepository.On("FindByID", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("string")).Return(domain, nil).Once()
		cityRepository.On("Delete", mock.Anything, mock.AnythingOfType("*city.Domain")).Return(city.Domain{}, errError).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		_, err := cityUsecase.Delete(ctx, &domain)

		assert.Equal(t, errError, err)
	})
}
