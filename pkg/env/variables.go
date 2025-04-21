package env

import (
	"sync"
)

var (
	variables *Variables
	once      sync.Once
)

type Variables struct {
	SERVER_PORT           string
	SERVER_HOST           string
	DATA_PATH             string
	DB_NAME               string
	JWKS_URL              string
	JWKS_KID              string
	SSH_PUBLIC_KEY        string
	SSH_PRIVATE_KEY       string
	JWT_CLAIM_ID          string
	PEM_PRIVATE_FILENAME  string
	PEM_PUBLIC_FILENAME   string
	PROCESSOR_CONFIG_PATH string
}

func GetVars() *Variables {
	once.Do(func() {
		variables = &Variables{
			SERVER_PORT:           readEnv("SERVER_PORT", "8080"),
			SERVER_HOST:           readEnv("SERVER_HOST", "localhost"),
			DATA_PATH:             readEnv("DATA_PATH", "./data"),
			DB_NAME:               readEnv("DB_NAME", "database.db"),
			JWKS_URL:              readEnv("JWKS_URL", ""),
			JWKS_KID:              readEnv("JWKS_KID", "main"),
			SSH_PUBLIC_KEY:        readEnv("SSH_PUBLIC_KEY", ""),
			SSH_PRIVATE_KEY:       readEnv("SSH_PRIVATE_KEY", ""),
			JWT_CLAIM_ID:          readEnv("JWT_CLAIM_ID", "id"),
			PEM_PRIVATE_FILENAME:  readEnv("PEM_PRIVATE_FILENAME", "key.pem"),
			PEM_PUBLIC_FILENAME:   readEnv("PEM_PUBLIC_FILENAME", "key_pub.pem"),
			PROCESSOR_CONFIG_PATH: readEnv("PROCESSOR_CONFIG_PATH", "./processors.cfg"),
		}
	})
	return variables
}
