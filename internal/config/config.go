package config

type Config struct {
}

func Load() Config {
	return Default()
}
