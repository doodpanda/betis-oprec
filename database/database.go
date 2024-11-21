package database

import (
	"fmt"
	"log"
	"strconv"

	"betis-oprec/config"
	"betis-oprec/internal/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Declare the variable for the database
var DB *gorm.DB

// ConnectDB connect to db
func ConnectDB() {
	var err error
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		log.Println("DB ERR")
	}

	// Connection URL to connect to Postgres Database
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"))
	// Connect to the DB and initialize the DB variable
	DB, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")

}

// MigrateDB connects to the database using configuration parameters and performs
// automatic migration for the specified models. It retrieves the database connection
// details from the configuration, constructs the DSN (Data Source Name), and opens
// a connection to the PostgreSQL database using GORM. If the connection or migration
// fails, it logs a fatal error and terminates the program. Upon successful migration,
// it logs a completion message.
func MigrateDB() {
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Automigrate all models
	err = db.AutoMigrate(&model.MagicBook{}, &model.Witch{}, &model.AccessPermission{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	log.Println("Database migration completed successfully.")
}

func MigrateEnums(db *gorm.DB) error {
	// Create the magic_type enum if it doesn't exist
	err := db.Exec(`
		DO $$
		BEGIN
			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'magic_type') THEN
				CREATE TYPE magic_type AS ENUM ('elemental', 'illusion', 'necromancy', 'healing');
			END IF;
		END $$;
	`).Error
	if err != nil {
		return err
	}

	// Create other enums if necessary
	err = db.Exec(`
		DO $$
		BEGIN
			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'magic_rank') THEN
				CREATE TYPE magic_rank AS ENUM ('apprentice', 'adept', 'master', 'archmage');
			END IF;
		END $$;
	`).Error
	return err
}
