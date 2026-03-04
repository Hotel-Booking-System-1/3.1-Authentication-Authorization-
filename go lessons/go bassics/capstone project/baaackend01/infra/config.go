package infra

// Configurations holds database config
type Configurations struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
}

// GetConfigurations returns default configuration for DB
func GetConfigurations() *Configurations {
	return &Configurations{
		DBHost:     "localhost",
		DBUser:     "postgres",
		DBPassword: "password",
		DBName:     "guestdb",
		DBPort:     "5432",
	}
}
