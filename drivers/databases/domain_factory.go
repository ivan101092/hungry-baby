package databases

import (
	categoryDomain "hungry-baby/businesses/category"
	categoryDB "hungry-baby/drivers/databases/category"

	"gorm.io/gorm"
)

type Base struct {
	Offset int
	Limit  int
	Order  string
	By     string
}

//NewCategoryRepository Factory with category domain
func NewCategoryRepository(conn *gorm.DB) categoryDomain.Repository {
	return categoryDB.NewPostgresRepository(conn)
}
