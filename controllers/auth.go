package controllers

import (
	"time"

	"github.com/AlfrinP/point_calculator/config"
	"github.com/AlfrinP/point_calculator/models"
	"github.com/AlfrinP/point_calculator/repository"
	"github.com/AlfrinP/point_calculator/storage"
	"github.com/AlfrinP/point_calculator/util"
	"github.com/gofiber/fiber/v2"
)

func SignUp(c *fiber.Ctx) error {

	role := c.Params("role")

	if role == "student" {
		params := &models.StudentCreate{}
		if err := c.BodyParser(params); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"msg": err.Error(),
			})
		}

		if err := params.Validate(); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"msg": err.Error(),
			})
		}

		student, err := params.Convert()
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"msg": err.Error(),
			})
		}

		studentRepo := repository.NewStudentRepository(storage.GetDB())
		if err := studentRepo.Create(student); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"msg": err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"user": student,
		})
	} else if role == "faculty" {
		params := &models.FacultyCreate{}
		if err := c.BodyParser(params); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"msg": err.Error(),
			})
		}

		if err := params.Validate(); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"msg": err.Error(),
			})
		}

		faculty, err := params.Convert()
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"msg": err.Error(),
			})
		}

		facultyRepo := repository.NewFacultyRepository(storage.GetDB())
		if err := facultyRepo.Create(faculty); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"msg": err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"user": faculty,
		})
	} else {
		return c.JSON(fiber.Map{
			"error": "ivalid role",
		})
	}
}

func SignIn(c *fiber.Ctx) error {
	params := &models.UserSignIn{}
	var id uint
	role := c.Params("role")

	if err := c.BodyParser(params); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	if role == "student" {
		studentRepo := repository.NewStudentRepository(storage.GetDB())
		student, err := studentRepo.Get(params.Email)
		id = student.ID
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"msg": err.Error(),
			})
		}
		if err := util.VerifyPassword(student.PasswordHash, params.Password); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"msg": "Invalid Email or Password",
			})
		}
	} else if role == "faculty" {
		facultyRepo := repository.NewFacultyRepository(storage.GetDB())
		faculty, err := facultyRepo.Get(params.Email)
		id = faculty.ID
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"msg": err.Error(),
			})
		}
		if err := util.VerifyPassword(faculty.Password, params.Password); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"msg": "Invalid Email or Password",
			})
		}
	} else {
		return c.JSON(fiber.Map{
			"error": "ivalid role",
		})
	}

	config, _ := config.LoadConfig(".")
	tokenString, err := util.GenerateToken(id, role, config)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"msg": "Generating JWT Token failed",
		})
	}
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    tokenString,
		Path:     "/",
		MaxAge:   config.JwtMaxAge * 60,
		Secure:   false,
		HTTPOnly: true,
		Domain:   "localhost",
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"token": tokenString})
}

func LogoutUser(c *fiber.Ctx) error {
	expired := time.Now().Add(-time.Hour * 24)
	c.Cookie(&fiber.Cookie{
		Name:    "token",
		Value:   "",
		Expires: expired,
	})
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success"})
}
