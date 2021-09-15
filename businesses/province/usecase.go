package province

import (
	"context"
	"database/sql"
	"hungry-baby/businesses"
	"time"
)

type provinceUsecase struct {
	provinceRepository Repository
	contextTimeout     time.Duration
}

func NewProvinceUsecase(timeout time.Duration, repo Repository) Usecase {
	return &provinceUsecase{
		provinceRepository: repo,
		contextTimeout:     timeout,
	}
}

func (uc *provinceUsecase) FindAll(ctx context.Context, search string, countryID int, status string) ([]Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	res, err := uc.provinceRepository.FindAll(ctx, search, countryID, status)
	if err != nil {
		return []Domain{}, err
	}

	return res, nil
}

func (uc *provinceUsecase) Find(ctx context.Context, search string, countryID int, status string, page, perpage int) ([]Domain, int, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	if page <= 0 {
		page = 1
	}
	if perpage <= 0 {
		perpage = 25
	}

	res, total, err := uc.provinceRepository.Find(ctx, search, countryID, status, page, perpage)
	if err != nil {
		return []Domain{}, 0, err
	}

	return res, total, nil
}

func (uc *provinceUsecase) FindByID(ctx context.Context, provinceId int, status string) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	if provinceId <= 0 {
		return Domain{}, businesses.ErrIDNotFound
	}
	res, err := uc.provinceRepository.FindByID(ctx, provinceId, status)
	if err != nil {
		return Domain{}, err
	}

	return res, nil
}

func (uc *provinceUsecase) FindByCode(ctx context.Context, code, status string) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	res, err := uc.provinceRepository.FindByCode(ctx, code, status)
	if err != nil {
		return Domain{}, err
	}

	return res, nil
}

func (uc *provinceUsecase) Store(ctx context.Context, provinceDomain *Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	exist, err := uc.provinceRepository.FindByCode(ctx, provinceDomain.Code, "")
	if err != nil {
		if err != sql.ErrNoRows && err.Error() != "record not found" {
			return Domain{}, err
		}
	}
	if exist != (Domain{}) {
		return Domain{}, businesses.ErrDuplicateData
	}

	result, err := uc.provinceRepository.Store(ctx, provinceDomain)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (uc *provinceUsecase) Update(ctx context.Context, provinceDomain *Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	exist, err := uc.provinceRepository.FindByCode(ctx, provinceDomain.Code, "")
	if err != nil {
		if err != sql.ErrNoRows && err.Error() != "record not found" {
			return Domain{}, err
		}
	}
	if exist != (Domain{}) && exist.ID != provinceDomain.ID {
		return Domain{}, businesses.ErrDuplicateData
	}

	result, err := uc.provinceRepository.Update(ctx, provinceDomain)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (uc *provinceUsecase) Delete(ctx context.Context, provinceDomain *Domain) (Domain, error) {
	existedProvince, err := uc.provinceRepository.FindByID(ctx, provinceDomain.ID, "true")
	if err != nil {
		return Domain{}, err
	}
	provinceDomain.ID = existedProvince.ID

	result, err := uc.provinceRepository.Delete(ctx, provinceDomain)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}
