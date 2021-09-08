package postgres

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	CreatedAt time.Time      `json:"created_at,omitempty" gorm:"column:created_at;<-:false"`
	UpdatedAt time.Time      `json:"updated_at,omitempty" gorm:"column:updated_at;<-:false"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"column:deleted_at;<-:false"`
}

type BasePagination struct {
	Page      int    `query:"page"`
	Limit     int    `query:"limit"`
	Sort      string `validate:"omitempty,oneof=id createdAt" json:"sort" query:"sort"`
	Direction string `validate:"omitempty,oneof=ASC DESC" json:"direction" query:"direction"`
}
