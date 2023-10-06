package config

type Config struct {
	GoogleKey      string
	TripAdvisorKey string
	LogLevel       int
}

func GetConfig() (Config, error) {
	c := Config{}
	return c, nil
}
