package user

type User struct {
	ID         int32  `gorm:"primaryKey"`
	FullName   string `gorm:"column:full_name"`
	UserName   string `gorm:"column:user_name"`
	EmailId    string `gorm:"column:email_id"`
	Password   string `gorm:"column:password"`
	CreatedAt  int64  `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt  int64  `gorm:"column:updated_at;autoUpdateTime:milli"`
	Last_login int64  `gorm:"column:last_login"`
	IsDeleted  bool   `gorm:"column:is_deleted"`
}
