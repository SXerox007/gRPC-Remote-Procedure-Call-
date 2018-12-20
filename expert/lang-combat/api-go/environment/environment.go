package environment

import (
	"os"
)

func GetEnv() string {
	env, isEnv := os.LookupEnv("environment")
	if !isEnv {
		env = "tcp"
	}
	return env
}

func GetPort() string {
	port, isPort := os.LookupEnv("port")
	if !isPort {
		port = ":5051"
	}
	return port
}
