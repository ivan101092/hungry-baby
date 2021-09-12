package mealPlan

import (
	"context"
	"hungry-baby/businesses"
	"time"
)

type mealPlanUsecase struct {
	mealPlanRepository Repository
	contextTimeout     time.Duration
}

func NewMealPlanUsecase(timeout time.Duration, repo Repository) Usecase {
	return &mealPlanUsecase{
		mealPlanRepository: repo,
		contextTimeout:     timeout,
	}
}

func (uc *mealPlanUsecase) FindAll(ctx context.Context, search string, userID int, status string) ([]Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	res, err := uc.mealPlanRepository.FindAll(ctx, search, userID, status)
	if err != nil {
		return []Domain{}, err
	}

	return res, nil
}

func (uc *mealPlanUsecase) Find(ctx context.Context, search string, userID int, status string, page, perpage int) ([]Domain, int, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	if page <= 0 {
		page = 1
	}
	if perpage <= 0 {
		perpage = 25
	}

	res, total, err := uc.mealPlanRepository.Find(ctx, search, userID, status, page, perpage)
	if err != nil {
		return []Domain{}, 0, err
	}

	return res, total, nil
}

func (uc *mealPlanUsecase) FindByID(ctx context.Context, mealPlanId int, status string) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	if mealPlanId <= 0 {
		return Domain{}, businesses.ErrIDNotFound
	}
	res, err := uc.mealPlanRepository.FindByID(ctx, mealPlanId, status)
	if err != nil {
		return Domain{}, err
	}

	return res, nil
}

func (uc *mealPlanUsecase) Store(ctx context.Context, mealPlanDomain *Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	result, err := uc.mealPlanRepository.Store(ctx, mealPlanDomain)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (uc *mealPlanUsecase) Update(ctx context.Context, mealPlanDomain *Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	result, err := uc.mealPlanRepository.Update(ctx, mealPlanDomain)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (uc *mealPlanUsecase) Delete(ctx context.Context, mealPlanDomain *Domain) (Domain, error) {
	existedMealPlan, err := uc.mealPlanRepository.FindByID(ctx, mealPlanDomain.ID, "true")
	if err != nil {
		return Domain{}, err
	}
	mealPlanDomain.ID = existedMealPlan.ID

	result, err := uc.mealPlanRepository.Delete(ctx, mealPlanDomain)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}
