package configs

import "os"

type Config struct {
	DBDSN            string
	GoogleMapsAPIKey string
	Addr             string
}

func Load() *Config {
	c := &Config{
		DBDSN: os.Getenv("DATABASE_DSN"),
		GoogleMapsAPIKey: os.Getenv("GOOGLE_MAPS_API_KEY"),
		Addr: os.Getenv("ADDR"),
	}
	if c.Addr == ""{
		c.Addr = ":8080"
	}
	return c
}
