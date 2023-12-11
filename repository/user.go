package repository

import (
	"errors"
	"strings"

	"github.com/AlfrinP/point_calculator/models"
	"gorm.io/gorm"
)

// type User interface{
// 	Get(username string)
// }

type StudentRepositry struct {
	db *gorm.DB
}

func NewStudentRepository(db *gorm.DB) *StudentRepositry {
	return &StudentRepositry{db}
}

func (repo *StudentRepositry) Create(u *models.Student) error {
	if err := repo.db.Create(u).Error; err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique") {
			return errors.New("username already exists")
		}
		return err
	}
	return nil
}

func (repo *StudentRepositry) Get(username string) (models.Student, error) {
	var student models.Student
	if err := repo.db.Where("username = ?", username).First(&student).Error; err != nil {
		return student, err
	}
	return student, nil
}

func (repo *StudentRepositry) All() ([]models.Student, error) {
	var students []models.Student
	if err := repo.db.Find(&students).Error; err != nil {
		return students, err
	}
	return students, nil
}
