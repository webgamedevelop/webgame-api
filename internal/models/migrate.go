package models

// Migrate AutoMigrate models
func Migrate() error {
	return db.AutoMigrate(
		&IngressClass{},
		&ImagePullSecret{},
		&ResourceSpec{},
		&Repository{},
		&Tag{},
		&GameType{},
		&GameTypeVersion{},
		&Instance{},
		&User{},
	)
}
