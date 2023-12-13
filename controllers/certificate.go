package controllers

import (
	"github.com/AlfrinP/point_calculator/models"
	"github.com/AlfrinP/point_calculator/repository"
	"github.com/AlfrinP/point_calculator/storage"
	"log"
	"github.com/AlfrinP/point_calculator/util"
	"github.com/gofiber/fiber/v2"
)

func PostCertificate(c *fiber.Ctx) error {

	role, id := util.GetRoleAndID()
	if role == "student" {
		params := &models.CertificateCreate{}
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

		file,err:=c.FormFile("upload_certificate")
		if err!= nil {
            return err
        }
		log.Println(file.Filename)
		c.SaveFile(file,"certificates/"+file.Filename)

		certificate := &models.Certificate{
			StudentID: id,
			Name:      params.Name,
			Category:  params.Category,
			Level:     params.Level,
		}

		certificateRepo := repository.NewCertificateRepository(storage.GetDB())

		if err := certificateRepo.Create(certificate); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"msg": err.Error(),
			})
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Certificate successfully uploaded",
		})
	} else {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid role",
		})
	}
	
}