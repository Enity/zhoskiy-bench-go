package gorm

import (
	"enity/zhoskiy-bench-go/models"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Db *gorm.DB

func Init() {
	mysqlDsn := os.Getenv("MYSQL_DSN")

	db, err := gorm.Open(mysql.Open(mysqlDsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})

	if err != nil {
		panic("failed to connect database")
	}

	Db = db

	migrirovat := os.Getenv("ZHOSKA_MIGRIROVAT_DB")

	if migrirovat == "true" {
		fmt.Println("Начинаю жоска мигрировать бд")
		Db.AutoMigrate(&models.Bear{})
		fmt.Println("Жоска мигрировал бд")
	}

	napolnit := os.Getenv("ZHOSKA_NAPOLNIT_DB")

	if napolnit == "true" {
		fmt.Println("Начинаю жоска наполнять бд")

		var bears [5000]models.Bear
		for i := 0; i < len(bears); i++ {
			bears[i] = models.Bear{
				Name:            "sanes",
				KdRatio:         "5/1",
				LoveToSuckCocks: true,
			}
		}
		Db.Create(&bears)
		fmt.Println("Жоска наполнил бд")
	}
}
