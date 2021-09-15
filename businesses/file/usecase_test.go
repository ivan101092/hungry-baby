package file_test

import (
	"context"
	"errors"
	"hungry-baby/businesses"
	file "hungry-baby/businesses/file"
	fileMock "hungry-baby/mocks/file"
	minioMock "hungry-baby/mocks/minio"
	"mime/multipart"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestFindByID(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		var (
			fileRepository  fileMock.Repository
			minioRepository minioMock.Repository
		)
		fileUsecase := file.NewFileUsecase(2, &fileRepository, &minioRepository)

		domain := file.Domain{
			ID:         1,
			Type:       "file",
			URL:        "file.jpg",
			FullURL:    "https://s3.hungrybaby.com/file.jpg",
			UserUpload: "1",
		}
		fileRepository.On("FindByID", mock.Anything, mock.AnythingOfType("int")).Return(domain, nil).Once()
		minioRepository.On("GetFile", mock.AnythingOfType("string")).Return("https://s3.hungrybaby.com/file.jpg", nil).Once()

		result, err := fileUsecase.FindByID(context.Background(), 1)

		assert.Nil(t, err)
		assert.Equal(t, domain.ID, result.ID)
	})

	t.Run("test case 2, invalid id", func(t *testing.T) {
		var (
			fileRepository  fileMock.Repository
			minioRepository minioMock.Repository
		)
		fileUsecase := file.NewFileUsecase(2, &fileRepository, &minioRepository)

		result, err := fileUsecase.FindByID(context.Background(), -1)

		assert.Equal(t, result, file.Domain{})
		assert.Equal(t, err, businesses.ErrFileIDResource)
	})

	t.Run("test case 3, repository error", func(t *testing.T) {
		var (
			fileRepository  fileMock.Repository
			minioRepository minioMock.Repository
		)
		fileUsecase := file.NewFileUsecase(2, &fileRepository, &minioRepository)

		errNotFound := errors.New("(Repo) ID Not Found")
		fileRepository.On("FindByID", mock.Anything, mock.AnythingOfType("int")).Return(file.Domain{}, errNotFound).Once()
		minioRepository.On("GetFile", mock.AnythingOfType("string")).Return("https://s3.hungrybaby.com/file.jpg", nil).Once()
		result, err := fileUsecase.FindByID(context.Background(), 10)

		assert.Equal(t, result, file.Domain{})
		assert.Equal(t, err, errNotFound)
	})

	t.Run("test case 4, minio error", func(t *testing.T) {
		var (
			fileRepository  fileMock.Repository
			minioRepository minioMock.Repository
		)
		fileUsecase := file.NewFileUsecase(2, &fileRepository, &minioRepository)

		domain := file.Domain{
			ID:         1,
			Type:       "file",
			URL:        "file.jpg",
			FullURL:    "https://s3.hungrybaby.com/file.jpg",
			UserUpload: "1",
		}
		errNotFound := errors.New("(Minio) ID Not Found")
		fileRepository.On("FindByID", mock.Anything, mock.AnythingOfType("int")).Return(domain, nil).Once()
		minioRepository.On("GetFile", mock.AnythingOfType("string")).Return("", errNotFound).Once()
		result, err := fileUsecase.FindByID(context.Background(), 1)

		assert.Equal(t, result, file.Domain{})
		assert.Equal(t, err, errNotFound)
	})
}

func TestStore(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		var (
			fileRepository  fileMock.Repository
			minioRepository minioMock.Repository
		)
		fileUsecase := file.NewFileUsecase(2, &fileRepository, &minioRepository)

		domain := file.Domain{
			ID:         1,
			Type:       "file",
			URL:        "file.jpg",
			FullURL:    "https://s3.hungrybaby.com/file.jpg",
			UserUpload: "1",
		}
		ctx := context.WithValue(context.Background(), "userID", 1)
		fileRepository.On("Store", mock.Anything, &domain).Return(domain, nil).Once()
		minioRepository.On("GetFile", mock.AnythingOfType("string")).Return("https://s3.hungrybaby.com/file.jpg", nil).Once()
		result, err := fileUsecase.Store(ctx, &domain)

		assert.Equal(t, result, domain)
		assert.Nil(t, err)
	})

	t.Run("test case 2, repository error", func(t *testing.T) {
		var (
			fileRepository  fileMock.Repository
			minioRepository minioMock.Repository
		)
		fileUsecase := file.NewFileUsecase(2, &fileRepository, &minioRepository)

		domain := file.Domain{
			ID:         1,
			Type:       "file",
			URL:        "file.jpg",
			FullURL:    "https://s3.hungrybaby.com/file.jpg",
			UserUpload: "1",
		}
		ctx := context.WithValue(context.Background(), "userID", 1)
		errNotFound := errors.New("(Repo) ID Not Found")
		fileRepository.On("Store", mock.Anything, &domain).Return(file.Domain{}, errNotFound).Once()
		minioRepository.On("GetFile", mock.AnythingOfType("string")).Return("https://s3.hungrybaby.com/file.jpg", nil).Once()
		result, err := fileUsecase.Store(ctx, &domain)

		assert.Equal(t, result, file.Domain{})
		assert.Equal(t, err, errNotFound)
	})

	t.Run("test case 3, minio error", func(t *testing.T) {
		var (
			fileRepository  fileMock.Repository
			minioRepository minioMock.Repository
		)
		fileUsecase := file.NewFileUsecase(2, &fileRepository, &minioRepository)

		domain := file.Domain{
			ID:         1,
			Type:       "file",
			URL:        "file.jpg",
			FullURL:    "",
			UserUpload: "1",
		}
		ctx := context.WithValue(context.Background(), "userID", 1)
		errError := errors.New("(Minio) Error")
		fileRepository.On("Store", mock.Anything, mock.AnythingOfType("*file.Domain")).Return(domain, nil).Once()
		minioRepository.On("GetFile", mock.AnythingOfType("string")).Return("", errError).Once()
		result, err := fileUsecase.Store(ctx, &domain)

		assert.Equal(t, result, file.Domain{})
		assert.Equal(t, errError, err)
	})
}

func TestUpload(t *testing.T) {
	t.Run("test case 3, repository error", func(t *testing.T) {
		var (
			fileRepository  fileMock.Repository
			minioRepository minioMock.Repository
		)
		fileUsecase := file.NewFileUsecase(2, &fileRepository, &minioRepository)

		errNotFound := errors.New("(Repo) ID Not Found")
		var fileHeader *multipart.FileHeader
		domain := file.Domain{
			Type:       "file",
			URL:        "file.jpg",
			FullURL:    "https://s3.hungrybaby.com/file.jpg",
			UserUpload: "1",
		}
		ctx := context.WithValue(context.Background(), "userID", 1)
		minioRepository.On("Upload", mock.AnythingOfType("string"), mock.Anything).Return("file.jpg", nil).Once()
		domainStore := file.Domain{
			Type:       "file",
			URL:        "file.jpg",
			FullURL:    "",
			UserUpload: "1",
		}
		fileRepository.On("Store", mock.Anything, &domainStore).Return(file.Domain{}, errNotFound).Once()
		minioRepository.On("GetFile", mock.AnythingOfType("string")).Return("https://s3.hungrybaby.com/file.jpg", nil).Once()
		result, err := fileUsecase.Upload(ctx, domain.Type, "", fileHeader)

		assert.Equal(t, result, file.Domain{})
		assert.Equal(t, errNotFound, err)
	})

	t.Run("test case 1, valid test", func(t *testing.T) {
		var (
			fileRepository  fileMock.Repository
			minioRepository minioMock.Repository
		)
		fileUsecase := file.NewFileUsecase(2, &fileRepository, &minioRepository)

		var fileHeader *multipart.FileHeader
		domain := file.Domain{
			Type:       "file",
			URL:        "file.jpg",
			FullURL:    "https://s3.hungrybaby.com/file.jpg",
			UserUpload: "1",
		}
		ctx := context.WithValue(context.Background(), "userID", 1)
		minioRepository.On("Upload", mock.AnythingOfType("string"), mock.Anything).Return("file.jpg", nil).Once()
		domainStore := file.Domain{
			Type:       "file",
			URL:        "file.jpg",
			FullURL:    "",
			UserUpload: "1",
		}
		fileRepository.On("Store", mock.Anything, &domainStore).Return(domain, nil).Once()
		minioRepository.On("GetFile", mock.AnythingOfType("string")).Return("https://s3.hungrybaby.com/file.jpg", nil).Once()
		result, err := fileUsecase.Upload(ctx, domain.Type, "", fileHeader)

		assert.Equal(t, result, domain)
		assert.Nil(t, err)
	})

	t.Run("test case 2, minio error", func(t *testing.T) {
		var (
			fileRepository  fileMock.Repository
			minioRepository minioMock.Repository
		)
		fileUsecase := file.NewFileUsecase(2, &fileRepository, &minioRepository)

		errError := errors.New("(Minio) Error")
		var fileHeader *multipart.FileHeader
		domain := file.Domain{
			Type:       "file",
			URL:        "file.jpg",
			FullURL:    "https://s3.hungrybaby.com/file.jpg",
			UserUpload: "1",
		}
		ctx := context.WithValue(context.Background(), "userID", 1)
		minioRepository.On("Upload", mock.AnythingOfType("string"), mock.Anything).Return("", errError).Once()
		domainStore := file.Domain{
			Type:       "file",
			URL:        "file.jpg",
			FullURL:    "",
			UserUpload: "1",
		}
		fileRepository.On("Store", mock.Anything, &domainStore).Return(domainStore, nil).Once()
		minioRepository.On("GetFile", mock.AnythingOfType("string")).Return("https://s3.hungrybaby.com/file.jpg", nil).Once()
		result, err := fileUsecase.Upload(ctx, domain.Type, "", fileHeader)

		assert.Equal(t, result, file.Domain{})
		assert.Equal(t, err, errError)
	})
}

func TestDelete(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		var (
			fileRepository  fileMock.Repository
			minioRepository minioMock.Repository
		)
		fileUsecase := file.NewFileUsecase(2, &fileRepository, &minioRepository)

		domain := file.Domain{
			ID:         1,
			Type:       "file",
			URL:        "file.jpg",
			FullURL:    "https://s3.hungrybaby.com/file.jpg",
			UserUpload: "1",
		}
		ctx := context.WithValue(context.Background(), "userID", 1)
		fileRepository.On("FindByID", mock.Anything, mock.AnythingOfType("int")).Return(domain, nil).Once()
		minioRepository.On("Delete", mock.AnythingOfType("string")).Return(nil).Once()
		fileRepository.On("Delete", mock.Anything, mock.AnythingOfType("*file.Domain")).Return(domain, nil).Once()
		result, err := fileUsecase.Delete(ctx, &domain)

		assert.Equal(t, result, domain)
		assert.Nil(t, err)
	})

	t.Run("test case 2, repo error", func(t *testing.T) {
		var (
			fileRepository  fileMock.Repository
			minioRepository minioMock.Repository
		)
		fileUsecase := file.NewFileUsecase(2, &fileRepository, &minioRepository)

		errError := errors.New("(Repo) Error")
		domain := file.Domain{
			ID:         1,
			Type:       "file",
			URL:        "file.jpg",
			FullURL:    "https://s3.hungrybaby.com/file.jpg",
			UserUpload: "1",
		}
		ctx := context.WithValue(context.Background(), "userID", 1)
		fileRepository.On("FindByID", mock.Anything, mock.AnythingOfType("int")).Return(file.Domain{}, errError).Once()
		minioRepository.On("Delete", mock.AnythingOfType("string")).Return(nil).Once()
		fileRepository.On("Delete", mock.Anything, mock.AnythingOfType("*file.Domain")).Return(domain, nil).Once()
		result, err := fileUsecase.Delete(ctx, &domain)

		assert.Equal(t, result, file.Domain{})
		assert.Equal(t, err, errError)
	})

	t.Run("test case 3, repository error", func(t *testing.T) {
		var (
			fileRepository  fileMock.Repository
			minioRepository minioMock.Repository
		)
		fileUsecase := file.NewFileUsecase(2, &fileRepository, &minioRepository)

		errError := errors.New("(Minio) Error")
		domain := file.Domain{
			ID:         1,
			Type:       "file",
			URL:        "file.jpg",
			FullURL:    "https://s3.hungrybaby.com/file.jpg",
			UserUpload: "1",
		}
		ctx := context.WithValue(context.Background(), "userID", 1)
		fileRepository.On("FindByID", mock.Anything, mock.AnythingOfType("int")).Return(domain, nil).Once()
		minioRepository.On("Delete", mock.AnythingOfType("string")).Return(errError).Once()
		fileRepository.On("Delete", mock.Anything, mock.AnythingOfType("*file.Domain")).Return(domain, nil).Once()
		result, err := fileUsecase.Delete(ctx, &domain)

		assert.Equal(t, result, file.Domain{})
		assert.Equal(t, err, errError)
	})

	t.Run("test case 4, repository delete error", func(t *testing.T) {
		var (
			fileRepository  fileMock.Repository
			minioRepository minioMock.Repository
		)
		fileUsecase := file.NewFileUsecase(2, &fileRepository, &minioRepository)

		errError := errors.New("(Repo) Delete Error")
		domain := file.Domain{
			ID:         1,
			Type:       "file",
			URL:        "file.jpg",
			FullURL:    "https://s3.hungrybaby.com/file.jpg",
			UserUpload: "1",
		}
		ctx := context.WithValue(context.Background(), "userID", 1)
		fileRepository.On("FindByID", mock.Anything, mock.AnythingOfType("int")).Return(domain, nil).Once()
		minioRepository.On("Delete", mock.AnythingOfType("string")).Return(nil).Once()
		fileRepository.On("Delete", mock.Anything, mock.AnythingOfType("*file.Domain")).Return(domain, errError).Once()
		result, err := fileUsecase.Delete(ctx, &domain)

		assert.Equal(t, result, file.Domain{})
		assert.Equal(t, err, errError)
	})
}
