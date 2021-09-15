package country_test

import (
	"context"
	"errors"
	"hungry-baby/businesses"
	country "hungry-baby/businesses/country"
	countryMock "hungry-baby/mocks/country"
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
			countryRepository countryMock.Repository
		)
		countryUsecase := country.NewCountryUsecase(2, &countryRepository)

		domain := []country.Domain{
			{
				ID:          1,
				CountryCode: "code",
				Name:        "name",
				Status:      true,
			},
		}
		countryRepository.On("FindAll", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(domain, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		result, err := countryUsecase.FindAll(ctx, "1", "true")

		assert.Nil(t, err)
		assert.Equal(t, domain, result)
	})

	t.Run("test case 2, error", func(t *testing.T) {
		var (
			countryRepository countryMock.Repository
		)
		countryUsecase := country.NewCountryUsecase(2, &countryRepository)

		errError := errors.New("Error Repo")
		countryRepository.On("FindAll", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(nil, errError).Once()
		ctx := context.WithValue(context.Background(), "userID", 1)
		_, err := countryUsecase.FindAll(ctx, "1", "true")

		assert.Equal(t, errError, err)
	})
}

func TestFind(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		var (
			countryRepository countryMock.Repository
		)
		countryUsecase := country.NewCountryUsecase(2, &countryRepository)

		domain := []country.Domain{
			{
				ID:          1,
				CountryCode: "code",
				Name:        "name",
				Status:      true,
			},
		}
		countryRepository.On("Find", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("string"), mock.AnythingOfType("int"),
			mock.AnythingOfType("int")).Return(domain, 1, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		result, total, err := countryUsecase.Find(ctx, "1", "true", 0, 0)

		assert.Nil(t, err)
		assert.Equal(t, domain, result)
		assert.Equal(t, 1, total)
	})

	t.Run("test case 2, error", func(t *testing.T) {
		var (
			countryRepository countryMock.Repository
		)
		countryUsecase := country.NewCountryUsecase(2, &countryRepository)

		errError := errors.New("Error Repo")
		countryRepository.On("Find", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("string"), mock.AnythingOfType("int"),
			mock.AnythingOfType("int")).Return(nil, 0, errError).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		_, _, err := countryUsecase.Find(ctx, "1", "true", 0, 0)

		assert.Equal(t, errError, err)
	})
}

func TestFindByID(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		var (
			countryRepository countryMock.Repository
		)
		countryUsecase := country.NewCountryUsecase(2, &countryRepository)

		domain := country.Domain{
			ID:          1,
			CountryCode: "code",
			Name:        "name",
			Status:      true,
		}
		countryRepository.On("FindByID", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("string")).Return(domain, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		result, err := countryUsecase.FindByID(ctx, 1, "true")

		assert.Nil(t, err)
		assert.Equal(t, domain, result)
	})

	t.Run("test case 2, invalid id", func(t *testing.T) {
		var (
			countryRepository countryMock.Repository
		)
		countryUsecase := country.NewCountryUsecase(2, &countryRepository)

		domain := country.Domain{
			ID:          1,
			CountryCode: "code",
			Name:        "name",
			Status:      true,
		}
		countryRepository.On("FindByID", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("string")).Return(domain, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		_, err := countryUsecase.FindByID(ctx, -1, "true")

		assert.Equal(t, businesses.ErrIDNotFound, err)
	})

	t.Run("test case 3, error", func(t *testing.T) {
		var (
			countryRepository countryMock.Repository
		)
		countryUsecase := country.NewCountryUsecase(2, &countryRepository)

		errError := errors.New("Error Repo")
		countryRepository.On("FindByID", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("string")).Return(country.Domain{}, errError).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		_, err := countryUsecase.FindByID(ctx, 1, "true")

		assert.Equal(t, errError, err)
	})
}

func TestFindByCode(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		var (
			countryRepository countryMock.Repository
		)
		countryUsecase := country.NewCountryUsecase(2, &countryRepository)

		domain := country.Domain{
			ID:          1,
			CountryCode: "code",
			Name:        "name",
			Status:      true,
		}
		countryRepository.On("FindByCode", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(domain, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		result, err := countryUsecase.FindByCode(ctx, "code", "true")

		assert.Nil(t, err)
		assert.Equal(t, domain, result)
	})

	t.Run("test case 2, error", func(t *testing.T) {
		var (
			countryRepository countryMock.Repository
		)
		countryUsecase := country.NewCountryUsecase(2, &countryRepository)

		errError := errors.New("Error Repo")
		countryRepository.On("FindByCode", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(country.Domain{}, errError).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		_, err := countryUsecase.FindByCode(ctx, "code", "true")

		assert.Equal(t, errError, err)
	})
}

func TestStore(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		var (
			countryRepository countryMock.Repository
		)
		countryUsecase := country.NewCountryUsecase(2, &countryRepository)

		domain := country.Domain{
			ID:          1,
			CountryCode: "code",
			Name:        "name",
			Status:      true,
		}
		errError := errors.New("record not found")
		countryRepository.On("FindByCode", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(country.Domain{}, errError).Once()
		countryRepository.On("Store", mock.Anything, mock.AnythingOfType("*country.Domain")).Return(domain, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		result, err := countryUsecase.Store(ctx, &domain)

		assert.Nil(t, err)
		assert.Equal(t, domain, result)
	})

	t.Run("test case 2, find error", func(t *testing.T) {
		var (
			countryRepository countryMock.Repository
		)
		countryUsecase := country.NewCountryUsecase(2, &countryRepository)

		domain := country.Domain{
			ID:          1,
			CountryCode: "code",
			Name:        "name",
			Status:      true,
		}
		errError := errors.New("internal server")
		countryRepository.On("FindByCode", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(country.Domain{}, errError).Once()
		countryRepository.On("Store", mock.Anything, mock.AnythingOfType("*country.Domain")).Return(domain, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		_, err := countryUsecase.Store(ctx, &domain)

		assert.Equal(t, errError, err)
	})

	t.Run("test case 3, find exist", func(t *testing.T) {
		var (
			countryRepository countryMock.Repository
		)
		countryUsecase := country.NewCountryUsecase(2, &countryRepository)

		domain := country.Domain{
			ID:          1,
			CountryCode: "code",
			Name:        "name",
			Status:      true,
		}
		countryRepository.On("FindByCode", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(domain, nil).Once()
		countryRepository.On("Store", mock.Anything, mock.AnythingOfType("*country.Domain")).Return(domain, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		_, err := countryUsecase.Store(ctx, &domain)

		assert.Equal(t, businesses.ErrDuplicateData, err)
	})

	t.Run("test case 4, store error", func(t *testing.T) {
		var (
			countryRepository countryMock.Repository
		)
		countryUsecase := country.NewCountryUsecase(2, &countryRepository)

		domain := country.Domain{
			ID:          1,
			CountryCode: "code",
			Name:        "name",
			Status:      true,
		}
		errError := errors.New("record not found")
		errRepo := errors.New("repo error")
		countryRepository.On("FindByCode", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(country.Domain{}, errError).Once()
		countryRepository.On("Store", mock.Anything, mock.AnythingOfType("*country.Domain")).Return(country.Domain{}, errRepo).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		_, err := countryUsecase.Store(ctx, &domain)

		assert.Equal(t, errRepo, err)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		var (
			countryRepository countryMock.Repository
		)
		countryUsecase := country.NewCountryUsecase(2, &countryRepository)

		domain := country.Domain{
			ID:          1,
			CountryCode: "code",
			Name:        "name",
			Status:      true,
		}
		errError := errors.New("record not found")
		countryRepository.On("FindByCode", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(country.Domain{}, errError).Once()
		countryRepository.On("Update", mock.Anything, mock.AnythingOfType("*country.Domain")).Return(domain, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		result, err := countryUsecase.Update(ctx, &domain)

		assert.Nil(t, err)
		assert.Equal(t, domain, result)
	})

	t.Run("test case 2, find error", func(t *testing.T) {
		var (
			countryRepository countryMock.Repository
		)
		countryUsecase := country.NewCountryUsecase(2, &countryRepository)

		domain := country.Domain{
			ID:          1,
			CountryCode: "code",
			Name:        "name",
			Status:      true,
		}
		errError := errors.New("internal server")
		countryRepository.On("FindByCode", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(country.Domain{}, errError).Once()
		countryRepository.On("Update", mock.Anything, mock.AnythingOfType("*country.Domain")).Return(domain, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		_, err := countryUsecase.Update(ctx, &domain)

		assert.Equal(t, errError, err)
	})

	t.Run("test case 3, find exist", func(t *testing.T) {
		var (
			countryRepository countryMock.Repository
		)
		countryUsecase := country.NewCountryUsecase(2, &countryRepository)

		domain := country.Domain{
			ID:          1,
			CountryCode: "code",
			Name:        "name",
			Status:      true,
		}
		countryRepository.On("FindByCode", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(domain, nil).Once()
		countryRepository.On("Update", mock.Anything, mock.AnythingOfType("*country.Domain")).Return(domain, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		domain.ID = 2
		_, err := countryUsecase.Update(ctx, &domain)

		assert.Equal(t, businesses.ErrDuplicateData, err)
	})

	t.Run("test case 4, store error", func(t *testing.T) {
		var (
			countryRepository countryMock.Repository
		)
		countryUsecase := country.NewCountryUsecase(2, &countryRepository)

		domain := country.Domain{
			ID:          1,
			CountryCode: "code",
			Name:        "name",
			Status:      true,
		}
		errError := errors.New("record not found")
		errRepo := errors.New("repo error")
		countryRepository.On("FindByCode", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(country.Domain{}, errError).Once()
		countryRepository.On("Update", mock.Anything, mock.AnythingOfType("*country.Domain")).Return(country.Domain{}, errRepo).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		_, err := countryUsecase.Update(ctx, &domain)

		assert.Equal(t, errRepo, err)
	})
}

func TestDelete(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		var (
			countryRepository countryMock.Repository
		)
		countryUsecase := country.NewCountryUsecase(2, &countryRepository)

		domain := country.Domain{
			ID:          1,
			CountryCode: "code",
			Name:        "name",
			Status:      true,
		}
		countryRepository.On("FindByID", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("string")).Return(domain, nil).Once()
		countryRepository.On("Delete", mock.Anything, mock.AnythingOfType("*country.Domain")).Return(domain, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		result, err := countryUsecase.Delete(ctx, &domain)

		assert.Nil(t, err)
		assert.Equal(t, domain, result)
	})

	t.Run("test case 2, find error", func(t *testing.T) {
		var (
			countryRepository countryMock.Repository
		)
		countryUsecase := country.NewCountryUsecase(2, &countryRepository)

		domain := country.Domain{
			ID:          1,
			CountryCode: "code",
			Name:        "name",
			Status:      true,
		}
		errError := errors.New("record not found")
		countryRepository.On("FindByID", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("string")).Return(country.Domain{}, errError).Once()
		countryRepository.On("Delete", mock.Anything, mock.AnythingOfType("*country.Domain")).Return(country.Domain{}, errError).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		_, err := countryUsecase.Delete(ctx, &domain)

		assert.Equal(t, errError, err)
	})

	t.Run("test case 3, delete error", func(t *testing.T) {
		var (
			countryRepository countryMock.Repository
		)
		countryUsecase := country.NewCountryUsecase(2, &countryRepository)

		domain := country.Domain{
			ID:          1,
			CountryCode: "code",
			Name:        "name",
			Status:      true,
		}
		errError := errors.New("repo error")
		countryRepository.On("FindByID", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("string")).Return(domain, nil).Once()
		countryRepository.On("Delete", mock.Anything, mock.AnythingOfType("*country.Domain")).Return(country.Domain{}, errError).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		_, err := countryUsecase.Delete(ctx, &domain)

		assert.Equal(t, errError, err)
	})
}
