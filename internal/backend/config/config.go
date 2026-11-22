package config

type Config struct{}

func Load() (Config, error)

func Save(c Config) error
