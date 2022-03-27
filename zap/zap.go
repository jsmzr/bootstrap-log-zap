package zap

import (
	"github.com/jsmzr/bootstrap-log/log"

	"go.uber.org/zap"
)

type ZapConfig struct{}

type ZapContainer struct {
	logger *zap.Logger
}

func (c *ZapConfig) Load() (log.Logger, error) {
	// TODO 提供配置支持
	return &ZapContainer{
		logger: zap.NewExample(),
	}, nil
}

// 牺牲些许性能使用 sugar 简化代码
func (c *ZapContainer) Debug(msg string, params ...interface{}) {
	c.logger.Sugar().Debugf(msg, params...)
}

func (c *ZapContainer) Info(msg string, params ...interface{}) {
	c.logger.Sugar().Infof(msg, params...)
}

func (c *ZapContainer) Warn(msg string, params ...interface{}) {
	c.logger.Sugar().Warnf(msg, params...)
}

func (c *ZapContainer) Error(msg string, params ...interface{}) {
	c.logger.Sugar().Errorf(msg, params...)
}

func init() {
	log.Register("zap", &ZapConfig{})
}
