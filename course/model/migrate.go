package model

import "gorm.io/gorm"

func migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&Course{},
		&Elective{},
		// TODO: models' instances here
	)

	return err
}