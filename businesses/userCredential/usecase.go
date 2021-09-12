package userCredential

import (
	"context"
	"hungry-baby/businesses"
	"strings"
	"time"
)

type userCredentialUsecase struct {
	userCredentialRepository Repository
	contextTimeout           time.Duration
}

func NewUserCredentialUsecase(timeout time.Duration, repo Repository) Usecase {
	return &userCredentialUsecase{
		userCredentialRepository: repo,
		contextTimeout:           timeout,
	}
}

func (uc *userCredentialUsecase) FindAll(ctx context.Context, search string, status string) ([]Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	res, err := uc.userCredentialRepository.FindAll(ctx, search, status)
	if err != nil {
		return []Domain{}, err
	}

	return res, nil
}

func (uc *userCredentialUsecase) Find(ctx context.Context, search string, status string, page, perpage int) ([]Domain, int, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	if page <= 0 {
		page = 1
	}
	if perpage <= 0 {
		perpage = 25
	}

	res, total, err := uc.userCredentialRepository.Find(ctx, search, status, page, perpage)
	if err != nil {
		return []Domain{}, 0, err
	}

	return res, total, nil
}

func (uc *userCredentialUsecase) FindByID(ctx context.Context, id int, status string) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	if id <= 0 {
		return Domain{}, businesses.ErrIDNotFound
	}
	res, err := uc.userCredentialRepository.FindByID(ctx, id, status)
	if err != nil {
		return Domain{}, err
	}

	return res, nil
}

func (uc *userCredentialUsecase) FindByCode(ctx context.Context, code, status string) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	res, err := uc.userCredentialRepository.FindByCode(ctx, code, status)
	if err != nil {
		return Domain{}, err
	}

	return res, nil
}

func (uc *userCredentialUsecase) FindByEmail(ctx context.Context, email, status string) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	res, err := uc.userCredentialRepository.FindByEmail(ctx, email, status)
	if err != nil {
		return Domain{}, err
	}

	return res, nil
}

func (uc *userCredentialUsecase) Store(ctx context.Context, userCredentialDomain *Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	exist, err := uc.userCredentialRepository.FindByEmail(ctx, userCredentialDomain.Email, "")
	if err != nil {
		if !strings.Contains(err.Error(), "not found") {
			return Domain{}, err
		}
	}
	if exist != (Domain{}) {
		return Domain{}, businesses.ErrDuplicateData
	}

	result, err := uc.userCredentialRepository.Store(ctx, userCredentialDomain)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (uc *userCredentialUsecase) Update(ctx context.Context, userCredentialDomain *Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	exist, err := uc.userCredentialRepository.FindByEmail(ctx, userCredentialDomain.Email, "")
	if err != nil {
		if !strings.Contains(err.Error(), "not found") {
			return Domain{}, err
		}
	}
	if exist != (Domain{}) && exist.ID != userCredentialDomain.ID {
		return Domain{}, businesses.ErrDuplicateData
	}

	result, err := uc.userCredentialRepository.Update(ctx, userCredentialDomain)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (uc *userCredentialUsecase) Delete(ctx context.Context, userCredentialDomain *Domain) (Domain, error) {
	existedUserCredential, err := uc.userCredentialRepository.FindByID(ctx, userCredentialDomain.ID, "true")
	if err != nil {
		return Domain{}, err
	}
	userCredentialDomain.ID = existedUserCredential.ID

	result, err := uc.userCredentialRepository.Delete(ctx, userCredentialDomain)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}
