package databases

import (
	countryDomain "hungry-baby/businesses/country"
	countryDB "hungry-baby/drivers/databases/country"

	fileDomain "hungry-baby/businesses/file"
	fileDB "hungry-baby/drivers/databases/file"

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
