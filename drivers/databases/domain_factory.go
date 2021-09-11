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
