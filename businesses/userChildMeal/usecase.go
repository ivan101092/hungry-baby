package userChildMeal

import (
	"context"
	"errors"
	"hungry-baby/businesses"
	mealPlanBusiness "hungry-baby/businesses/mealPlan"
	userChildBusiness "hungry-baby/businesses/userChild"
	"time"
)

type userChildMealUsecase struct {
	userChildMealRepository Repository
	userChildUsecase        userChildBusiness.Usecase
	mealPlanUsecase         mealPlanBusiness.Usecase
	contextTimeout          time.Duration
}

func NewUserChildMealUsecase(timeout time.Duration, repo Repository,
	userChildUsecase userChildBusiness.Usecase, mealPlanUsecase mealPlanBusiness.Usecase) Usecase {
	return &userChildMealUsecase{
		userChildMealRepository: repo,
		userChildUsecase:        userChildUsecase,
		mealPlanUsecase:         mealPlanUsecase,
		contextTimeout:          timeout,
	}
}

func (uc *userChildMealUsecase) FindAll(ctx context.Context, search string, userChildID int) ([]Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	res, err := uc.userChildMealRepository.FindAll(ctx, search, userChildID)
	if err != nil {
		return []Domain{}, err
	}

	return res, nil
}

func (uc *userChildMealUsecase) Find(ctx context.Context, search string, userChildID int, page, perpage int) ([]Domain, int, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	if page <= 0 {
		page = 1
	}
	if perpage <= 0 {
		perpage = 25
	}

	res, total, err := uc.userChildMealRepository.Find(ctx, search, userChildID, page, perpage)
	if err != nil {
		return []Domain{}, 0, err
	}

	return res, total, nil
}

func (uc *userChildMealUsecase) FindByID(ctx context.Context, userChildMealId int) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	if userChildMealId <= 0 {
		return Domain{}, businesses.ErrIDNotFound
	}
	res, err := uc.userChildMealRepository.FindByID(ctx, userChildMealId)
	if err != nil {
		return Domain{}, err
	}

	return res, nil
}

func (uc *userChildMealUsecase) FindByChildMeal(ctx context.Context, userChildID, mealPlanID int) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	res, err := uc.userChildMealRepository.FindByChildMeal(ctx, userChildID, mealPlanID)
	if err != nil {
		return Domain{}, err
	}

	return res, nil
}

func (uc *userChildMealUsecase) Store(ctx context.Context, userChildMealDomain *Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	userChild, err := uc.userChildUsecase.FindByID(ctx, userChildMealDomain.UserChildID)
	if err != nil {
		return Domain{}, err
	}
	if userChild.UserID != userChildMealDomain.UserID {
		return Domain{}, errors.New("Invalid access")
	}

	mealPlan, err := uc.mealPlanUsecase.FindByID(ctx, userChildMealDomain.MealPlanID, "true")
	if err != nil {
		return Domain{}, err
	}
	userChildMealDomain.Name = mealPlan.Name
	userChildMealDomain.SuggestionQuantity = mealPlan.SuggestionQuantity
	userChildMealDomain.Unit = mealPlan.Unit

	if userChildMealDomain.FinishAt != "" {
		finishAt, err := time.Parse(time.RFC3339, userChildMealDomain.FinishAt)
		if err != nil {
			return Domain{}, err
		}
		if finishAt.After(time.Now()) {
			userChildMealDomain.Status = "pending"
		} else {
			userChildMealDomain.Status = "done"
		}
	}

	result, err := uc.userChildMealRepository.Store(ctx, userChildMealDomain)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (uc *userChildMealUsecase) Update(ctx context.Context, userChildMealDomain *Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	if userChildMealDomain.Status == "done" && userChildMealDomain.Quantity == 0 {
		return Domain{}, errors.New("Invalid quantity")
	} else if userChildMealDomain.Status == "done" && userChildMealDomain.FinishAt == "" {
		return Domain{}, errors.New("Invalid finish at")
	}

	userChild, err := uc.userChildUsecase.FindByID(ctx, userChildMealDomain.UserChildID)
	if err != nil {
		return Domain{}, err
	}
	if userChild.UserID != userChildMealDomain.UserID {
		return Domain{}, errors.New("Invalid access")
	}

	mealPlan, err := uc.mealPlanUsecase.FindByID(ctx, userChildMealDomain.MealPlanID, "true")
	if err != nil {
		return Domain{}, err
	}
	userChildMealDomain.Name = mealPlan.Name
	userChildMealDomain.SuggestionQuantity = mealPlan.SuggestionQuantity
	userChildMealDomain.Unit = mealPlan.Unit

	if userChildMealDomain.FinishAt != "" {
		finishAt, err := time.Parse(time.RFC3339, userChildMealDomain.FinishAt)
		if err != nil {
			return Domain{}, err
		}
		if finishAt.After(time.Now()) {
			userChildMealDomain.Status = "pending"
		} else {
			userChildMealDomain.Status = "done"
		}
	}

	result, err := uc.userChildMealRepository.Update(ctx, userChildMealDomain)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (uc *userChildMealUsecase) Delete(ctx context.Context, userChildMealDomain *Domain) (Domain, error) {
	existedUserChildMeal, err := uc.userChildMealRepository.FindByID(ctx, userChildMealDomain.ID)
	if err != nil {
		return Domain{}, err
	}

	userChild, err := uc.userChildUsecase.FindByID(ctx, existedUserChildMeal.UserChildID)
	if err != nil {
		return Domain{}, err
	}
	if userChild.UserID != userChildMealDomain.UserID {
		return Domain{}, errors.New("Invalid access")
	}

	result, err := uc.userChildMealRepository.Delete(ctx, userChildMealDomain)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}
