package database

type DbConfig struct {
	dbUser     string
	dbPassword string
	dbName     string
	dbHost     string
	dbPort     string
	dbSslmode  string
	dbUrl      string
	// optional attribute if not set will be nil
	dbMigrationPath *string
}

// create a new db config type
func NewDbConfig(user string, password string, name string, host string, port string, sslmode string, url string) *DbConfig {
	return &DbConfig{
		dbUser:     user,
		dbPassword: password,
		dbName:     name,
		dbHost:     host,
		dbPort:     port,
		dbSslmode:  sslmode,
		dbUrl:      url,
	}
}

// we use this function to set the path if migration path is not nil
func (config *DbConfig) WithMigration(path string) *DbConfig {
	config.dbMigrationPath = &path
	return config
}
