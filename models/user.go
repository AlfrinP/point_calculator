package models

import (
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email        string        `gorm:"unique" json:"email"`
	UserName     string        `json:"userName"`
	PasswordHash string        `json:"passwordHash"`
	Department   string        `json:"department"`
	Batch        string        `json:"batch"`
	FacultyID    uint          `json:"faculty_id" gorm:"index;default:null"`
	Faculty      Faculty       `json:"faculty" gorm:"foreignKey:FacultyID"`
	Certificates []Certificate `json:"certificates" gorm:"foreignKey:StudentID"`
}

type UserCreate struct {
	Email      string `validate:"required" json:"email"`
	UserName   string `validate:"required" json:"userName"`
	Password   string `validate:"required" json:"password"`
	Department string `json:"department"`
	Batch      string `json:"batch"`
}

type UserSignIn struct {
	Email    string `validate:"required" json:"email"`
	Password string `validate:"required" json:"password"`
}

func (bc *UserCreate) Validate() error {
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

func (bc *UserCreate) Convert() (*User, error) {
	hashedPasswd, err := bcrypt.GenerateFromPassword([]byte(bc.Password), bcrypt.DefaultCost)
	if err != nil {
		return &User{}, err
	}

	return &User{
		Email:        bc.Email,
		UserName:     bc.UserName,
		PasswordHash: string(hashedPasswd),
		Department:   bc.Department,
		Batch:        bc.Batch,
	}, nil
}
