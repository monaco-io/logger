# logger

## Basic Usage
```go
import "github.com/monaco-io/logger"

func run(){
	logger.D("msg", "key", "val")
	logger.I("msg", "key", "val")
	logger.W("msg", "key", "val")
	logger.D("msg", "key", "val")
}

```

## New writer
```go
import "github.com/monaco-io/logger"

type slave struct{}

func (slave) Write(p []byte) (n int, err error) {
	fmt.Println("slave: ", string(p))
	return
}

func init(){
	logger.RegisterWriter(new(slave))
}

```
