package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"

	"github.com/gorilla/mux"
	"hsaaku.com/auth-api/app/database/migrate"
	"hsaaku.com/auth-api/app/routes"
	"hsaaku.com/auth-api/config"
)

// App holds app instance
type App struct {
	DB     *sqlx.DB
	Router *mux.Router
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
	a.DB = sqlx.MustConnect("postgres", dsn)
	a.DB.SetMaxIdleConns(10)
	a.DB.SetMaxOpenConns(10)

	// Resetting database
	migrate.DropAllTables(a.DB)
	migrate.ImportTables(a.DB)

	a.Router = mux.NewRouter()

	routes.InitializeRoutes(a.DB, a.Router)
}

// Run the app
func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}
