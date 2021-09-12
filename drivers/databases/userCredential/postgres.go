package userCredential

import (
	"context"
	"hungry-baby/businesses"
	"hungry-baby/businesses/userCredential"
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

func (cr *PostgresRepository) FindAll(ctx context.Context, search string, status string) ([]userCredential.Domain, error) {
	rec := []UserCredential{}

	query := cr.conn.Debug().Table("user_credentials")
	if search != "" {
		query = query.Where("LOWER(code) LIKE ? AND LOWER(name) LIKE ?",
			`%`+strings.ToLower(search)+`%`, `%`+strings.ToLower(search)+`%`)
	}
	if str.CheckBool(status) {
		query = query.Where("status = ?", status)
	}
	err := query.Find(&rec).Error
	if err != nil {
		return []userCredential.Domain{}, err
	}

	userCredentialDomain := []userCredential.Domain{}
	for _, value := range rec {
		userCredentialDomain = append(userCredentialDomain, value.ToDomain())
	}

	return userCredentialDomain, nil
}

func (cr *PostgresRepository) Find(ctx context.Context, search string, status string, page, perpage int) ([]userCredential.Domain, int, error) {
	rec := []UserCredential{}

	offset := (page - 1) * perpage
	query := cr.conn.Debug().Table("user_credentials")
	if search != "" {
		query = query.Where("LOWER(code) LIKE ? AND LOWER(name) LIKE ?",
			`%`+strings.ToLower(search)+`%`, `%`+strings.ToLower(search)+`%`)
	}
	if str.CheckBool(status) {
		query = query.Where("status = ?", status)
	}
	err := query.Offset(offset).Limit(perpage).Find(&rec).Error
	if err != nil {
		return []userCredential.Domain{}, 0, err
	}

	var totalData int64
	err = cr.conn.Table("user_credentials").Count(&totalData).Error
	if err != nil {
		return []userCredential.Domain{}, 0, err
	}

	var domainUserCredential []userCredential.Domain
	for _, value := range rec {
		domainUserCredential = append(domainUserCredential, value.ToDomain())
	}
	return domainUserCredential, int(totalData), nil
}

func (cr *PostgresRepository) FindByID(ctx context.Context, id int, status string) (userCredential.Domain, error) {
	rec := UserCredential{}

	query := cr.conn.Debug().Table("user_credentials")
	query = query.Select("userCredentials.*, c.name as city_name, f.url as profile_image_url")
	query = query.Joins("LEFT JOIN cities c ON c.id = userCredentials.city_id")
	query = query.Joins("LEFT JOIN files f ON f.id = userCredentials.profile_image_id")
	if str.CheckBool(status) {
		query = query.Where("userCredentials.status = ?", status)
	}
	if err := query.Where("userCredentials.id = ?", id).First(&rec).Error; err != nil {
		return userCredential.Domain{}, err
	}
	return rec.ToDomain(), nil
}

func (cr *PostgresRepository) FindByEmail(ctx context.Context, email, status string) (userCredential.Domain, error) {
	rec := UserCredential{}

	query := cr.conn.Debug().Table("user_credentials")
	if str.CheckBool(status) {
		query = query.Where("status = ?", status)
	}
	if err := query.Where("LOWER(email) = ?", strings.ToLower(email)).First(&rec).Error; err != nil {
		return userCredential.Domain{}, err
	}
	return rec.ToDomain(), nil
}

func (cr *PostgresRepository) FindByUserType(ctx context.Context, userID int, types, status string) (userCredential.Domain, error) {
	rec := UserCredential{}

	query := cr.conn.Debug().Table("user_credentials")
	if str.CheckBool(status) {
		query = query.Where("status = ?", status)
	}
	if userID != 0 {
		query = query.Where("user_id = ?", userID)
	}
	if types != "" {
		query = query.Where("type = ?", types)
	}
	if err := query.First(&rec).Error; err != nil {
		return userCredential.Domain{}, err
	}
	return rec.ToDomain(), nil
}

func (cr *PostgresRepository) Store(ctx context.Context, userCredentialDomain *userCredential.Domain) (userCredential.Domain, error) {
	rec := FromDomain(userCredentialDomain)

	result := cr.conn.Table("user_credentials").Create(&rec)
	if result.Error != nil {
		return userCredential.Domain{}, result.Error
	}

	err := cr.conn.Table("user_credentials").First(&rec, rec.ID).Error
	if err != nil {
		return userCredential.Domain{}, err
	}

	return rec.ToDomain(), nil
}

func (cr *PostgresRepository) Update(ctx context.Context, userCredentialDomain *userCredential.Domain) (userCredential.Domain, error) {
	rec := FromDomain(userCredentialDomain)

	result := cr.conn.Table("user_credentials").Updates(&rec)
	if result.Error != nil {
		return userCredential.Domain{}, result.Error
	}
	if result.RowsAffected == 0 {
		return userCredential.Domain{}, businesses.ErrIDNotFound
	}

	err := cr.conn.Table("user_credentials").First(&rec, rec.ID).Error
	if err != nil {
		return userCredential.Domain{}, err
	}

	return rec.ToDomain(), nil
}

func (cr *PostgresRepository) Delete(ctx context.Context, userCredentialDomain *userCredential.Domain) (userCredential.Domain, error) {
	rec := FromDomain(userCredentialDomain)

	result := cr.conn.Table("user_credentials").Where("id", rec.ID).Delete(&rec)
	if result.Error != nil {
		return userCredential.Domain{}, result.Error
	}

	return rec.ToDomain(), nil
}
