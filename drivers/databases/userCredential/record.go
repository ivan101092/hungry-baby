package userCredential

import (
	"hungry-baby/businesses/userCredential"
	userCredentialUsecase "hungry-baby/businesses/userCredential"
	"hungry-baby/drivers/postgres"
	"hungry-baby/helpers/interfacepkg"
)

type UserCredential struct {
	ID                  int
	UserID              int
	Type                string
	Email               string
	EmailValidAt        string
	Password            string
	RegistrationDetails string
	Status              bool
	postgres.BaseModel
}

func FromDomain(domain *userCredentialUsecase.Domain) *UserCredential {
	return &UserCredential{
		ID:                  domain.ID,
		UserID:              domain.UserID,
		Type:                domain.Type,
		Email:               domain.Email,
		EmailValidAt:        domain.EmailValidAt,
		Password:            domain.Password,
		RegistrationDetails: interfacepkg.Marshal(domain.RegistrationDetails),
		Status:              domain.Status,
	}
}

func (rec *UserCredential) ToDomain() userCredential.Domain {
	var registrationDetails userCredential.DomainRegistrationDetails
	interfacepkg.UnmarshalCb(rec.RegistrationDetails, &registrationDetails)

	return userCredential.Domain{
		ID:                  rec.ID,
		UserID:              rec.UserID,
		Type:                rec.Type,
		Email:               rec.Email,
		EmailValidAt:        rec.EmailValidAt,
		Password:            rec.Password,
		RegistrationDetails: registrationDetails,
		Status:              rec.Status,
	}
}
