package userChildMeal

import (
	"context"
	"fmt"
	"hungry-baby/businesses"
	"hungry-baby/businesses/userChildMeal"
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

func (cr *PostgresRepository) FindAll(ctx context.Context, search string, userChildID int) ([]userChildMeal.Domain, error) {
	rec := []UserChildMeal{}

	query := cr.conn.Debug().Table("user_child_meals")
	if search != "" {
		query = query.Where("LOWER(name) LIKE ?", `%`+strings.ToLower(search)+`%`)
	}
	if userChildID != 0 {
		query = query.Where("user_child_id = ?", userChildID)
	}
	err := query.Find(&rec).Error
	if err != nil {
		return []userChildMeal.Domain{}, err
	}

	userChildMealDomain := []userChildMeal.Domain{}
	for _, value := range rec {
		userChildMealDomain = append(userChildMealDomain, value.ToDomain())
	}

	return userChildMealDomain, nil
}

func (cr *PostgresRepository) Find(ctx context.Context, search string, userChildID, page, perpage int) ([]userChildMeal.Domain, int, error) {
	rec := []UserChildMeal{}

	offset := (page - 1) * perpage
	query := cr.conn.Debug().Table("user_child_meals")
	if search != "" {
		query = query.Where("LOWER(name) LIKE ?", `%`+strings.ToLower(search)+`%`)
	}
	if userChildID != 0 {
		query = query.Where("user_child_id = ?", userChildID)
	}
	err := query.Offset(offset).Limit(perpage).Find(&rec).Error
	if err != nil {
		return []userChildMeal.Domain{}, 0, err
	}

	var totalData int64
	err = cr.conn.Table("user_child_meals").Count(&totalData).Error
	if err != nil {
		return []userChildMeal.Domain{}, 0, err
	}

	var domainUserChildMeal []userChildMeal.Domain
	for _, value := range rec {
		domainUserChildMeal = append(domainUserChildMeal, value.ToDomain())
	}
	return domainUserChildMeal, int(totalData), nil
}

func (cr *PostgresRepository) FindByID(ctx context.Context, id int) (userChildMeal.Domain, error) {
	rec := UserChildMeal{}

	query := cr.conn.Debug().Table("user_child_meals")
	if err := query.Where("id = ?", id).First(&rec).Error; err != nil {
		return userChildMeal.Domain{}, err
	}
	return rec.ToDomain(), nil
}

func (cr *PostgresRepository) FindByChildMeal(ctx context.Context, userChildID, mealPlanID int) (userChildMeal.Domain, error) {
	rec := UserChildMeal{}

	query := cr.conn.Debug().Table("user_child_meals")
	if err := query.Where("user_child_id = ? AND meal_plan_id = ?", userChildID, mealPlanID).First(&rec).Error; err != nil {
		return userChildMeal.Domain{}, err
	}
	return rec.ToDomain(), nil
}

func (cr *PostgresRepository) Store(ctx context.Context, userChildMealDomain *userChildMeal.Domain) (userChildMeal.Domain, error) {
	rec := FromDomain(userChildMealDomain)

	result := cr.conn.Table("user_child_meals").Create(&rec)
	if result.Error != nil {
		return userChildMeal.Domain{}, result.Error
	}

	err := cr.conn.Table("user_child_meals").First(&rec, rec.ID).Error
	if err != nil {
		return userChildMeal.Domain{}, err
	}

	return rec.ToDomain(), nil
}

func (cr *PostgresRepository) Update(ctx context.Context, userChildMealDomain *userChildMeal.Domain) (userChildMeal.Domain, error) {
	rec := FromDomain(userChildMealDomain)
	fmt.Println("calendar", rec.CalendarID)

	result := cr.conn.Table("user_child_meals").Updates(&rec)
	if result.Error != nil {
		return userChildMeal.Domain{}, result.Error
	}
	if result.RowsAffected == 0 {
		return userChildMeal.Domain{}, businesses.ErrIDNotFound
	}

	err := cr.conn.Table("user_child_meals").First(&rec, rec.ID).Error
	if err != nil {
		return userChildMeal.Domain{}, err
	}

	return rec.ToDomain(), nil
}

func (cr *PostgresRepository) Delete(ctx context.Context, userChildMealDomain *userChildMeal.Domain) (userChildMeal.Domain, error) {
	rec := FromDomain(userChildMealDomain)

	result := cr.conn.Table("user_child_meals").Where("id", rec.ID).Delete(&rec)
	if result.Error != nil {
		return userChildMeal.Domain{}, result.Error
	}

	return rec.ToDomain(), nil
}
