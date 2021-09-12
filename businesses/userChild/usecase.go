package userChild

import (
	"context"
	"hungry-baby/businesses"
	"time"
)

type userChildUsecase struct {
	userChildRepository Repository
	contextTimeout      time.Duration
}

func NewUserChildUsecase(timeout time.Duration, repo Repository) Usecase {
	return &userChildUsecase{
		userChildRepository: repo,
		contextTimeout:      timeout,
	}
}

func (uc *userChildUsecase) FindAll(ctx context.Context, search string, userID int) ([]Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	res, err := uc.userChildRepository.FindAll(ctx, search, userID)
	if err != nil {
		return []Domain{}, err
	}

	return res, nil
}

func (uc *userChildUsecase) Find(ctx context.Context, search string, userID int, page, perpage int) ([]Domain, int, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	if page <= 0 {
		page = 1
	}
	if perpage <= 0 {
		perpage = 25
	}

	res, total, err := uc.userChildRepository.Find(ctx, search, userID, page, perpage)
	if err != nil {
		return []Domain{}, 0, err
	}

	return res, total, nil
}

func (uc *userChildUsecase) FindByID(ctx context.Context, userChildId int) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	if userChildId <= 0 {
		return Domain{}, businesses.ErrIDNotFound
	}
	res, err := uc.userChildRepository.FindByID(ctx, userChildId)
	if err != nil {
		return Domain{}, err
	}

	return res, nil
}

func (uc *userChildUsecase) Store(ctx context.Context, userChildDomain *Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	result, err := uc.userChildRepository.Store(ctx, userChildDomain)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (uc *userChildUsecase) Update(ctx context.Context, userChildDomain *Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	result, err := uc.userChildRepository.Update(ctx, userChildDomain)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (uc *userChildUsecase) Delete(ctx context.Context, userChildDomain *Domain) (Domain, error) {
	existedUserChild, err := uc.userChildRepository.FindByID(ctx, userChildDomain.ID)
	if err != nil {
		return Domain{}, err
	}
	userChildDomain.ID = existedUserChild.ID

	result, err := uc.userChildRepository.Delete(ctx, userChildDomain)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}
