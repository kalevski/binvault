package env

import (
	"os"
	"path/filepath"
	"reflect"
)

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

func EnvExists(key string) bool {
	return GetVar(key) != ""
}
