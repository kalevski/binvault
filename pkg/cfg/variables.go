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
	SERVER_PORT          string
	SERVER_HOST          string
	DATA_PATH            string
	TEMP_DIR_NAME        string
	DB_NAME              string
	JWKS_URL             string
	JWKS_KID             string
	SSH_PUBLIC_KEY       string
	SSH_PRIVATE_KEY      string
	JWT_CLAIM_ID         string
	PEM_PRIVATE_FILENAME string
	PEM_PUBLIC_FILENAME  string
}

func GetVars() *Variables {
	once.Do(func() {
		variables = &Variables{
			SERVER_PORT:          readEnv("SERVER_PORT", "8080"),
			SERVER_HOST:          readEnv("SERVER_HOST", "localhost"),
			DATA_PATH:            readEnv("DATA_PATH", "./data"),
			DB_NAME:              readEnv("DB_NAME", "database.db"),
			JWKS_URL:             readEnv("JWKS_URL", ""),
			JWKS_KID:             readEnv("JWKS_KID", "main"),
			SSH_PUBLIC_KEY:       readEnv("SSH_PUBLIC_KEY", ""),
			SSH_PRIVATE_KEY:      readEnv("SSH_PRIVATE_KEY", ""),
			JWT_CLAIM_ID:         readEnv("JWT_CLAIM_ID", "id"),
			PEM_PRIVATE_FILENAME: readEnv("PEM_PRIVATE_FILENAME", "key.pem"),
			PEM_PUBLIC_FILENAME:  readEnv("PEM_PUBLIC_FILENAME", "key_pub.pem"),
			TEMP_DIR_NAME:        readEnv("TEMP_DIR_NAME", "temp"),
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

func EnvExists(key string) bool {
	return GetVar(key) != ""
}
