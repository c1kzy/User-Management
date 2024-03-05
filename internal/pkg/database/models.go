package database

type Config struct {
	Host       string `env:"HOST"`
	Port       int    `env:"PORT"`
	User       string `env:"USER"`
	Password   string `env:"PASSWORD"`
	Name       string `env:"NAME"`
	Salt       string `env:"SALT"`
	SigningKey string `env:"SIGNINGKEY"`
}
