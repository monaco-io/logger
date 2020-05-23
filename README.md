# logger

## Register

- Register once, ues anyware.
- switch debug level on code running.

## Basic Usage

```go
import "github.com/monaco-io/logger"

func run(){
    logger.I("msg", "key", "val")
    logger.D("msg", "key", "val")
    logger.W("msg", "key", "val")
    logger.E("msg", "key", "val")
}
// 2020-05-23T19:57:03.903+0800    INFO    logger/level_test.go:9     msg    {"key": "val"}
// 2020-05-23T19:57:03.903+0800    WARN    logger/level_test.go:11    msg    {"key": "val"}
// 2020-05-23T19:57:03.903+0800    ERROR   logger/level_test.go:12    msg    {"key": "val"}
// github.com/monaco-io/logger._handler
//     /root/code/logger/core.go:13
// github.com/monaco-io/logger.E
//     /root/code/logger/level.go:16
// github.com/monaco-io/logger.TestD
//     /root/code/logger/level_test.go:12
// testing.tRunner
//     /usr/local/go/src/testing/testing.go:991
```

## Define Service Name

```go
import "github.com/monaco-io/logger"

func init(){
    logger.RegisterServiceName("Monaco")
    I("msg", "key", "val")
}

// 2020-05-23T23:22:01.718+0800   INFO   Monaco   logger/level_test.go:80   msg   {"key": "val"}
// {"L":"INFO","T":"2020-05-23T23:22:01.718+0800","N":"Monaco","C":"logger/level_test.go:80","M":"msg","key":"val"}
```

## Enable Debug dimaic

```go
import "github.com/monaco-io/logger"

func init(){
    logger.D("msg", "key", "val")
    logger.RegisterDebug(true)
    logger.D("msg", "key2", "val2")
}

//  {"L":"DEBUG","T":"2020-05-23T19:58:59.323+0800","C":"logger/level_test.go:28","M":"msg","key2":"val2"}
```

## New writer

```go
import "github.com/monaco-io/logger"

func init(){
    logger.RegisterWriter(new(slave))
    logger.I("msg", "key", "val")
}

type slave struct{}

func (slave) Write(p []byte) (n int, err error) {
    // do some with log data
    fmt.Println("slave: ", string(p))
    return
}

// 2020-05-23T21:04:16.065+0800    INFO    logger/level_test.go:49    msg    {"key": "val"}
// slave:  {"L":"INFO","T":"2020-05-23T21:04:16.065+0800","C":"logger/level_test.go:49","M":"msg","key":"val"}
```
