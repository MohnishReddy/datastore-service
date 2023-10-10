package models

type Config struct {
	Database *DatabaseConfig `yaml:"database"`
}

type DatabaseConfig struct {
	DBUrl    string `yaml:"db_url"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}
