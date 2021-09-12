package userChild

import (
	"context"
	"hungry-baby/businesses"
	"hungry-baby/businesses/userChild"
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

func (cr *PostgresRepository) FindAll(ctx context.Context, search string, userID int) ([]userChild.Domain, error) {
	rec := []UserChild{}

	query := cr.conn.Debug().Table("user_childs")
	if search != "" {
		query = query.Where("LOWER(name) LIKE ?", `%`+strings.ToLower(search)+`%`)
	}
	if userID != 0 {
		query = query.Where("user_id IN(?,?)", userID, 0)
	} else {
		query = query.Where("user_id = ?", userID)
	}
	err := query.Find(&rec).Error
	if err != nil {
		return []userChild.Domain{}, err
	}

	userChildDomain := []userChild.Domain{}
	for _, value := range rec {
		userChildDomain = append(userChildDomain, value.ToDomain())
	}

	return userChildDomain, nil
}

func (cr *PostgresRepository) Find(ctx context.Context, search string, userID, page, perpage int) ([]userChild.Domain, int, error) {
	rec := []UserChild{}

	offset := (page - 1) * perpage
	query := cr.conn.Debug().Table("user_childs")
	if search != "" {
		query = query.Where("LOWER(name) LIKE ?", `%`+strings.ToLower(search)+`%`)
	}
	if userID != 0 {
		query = query.Where("user_id IN(?,?)", userID, 0)
	} else {
		query = query.Where("user_id = ?", userID)
	}
	err := query.Offset(offset).Limit(perpage).Find(&rec).Error
	if err != nil {
		return []userChild.Domain{}, 0, err
	}

	var totalData int64
	err = cr.conn.Table("user_childs").Count(&totalData).Error
	if err != nil {
		return []userChild.Domain{}, 0, err
	}

	var domainUserChild []userChild.Domain
	for _, value := range rec {
		domainUserChild = append(domainUserChild, value.ToDomain())
	}
	return domainUserChild, int(totalData), nil
}

func (cr *PostgresRepository) FindByID(ctx context.Context, id int) (userChild.Domain, error) {
	rec := UserChild{}

	query := cr.conn.Debug().Table("user_childs")
	if err := query.Where("id = ?", id).First(&rec).Error; err != nil {
		return userChild.Domain{}, err
	}
	return rec.ToDomain(), nil
}

func (cr *PostgresRepository) Store(ctx context.Context, userChildDomain *userChild.Domain) (userChild.Domain, error) {
	rec := FromDomain(userChildDomain)

	result := cr.conn.Table("user_childs").Create(&rec)
	if result.Error != nil {
		return userChild.Domain{}, result.Error
	}

	err := cr.conn.Table("user_childs").First(&rec, rec.ID).Error
	if err != nil {
		return userChild.Domain{}, err
	}

	return rec.ToDomain(), nil
}

func (cr *PostgresRepository) Update(ctx context.Context, userChildDomain *userChild.Domain) (userChild.Domain, error) {
	rec := FromDomain(userChildDomain)

	result := cr.conn.Table("user_childs").Updates(&rec)
	if result.Error != nil {
		return userChild.Domain{}, result.Error
	}
	if result.RowsAffected == 0 {
		return userChild.Domain{}, businesses.ErrIDNotFound
	}

	err := cr.conn.Table("user_childs").First(&rec, rec.ID).Error
	if err != nil {
		return userChild.Domain{}, err
	}

	return rec.ToDomain(), nil
}

func (cr *PostgresRepository) Delete(ctx context.Context, userChildDomain *userChild.Domain) (userChild.Domain, error) {
	rec := FromDomain(userChildDomain)

	result := cr.conn.Table("user_childs").Where("id", rec.ID).Delete(&rec)
	if result.Error != nil {
		return userChild.Domain{}, result.Error
	}

	return rec.ToDomain(), nil
}
