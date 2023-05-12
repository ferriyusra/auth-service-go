package logger

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"skegsTech/auth-service-go/config"

	"github.com/dlclark/regexp2"
	"github.com/getsentry/sentry-go"
	"github.com/sirupsen/logrus"
)

var (
	logger *logrus.Logger
)

type JsonFormatter struct{}

func (f *JsonFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := entry.Time.Format("2006-01-02T15:04:05+07:00")
	var msgRegexp *regexp2.Regexp
	msg := entry.Message

	// remove multitab
	msgRegexp = regexp2.MustCompile(`\t+(?=\t)`, 0)
	msg, _ = msgRegexp.Replace(msg, "", -1, -1)

	// remove multi newline
	msgRegexp = regexp2.MustCompile(`(\r\n|\n|\r)`, 0)
	msg, _ = msgRegexp.Replace(msg, "", -1, -1)

	// tab to space
	msgRegexp = regexp2.MustCompile(`\t`, 0)
	msg, _ = msgRegexp.Replace(msg, " ", -1, -1)

	// remove multispaces
	msgRegexp = regexp2.MustCompile(` +(?= )`, 0)
	msg, _ = msgRegexp.Replace(msg, "", -1, -1)

	output := fmt.Sprintf("%s | %s | %s | %s", timestamp, entry.Data["requestID"], entry.Level, msg)
	return append([]byte(output), '\n'), nil
}

func Init(config *config.Config) {
	logger = logrus.New()
	logger.Out = os.Stdout
	logger.Formatter = &JsonFormatter{}

	// init sentry
	err := sentry.Init(sentry.ClientOptions{
		Dsn:              config.SentryDsn,
		Environment:      config.AppEnv,
		Release:          fmt.Sprintf("%s@%s", config.AppName, config.AppVersion),
		Debug:            config.AppDebug,
		TracesSampleRate: config.SentrySampleRate,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	defer sentry.Flush(2 * time.Second)
}

func Info(ctx context.Context, message string) {
	logger.WithFields(logrus.Fields{
		"requestID": ctx.Value("requestID"),
	}).Info(message)
}

func Infof(ctx context.Context, message string, args ...interface{}) {
	logger.WithFields(logrus.Fields{
		"requestID": ctx.Value("requestID"),
	}).Infof(message, args...)
}

func Error(ctx context.Context, err error) {
	// log and capture error for non "test" environment
	if os.Getenv("APP_ENV") != "test" {
		sentry.CaptureException(err)

		logger.WithFields(logrus.Fields{
			"requestID": ctx.Value("requestID"),
		}).Error(err)
	}
}
