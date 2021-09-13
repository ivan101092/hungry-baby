package calendar

import (
	"context"
	userCredentialBusiness "hungry-baby/businesses/userCredential"
	"hungry-baby/helpers/interfacepkg"
	"time"
)

type calendarUsecase struct {
	calendarRepository    Repository
	userCredentialUsecase userCredentialBusiness.Usecase
	contextTimeout        time.Duration
}

func NewCalendarUsecase(timeout time.Duration, repo Repository, userCredentialUsecase userCredentialBusiness.Usecase) Usecase {
	return &calendarUsecase{
		calendarRepository:    repo,
		userCredentialUsecase: userCredentialUsecase,
		contextTimeout:        timeout,
	}
}

func (uc *calendarUsecase) FindAll(ctx context.Context, search, startAt, endAt, pageToken string, limit int) ([]Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	userCredential, err := uc.userCredentialUsecase.FindByUserType(ctx, ctx.Value("userID").(int), "gmail", "true")
	if err != nil {
		return nil, err
	}

	res, err := uc.calendarRepository.FindAll(ctx, interfacepkg.Marshal(userCredential.RegistrationDetails), search, startAt, endAt, pageToken, limit)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (uc *calendarUsecase) FindByID(ctx context.Context, id string) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	userCredential, err := uc.userCredentialUsecase.FindByUserType(ctx, ctx.Value("userID").(int), "gmail", "true")
	if err != nil {
		return Domain{}, err
	}

	res, err := uc.calendarRepository.FindByID(ctx, interfacepkg.Marshal(userCredential.RegistrationDetails), id)
	if err != nil {
		return Domain{}, err
	}

	return res, nil
}

func (uc *calendarUsecase) Store(ctx context.Context, calendarDomain *Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	userCredential, err := uc.userCredentialUsecase.FindByUserType(ctx, ctx.Value("userID").(int), "gmail", "true")
	if err != nil {
		return Domain{}, err
	}

	res, err := uc.calendarRepository.Add(ctx, interfacepkg.Marshal(userCredential.RegistrationDetails), calendarDomain)
	if err != nil {
		return Domain{}, err
	}

	return res, nil
}

func (uc *calendarUsecase) Delete(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	userCredential, err := uc.userCredentialUsecase.FindByUserType(ctx, ctx.Value("userID").(int), "gmail", "true")
	if err != nil {
		return err
	}

	err = uc.calendarRepository.Delete(ctx, interfacepkg.Marshal(userCredential.RegistrationDetails), id)
	if err != nil {
		return err
	}

	return nil
}
