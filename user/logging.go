package user

import (
	"time"

	"github.com/go-kit/kit/log"
)

func LoggingMiddleware(logger log.Logger) ServiceMiddleware {
	return func(next Service) Service {
		return logmw{logger, next}
	}
}

type logmw struct {
	logger log.Logger
	Service
}

func (mw logmw) AddUser(firstName, lastName string) (output int, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "addUser",
			"input", firstName+lastName,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.Service.AddUser(firstName, lastName)
	return
}
