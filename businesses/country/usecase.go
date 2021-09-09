package country

import (
	"context"
	"hungry-baby/businesses"
	"strings"
	"time"
)

type countryUsecase struct {
	countryRepository Repository
	contextTimeout    time.Duration
}

func NewCountryUsecase(timeout time.Duration, repo Repository) Usecase {
	return &countryUsecase{
		countryRepository: repo,
		contextTimeout:    timeout,
	}
}

func (uc *countryUsecase) FindAll(ctx context.Context, search, status string) ([]Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	res, err := uc.countryRepository.FindAll(ctx, search, status)
	if err != nil {
		return []Domain{}, err
	}

	return res, nil
}

func (uc *countryUsecase) Find(ctx context.Context, search, status string, page, perpage int) ([]Domain, int, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	if page <= 0 {
		page = 1
	}
	if perpage <= 0 {
		perpage = 25
	}

	res, total, err := uc.countryRepository.Find(ctx, search, status, page, perpage)
	if err != nil {
		return []Domain{}, 0, err
	}

	return res, total, nil
}

func (uc *countryUsecase) FindByID(ctx context.Context, countryId int, status string) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	if countryId <= 0 {
		return Domain{}, businesses.ErrIDNotFound
	}
	res, err := uc.countryRepository.FindByID(ctx, countryId, status)
	if err != nil {
		return Domain{}, err
	}

	return res, nil
}

func (uc *countryUsecase) FindByCode(ctx context.Context, code, status string) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	res, err := uc.countryRepository.FindByCode(ctx, code, status)
	if err != nil {
		return Domain{}, err
	}

	return res, nil
}

func (uc *countryUsecase) Store(ctx context.Context, countryDomain *Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	exist, err := uc.countryRepository.FindByCode(ctx, countryDomain.CountryCode, "")
	if err != nil {
		if !strings.Contains(err.Error(), "not found") {
			return Domain{}, err
		}
	}
	if exist != (Domain{}) {
		return Domain{}, businesses.ErrDuplicateData
	}

	result, err := uc.countryRepository.Store(ctx, countryDomain)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (uc *countryUsecase) Update(ctx context.Context, countryDomain *Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	exist, err := uc.countryRepository.FindByCode(ctx, countryDomain.CountryCode, "")
	if err != nil {
		if !strings.Contains(err.Error(), "not found") {
			return Domain{}, err
		}
	}
	if exist != (Domain{}) && exist.ID != countryDomain.ID {
		return Domain{}, businesses.ErrDuplicateData
	}

	result, err := uc.countryRepository.Update(ctx, countryDomain)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (uc *countryUsecase) Delete(ctx context.Context, countryDomain *Domain) (Domain, error) {
	existedCountry, err := uc.countryRepository.FindByID(ctx, countryDomain.ID, "true")
	if err != nil {
		return Domain{}, err
	}
	countryDomain.ID = existedCountry.ID

	result, err := uc.countryRepository.Delete(ctx, countryDomain)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}
