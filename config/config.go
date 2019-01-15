package config

//Config que erda DBConfig
type Config struct {
	DB *DBConfig
}

//Estrutura
type DBConfig struct {
	Dialect  string
	Username string
	Password string
	Name     string
	Charset  string
}

//Retorna o endero do config que erda a estrutura do DBConfig
func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect:  "mysql",
			Username: "root",
			Password: "123",
			Name:     "codes2",
			Charset:  "utf8",
		},
	}
}
