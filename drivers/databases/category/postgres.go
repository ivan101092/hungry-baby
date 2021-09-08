package category

import (
	"context"
	"hungry-baby/businesses/category"

	"gorm.io/gorm"
)

type PostgresRepository struct {
	conn *gorm.DB
}

//NewPostgresRepository we need this to work around the repository test
func NewPostgresRepository(conn *gorm.DB) *PostgresRepository {
	return &PostgresRepository{
		conn: conn,
	}
}

func (cr *PostgresRepository) Find(ctx context.Context, active string) ([]category.Domain, error) {
	rec := []Category{}

	query := cr.conn.Debug().Where("archive = ?", false)

	if active != "" {
		if active == "false" {
			query = query.Where("active = ?", false)
		} else {
			query = query.Where("active = ?", true)
		}
	}

	err := query.Find(&rec).Error
	if err != nil {
		return []category.Domain{}, err
	}

	categoryDomain := []category.Domain{}
	for _, value := range rec {
		categoryDomain = append(categoryDomain, value.ToDomain())
	}

	return categoryDomain, nil
}

func (cr *PostgresRepository) FindByID(id int) (category.Domain, error) {
	rec := Category{}

	if err := cr.conn.Where("id = ?", id).First(&rec).Error; err != nil {
		return category.Domain{}, err
	}
	return rec.ToDomain(), nil
}
