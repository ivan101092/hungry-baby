package city

import (
	"context"
	"hungry-baby/businesses"
	"strings"
	"time"
)

type cityUsecase struct {
	cityRepository Repository
	contextTimeout time.Duration
}

func NewCityUsecase(timeout time.Duration, repo Repository) Usecase {
	return &cityUsecase{
		cityRepository: repo,
		contextTimeout: timeout,
	}
}

func (uc *cityUsecase) FindAll(ctx context.Context, search string, provinceID int, status string) ([]Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	res, err := uc.cityRepository.FindAll(ctx, search, provinceID, status)
	if err != nil {
		return []Domain{}, err
	}

	return res, nil
}

func (uc *cityUsecase) Find(ctx context.Context, search string, provinceID int, status string, page, perpage int) ([]Domain, int, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	if page <= 0 {
		page = 1
	}
	if perpage <= 0 {
		perpage = 25
	}

	res, total, err := uc.cityRepository.Find(ctx, search, provinceID, status, page, perpage)
	if err != nil {
		return []Domain{}, 0, err
	}

	return res, total, nil
}

func (uc *cityUsecase) FindByID(ctx context.Context, cityId int, status string) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	if cityId <= 0 {
		return Domain{}, businesses.ErrIDNotFound
	}
	res, err := uc.cityRepository.FindByID(ctx, cityId, status)
	if err != nil {
		return Domain{}, err
	}

	return res, nil
}

func (uc *cityUsecase) FindByCode(ctx context.Context, code, status string) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	res, err := uc.cityRepository.FindByCode(ctx, code, status)
	if err != nil {
		return Domain{}, err
	}

	return res, nil
}

func (uc *cityUsecase) Store(ctx context.Context, cityDomain *Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	exist, err := uc.cityRepository.FindByCode(ctx, cityDomain.Code, "")
	if err != nil {
		if !strings.Contains(err.Error(), "not found") {
			return Domain{}, err
		}
	}
	if exist != (Domain{}) {
		return Domain{}, businesses.ErrDuplicateData
	}

	result, err := uc.cityRepository.Store(ctx, cityDomain)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (uc *cityUsecase) Update(ctx context.Context, cityDomain *Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	exist, err := uc.cityRepository.FindByCode(ctx, cityDomain.Code, "")
	if err != nil {
		if !strings.Contains(err.Error(), "not found") {
			return Domain{}, err
		}
	}
	if exist != (Domain{}) && exist.ID != cityDomain.ID {
		return Domain{}, businesses.ErrDuplicateData
	}

	result, err := uc.cityRepository.Update(ctx, cityDomain)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (uc *cityUsecase) Delete(ctx context.Context, cityDomain *Domain) (Domain, error) {
	existedCity, err := uc.cityRepository.FindByID(ctx, cityDomain.ID, "true")
	if err != nil {
		return Domain{}, err
	}
	cityDomain.ID = existedCity.ID

	result, err := uc.cityRepository.Delete(ctx, cityDomain)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}
