package zap

import (
	"fmt"

	"github.com/jiurongzhao/bootstrap-global/log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapConfig struct{}

type ZapContainer struct {
	logger *zap.Logger
}

func (c *ZapConfig) Load() (log.Logger, error) {
	return &ZapContainer{
		logger: zap.NewExample(),
	}, nil
}
func (c *ZapContainer) Debug(msg string, params ...interface{}) {
	logs(c.logger.Debug, msg, params...)
}

func (c *ZapContainer) Info(msg string, params ...interface{}) {
	logs(c.logger.Info, msg, params...)
}

func (c *ZapContainer) Warn(msg string, params ...interface{}) {
	logs(c.logger.Warn, msg, params...)
}

func (c *ZapContainer) Error(msg string, params ...interface{}) {
	logs(c.logger.Error, msg, params...)
}

func logs(f func(msg string, params ...zapcore.Field), msg string, params ...interface{}) {
	if len(params) == 0 {
		f(msg)
	} else {
		f(msg, convertParam(params...)...)
	}
}

func convertParam(params ...interface{}) []zapcore.Field {
	size := int(len(params) / 2)
	fields := make([]zapcore.Field, size)
	for index := 0; index < size; index++ {
		key := fmt.Sprintf("%s", params[2*index])
		value := params[2*index+1]
		switch t := value.(type) {
		case int:
			fields[index] = zap.Int(key, t)
		case string:
			fields[index] = zap.String(key, t)
		case float64:
			fields[index] = zap.Float64(key, t)
		case bool:
			fields[index] = zap.Bool(key, t)
		}
	}
	return fields
}

func init() {
	log.Register("zap", &ZapConfig{})
}
