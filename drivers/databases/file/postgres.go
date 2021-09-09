package file

import (
	"context"
	"hungry-baby/businesses/file"

	"gorm.io/gorm"
)

type PostgresRepository struct {
	conn *gorm.DB
}

// NewPostgresRepository we need this to work around the repository test
func NewPostgresRepository(conn *gorm.DB) *PostgresRepository {
	return &PostgresRepository{
		conn: conn,
	}
}

func (cr *PostgresRepository) FindByID(ctx context.Context, id int) (file.Domain, error) {
	rec := File{}

	if err := cr.conn.Where("id = ?", id).First(&rec).Error; err != nil {
		return file.Domain{}, err
	}
	return rec.ToDomain(), nil
}

func (cr *PostgresRepository) Store(ctx context.Context, fileDomain *file.Domain) (file.Domain, error) {
	rec := FromDomain(fileDomain)

	result := cr.conn.Create(&rec)
	if result.Error != nil {
		return file.Domain{}, result.Error
	}

	err := cr.conn.First(&rec, rec.ID).Error
	if err != nil {
		return file.Domain{}, result.Error
	}

	return rec.ToDomain(), nil
}

func (cr *PostgresRepository) Delete(ctx context.Context, fileDomain *file.Domain) (file.Domain, error) {
	rec := FromDomain(fileDomain)

	result := cr.conn.Where("id", rec.ID).Delete(&rec)
	if result.Error != nil {
		return file.Domain{}, result.Error
	}

	return rec.ToDomain(), nil
}
