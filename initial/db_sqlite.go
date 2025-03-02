package initial

import (
	"path/filepath"

	"github.com/Yosorable/gomic/internal/global"
	"github.com/Yosorable/gomic/internal/model/database"
	"github.com/Yosorable/gomic/internal/utils"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDBSqlite() {
	db, err := gorm.Open(sqlite.Open(filepath.Join(global.CONFIG.DataPath, "data.db")), &gorm.Config{})
	if err != nil {
		logrus.Fatalln("Failed to connect database", err)
	}
	db.AutoMigrate(
		&database.Author{},
		&database.Archive{},
		&database.ArchiveFile{},
		&database.User{},
		&database.CacheFile{},
	)
	var usersCount int64 = 0
	err = db.Model(&database.User{}).Count(&usersCount).Error
	if err != nil {
		logrus.Fatalln("Failed to count users table", err)
	}
	if usersCount == 0 {
		hash, err := utils.HashPassword("admin")
		if err != nil {
			logrus.Fatalln("Failed to generate password", err)
		}
		user := database.User{
			Name:    "admin",
			PWDHash: hash,
			IsAdmin: true,
		}
		err = db.Create(&user).Error
		if err != nil {
			logrus.Fatalln("Failed to create user", err)
		}
	}
	global.DB = db
}
