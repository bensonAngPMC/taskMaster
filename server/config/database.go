package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	host     = "127.0.0.1"
	port     = 5432
	username = "postgres"
	password = "123456"
	dbname   = "task-test"
)

var postgresLogger logger.Interface

func DatabaseConnection() *gorm.DB {
	postgresLogger = logger.Default.LogMode(logger.Info)
	// dsn := "host=localhost user=postgres password=123456 dbname=test-gorm port=5432 sslmode=disable"
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Kuala_Lumpur connect_timeout=10", host, username, password, dbname, port)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DriverName: "pgx", // 指定 pgx 驱动
		DSN:        dsn,   // 数据源名称
	}), &gorm.Config{
		Logger: postgresLogger,
	})

	if err != nil {
		panic(fmt.Sprintf("Failed to connect to the database: %v", err))
	}
	return db
}
