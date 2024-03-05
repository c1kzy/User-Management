package cache

type RedisConfig struct {
	Addr string `env:"ADDR"`
	Pass string `env:"REDISPASS"`
	DB   int    `env:"DB"`
}
