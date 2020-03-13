package log

import (
	"context"
	"github.com/gnehcaij/zeus/constant"
	"github.com/sirupsen/logrus"
)

func CtxError(ctx context.Context, format string, args ...interface{}) {
	logrus.WithField("logId", ctx.Value(constant.LOG_ID).(string)).Errorf(format, args...)
}

func CtxInfo(ctx context.Context, format string, args ...interface{}) {
	logrus.WithField("logId", ctx.Value(constant.LOG_ID).(string)).Infof(format , args...)
}

func CtxWarning(ctx context.Context, format string, args ...interface{}) {
	logrus.WithField("logId", ctx.Value(constant.LOG_ID).(string)).Warnf(format, args...)
}
