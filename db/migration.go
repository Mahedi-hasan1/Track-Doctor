package database

import (
	Models "track_doctor/models"
	"gorm.io/gorm"
)

func AutoMigrate(connection *gorm.DB) {
	connection.Debug().AutoMigrate(
		&Models.Doctor{},
		&Models.DoctorAuth{},
		&Models.User{},
		&Models.UserAuth{},
		&Models.Chamber{},
		&Models.Schedule{},
		&Models.Review{},
		&Models.Alarm{},
	)
}
