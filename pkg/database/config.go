package database

type DbConfig struct {
	dbUser     string
	dbPassword string
	dbName     string
	dbHost     string
	dbPort     string
	// optional attribute if not set will be nil
	dbMigrationPath *string
}

// create a new db config type
func NewDbConfig(user string, password string, name string, host string, port string) *DbConfig {
	return &DbConfig{
		dbUser:     user,
		dbPassword: password,
		dbName:     name,
		dbHost:     host,
		dbPort:     port,
	}
}

// we use this function to set the path if migration path is not nil
func (config *DbConfig) WithMigration(path string) *DbConfig {
	config.dbMigrationPath = &path
	return config
}
