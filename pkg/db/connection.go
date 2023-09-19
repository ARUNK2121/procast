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

	db.AutoMigrate(domain.State{})
	db.AutoMigrate(domain.Provider{})
	db.AutoMigrate(domain.Admin{})
	db.AutoMigrate(domain.District{})
	db.AutoMigrate(domain.Category{})
	db.AutoMigrate(domain.Profession{})
	db.AutoMigrate(domain.Probook{})
	db.AutoMigrate(domain.PreferredLocation{})
	db.AutoMigrate(domain.Rating{})
	db.AutoMigrate(domain.ProviderNotification{})
	db.AutoMigrate(domain.Post{})
	db.AutoMigrate(domain.PostImages{})
	db.AutoMigrate(domain.IdProof{})
	db.AutoMigrate(domain.Bid{})
	db.AutoMigrate(domain.Work{})
	db.AutoMigrate(domain.CompletedImages{})
	db.AutoMigrate(domain.WorkspaceImages{})
	db.AutoMigrate(domain.User{})
	db.AutoMigrate(domain.UserNotification{})

	//if the database doesnot have an admin then create a super admin with credentials in env file
	utils.CheckAndCreateAdmin(db, cfg)

	return db, dbErr
}
