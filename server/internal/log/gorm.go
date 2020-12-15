package log

import (
	"context"
	"time"

	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
)

const (
	slowThreshold time.Duration = 0
)

// GormLogger is a custom logger for Gorm, making it use logrus.
type GormLogger struct{}

func (l *GormLogger) LogMode(logger.LogLevel) logger.Interface {
	return l
}

func (*GormLogger) Info(ctx context.Context, template string, args ...interface{}) {
	WithContext(ctx).Infof(template, args...)
}

func (*GormLogger) Warn(ctx context.Context, template string, args ...interface{}) {
	WithContext(ctx).Warnf(template, args...)
}

func (*GormLogger) Error(ctx context.Context, template string, args ...interface{}) {
	WithContext(ctx).Errorf(template, args...)
}

func (*GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)

	sql, rows := fc()

	lg := WithContext(ctx).WithFields(Fields{
		"elapsed": float64(elapsed.Nanoseconds()) / 1e6,
		"file": utils.FileWithLineNum(),
		"rows": rows,
		"sql":  sql,
	})
	if err != nil {
		lg.Warnf("error: %#v", err)
		lg.WithError(err).Error("SQL error")
	} else if elapsed > slowThreshold && slowThreshold > 0 {
		lg.Warnf("SLOW SQL >= %v", slowThreshold)
	} else {
		lg.Infof("SQL executed successfully")
	}
}
