package config

import (
	"math"
	"os"
	"strconv"
	"testing"
)

const float64EqualityThreshold = 1e-4

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) <= float64EqualityThreshold
}

func TestNew(t *testing.T) {
	t.Run("config not initialized", func(t *testing.T) {
		_, err := GetConfig()
		if err == nil {
			t.Errorf(`want: %v`, err)
		}
	})

	t.Run("APP_NAME not set", func(t *testing.T) {
		_, err := New()
		if err == nil {
			t.Errorf(`Value: %v`, err)
		}
	})

	t.Run("APP_ENV not set", func(t *testing.T) {
		os.Setenv("APP_NAME", "super-app")

		_, err := New()
		if err == nil {
			t.Errorf(`Value: %v`, err)
		}
	})

	t.Run("APP_VERSION not set", func(t *testing.T) {
		os.Setenv("APP_NAME", "super-app")
		os.Setenv("APP_ENV", "super-env")

		_, err := New()
		if err == nil {
			t.Errorf(`Value: %v`, err)
		}
	})

	t.Run("AUTH_JWT_SECRET not set", func(t *testing.T) {
		os.Setenv("APP_NAME", "super-app")
		os.Setenv("APP_ENV", "super-env")
		os.Setenv("APP_VERSION", "12.34.5")

		_, err := New()
		if err == nil {
			t.Errorf(`Value: %v`, err)
		}
	})

	t.Run("AUTH_JWT_SECRET_CRM not set", func(t *testing.T) {
		os.Setenv("APP_NAME", "super-app")
		os.Setenv("APP_ENV", "super-env")
		os.Setenv("APP_VERSION", "12.34.5")
		os.Setenv("AUTH_JWT_SECRET", "秘密")

		_, err := New()
		if err == nil {
			t.Errorf(`Value: %v`, err)
		}
	})

	t.Run("KAFKA_BROKERS not set", func(t *testing.T) {
		os.Setenv("APP_NAME", "super-app")
		os.Setenv("APP_ENV", "super-env")
		os.Setenv("APP_VERSION", "12.34.5")
		os.Setenv("AUTH_JWT_SECRET", "秘密")
		os.Setenv("AUTH_JWT_SECRET_CRM", "비밀")

		_, err := New()
		if err == nil {
			t.Errorf(`Value: %v`, err)
		}
	})

	t.Run("RDS_URL not set", func(t *testing.T) {
		os.Setenv("APP_NAME", "super-app")
		os.Setenv("APP_ENV", "super-env")
		os.Setenv("APP_VERSION", "12.34.5")
		os.Setenv("AUTH_JWT_SECRET", "秘密")
		os.Setenv("AUTH_JWT_SECRET_CRM", "비밀")
		os.Setenv("KAFKA_BROKERS", "127.0.0.1:29092")

		_, err := New()
		if err == nil {
			t.Errorf(`Value: %v`, err)
		}
	})

	t.Run(" not set", func(t *testing.T) {
		os.Setenv("APP_NAME", "super-app")
		os.Setenv("APP_ENV", "super-env")
		os.Setenv("APP_VERSION", "12.34.5")
		os.Setenv("AUTH_JWT_SECRET", "秘密")
		os.Setenv("AUTH_JWT_SECRET_CRM", "비밀")
		os.Setenv("KAFKA_BROKERS", "127.0.0.1:29092")
		os.Setenv("RDS_URL", "http://rds.com/aaa")

		_, err := New()
		if err == nil {
			t.Errorf(`Value: %v`, err)
		}
	})

	t.Run("TOPIC_MODULE_COMPLETED not set", func(t *testing.T) {
		os.Setenv("APP_NAME", "super-app")
		os.Setenv("APP_ENV", "super-env")
		os.Setenv("APP_VERSION", "12.34.5")
		os.Setenv("AUTH_JWT_SECRET", "秘密")
		os.Setenv("AUTH_JWT_SECRET_CRM", "비밀")
		os.Setenv("KAFKA_BROKERS", "127.0.0.1:29092")
		os.Setenv("RDS_URL", "http://rds.com/aaa")
		os.Setenv("SENTRY_DSN", "http://sentry.com/bbb")

		_, err := New()
		if err == nil {
			t.Errorf(`Value: %v`, err)
		}
	})

	appName := "super-app"
	appEnv := "super-env"
	appDebug := "true"
	appVersion := "12.34.5"
	authJwtSecret := "秘密"
	authJwtSecretCrm := "비밀"
	kafkaBrokers := "127.0."
	rdsUrl := "http://rds.com/aaa"
	sentryDsn := "http://sentry.com/bbb"
	topicModuleCompleted := "agent-activity"

	os.Setenv("APP_NAME", appName)
	os.Setenv("APP_ENV", appEnv)
	os.Setenv("APP_DEBUG", appDebug)
	os.Setenv("APP_VERSION", appVersion)
	os.Setenv("AUTH_JWT_SECRET", authJwtSecret)
	os.Setenv("AUTH_JWT_SECRET_CRM", authJwtSecretCrm)
	os.Setenv("KAFKA_BROKERS", kafkaBrokers)
	os.Setenv("RDS_URL", rdsUrl)
	os.Setenv("SENTRY_DSN", sentryDsn)
	os.Setenv("TOPIC_MODULE_COMPLETED", topicModuleCompleted)

	cfg, _ := New()

	tests := []struct {
		name string
		env  string
		want string
	}{
		{name: "APP_NAME", env: cfg.AppName, want: appName},
		{name: "APP_ENV", env: cfg.AppEnv, want: appEnv},
		{name: "APP_VERSION", env: cfg.AppVersion, want: appVersion},
		{name: "AUTH_JWT_SECRET", env: cfg.AuthJwtSecret, want: authJwtSecret},
		{name: "AUTH_JWT_SECRET", env: cfg.AuthJwtSecret, want: authJwtSecret},
		{name: "RDS_URL", env: cfg.RdsUrl, want: rdsUrl},
		{name: "SENTRY_DSN", env: cfg.SentryDsn, want: sentryDsn},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.env != tt.want {
				t.Errorf(`Value: %s. want: %s`, tt.env, tt.want)
			}
		})
	}

	appDebugBool, _ := strconv.ParseBool(appDebug)
	if cfg.AppDebug != appDebugBool {
		t.Errorf("Value: %t. want: %t", cfg.AppDebug, appDebugBool)
	}
}

func TestGetConfig(t *testing.T) {
	t.Run("GetConfig is working", func(t *testing.T) {
		// create new config
		cfg, _ := New()

		// get the config
		config, _ := GetConfig()

		if cfg != config {
			t.Errorf(`got: %v. want: %v`, cfg, config)
		}
	})
}

func TestGetEnv(t *testing.T) {

	a := "aaa"
	defaultValue := "default-value"

	os.Setenv("ENV_A", a)

	tests := []struct {
		name string
		env  string
		def  string
		want string
	}{
		{name: "ENV_A", env: "ENV_A", def: defaultValue, want: a},
		{name: "DONTEXIST", env: "DONTEXIST", def: defaultValue, want: defaultValue},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			res := getEnv(tt.env, tt.def)

			if res != tt.want {
				t.Errorf(`Value: %s. want: %s`, res, tt.want)
			}
		})
	}
}

func TestGetEnvAsBool(t *testing.T) {

	a := true
	b := false
	defaultTrueValue := true
	defaultFalseValue := false

	os.Setenv("ENV_TRUE", "true")
	os.Setenv("ENV_FALSE", "false")

	tests := []struct {
		name string
		env  string
		def  bool
		want bool
	}{
		{name: "ENV_TRUE", env: "ENV_TRUE", def: defaultTrueValue, want: a},
		{name: "ENV_FALSE", env: "ENV_FALSE", def: defaultFalseValue, want: b},
		{name: "DONTEXIST", env: "DONTEXIST", def: defaultTrueValue, want: defaultTrueValue},
		{name: "DONTEXIST", env: "DONTEXIST", def: defaultFalseValue, want: defaultFalseValue},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			res := getEnvAsBool(tt.env, tt.def)

			if res != tt.want {
				t.Errorf(`Value: %t. want: %t`, res, tt.want)
			}
		})
	}
}

func TestGetEnvAsFloat64(t *testing.T) {
	os.Setenv("ENV_FLOAT64", "0.7")
	defaultValue := 0.5

	tests := []struct {
		env  string
		def  float64
		want float64
	}{
		{env: "ENV_FLOAT64", def: defaultValue, want: 0.7},
		{env: "DONTEXIST", def: defaultValue, want: defaultValue},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			res := getEnvAsFloat64(tt.env, tt.def)

			if !almostEqual(res, tt.want) {
				t.Errorf(`got: %v. want: %v`, res, tt.want)
			}
		})
	}
}

func TestGetEnvAsInt(t *testing.T) {
	os.Setenv("ENV_INT", "100")
	defaultValue := 100

	tests := []struct {
		env  string
		def  int
		want int
	}{
		{env: "ENV_INT", def: defaultValue, want: 100},
		{env: "DONTEXIST", def: defaultValue, want: defaultValue},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			res := getEnvAsInt(tt.env, tt.def)

			if res != tt.want {
				t.Errorf(`got: %v. want: %v`, res, tt.want)
			}
		})
	}
}
