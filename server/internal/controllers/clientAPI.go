package controllers

import (
	"github.com/1boombacks1/botInsurance/internal/services"
	"github.com/gofiber/fiber/v2"
)

type clientRoutes struct {
	carInsuranceService services.CarInsuranceApplicationBuilder
}

func newClientRoutes(group fiber.Router, carInsuranceService services.CarInsuranceApplicationBuilder) {
	r := clientRoutes{carInsuranceService}

	group.Post("/setDesc", r.setDescription)
	group.Post("/checkVIN", r.checkVIN)
	// group.Post("/checkMainOutsidePhotos")
	// group.Post("/checkWindshieldPhoto")
	// group.Post("/checkVIN")
	// group.Post("/checkVIN")
	// group.Post("/checkVIN")
	// group.Post("/checkVIN")
}

// json: {name string, lastname string, pastronomyc string, phone string}
func (r *clientRoutes) setDescription(c *fiber.Ctx) error {
	type DescInput struct {
		Name       string `json:"name"`
		LastName   string `json:"last_name"`
		Patronymic string `json:"patronymic"`
		Phone      string `json:"phone"`
	}
	var input DescInput
	err := c.BodyParser(&input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "bad request. Ты что-то не то отправил",
			"data":    err,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "данные были получены и сохранены",
		"data":    input,
	})
}

// formdata с фото VIN
func (r *clientRoutes) checkVIN(c *fiber.Ctx) error {
	photo, err := c.FormFile("VIN")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "bad request. Ты что-то не то отправил, нужен form-data: key -> VIN",
			"data":    err,
		})
	}
	file, err := photo.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "server error. В серваке ошибка, не удалось открыть файл",
			"data":    err,
		})
	}

	err = r.carInsuranceService.SetVINPhoto(c.Context(), file)
	if err != nil {
		if err == services.ErrBadResolution {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status":  "not valid resolution",
				"message": "bad resolution. Неправильное разрешение",
				"data":    err,
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "internal error",
			"message": "server error. В серваке ошибка, не удалось декодировать/сохранить файл",
			"data":    err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "фото валидна и была сохранена успешна",
	})
}
