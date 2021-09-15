package userChildMeal

import (
	"context"
	"errors"
	"hungry-baby/businesses"
	calendarBusiness "hungry-baby/businesses/calendar"
	mealPlanBusiness "hungry-baby/businesses/mealPlan"
	userBusiness "hungry-baby/businesses/user"
	userChildBusiness "hungry-baby/businesses/userChild"
	"time"
)

type userChildMealUsecase struct {
	userChildMealRepository Repository
	userChildUsecase        userChildBusiness.Usecase
	mealPlanUsecase         mealPlanBusiness.Usecase
	calendarUsecase         calendarBusiness.Usecase
	userUsecase             userBusiness.Usecase
	contextTimeout          time.Duration
}

func NewUserChildMealUsecase(timeout time.Duration, repo Repository,
	userChildUsecase userChildBusiness.Usecase, mealPlanUsecase mealPlanBusiness.Usecase,
	calendarUsecase calendarBusiness.Usecase, userUsecase userBusiness.Usecase) Usecase {
	return &userChildMealUsecase{
		userChildMealRepository: repo,
		userChildUsecase:        userChildUsecase,
		mealPlanUsecase:         mealPlanUsecase,
		calendarUsecase:         calendarUsecase,
		userUsecase:             userUsecase,
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

	if userChildMealDomain.ScheduledAt != "" {
		scheduleAt, err := time.Parse(time.RFC3339, userChildMealDomain.ScheduledAt)
		if err != nil {
			return Domain{}, err
		}
		if scheduleAt.After(time.Now()) {
			userChildMealDomain.Status = "pending"
		} else {
			userChildMealDomain.Status = "done"
		}
	}
	if userChildMealDomain.Status == "pending" {
		userChildMealDomain.FinishAt = ""
	}

	if userChildMealDomain.FinishAt != "" {
		finishAt, err := time.Parse(time.RFC3339, userChildMealDomain.FinishAt)
		if err != nil {
			return Domain{}, err
		}
		scheduleAt, err := time.Parse(time.RFC3339, userChildMealDomain.ScheduledAt)
		if err != nil {
			return Domain{}, err
		}
		if finishAt.Before(scheduleAt) {
			return Domain{}, errors.New("Invalid Finish")
		}
	} else if userChildMealDomain.Status != "pending" {
		userChildMealDomain.FinishAt = time.Now().Format(time.RFC3339)
	}

	if userChildMealDomain.Status == "pending" {
		calendar, err := uc.calendarUsecase.Store(ctx, &calendarBusiness.Domain{
			Title:   mealPlan.Name,
			StartAt: userChildMealDomain.ScheduledAt,
			EndAt:   userChildMealDomain.ScheduledAt,
		})
		if err != nil {
			return Domain{}, err
		}
		userChildMealDomain.CalendarID = calendar.ID
	}

	result, err := uc.userChildMealRepository.Store(ctx, userChildMealDomain)
	if err != nil {
		return Domain{}, err
	}

	if userChildMealDomain.Status == "done" {
		nextPending, err := uc.userChildMealRepository.FindNextPending(ctx, userChildMealDomain.UserChildID, userChildMealDomain.MealPlanID)
		if nextPending.ID != 0 {
			return Domain{}, nil
		}

		user, err := uc.userUsecase.FindByID(ctx, ctx.Value("userID").(int), "")
		if err != nil {
			return Domain{}, err
		}

		if user.Settings.AutoNotification {
			finishAt, err := time.Parse(time.RFC3339, userChildMealDomain.FinishAt)
			if err != nil {
				return Domain{}, err
			}

			var scheduleAt time.Time
			if finishAt.Before(time.Now()) && finishAt.Add(time.Duration(mealPlan.Interval)*time.Minute).After(time.Now()) {
				scheduleAt = finishAt.Add(time.Duration(mealPlan.Interval) * time.Minute)
			} else {
				scheduleAt = time.Now().Add(time.Duration(mealPlan.Interval) * time.Minute)
			}

			_, err = uc.userChildMealRepository.Store(ctx, &Domain{
				UserID:             userChildMealDomain.UserID,
				UserChildID:        userChildMealDomain.UserChildID,
				MealPlanID:         userChildMealDomain.MealPlanID,
				Name:               userChildMealDomain.Name,
				SuggestionQuantity: userChildMealDomain.SuggestionQuantity,
				Quantity:           0,
				Unit:               userChildMealDomain.Unit,
				ScheduledAt:        scheduleAt.Format(time.RFC3339),
			})
			if err != nil {
				return Domain{}, err
			}
		}
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

	exist, err := uc.userChildMealRepository.FindByID(ctx, userChildMealDomain.ID)
	if err != nil {
		return Domain{}, err
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

	if userChildMealDomain.ScheduledAt != "" {
		scheduleAt, err := time.Parse(time.RFC3339, userChildMealDomain.ScheduledAt)
		if err != nil {
			return Domain{}, err
		}
		if scheduleAt.After(time.Now()) {
			userChildMealDomain.Status = "pending"
		} else {
			userChildMealDomain.Status = "done"
		}
	}
	if userChildMealDomain.Status == "pending" {
		userChildMealDomain.FinishAt = ""
	}

	if userChildMealDomain.FinishAt != "" {
		finishAt, err := time.Parse(time.RFC3339, userChildMealDomain.FinishAt)
		if err != nil {
			return Domain{}, err
		}
		scheduleAt, err := time.Parse(time.RFC3339, userChildMealDomain.ScheduledAt)
		if err != nil {
			return Domain{}, err
		}
		if finishAt.Before(scheduleAt) {
			return Domain{}, errors.New("Invalid Finish")
		}
	} else if userChildMealDomain.Status != "pending" {
		userChildMealDomain.FinishAt = time.Now().Format(time.RFC3339)
	}

	if userChildMealDomain.Status == "pending" && exist.ScheduledAt != userChildMealDomain.ScheduledAt {
		existScheduleAt, err := time.Parse(time.RFC3339, exist.ScheduledAt)
		if err != nil {
			return Domain{}, err
		}
		scheduleAt, err := time.Parse(time.RFC3339, userChildMealDomain.ScheduledAt)
		if err != nil {
			return Domain{}, err
		}
		if !existScheduleAt.Equal(scheduleAt) {
			if exist.CalendarID != "" {
				uc.calendarUsecase.Delete(ctx, exist.CalendarID)
			}

			calendar, err := uc.calendarUsecase.Store(ctx, &calendarBusiness.Domain{
				Title:   mealPlan.Name,
				StartAt: userChildMealDomain.ScheduledAt,
				EndAt:   userChildMealDomain.ScheduledAt,
			})
			if err != nil {
				return Domain{}, err
			}
			userChildMealDomain.CalendarID = calendar.ID
		}
	}
	if userChildMealDomain.CalendarID == "" {
		userChildMealDomain.CalendarID = exist.CalendarID
	}

	if userChildMealDomain.Status == "done" && exist.CalendarID != "" {
		existScheduleAt, err := time.Parse(time.RFC3339, exist.ScheduledAt)
		if err != nil {
			return Domain{}, err
		}
		if existScheduleAt.After(time.Now()) {
			uc.calendarUsecase.Delete(ctx, exist.CalendarID)
			userChildMealDomain.CalendarID = ""
		}
	}

	result, err := uc.userChildMealRepository.Update(ctx, userChildMealDomain)
	if err != nil {
		return Domain{}, err
	}

	if userChildMealDomain.Status == "done" {
		nextPending, err := uc.userChildMealRepository.FindNextPending(ctx, userChildMealDomain.UserChildID, userChildMealDomain.MealPlanID)
		if nextPending.ID != 0 {
			return Domain{}, nil
		}

		user, err := uc.userUsecase.FindByID(ctx, ctx.Value("userID").(int), "")
		if err != nil {
			return Domain{}, err
		}

		if user.Settings.AutoNotification {
			finishAt, err := time.Parse(time.RFC3339, userChildMealDomain.FinishAt)
			if err != nil {
				return Domain{}, err
			}

			var scheduleAt time.Time
			if finishAt.Before(time.Now()) && finishAt.Add(time.Duration(mealPlan.Interval)*time.Minute).After(time.Now()) {
				scheduleAt = finishAt.Add(time.Duration(mealPlan.Interval) * time.Minute)
			} else {
				scheduleAt = time.Now().Add(time.Duration(mealPlan.Interval) * time.Minute)
			}

			_, err = uc.userChildMealRepository.Store(ctx, &Domain{
				UserID:             userChildMealDomain.UserID,
				UserChildID:        userChildMealDomain.UserChildID,
				MealPlanID:         userChildMealDomain.MealPlanID,
				Name:               userChildMealDomain.Name,
				SuggestionQuantity: userChildMealDomain.SuggestionQuantity,
				Quantity:           0,
				Unit:               userChildMealDomain.Unit,
				ScheduledAt:        scheduleAt.Format(time.RFC3339),
			})
			if err != nil {
				return Domain{}, err
			}
		}
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
