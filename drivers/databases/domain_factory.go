package databases

import (
	fileDomain "hungry-baby/businesses/file"
	fileDB "hungry-baby/drivers/databases/file"

	countryDomain "hungry-baby/businesses/country"
	countryDB "hungry-baby/drivers/databases/country"

	provinceDomain "hungry-baby/businesses/province"
	provinceDB "hungry-baby/drivers/databases/province"

	cityDomain "hungry-baby/businesses/city"
	cityDB "hungry-baby/drivers/databases/city"

	userDomain "hungry-baby/businesses/user"
	userDB "hungry-baby/drivers/databases/user"

	userCredentialDomain "hungry-baby/businesses/userCredential"
	userCredentialDB "hungry-baby/drivers/databases/userCredential"

	mealPlanDomain "hungry-baby/businesses/mealPlan"
	mealPlanDB "hungry-baby/drivers/databases/mealPlan"

	userChildDomain "hungry-baby/businesses/userChild"
	userChildDB "hungry-baby/drivers/databases/userChild"

	userChildMealDomain "hungry-baby/businesses/userChildMeal"
	userChildMealDB "hungry-baby/drivers/databases/userChildMeal"

	"gorm.io/gorm"
)

//NewFileRepository Factory with country domain
func NewFileRepository(conn *gorm.DB) fileDomain.Repository {
	return fileDB.NewPostgresRepository(conn)
}

//NewCountryRepository Factory with country domain
func NewCountryRepository(conn *gorm.DB) countryDomain.Repository {
	return countryDB.NewPostgresRepository(conn)
}

//NewProvinceRepository Factory with province domain
func NewProvinceRepository(conn *gorm.DB) provinceDomain.Repository {
	return provinceDB.NewPostgresRepository(conn)
}

//NewCityRepository Factory with city domain
func NewCityRepository(conn *gorm.DB) cityDomain.Repository {
	return cityDB.NewPostgresRepository(conn)
}

//NewUserRepository Factory with user domain
func NewUserRepository(conn *gorm.DB) userDomain.Repository {
	return userDB.NewPostgresRepository(conn)
}

//NewUserCredentialRepository Factory with user domain
func NewUserCredentialRepository(conn *gorm.DB) userCredentialDomain.Repository {
	return userCredentialDB.NewPostgresRepository(conn)
}

//NewMealPlanRepository Factory with user domain
func NewMealPlanRepository(conn *gorm.DB) mealPlanDomain.Repository {
	return mealPlanDB.NewPostgresRepository(conn)
}

//NewUserChildRepository Factory with user domain
func NewUserChildRepository(conn *gorm.DB) userChildDomain.Repository {
	return userChildDB.NewPostgresRepository(conn)
}

//NewUserChildMealRepository Factory with user domain
func NewUserChildMealRepository(conn *gorm.DB) userChildMealDomain.Repository {
	return userChildMealDB.NewPostgresRepository(conn)
}
