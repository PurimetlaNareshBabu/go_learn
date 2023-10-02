package company

type Company struct {
	ID        int32 `gorm:"primaryKey"`
	CreatedBy int32 `gorm:"column:created_by"`
	CreatedAt int64 `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt int64 `gorm:"column:updated_at;autoUpdateTime:milli"`
	IsDeleted bool  `gorm:"column:is_deleted"`
}

type Role struct {
	ID        int32   `gorm:"primaryKey"`
	CreatedBy int32   `gorm:"column:created_by"`
	CreatedAt int64   `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt int64   `gorm:"column:updated_at;autoUpdateTime:milli"`
	IsDeleted bool    `gorm:"column:is_deleted"`
	CompanyID uint    `gorm:"column:company_id;not null;index;foreignKey:company_id"`
	Company   Company `gorm:"foreignKey:company_id"`
}
