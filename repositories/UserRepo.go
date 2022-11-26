package repositories

import (
	"gorm.io/gorm"
	"shopping_service_go/models"
)

type UserModel struct {
	Db *gorm.DB
}

func (m *UserModel) Create(user models.User) error {
	result := m.Db.Create(&user)
	return result.Error
}

func (m *UserModel) Find(user models.User) error {
	result := m.Db.Find(&user)
	return result.Error
}
