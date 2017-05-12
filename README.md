支持分级处理的日志模块, 默认输出到日志文件 ./output.log 同时在控制台输出

```
import "github.com/ruandao/log"

func main() {
    log.SetEnableFileLog(false) // 不输出到日志文件
    log.Debug("nihaoshijie")
    log.Error("hello world")
    log.SetLogLevel(log.LevelError) // Error 及以上级别的日志记录
    log.Debug("不显示")
    log.Error("我可以显示")

}
```