package app

import (
	"database/sql"
	"fmt"
	"log"

	"hsaaku.com/auth-api/app/database/migrate"
	"hsaaku.com/auth-api/config"
)

// App holds app instance
type App struct {
	DB *sql.DB
}

// Initialize will open the database connection
// and set up routes
func (a *App) Initialize(config *config.Config) {
	dsn := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%d sslmode=disable",
		config.DB.User,
		config.DB.Pass,
		config.DB.Name,
		config.DB.Host,
		config.DB.Port,
	)

	var err error

	a.DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	defer a.DB.Close()

	err = a.DB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Running database migrations...")
	migrate.DropAllTables(a.DB)
	migrate.ImportTables(a.DB)
	fmt.Println("Database migrated successfully")

	fmt.Println("DB connected!")
}
