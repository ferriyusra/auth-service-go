package config

import (
	"errors"
	"os"
	"strconv"
)

type Config struct {
	AppName          string
	AppEnv           string
	AppDebug         bool
	AppVersion       string
	AuthJwtSecret    string
	AuthJwtSecretCrm string
	RdsUrl           string
	SentryDsn        string
	SentrySampleRate float64
}

var configImpl *Config

func New() (*Config, error) {

	appName := getEnv("APP_NAME", "")
	if appName == "" {
		return nil, errors.New("APP_NAME env is required")
	}

	appEnv := getEnv("APP_ENV", "")
	if appEnv == "" {
		return nil, errors.New("APP_ENV env is required")
	}

	appDebug := getEnvAsBool("APP_DEBUG", false)

	appVersion := getEnv("APP_VERSION", "")
	if appVersion == "" {
		return nil, errors.New("APP_VERSION env is required")
	}

	authJwtSecret := getEnv("AUTH_JWT_SECRET", "")
	if authJwtSecret == "" {
		return nil, errors.New("AUTH_JWT_SECRET env is required")
	}

	authJwtSecretCrm := getEnv("AUTH_JWT_SECRET_CRM", "")
	if authJwtSecretCrm == "" {
		return nil, errors.New("AUTH_JWT_SECRET_CRM env is required")
	}

	rdsUrl := getEnv("RDS_URL", "")
	if rdsUrl == "" {
		return nil, errors.New("RDS_URL env is required")
	}

	// sentryDsn := getEnv("SENTRY_DSN", "")
	// if sentryDsn == "" {
	// 	return nil, errors.New("SENTRY_DSN env is required")
	// }

	// sentrySampleRate := getEnvAsFloat64("SENTRY_SAMPLE_RATE", 0.5)

	configImpl = &Config{
		AppName:          appName,
		AppEnv:           appEnv,
		AppDebug:         appDebug,
		AppVersion:       appVersion,
		AuthJwtSecret:    authJwtSecret,
		AuthJwtSecretCrm: authJwtSecretCrm,
		RdsUrl:           rdsUrl,
		// SentryDsn:        sentryDsn,
		// SentrySampleRate: sentrySampleRate,
	}
	return configImpl, nil
}

func GetConfig() (*Config, error) {
	if configImpl == nil {
		return nil, errors.New("config has not been initialized")
	}

	return configImpl, nil
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func getEnvAsBool(name string, defaultVal bool) bool {
	valStr := getEnv(name, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}

	return defaultVal
}

func getEnvAsFloat64(name string, defaultVal float64) float64 {
	valueStr := getEnv(name, "")

	if val, err := strconv.ParseFloat(valueStr, 32); err == nil {
		return val
	}

	return defaultVal
}

func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")

	if val, err := strconv.Atoi(valueStr); err == nil {
		return val
	}

	return defaultVal
}
