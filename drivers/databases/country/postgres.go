package country

import (
	"context"
	"hungry-baby/businesses/country"
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

func (cr *PostgresRepository) FindAll(ctx context.Context, search, status string) ([]country.Domain, error) {
	rec := []Country{}

	query := cr.conn.Debug()
	if search != "" {
		query = query.Where("LOWER(country_code) LIKE ? AND LOWER(name) LIKE ?",
			`%`+strings.ToLower(search)+`%`, `%`+strings.ToLower(search)+`%`)
	}
	if str.CheckBool(status) {
		query = query.Where("status = ?", status)
	}
	err := query.Find(&rec).Error
	if err != nil {
		return []country.Domain{}, err
	}

	countryDomain := []country.Domain{}
	for _, value := range rec {
		countryDomain = append(countryDomain, value.ToDomain())
	}

	return countryDomain, nil
}

func (cr *PostgresRepository) Find(ctx context.Context, search, status string, page, perpage int) ([]country.Domain, int, error) {
	rec := []Country{}

	offset := (page - 1) * perpage
	query := cr.conn.Debug()
	if search != "" {
		query = query.Where("LOWER(country_code) LIKE ? AND LOWER(name) LIKE ?",
			`%`+strings.ToLower(search)+`%`, `%`+strings.ToLower(search)+`%`)
	}
	if str.CheckBool(status) {
		query = query.Where("status = ?", status)
	}
	err := query.Offset(offset).Limit(perpage).Find(&rec).Error
	if err != nil {
		return []country.Domain{}, 0, err
	}

	var totalData int64
	err = cr.conn.Model(&Country{}).Count(&totalData).Error
	if err != nil {
		return []country.Domain{}, 0, err
	}

	var domainCountry []country.Domain
	for _, value := range rec {
		domainCountry = append(domainCountry, value.ToDomain())
	}
	return domainCountry, int(totalData), nil
}

func (cr *PostgresRepository) FindByID(ctx context.Context, id int, status string) (country.Domain, error) {
	rec := Country{}

	query := cr.conn.Debug()
	if str.CheckBool(status) {
		query = query.Where("status = ?", status)
	}
	if err := query.Where("id = ?", id).First(&rec).Error; err != nil {
		return country.Domain{}, err
	}
	return rec.ToDomain(), nil
}

func (cr *PostgresRepository) FindByCode(ctx context.Context, code, status string) (country.Domain, error) {
	rec := Country{}

	query := cr.conn.Debug()
	if str.CheckBool(status) {
		query = query.Where("status = ?", status)
	}
	if err := query.Where("country_code = ?", code).First(&rec).Error; err != nil {
		return country.Domain{}, err
	}
	return rec.ToDomain(), nil
}

func (cr *PostgresRepository) Store(ctx context.Context, countryDomain *country.Domain) (country.Domain, error) {
	rec := FromDomain(countryDomain)

	result := cr.conn.Create(&rec)
	if result.Error != nil {
		return country.Domain{}, result.Error
	}

	err := cr.conn.First(&rec, rec.ID).Error
	if err != nil {
		return country.Domain{}, result.Error
	}

	return rec.ToDomain(), nil
}

func (cr *PostgresRepository) Update(ctx context.Context, countryDomain *country.Domain) (country.Domain, error) {
	rec := FromDomain(countryDomain)

	result := cr.conn.Save(&rec)
	if result.Error != nil {
		return country.Domain{}, result.Error
	}

	err := cr.conn.First(&rec, rec.ID).Error
	if err != nil {
		return country.Domain{}, result.Error
	}

	return rec.ToDomain(), nil
}

func (cr *PostgresRepository) Delete(ctx context.Context, countryDomain *country.Domain) (country.Domain, error) {
	rec := FromDomain(countryDomain)

	result := cr.conn.Where("id", rec.ID).Delete(&rec)
	if result.Error != nil {
		return country.Domain{}, result.Error
	}

	return rec.ToDomain(), nil
}
