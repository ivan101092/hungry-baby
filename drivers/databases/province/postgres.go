package province

import (
	"context"
	"hungry-baby/businesses"
	"hungry-baby/businesses/province"
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

func (cr *PostgresRepository) FindAll(ctx context.Context, search string, countryID int, status string) ([]province.Domain, error) {
	rec := []Province{}

	query := cr.conn.Debug()
	if search != "" {
		query = query.Where("LOWER(code) LIKE ? AND LOWER(name) LIKE ?",
			`%`+strings.ToLower(search)+`%`, `%`+strings.ToLower(search)+`%`)
	}
	if countryID != 0 {
		query = query.Where("country_id = ?", countryID)
	}
	if str.CheckBool(status) {
		query = query.Where("status = ?", status)
	}
	err := query.Find(&rec).Error
	if err != nil {
		return []province.Domain{}, err
	}

	provinceDomain := []province.Domain{}
	for _, value := range rec {
		provinceDomain = append(provinceDomain, value.ToDomain())
	}

	return provinceDomain, nil
}

func (cr *PostgresRepository) Find(ctx context.Context, search string, countryID int, status string, page, perpage int) ([]province.Domain, int, error) {
	rec := []Province{}

	offset := (page - 1) * perpage
	query := cr.conn.Debug()
	if search != "" {
		query = query.Where("LOWER(code) LIKE ? AND LOWER(name) LIKE ?",
			`%`+strings.ToLower(search)+`%`, `%`+strings.ToLower(search)+`%`)
	}
	if countryID != 0 {
		query = query.Where("country_id = ?", countryID)
	}
	if str.CheckBool(status) {
		query = query.Where("status = ?", status)
	}
	err := query.Offset(offset).Limit(perpage).Find(&rec).Error
	if err != nil {
		return []province.Domain{}, 0, err
	}

	var totalData int64
	err = cr.conn.Model(&Province{}).Count(&totalData).Error
	if err != nil {
		return []province.Domain{}, 0, err
	}

	var domainProvince []province.Domain
	for _, value := range rec {
		domainProvince = append(domainProvince, value.ToDomain())
	}
	return domainProvince, int(totalData), nil
}

func (cr *PostgresRepository) FindByID(ctx context.Context, id int, status string) (province.Domain, error) {
	rec := Province{}

	query := cr.conn.Debug()
	if str.CheckBool(status) {
		query = query.Where("status = ?", status)
	}
	if err := query.Where("id = ?", id).First(&rec).Error; err != nil {
		return province.Domain{}, err
	}
	return rec.ToDomain(), nil
}

func (cr *PostgresRepository) FindByCode(ctx context.Context, code, status string) (province.Domain, error) {
	rec := Province{}

	query := cr.conn.Debug()
	if str.CheckBool(status) {
		query = query.Where("status = ?", status)
	}
	if err := query.Where("code = ?", code).First(&rec).Error; err != nil {
		return province.Domain{}, err
	}
	return rec.ToDomain(), nil
}

func (cr *PostgresRepository) Store(ctx context.Context, provinceDomain *province.Domain) (province.Domain, error) {
	rec := FromDomain(provinceDomain)

	result := cr.conn.Create(&rec)
	if result.Error != nil {
		return province.Domain{}, result.Error
	}

	err := cr.conn.First(&rec, rec.ID).Error
	if err != nil {
		return province.Domain{}, err
	}

	return rec.ToDomain(), nil
}

func (cr *PostgresRepository) Update(ctx context.Context, provinceDomain *province.Domain) (province.Domain, error) {
	rec := FromDomain(provinceDomain)

	result := cr.conn.Updates(&rec)
	if result.Error != nil {
		return province.Domain{}, result.Error
	}
	if result.RowsAffected == 0 {
		return province.Domain{}, businesses.ErrIDNotFound
	}

	err := cr.conn.First(&rec, rec.ID).Error
	if err != nil {
		return province.Domain{}, err
	}

	return rec.ToDomain(), nil
}

func (cr *PostgresRepository) Delete(ctx context.Context, provinceDomain *province.Domain) (province.Domain, error) {
	rec := FromDomain(provinceDomain)

	result := cr.conn.Where("id", rec.ID).Delete(&rec)
	if result.Error != nil {
		return province.Domain{}, result.Error
	}

	return rec.ToDomain(), nil
}
