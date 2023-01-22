package database

import (
	"fmt"
	"github.com/singhbrijesh540/fampay-assignment/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
)

func NewDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable TimeZone=UTC",
		config.GetEnv().DB.Host, config.GetEnv().DB.Port, config.GetEnv().DB.User,
		config.GetEnv().DB.DBName)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: false,       // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Disable color
		},
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   config.GetEnv().DB.TableNamePrefix, // schema name
			SingularTable: true,
		},
		Logger: newLogger,
	})
	if err != nil {
		//TODO
		fmt.Printf("Couldn't connect to the database %e", err)
		return nil, err
	} else {
		fmt.Println("Connected to database")
	}
	return db, nil
}
