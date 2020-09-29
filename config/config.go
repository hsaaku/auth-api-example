package config

// Config stores all configurations
type Config struct {
	DB *DBConfig
}

// DBConfig hold PostgreSQL connection
type DBConfig struct {
	Host string
	Port int
	User string
	Pass string
	Name string
}

// GetConfig is used to initialize configurations
// and return it
func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Host: "localhost",
			Port: 5432,
			User: "postgres",
			Pass: "Uk45y4h!",
			Name: "go_auth",
		},
	}
}
