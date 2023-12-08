package config

import "os"

type Config struct {
	TripAdvisorKey string
	OpenCageKey    string
}

func GetConfig() Config {
	c := Config{
		TripAdvisorKey: os.Getenv("TRIPADVISORKEY"),
		OpenCageKey:    os.Getenv("OPENCAGEKEY"),
	}
	return c
}
