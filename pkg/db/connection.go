package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	config "github.com/ARUNK2121/procast/pkg/config"
	"github.com/ARUNK2121/procast/pkg/domain"
	"github.com/ARUNK2121/procast/pkg/utils"
)

func ConnectDatabase(cfg config.Config) (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", cfg.DBHost, cfg.DBUser, cfg.DBName, cfg.DBPort, cfg.DBPassword)
	db, dbErr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{SkipDefaultTransaction: true})

	err := db.AutoMigrate(domain.State{})
	if err != nil {
		return &gorm.DB{}, err
	}
	err = db.AutoMigrate(domain.Provider{})
	if err != nil {
		return &gorm.DB{}, err
	}
	err = db.AutoMigrate(domain.Admin{})
	if err != nil {
		return &gorm.DB{}, err
	}
	err = db.AutoMigrate(domain.District{})
	if err != nil {
		return &gorm.DB{}, err
	}
	err = db.AutoMigrate(domain.Category{})
	if err != nil {
		return &gorm.DB{}, err
	}
	err = db.AutoMigrate(domain.Profession{})
	if err != nil {
		return &gorm.DB{}, err
	}
	err = db.AutoMigrate(domain.Probook{})
	if err != nil {
		return &gorm.DB{}, err
	}
	err = db.AutoMigrate(domain.PreferredLocation{})
	if err != nil {
		return &gorm.DB{}, err
	}
	err = db.AutoMigrate(domain.Rating{})
	if err != nil {
		return &gorm.DB{}, err
	}
	err = db.AutoMigrate(domain.ProviderNotification{})
	if err != nil {
		return &gorm.DB{}, err
	}
	err = db.AutoMigrate(domain.Post{})
	if err != nil {
		return &gorm.DB{}, err
	}
	err = db.AutoMigrate(domain.PostImages{})
	if err != nil {
		return &gorm.DB{}, err
	}
	err = db.AutoMigrate(domain.IdProof{})
	if err != nil {
		return &gorm.DB{}, err
	}
	err = db.AutoMigrate(domain.Bid{})
	if err != nil {
		return &gorm.DB{}, err
	}
	err = db.AutoMigrate(domain.Work{})
	if err != nil {
		return &gorm.DB{}, err
	}
	err = db.AutoMigrate(domain.CompletedImages{})
	if err != nil {
		return &gorm.DB{}, err
	}
	err = db.AutoMigrate(domain.WorkspaceImages{})
	if err != nil {
		return &gorm.DB{}, err
	}
	err = db.AutoMigrate(domain.User{})
	if err != nil {
		return &gorm.DB{}, err
	}
	err = db.AutoMigrate(domain.UserNotification{})
	if err != nil {
		return &gorm.DB{}, err
	}

	//if the database doesnot have an admin then create a super admin with credentials in env file
	utils.CheckAndCreateAdmin(db, cfg)

	return db, dbErr
}
