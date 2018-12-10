package log

import (
	"fmt"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// Logger is any type that is capable of logging rfs application output
type Logger interface {
	Debugw(msg string, keysAndValues ...interface{})
	Error(msg string, err error)
	// Info(args ...interface{})
	Infow(msg string, keysAndValues ...interface{})
	Warnw(msg string, keysAndValues ...interface{})
	Sync() error
}

// zapLogger is an implementation of Logger that wraps the uber-zap logger
type zapLogger struct {
	*zap.SugaredLogger
}

// NewZapSugaredLogger returns a Logger implemented by uber's zap logger
// The strings 'p', 'prod', and 'production' are the only valid lvl arg
// that will change
func NewZapSugaredLogger(lvl string) (Logger, error) {
	zl := zapLogger{}
	switch lvl {
	case "production":
		pl, err := zap.NewProduction()
		if err != nil {
			return nil, errors.Wrap(err, "could not instantiate prod sugared logger")
		}

		zl.SugaredLogger = pl.Sugar()
		zl.SugaredLogger.Info("Logging level set at 'Production'")
	default:
		// assume development unless told production
		dl, err := zap.NewDevelopment()
		if err != nil {
			return nil, errors.Wrap(err, "could not instantiate dev sugared logger")
		}

		zl.SugaredLogger = dl.Sugar()
		zl.SugaredLogger.Info("Logging level set at 'Development'")
	}

	return &zl, nil
}

func (zl *zapLogger) Error(msg string, err error) {
	zl.Errorw(msg,
		"StackTrace", fmt.Sprintf("%+s", err),
	)
}
