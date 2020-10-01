package server

type Config struct {
	BindAddr    string `json:"bind_addr"`
	DatabaseURL string `json:"database_url"`
}
