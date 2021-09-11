package city

import (
	"context"
	"hungry-baby/businesses"
	"hungry-baby/businesses/city"
	"hungry-baby/helpers/str"
	"strings"

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

func (cr *PostgresRepository) FindAll(ctx context.Context, search string, countryID int, status string) ([]city.Domain, error) {
	rec := []City{}

	query := cr.conn.Debug()
	if search != "" {
		query = query.Where("LOWER(code) LIKE ? AND LOWER(name) LIKE ?",
			`%`+strings.ToLower(search)+`%`, `%`+strings.ToLower(search)+`%`)
	}
	if str.CheckBool(status) {
		query = query.Where("status = ?", status)
	}
	err := query.Find(&rec).Error
	if err != nil {
		return []city.Domain{}, err
	}

	cityDomain := []city.Domain{}
	for _, value := range rec {
		cityDomain = append(cityDomain, value.ToDomain())
	}

	return cityDomain, nil
}

func (cr *PostgresRepository) Find(ctx context.Context, search string, countryID int, status string, page, perpage int) ([]city.Domain, int, error) {
	rec := []City{}

	offset := (page - 1) * perpage
	query := cr.conn.Debug()
	if search != "" {
		query = query.Where("LOWER(code) LIKE ? AND LOWER(name) LIKE ?",
			`%`+strings.ToLower(search)+`%`, `%`+strings.ToLower(search)+`%`)
	}
	if str.CheckBool(status) {
		query = query.Where("status = ?", status)
	}
	err := query.Offset(offset).Limit(perpage).Find(&rec).Error
	if err != nil {
		return []city.Domain{}, 0, err
	}

	var totalData int64
	err = cr.conn.Model(&City{}).Count(&totalData).Error
	if err != nil {
		return []city.Domain{}, 0, err
	}

	var domainCity []city.Domain
	for _, value := range rec {
		domainCity = append(domainCity, value.ToDomain())
	}
	return domainCity, int(totalData), nil
}

func (cr *PostgresRepository) FindByID(ctx context.Context, id int, status string) (city.Domain, error) {
	rec := City{}

	query := cr.conn.Debug()
	if str.CheckBool(status) {
		query = query.Where("status = ?", status)
	}
	if err := query.Where("id = ?", id).First(&rec).Error; err != nil {
		return city.Domain{}, err
	}
	return rec.ToDomain(), nil
}

func (cr *PostgresRepository) FindByCode(ctx context.Context, code, status string) (city.Domain, error) {
	rec := City{}

	query := cr.conn.Debug()
	if str.CheckBool(status) {
		query = query.Where("status = ?", status)
	}
	if err := query.Where("code = ?", code).First(&rec).Error; err != nil {
		return city.Domain{}, err
	}
	return rec.ToDomain(), nil
}

func (cr *PostgresRepository) Store(ctx context.Context, cityDomain *city.Domain) (city.Domain, error) {
	rec := FromDomain(cityDomain)

	result := cr.conn.Create(&rec)
	if result.Error != nil {
		return city.Domain{}, result.Error
	}

	err := cr.conn.First(&rec, rec.ID).Error
	if err != nil {
		return city.Domain{}, err
	}

	return rec.ToDomain(), nil
}

func (cr *PostgresRepository) Update(ctx context.Context, cityDomain *city.Domain) (city.Domain, error) {
	rec := FromDomain(cityDomain)

	result := cr.conn.Updates(&rec)
	if result.Error != nil {
		return city.Domain{}, result.Error
	}
	if result.RowsAffected == 0 {
		return city.Domain{}, businesses.ErrIDNotFound
	}

	err := cr.conn.First(&rec, rec.ID).Error
	if err != nil {
		return city.Domain{}, err
	}

	return rec.ToDomain(), nil
}

func (cr *PostgresRepository) Delete(ctx context.Context, cityDomain *city.Domain) (city.Domain, error) {
	rec := FromDomain(cityDomain)

	result := cr.conn.Where("id", rec.ID).Delete(&rec)
	if result.Error != nil {
		return city.Domain{}, result.Error
	}

	return rec.ToDomain(), nil
}
