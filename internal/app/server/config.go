package server

//Config is a representation of a server config
//It contains the connection port and a database URL
type Config struct {
	BindAddr    string `json:"bind_addr"`
	DatabaseURL string `json:"database_url"`
	LogLevel    string `json:"log_level"`
}

//NewConfig returns a default config
func NewConfig() *Config {
	return &Config{
		BindAddr:    ":8080",
		DatabaseURL: "",
		LogLevel:    "debug",
	}
}
