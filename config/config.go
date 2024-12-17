package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Port int
}

func get_val(val string) string {
	value, ok := os.LookupEnv(val)
	if !ok {
		panic(fmt.Sprintf("required env var not set: %s", val))
	}
	return value
}

func New() *Config {
	portStr := get_val("PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		panic(fmt.Sprintf("unable to parse string to int: %s", portStr))
	}

	return &Config{
		Port: port,
	}
}
