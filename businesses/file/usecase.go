package file

import (
	"context"
	"hungry-baby/businesses"
	minioBusiness "hungry-baby/businesses/minio"
	"mime/multipart"
	"strconv"
	"time"
)

type fileUsecase struct {
	fileRepository  Repository
	contextTimeout  time.Duration
	minioRepository minioBusiness.Repository
}

func NewFileUsecase(timeout time.Duration, repo Repository, minioRepo minioBusiness.Repository) Usecase {
	return &fileUsecase{
		fileRepository:  repo,
		contextTimeout:  timeout,
		minioRepository: minioRepo,
	}
}

func (uc *fileUsecase) FindByID(ctx context.Context, fileId int) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	if fileId <= 0 {
		return Domain{}, businesses.ErrFileIDResource
	}
	res, err := uc.fileRepository.FindByID(ctx, fileId)
	if err != nil {
		return Domain{}, err
	}

	// Get temporary url
	res.FullURL, err = uc.minioRepository.GetFile(res.URL)
	if err != nil {
		return Domain{}, err
	}

	return res, nil
}

func (uc *fileUsecase) Store(ctx context.Context, fileDomain *Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	fileDomain.UserUpload = strconv.Itoa(ctx.Value("userID").(int))
	result, err := uc.fileRepository.Store(ctx, fileDomain)
	if err != nil {
		return Domain{}, err
	}

	// Get temporary url
	result.FullURL, err = uc.minioRepository.GetFile(result.URL)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (uc *fileUsecase) Upload(ctx context.Context, types, filePath string, f *multipart.FileHeader) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	// Upload file to minio
	minioURL, err := uc.minioRepository.Upload(filePath, f)
	if err != nil {
		return Domain{}, err
	}

	fileDomain := Domain{
		Type:       types,
		URL:        minioURL,
		UserUpload: strconv.Itoa(ctx.Value("userID").(int)),
	}
	result, err := uc.Store(ctx, &fileDomain)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (uc *fileUsecase) Delete(ctx context.Context, fileDomain *Domain) (Domain, error) {
	existedFile, err := uc.fileRepository.FindByID(ctx, fileDomain.ID)
	if err != nil {
		return Domain{}, err
	}
	fileDomain.ID = existedFile.ID

	// Delete from minio
	err = uc.minioRepository.Delete(existedFile.URL)
	if err != nil {
		return Domain{}, err
	}

	result, err := uc.fileRepository.Delete(ctx, fileDomain)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}
