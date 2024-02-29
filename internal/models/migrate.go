package models

// Migrate AutoMigrate models
func Migrate() error {
	return db.AutoMigrate(&Webgame{})
}
