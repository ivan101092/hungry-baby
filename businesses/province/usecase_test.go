package province_test

import (
	"context"
	"errors"
	"hungry-baby/businesses"
	province "hungry-baby/businesses/province"
	provinceMock "hungry-baby/mocks/province"
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
			provinceRepository provinceMock.Repository
		)
		provinceUsecase := province.NewProvinceUsecase(2, &provinceRepository)

		domain := []province.Domain{
			{
				ID:        1,
				CountryID: 1,
				Code:      "code",
				Name:      "name",
				Status:    true,
			},
		}
		provinceRepository.On("FindAll", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("int"), mock.AnythingOfType("string")).Return(domain, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		result, err := provinceUsecase.FindAll(ctx, "1", 1, "true")

		assert.Nil(t, err)
		assert.Equal(t, domain, result)
	})

	t.Run("test case 2, error", func(t *testing.T) {
		var (
			provinceRepository provinceMock.Repository
		)
		provinceUsecase := province.NewProvinceUsecase(2, &provinceRepository)

		errError := errors.New("Error Repo")
		provinceRepository.On("FindAll", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("int"), mock.AnythingOfType("string")).Return(nil, errError).Once()
		ctx := context.WithValue(context.Background(), "userID", 1)
		_, err := provinceUsecase.FindAll(ctx, "1", 1, "true")

		assert.Equal(t, errError, err)
	})
}

func TestFind(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		var (
			provinceRepository provinceMock.Repository
		)
		provinceUsecase := province.NewProvinceUsecase(2, &provinceRepository)

		domain := []province.Domain{
			{
				ID:        1,
				CountryID: 1,
				Code:      "code",
				Name:      "name",
				Status:    true,
			},
		}
		provinceRepository.On("Find", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("int"),
			mock.AnythingOfType("int")).Return(domain, 1, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		result, total, err := provinceUsecase.Find(ctx, "1", 1, "true", 0, 0)

		assert.Nil(t, err)
		assert.Equal(t, domain, result)
		assert.Equal(t, 1, total)
	})

	t.Run("test case 2, error", func(t *testing.T) {
		var (
			provinceRepository provinceMock.Repository
		)
		provinceUsecase := province.NewProvinceUsecase(2, &provinceRepository)

		errError := errors.New("Error Repo")
		provinceRepository.On("Find", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("int"),
			mock.AnythingOfType("int")).Return(nil, 0, errError).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		_, _, err := provinceUsecase.Find(ctx, "1", 1, "true", 0, 0)

		assert.Equal(t, errError, err)
	})
}

func TestFindByID(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		var (
			provinceRepository provinceMock.Repository
		)
		provinceUsecase := province.NewProvinceUsecase(2, &provinceRepository)

		domain := province.Domain{
			ID:        1,
			CountryID: 1,
			Code:      "code",
			Name:      "name",
			Status:    true,
		}
		provinceRepository.On("FindByID", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("string")).Return(domain, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		result, err := provinceUsecase.FindByID(ctx, 1, "true")

		assert.Nil(t, err)
		assert.Equal(t, domain, result)
	})

	t.Run("test case 2, invalid id", func(t *testing.T) {
		var (
			provinceRepository provinceMock.Repository
		)
		provinceUsecase := province.NewProvinceUsecase(2, &provinceRepository)

		domain := province.Domain{
			ID:        1,
			CountryID: 1,
			Code:      "code",
			Name:      "name",
			Status:    true,
		}
		provinceRepository.On("FindByID", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("string")).Return(domain, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		_, err := provinceUsecase.FindByID(ctx, -1, "true")

		assert.Equal(t, businesses.ErrIDNotFound, err)
	})

	t.Run("test case 3, error", func(t *testing.T) {
		var (
			provinceRepository provinceMock.Repository
		)
		provinceUsecase := province.NewProvinceUsecase(2, &provinceRepository)

		errError := errors.New("Error Repo")
		provinceRepository.On("FindByID", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("string")).Return(province.Domain{}, errError).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		_, err := provinceUsecase.FindByID(ctx, 1, "true")

		assert.Equal(t, errError, err)
	})
}

func TestFindByCode(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		var (
			provinceRepository provinceMock.Repository
		)
		provinceUsecase := province.NewProvinceUsecase(2, &provinceRepository)

		domain := province.Domain{
			ID:        1,
			CountryID: 1,
			Code:      "code",
			Name:      "name",
			Status:    true,
		}
		provinceRepository.On("FindByCode", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(domain, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		result, err := provinceUsecase.FindByCode(ctx, "code", "true")

		assert.Nil(t, err)
		assert.Equal(t, domain, result)
	})

	t.Run("test case 2, error", func(t *testing.T) {
		var (
			provinceRepository provinceMock.Repository
		)
		provinceUsecase := province.NewProvinceUsecase(2, &provinceRepository)

		errError := errors.New("Error Repo")
		provinceRepository.On("FindByCode", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(province.Domain{}, errError).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		_, err := provinceUsecase.FindByCode(ctx, "code", "true")

		assert.Equal(t, errError, err)
	})
}

func TestStore(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		var (
			provinceRepository provinceMock.Repository
		)
		provinceUsecase := province.NewProvinceUsecase(2, &provinceRepository)

		domain := province.Domain{
			ID:        1,
			CountryID: 1,
			Code:      "code",
			Name:      "name",
			Status:    true,
		}
		errError := errors.New("record not found")
		provinceRepository.On("FindByCode", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(province.Domain{}, errError).Once()
		provinceRepository.On("Store", mock.Anything, mock.AnythingOfType("*province.Domain")).Return(domain, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		result, err := provinceUsecase.Store(ctx, &domain)

		assert.Nil(t, err)
		assert.Equal(t, domain, result)
	})

	t.Run("test case 2, find error", func(t *testing.T) {
		var (
			provinceRepository provinceMock.Repository
		)
		provinceUsecase := province.NewProvinceUsecase(2, &provinceRepository)

		domain := province.Domain{
			ID:        1,
			CountryID: 1,
			Code:      "code",
			Name:      "name",
			Status:    true,
		}
		errError := errors.New("internal server")
		provinceRepository.On("FindByCode", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(province.Domain{}, errError).Once()
		provinceRepository.On("Store", mock.Anything, mock.AnythingOfType("*province.Domain")).Return(domain, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		_, err := provinceUsecase.Store(ctx, &domain)

		assert.Equal(t, errError, err)
	})

	t.Run("test case 3, find exist", func(t *testing.T) {
		var (
			provinceRepository provinceMock.Repository
		)
		provinceUsecase := province.NewProvinceUsecase(2, &provinceRepository)

		domain := province.Domain{
			ID:        1,
			CountryID: 1,
			Code:      "code",
			Name:      "name",
			Status:    true,
		}
		provinceRepository.On("FindByCode", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(domain, nil).Once()
		provinceRepository.On("Store", mock.Anything, mock.AnythingOfType("*province.Domain")).Return(domain, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		_, err := provinceUsecase.Store(ctx, &domain)

		assert.Equal(t, businesses.ErrDuplicateData, err)
	})

	t.Run("test case 4, store error", func(t *testing.T) {
		var (
			provinceRepository provinceMock.Repository
		)
		provinceUsecase := province.NewProvinceUsecase(2, &provinceRepository)

		domain := province.Domain{
			ID:        1,
			CountryID: 1,
			Code:      "code",
			Name:      "name",
			Status:    true,
		}
		errError := errors.New("record not found")
		errRepo := errors.New("repo error")
		provinceRepository.On("FindByCode", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(province.Domain{}, errError).Once()
		provinceRepository.On("Store", mock.Anything, mock.AnythingOfType("*province.Domain")).Return(province.Domain{}, errRepo).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		_, err := provinceUsecase.Store(ctx, &domain)

		assert.Equal(t, errRepo, err)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		var (
			provinceRepository provinceMock.Repository
		)
		provinceUsecase := province.NewProvinceUsecase(2, &provinceRepository)

		domain := province.Domain{
			ID:        1,
			CountryID: 1,
			Code:      "code",
			Name:      "name",
			Status:    true,
		}
		errError := errors.New("record not found")
		provinceRepository.On("FindByCode", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(province.Domain{}, errError).Once()
		provinceRepository.On("Update", mock.Anything, mock.AnythingOfType("*province.Domain")).Return(domain, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		result, err := provinceUsecase.Update(ctx, &domain)

		assert.Nil(t, err)
		assert.Equal(t, domain, result)
	})

	t.Run("test case 2, find error", func(t *testing.T) {
		var (
			provinceRepository provinceMock.Repository
		)
		provinceUsecase := province.NewProvinceUsecase(2, &provinceRepository)

		domain := province.Domain{
			ID:        1,
			CountryID: 1,
			Code:      "code",
			Name:      "name",
			Status:    true,
		}
		errError := errors.New("internal server")
		provinceRepository.On("FindByCode", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(province.Domain{}, errError).Once()
		provinceRepository.On("Update", mock.Anything, mock.AnythingOfType("*province.Domain")).Return(domain, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		_, err := provinceUsecase.Update(ctx, &domain)

		assert.Equal(t, errError, err)
	})

	t.Run("test case 3, find exist", func(t *testing.T) {
		var (
			provinceRepository provinceMock.Repository
		)
		provinceUsecase := province.NewProvinceUsecase(2, &provinceRepository)

		domain := province.Domain{
			ID:        1,
			CountryID: 1,
			Code:      "code",
			Name:      "name",
			Status:    true,
		}
		provinceRepository.On("FindByCode", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(domain, nil).Once()
		provinceRepository.On("Update", mock.Anything, mock.AnythingOfType("*province.Domain")).Return(domain, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		domain.ID = 2
		_, err := provinceUsecase.Update(ctx, &domain)

		assert.Equal(t, businesses.ErrDuplicateData, err)
	})

	t.Run("test case 4, store error", func(t *testing.T) {
		var (
			provinceRepository provinceMock.Repository
		)
		provinceUsecase := province.NewProvinceUsecase(2, &provinceRepository)

		domain := province.Domain{
			ID:        1,
			CountryID: 1,
			Code:      "code",
			Name:      "name",
			Status:    true,
		}
		errError := errors.New("record not found")
		errRepo := errors.New("repo error")
		provinceRepository.On("FindByCode", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(province.Domain{}, errError).Once()
		provinceRepository.On("Update", mock.Anything, mock.AnythingOfType("*province.Domain")).Return(province.Domain{}, errRepo).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		_, err := provinceUsecase.Update(ctx, &domain)

		assert.Equal(t, errRepo, err)
	})
}

func TestDelete(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		var (
			provinceRepository provinceMock.Repository
		)
		provinceUsecase := province.NewProvinceUsecase(2, &provinceRepository)

		domain := province.Domain{
			ID:        1,
			CountryID: 1,
			Code:      "code",
			Name:      "name",
			Status:    true,
		}
		provinceRepository.On("FindByID", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("string")).Return(domain, nil).Once()
		provinceRepository.On("Delete", mock.Anything, mock.AnythingOfType("*province.Domain")).Return(domain, nil).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		result, err := provinceUsecase.Delete(ctx, &domain)

		assert.Nil(t, err)
		assert.Equal(t, domain, result)
	})

	t.Run("test case 2, find error", func(t *testing.T) {
		var (
			provinceRepository provinceMock.Repository
		)
		provinceUsecase := province.NewProvinceUsecase(2, &provinceRepository)

		domain := province.Domain{
			ID:        1,
			CountryID: 1,
			Code:      "code",
			Name:      "name",
			Status:    true,
		}
		errError := errors.New("record not found")
		provinceRepository.On("FindByID", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("string")).Return(province.Domain{}, errError).Once()
		provinceRepository.On("Delete", mock.Anything, mock.AnythingOfType("*province.Domain")).Return(province.Domain{}, errError).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		_, err := provinceUsecase.Delete(ctx, &domain)

		assert.Equal(t, errError, err)
	})

	t.Run("test case 3, delete error", func(t *testing.T) {
		var (
			provinceRepository provinceMock.Repository
		)
		provinceUsecase := province.NewProvinceUsecase(2, &provinceRepository)

		domain := province.Domain{
			ID:        1,
			CountryID: 1,
			Code:      "code",
			Name:      "name",
			Status:    true,
		}
		errError := errors.New("repo error")
		provinceRepository.On("FindByID", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("string")).Return(domain, nil).Once()
		provinceRepository.On("Delete", mock.Anything, mock.AnythingOfType("*province.Domain")).Return(province.Domain{}, errError).Once()

		ctx := context.WithValue(context.Background(), "userID", 1)
		_, err := provinceUsecase.Delete(ctx, &domain)

		assert.Equal(t, errError, err)
	})
}
