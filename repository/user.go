package repository

import (
	"errors"
	"strings"

<<<<<<< HEAD
	"github.com/Levantate-Labs/sainterview-backend/auth-service/models"
=======
	"github.com/AlfrinP/point_calculator/models"
>>>>>>> 52a2cfba8417f30f47f3a85feb3c92850e82f352
	"gorm.io/gorm"
)

type UserRepositry struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepositry {
	return &UserRepositry{db}
}

func (repo *UserRepositry) Create(u *models.User) error {
	if err := repo.db.Create(u).Error; err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique") {
			return errors.New("email already exists")
		}
		return err
	}
	return nil
}

func (repo *UserRepositry) Get(email string) (models.User, error) {
	var user models.User
	if err := repo.db.Where("email = ?", email).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (repo *UserRepositry) All() ([]models.User, error) {
	var users []models.User
	if err := repo.db.Find(&users).Error; err != nil {
		return users, err
	}
	return users, nil
}
