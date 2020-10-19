package models

func migration() {
	DB.AutoMigrate(&User{})
}
