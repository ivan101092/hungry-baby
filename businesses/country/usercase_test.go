package country_test

import (
	"context"
	"errors"
	"hungry-baby/businesses"
	country "hungry-baby/businesses/country"
	countryMock "hungry-baby/businesses/country/mocks"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	countryRepository countryMock.Repository
	countryUsecase    country.Usecase
)

func setup() {
	countryUsecase = country.NewCountryUsecase(2, &countryRepository)
}

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestGetById(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		domain := country.Domain{
			ID:          1,
			CountryCode: "Sport",
			Name:        "TestCase1",
			Status:      true,
		}
		countryRepository.On("FindByID", mock.AnythingOfType("int")).Return(domain, nil).Once()

		result, err := countryUsecase.FindByID(context.Background(), 1, "true")

		assert.Nil(t, err)
		assert.Equal(t, domain.CountryCode, result.CountryCode)
	})

	t.Run("test case 2, invalid id", func(t *testing.T) {
		result, err := countryUsecase.FindByID(context.Background(), -1, "true")

		assert.Equal(t, result, country.Domain{})
		assert.Equal(t, err, businesses.ErrIDNotFound)
	})

	t.Run("test case 3, repository error", func(t *testing.T) {
		errNotFound := errors.New("(Repo) ID Not Found")
		countryRepository.On("FindByID", mock.AnythingOfType("int")).Return(country.Domain{}, errNotFound).Once()
		result, err := countryUsecase.FindByID(context.Background(), 10, "true")

		assert.Equal(t, result, country.Domain{})
		assert.Equal(t, err, errNotFound)
	})
}
