package user

import (
	"errors"
	"myapp/pkg/db"

	"gorm.io/gorm"
)

type userdao struct {
	db *gorm.DB
}

func New(dbClient *db.Client) *userdao {
	return &userdao{
		db: dbClient.DB,
	}
}

func (d *userdao) Createuser(newuser *User) (int32, error) {
	err := d.db.Create(newuser)
	if err != nil {
		return 0, errors.New("failed to create user: " + err.Error.Error())
	}
	return newuser.ID, nil
}

func (d *userdao) Updateuser(user_id int32, data map[string]interface{}) error {
	err := d.db.Model(&User{}).Where("id = ?", user_id).Updates(data)
	if err != nil {
		return errors.New("failed to create user: " + err.Error.Error())
	}
	return nil
}
