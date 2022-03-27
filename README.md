# bootstrap/log-zap

bootstrap-log 系列 zap 日志的适配器

## 使用说明

1. 通过 `import (_ "github.com/jsmzr/bootstrap-log-zap/zap)"` 导入适配器
2. 通过 `import "github.com/jsmzr/bootstrap-global/log"` 导入门面
4. 通过 `err := log.InitLogger("zap")` 进行初始化
3. 代码引入`import "github.com/jsmzr/bootstrap-log/log` 调用 `log.Info("message %s", key)` 即可


## 开发说明

bootstrap-log 系列是为了提供统一的门面和依赖管理而产生的，期望我们在项目中可以屏蔽底层的实现，可以通过切换适配器快速接入、切换到其他实现。

### TODO

- [ ] 日志配置





