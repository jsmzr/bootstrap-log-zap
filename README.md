# bootstrap/log-zap

bootstrap 系列 zap 日志的适配器

## 使用说明

1. 通过 `import (_ "github.com/jiurongzhao/bootstrap-log-zap/zap)"` 导入适配器
2. 通过 `import "github.com/jiurongzhao/bootstrap-global/log"` 导入门面
4. 通过 `err := log.InitGlobalLogger("zap")` 进行初始化
3. 代码中调用 `log.Info("message xxx", key, "foo", key2, 123)`


## 开发说明

bootstrap 系列是为了提供统一的门面和依赖管理而产生的，期望我们在项目中可以屏蔽底层的实现，可以通过切换适配器快速接入、切换到其他实现。

### TODO

- [ ] format 配置支持
- [ ] 各适配器输出格式统一
- [ ] 性能调优


### 问题

由于基于统一的门面进行适配，在适配时很难兼容各种实现的特性。

在 zap 中提供了对应的数据类型参数来避免底层大量使用反射而导致的性能消耗，而在业务代码中如使用 zapcore.Field 作为参数势必会让门面失去意义。 


