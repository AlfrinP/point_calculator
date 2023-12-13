package models

import (
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Email        string        `json:"email" gorm:"unique"`
	PasswordHash string        `json:"-"`
	Department   string        `json:"department"`
	Batch        string        `json:"batch"`
	FacultyID    uint          `json:"faculty_id" gorm:"index;default:null"`
	Faculty      Faculty       `json:"faculty" gorm:"foreignKey:FacultyID"`
	Certificates []Certificate `json:"certificates" gorm:"foreignKey:StudentID"`
}

type StudentCreate struct {
	Email      string `validate:"required" json:"email"`
	Password   string `validate:"required" json:"password"`
	Department string `json:"department"`
	Batch      string `json:"batch"`
	FacultyID  uint   `json:"faculty_id"`
}

type UserSignIn struct {
	Email    string `validate:"required" json:"email"`
	Password string `validate:"required" json:"password"`
}

func (bc *StudentCreate) Validate() error {
	validate := validator.New()
	if err := validate.Struct(bc); err != nil {
		return err
	}
	return nil
}

func (us *UserSignIn) Validate() error {
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
		Email:        bc.Email,
		PasswordHash: string(hashedPasswd),
		Department:   bc.Department,
		Batch:        bc.Batch,
		FacultyID:    bc.FacultyID,
	}, nil
}
