package models

import (
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Username     string        `json:"username" gorm:"unique"`
	PasswordHash string        `json:"passwordHash"`
	Department   string        `json:"department"`
	Batch        string        `json:"batch"`
	FacultyID    uint          `json:"faculty_id" gorm:"index;default:null"`
	Faculty      Faculty       `json:"faculty" gorm:"foreignKey:FacultyID"`
	Certificates []Certificate `json:"certificates" gorm:"foreignKey:StudentID"`
}

type StudentCreate struct {
	Username   string `validate:"required" json:"username"`
	Password   string `validate:"required" json:"password"`
	Department string `json:"department"`
	Batch      string `json:"batch"`
}

type StudentSignIn struct {
	Username string `validate:"required" json:"username"`
	Password string `validate:"required" json:"password"`
}

func (bc *StudentCreate) Validate() error {
	validate := validator.New()
	if err := validate.Struct(bc); err != nil {
		return err
	}
	return nil
}

func (us *StudentSignIn) Validate() error {
	validate := validator.New()
	if err := validate.Struct(us); err != nil {
		return err
	}
	return nil
}

func (bc *StudentCreate) Convert() (*Student, error) {
	hashedPasswd, err := bcrypt.GenerateFromPassword([]byte(bc.Password), bcrypt.DefaultCost)
	if err != nil {
		return &Student{}, err
	}

	return &Student{
		Username:     bc.Username,
		PasswordHash: string(hashedPasswd),
		Department:   bc.Department,
		Batch:        bc.Batch,
	}, nil
}
