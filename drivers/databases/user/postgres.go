package user

import (
	"context"
	"hungry-baby/businesses"
	"hungry-baby/businesses/user"
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

func (cr *PostgresRepository) FindAll(ctx context.Context, search string, status string) ([]user.Domain, error) {
	rec := []User{}

	query := cr.conn.Debug().Table("users")
	if search != "" {
		query = query.Where("LOWER(code) LIKE ? AND LOWER(name) LIKE ?",
			`%`+strings.ToLower(search)+`%`, `%`+strings.ToLower(search)+`%`)
	}
	if str.CheckBool(status) {
		query = query.Where("status = ?", status)
	}
	err := query.Find(&rec).Error
	if err != nil {
		return []user.Domain{}, err
	}

	userDomain := []user.Domain{}
	for _, value := range rec {
		userDomain = append(userDomain, value.ToDomain())
	}

	return userDomain, nil
}

func (cr *PostgresRepository) Find(ctx context.Context, search string, status string, page, perpage int) ([]user.Domain, int, error) {
	rec := []User{}

	offset := (page - 1) * perpage
	query := cr.conn.Debug().Table("users")
	if search != "" {
		query = query.Where("LOWER(code) LIKE ? AND LOWER(name) LIKE ?",
			`%`+strings.ToLower(search)+`%`, `%`+strings.ToLower(search)+`%`)
	}
	if str.CheckBool(status) {
		query = query.Where("status = ?", status)
	}
	err := query.Offset(offset).Limit(perpage).Find(&rec).Error
	if err != nil {
		return []user.Domain{}, 0, err
	}

	var totalData int64
	err = cr.conn.Table("users").Count(&totalData).Error
	if err != nil {
		return []user.Domain{}, 0, err
	}

	var domainUser []user.Domain
	for _, value := range rec {
		domainUser = append(domainUser, value.ToDomain())
	}
	return domainUser, int(totalData), nil
}

func (cr *PostgresRepository) FindByID(ctx context.Context, id int, status string) (user.Domain, error) {
	rec := User{}

	query := cr.conn.Debug().Table("users")
	query = query.Select("users.*, c.name as city_name, f.url as profile_image_url")
	query = query.Joins("LEFT JOIN cities c ON c.id = users.city_id")
	query = query.Joins("LEFT JOIN files f ON f.id = users.profile_image_id")
	if str.CheckBool(status) {
		query = query.Where("users.status = ?", status)
	}
	if err := query.Where("users.id = ?", id).First(&rec).Error; err != nil {
		return user.Domain{}, err
	}
	return rec.ToDomain(), nil
}

func (cr *PostgresRepository) FindByCode(ctx context.Context, code, status string) (user.Domain, error) {
	rec := User{}

	query := cr.conn.Debug().Table("users")
	if str.CheckBool(status) {
		query = query.Where("status = ?", status)
	}
	if err := query.Where("code = ?", code).First(&rec).Error; err != nil {
		return user.Domain{}, err
	}
	return rec.ToDomain(), nil
}

func (cr *PostgresRepository) FindByEmail(ctx context.Context, email, status string) (user.Domain, error) {
	rec := User{}

	query := cr.conn.Debug().Table("users")
	if str.CheckBool(status) {
		query = query.Where("status = ?", status)
	}
	if err := query.Where("LOWER(email) = ?", strings.ToLower(email)).First(&rec).Error; err != nil {
		return user.Domain{}, err
	}
	return rec.ToDomain(), nil
}

func (cr *PostgresRepository) Store(ctx context.Context, userDomain *user.Domain) (user.Domain, error) {
	rec := FromDomain(userDomain)

	result := cr.conn.Table("users").Create(&rec)
	if result.Error != nil {
		return user.Domain{}, result.Error
	}

	err := cr.conn.Table("users").First(&rec, rec.ID).Error
	if err != nil {
		return user.Domain{}, err
	}

	return rec.ToDomain(), nil
}

func (cr *PostgresRepository) Update(ctx context.Context, userDomain *user.Domain) (user.Domain, error) {
	rec := FromDomain(userDomain)

	result := cr.conn.Table("users").Updates(&rec)
	if result.Error != nil {
		return user.Domain{}, result.Error
	}
	if result.RowsAffected == 0 {
		return user.Domain{}, businesses.ErrIDNotFound
	}

	err := cr.conn.Table("users").First(&rec, rec.ID).Error
	if err != nil {
		return user.Domain{}, err
	}

	return rec.ToDomain(), nil
}

func (cr *PostgresRepository) Delete(ctx context.Context, userDomain *user.Domain) (user.Domain, error) {
	rec := FromDomain(userDomain)

	result := cr.conn.Table("users").Where("id", rec.ID).Delete(&rec)
	if result.Error != nil {
		return user.Domain{}, result.Error
	}

	return rec.ToDomain(), nil
}
