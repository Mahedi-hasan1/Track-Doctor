package routes

import (
	"track_doctor/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/doctor", controllers.CreateDoctor)
	app.Get("/doctors", controllers.DoctorsList)
	app.Get("/doctors/:doctorId", controllers.DoctorDetails)
	app.Delete("/doctors/:doctorId", controllers.DeleteDoctor)
	app.Put("/doctors/:doctorId", controllers.UpdateDoctor)
}
