package user

import (
	"context"
	"hungry-baby/businesses"
	"hungry-baby/drivers/minio"
	"strings"
	"time"
)

type userUsecase struct {
	userRepository  Repository
	contextTimeout  time.Duration
	minioRepository minio.IMinio
}

func NewUserUsecase(timeout time.Duration, repo Repository, minioRepo minio.IMinio) Usecase {
	return &userUsecase{
		userRepository:  repo,
		contextTimeout:  timeout,
		minioRepository: minioRepo,
	}
}

func (uc *userUsecase) FindAll(ctx context.Context, search string, status string) ([]Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	res, err := uc.userRepository.FindAll(ctx, search, status)
	if err != nil {
		return []Domain{}, err
	}

	return res, nil
}

func (uc *userUsecase) Find(ctx context.Context, search string, status string, page, perpage int) ([]Domain, int, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	if page <= 0 {
		page = 1
	}
	if perpage <= 0 {
		perpage = 25
	}

	res, total, err := uc.userRepository.Find(ctx, search, status, page, perpage)
	if err != nil {
		return []Domain{}, 0, err
	}

	return res, total, nil
}

func (uc *userUsecase) FindByID(ctx context.Context, id int, status string) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	if id <= 0 {
		return Domain{}, businesses.ErrIDNotFound
	}
	res, err := uc.userRepository.FindByID(ctx, id, status)
	if err != nil {
		return Domain{}, err
	}

	if res.ProfileImageURL != "" {
		res.ProfileImageURL, err = uc.minioRepository.GetFile(res.ProfileImageURL)
		if err != nil {
			return Domain{}, err
		}
	}

	return res, nil
}

func (uc *userUsecase) FindByCode(ctx context.Context, code, status string) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	res, err := uc.userRepository.FindByCode(ctx, code, status)
	if err != nil {
		return Domain{}, err
	}

	return res, nil
}

func (uc *userUsecase) FindByEmail(ctx context.Context, email, status string) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	res, err := uc.userRepository.FindByEmail(ctx, email, status)
	if err != nil {
		return Domain{}, err
	}

	return res, nil
}

func (uc *userUsecase) Store(ctx context.Context, userDomain *Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	exist, err := uc.userRepository.FindByEmail(ctx, userDomain.Email, "")
	if err != nil {
		if !strings.Contains(err.Error(), "not found") {
			return Domain{}, err
		}
	}
	if exist.ID != 0 {
		return Domain{}, businesses.ErrDuplicateData
	}

	result, err := uc.userRepository.Store(ctx, userDomain)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (uc *userUsecase) Update(ctx context.Context, userDomain *Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	exist, err := uc.userRepository.FindByID(ctx, userDomain.ID, "")
	if err != nil {
		return Domain{}, err
	}
	userDomain.Code = exist.Code
	userDomain.Email = exist.Email

	result, err := uc.userRepository.Update(ctx, userDomain)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (uc *userUsecase) Delete(ctx context.Context, userDomain *Domain) (Domain, error) {
	existedUser, err := uc.userRepository.FindByID(ctx, userDomain.ID, "true")
	if err != nil {
		return Domain{}, err
	}
	userDomain.ID = existedUser.ID

	result, err := uc.userRepository.Delete(ctx, userDomain)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}
