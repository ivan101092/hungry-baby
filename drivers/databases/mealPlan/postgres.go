package mealPlan

import (
	"context"
	"hungry-baby/businesses"
	"hungry-baby/businesses/mealPlan"
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

func (cr *PostgresRepository) FindAll(ctx context.Context, search string, userID int, status string) ([]mealPlan.Domain, error) {
	rec := []MealPlan{}

	query := cr.conn.Debug()
	if search != "" {
		query = query.Where("LOWER(name) LIKE ?", `%`+strings.ToLower(search)+`%`)
	}
	if userID != 0 {
		query = query.Where("user_id IN(?,?)", userID, 0)
	} else {
		query = query.Where("user_id = ?", userID)
	}
	if str.CheckBool(status) {
		query = query.Where("status = ?", status)
	}
	err := query.Find(&rec).Error
	if err != nil {
		return []mealPlan.Domain{}, err
	}

	mealPlanDomain := []mealPlan.Domain{}
	for _, value := range rec {
		mealPlanDomain = append(mealPlanDomain, value.ToDomain())
	}

	return mealPlanDomain, nil
}

func (cr *PostgresRepository) Find(ctx context.Context, search string, userID int, status string, page, perpage int) ([]mealPlan.Domain, int, error) {
	rec := []MealPlan{}

	offset := (page - 1) * perpage
	query := cr.conn.Debug()
	if search != "" {
		query = query.Where("LOWER(name) LIKE ?", `%`+strings.ToLower(search)+`%`)
	}
	if userID != 0 {
		query = query.Where("user_id IN(?,?)", userID, 0)
	} else {
		query = query.Where("user_id = ?", userID)
	}
	if str.CheckBool(status) {
		query = query.Where("status = ?", status)
	}
	err := query.Offset(offset).Limit(perpage).Find(&rec).Error
	if err != nil {
		return []mealPlan.Domain{}, 0, err
	}

	var totalData int64
	err = cr.conn.Model(&MealPlan{}).Count(&totalData).Error
	if err != nil {
		return []mealPlan.Domain{}, 0, err
	}

	var domainMealPlan []mealPlan.Domain
	for _, value := range rec {
		domainMealPlan = append(domainMealPlan, value.ToDomain())
	}
	return domainMealPlan, int(totalData), nil
}

func (cr *PostgresRepository) FindByID(ctx context.Context, id int, status string) (mealPlan.Domain, error) {
	rec := MealPlan{}

	query := cr.conn.Debug()
	if str.CheckBool(status) {
		query = query.Where("status = ?", status)
	}
	if err := query.Where("id = ?", id).First(&rec).Error; err != nil {
		return mealPlan.Domain{}, err
	}
	return rec.ToDomain(), nil
}

func (cr *PostgresRepository) Store(ctx context.Context, mealPlanDomain *mealPlan.Domain) (mealPlan.Domain, error) {
	rec := FromDomain(mealPlanDomain)

	result := cr.conn.Create(&rec)
	if result.Error != nil {
		return mealPlan.Domain{}, result.Error
	}

	err := cr.conn.First(&rec, rec.ID).Error
	if err != nil {
		return mealPlan.Domain{}, err
	}

	return rec.ToDomain(), nil
}

func (cr *PostgresRepository) Update(ctx context.Context, mealPlanDomain *mealPlan.Domain) (mealPlan.Domain, error) {
	rec := FromDomain(mealPlanDomain)

	result := cr.conn.Updates(&rec)
	if result.Error != nil {
		return mealPlan.Domain{}, result.Error
	}
	if result.RowsAffected == 0 {
		return mealPlan.Domain{}, businesses.ErrIDNotFound
	}

	err := cr.conn.First(&rec, rec.ID).Error
	if err != nil {
		return mealPlan.Domain{}, err
	}

	return rec.ToDomain(), nil
}

func (cr *PostgresRepository) Delete(ctx context.Context, mealPlanDomain *mealPlan.Domain) (mealPlan.Domain, error) {
	rec := FromDomain(mealPlanDomain)

	result := cr.conn.Where("id", rec.ID).Delete(&rec)
	if result.Error != nil {
		return mealPlan.Domain{}, result.Error
	}

	return rec.ToDomain(), nil
}
