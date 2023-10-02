package review

import (
	"myapp/pkg/db/company"
	"myapp/pkg/db/user"
)

type Review struct {
	ID          uint            `gorm:"primaryKey"`
	Content     string          `gorm:"column:content;not null"`
	CreatedAt   int64           `gorm:"column:created_at"`
	UpdatedAt   int64           `gorm:"column:updated_at"`
	CreatedByID uint            `gorm:"column:created_by_id;not null;index;foreignKey:created_by_id"` // Foreign key to User model's ID
	CreatedBy   user.User       `gorm:"foreignKey:created_by_id"`
	CompanyID   uint            `gorm:"column:company_id;not null;index;foreignKey:company_id"` // Foreign key to Company model's ID
	Company     company.Company `gorm:"foreignKey:company_id"`
	RoleID      uint            `gorm:"column:role_id;not null;index;foreignKey:role_id"` // Foreign key to Role model's ID
	Role        company.Role    `gorm:"foreignKey:role_id"`
	Title       string          `gorm:"column:title"`
}
