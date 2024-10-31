package init

import "os"

type Config struct {
	Port string
}

func GetConfig() Config {
	cfg := Config{}

	cfg.Port = os.Getenv("PORT")

	return cfg
}
