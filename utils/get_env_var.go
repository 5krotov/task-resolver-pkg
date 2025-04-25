package utils

import "os"

func GetEnvVar(name string) string {
	value, ok := os.LookupEnv(name)
	if ok {
		return value
	}
	return name
}
