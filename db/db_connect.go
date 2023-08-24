package db

import (
    "fmt"

    "github.com/team2/real_api/app/models"
    "github.com/team2/real_api/config"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
)

func Init(cfg *config.Config) *gorm.DB {
    masterDSN := fmt.Sprintf("host=%s user=%s password='%s' dbname=postgres port=%s sslmode=disable TimeZone=Asia/Ho_Chi_Minh",
        cfg.DB.Host,
        cfg.DB.User,
        cfg.DB.Password,
        cfg.DB.Port,
    )

    masterDB, err := gorm.Open(postgres.Open(masterDSN), &gorm.Config{
        Logger: logger.Default.LogMode(logger.Silent),
    })
    if err != nil {
        panic(err.Error())
    }

    // Check if the database exists
    var count int64
    result := masterDB.Raw("SELECT COUNT(datname) FROM pg_database WHERE datname = ?", cfg.DB.Name).Scan(&count)
    if result.Error != nil {
        panic(result.Error)
    }

    // If the database doesn't exist, create it
    if count == 0 {
        createDBQuery := fmt.Sprintf("CREATE DATABASE %s", cfg.DB.Name)
        result = masterDB.Exec(createDBQuery)
        if result.Error != nil {
            panic(result.Error)
        }
    }

    // Close the connection to the masterDB
    sqlDB, err := masterDB.DB()
    if err != nil {
        panic(err.Error())
    }
    sqlDB.Close()

    // Now, connect to the actual database
    dsn := fmt.Sprintf("host=%s user=%s password='%s' dbname=%s port=%s sslmode=disable TimeZone=Asia/Ho_Chi_Minh",
        cfg.DB.Host,
        cfg.DB.User,
        cfg.DB.Password,
        cfg.DB.Name,
        cfg.DB.Port,
    )

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

    db.AutoMigrate(
        &models.User{},
    )

    if err != nil {
        panic(err.Error())
    }

    return db
}
