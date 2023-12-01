package config

import "os"

type Config struct {
	OpenCage       string
	TripAdvisorKey string
}

func GetConfig() Config {
	c := Config{
		TripAdvisorKey: os.Getenv("TRIPADVISORKEY"),
	}
	return c
}
