package file

import (
	"context"
	"hungry-baby/businesses"
	"time"
)

type fileUsecase struct {
	fileRepository Repository
	contextTimeout time.Duration
}

func NewFileUsecase(nr Repository, timeout time.Duration) Usecase {
	return &fileUsecase{
		fileRepository: nr,
		contextTimeout: timeout,
	}
}

func (nu *fileUsecase) FindAll(ctx context.Context, page, perpage int) ([]Domain, int, error) {
	ctx, cancel := context.WithTimeout(ctx, nu.contextTimeout)
	defer cancel()

	if page <= 0 {
		page = 1
	}
	if perpage <= 0 {
		perpage = 25
	}

	res, total, err := nu.fileRepository.FindAll(ctx, page, perpage)
	if err != nil {
		return []Domain{}, 0, err
	}

	return res, total, nil
}

func (nu *fileUsecase) Find(ctx context.Context) ([]Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, nu.contextTimeout)
	defer cancel()

	res, err := nu.fileRepository.Find(ctx)
	if err != nil {
		return []Domain{}, err
	}

	return res, nil
}

func (nu *fileUsecase) FindByID(ctx context.Context, fileId int) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, nu.contextTimeout)
	defer cancel()

	if fileId <= 0 {
		return Domain{}, businesses.ErrFileIDResource
	}
	res, err := nu.fileRepository.FindByID(ctx, fileId)
	if err != nil {
		return Domain{}, err
	}

	return res, nil
}

func (nu *fileUsecase) Store(ctx context.Context, ip string, fileDomain *Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, nu.contextTimeout)
	defer cancel()

	result, err := nu.fileRepository.Store(ctx, fileDomain)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (nu *fileUsecase) Update(ctx context.Context, fileDomain *Domain) (*Domain, error) {
	existedFile, err := nu.fileRepository.FindByID(ctx, fileDomain.ID)
	if err != nil {
		return &Domain{}, err
	}
	fileDomain.ID = existedFile.ID

	result, err := nu.fileRepository.Update(ctx, fileDomain)
	if err != nil {
		return &Domain{}, err
	}

	return &result, nil
}

func (nu *fileUsecase) Delete(ctx context.Context, fileDomain *Domain) (*Domain, error) {
	existedFile, err := nu.fileRepository.FindByID(ctx, fileDomain.ID)
	if err != nil {
		return &Domain{}, err
	}
	fileDomain.ID = existedFile.ID

	result, err := nu.fileRepository.Delete(ctx, fileDomain)
	if err != nil {
		return &Domain{}, err
	}

	return &result, nil
}
