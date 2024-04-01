package controllers

import (
	"strconv"
	database "track_doctor/db"
	"track_doctor/models"
	"github.com/gofiber/fiber/v2"
)

func CreateDoctor(c *fiber.Ctx) error {
	var doctorData map[string]string
	err := c.BodyParser(&doctorData)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Invalid data to create doctor",
		})
	}
	if doctorData["name"] == "" {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Doctor name is required",
		})
	}
	doctor := models.Doctor{
		Name:    doctorData["name"],
		Digree:  doctorData["digree"],
		Specs:   doctorData["specs"],
		Picture: doctorData["picture"],
	}
	_ = doctor
	database.DB.Db.Create(&doctor)
	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "New doctor created successfully",
		"data":    doctorData,
	})
}

func DoctorsList(c *fiber.Ctx) error {
	limit, _ := strconv.Atoi(c.Query("limit"))
	skip, _ := strconv.Atoi(c.Query("skip"))
	var count int64
	var doctors []models.Doctor
	database.DB.Db.Select("*").Limit(limit).Offset(skip).Find(&doctors).Count(&count)
	metaMap := map[string]interface{}{
		"total": count,
		"limit": limit,
		"skip":  skip,
	}
	doctorsData := map[string]interface{}{
		"doctors": doctors,
		"meta":    metaMap,
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"Message": "Finding Doctors List Successfull",
		"data":    doctorsData,
	})
}

func DoctorDetails(c *fiber.Ctx) error {
	doctorId := c.Params("doctorId")
	var doctor models.Doctor
	database.DB.Db.Find(&doctor, "id=?", doctorId)
	if doctor.Name == "" {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Doctor Not Found",
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"Message": "Success",
		"data":    doctor,
	})
}

func UpdateDoctor(c *fiber.Ctx) error {
	doctorId := c.Params("doctorId")
	var doctor models.Doctor

	database.DB.Db.Find(&doctor, "id = ?", doctorId)
	if doctor.Name == "" {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Doctor Not Found",
		})
	}

	var updateDoctorData models.Doctor
	c.BodyParser(&updateDoctorData)
	if updateDoctorData.Name == "" && updateDoctorData.Digree == "" &&
		updateDoctorData.Specs == "" && updateDoctorData.Picture == "" {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Doctor information is required",
			"error":   map[string]interface{}{},
		})
	}
	if updateDoctorData.Name != ""{
		doctor.Name = updateDoctorData.Name
	}
	if updateDoctorData.Specs != ""{
		doctor.Specs = updateDoctorData.Specs
	}
	if updateDoctorData.Digree != ""{
		doctor.Digree = updateDoctorData.Digree
	}
	if updateDoctorData.Picture != ""{
		doctor.Picture = updateDoctorData.Picture
	}
	database.DB.Db.Save(&doctor)
	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Success",
		"data":    doctor,
	})
}

func DeleteDoctor(c *fiber.Ctx) error {
	doctorId := c.Params("doctorId")
	var doctor models.Doctor

	database.DB.Db.Delete(&doctor, "id = ?", doctorId)
	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Doctor deleted successfully",
		"data":    doctorId,
	})
}