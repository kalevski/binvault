package cfg

import (
	"os"
	"path/filepath"
	"reflect"
	"sync"
)

var (
	variables *Variables
	once      sync.Once
)

type Variables struct {
	SERVER_PORT string
	SERVER_HOST string
	DATA_PATH   string
	DB_NAME     string
	JWKS        string
}

func GetVars() *Variables {
	once.Do(func() {
		variables = &Variables{
			SERVER_PORT: readEnv("SERVER_PORT", "8080"),
			SERVER_HOST: readEnv("SERVER_HOST", "localhost"),
			DATA_PATH:   readEnv("DATA_PATH", "./data"),
			DB_NAME:     readEnv("DB_NAME", "_database"),
			JWKS:        readEnv("JWKS", ""),
		}
	})
	return variables
}

func GetVar(key string) string {
	vars := GetVars()
	v := reflect.ValueOf(vars)
	field := v.Elem().FieldByName(key)
	return field.String()
}

func GetPath(key string) string {
	path := GetVar(key)
	abs, err := filepath.Abs(path)
	if err != nil {
		return path
	}
	return abs
}

func readEnv(key string, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		value = fallback
	}
	return value
}
