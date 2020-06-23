package logger

import (
	"fmt"
	"testing"

	"go.uber.org/zap"
)

type slave struct{}

func (slave) Write(p []byte) (n int, err error) {
	// do something with log data
	fmt.Println("slave: ", string(p))
	return
}

type slave2 struct{}

func (slave2) Write(p []byte) (n int, err error) {
	// do something with log data
	fmt.Println("slave2: ", string(p))
	return
}

func TestLevel(t *testing.T) {
	I("msg", "key", "val")
	D("msg", "key", "val")
	W("msg", "key", "val")
	E("msg", "key", "val")
}

func TestLevelPanic(t *testing.T) {
	defer func() {
		e := recover()
		if e == "this is a panic" {
			t.Fatal()
		}
	}()
	P("this is a panic", "key", "val")
}

func TestFatal(t *testing.T) {
	defer func() {
		t.Fatal("you can not see this line.")
	}()
	F("msg", "key", "val")
}

func TestRegisterDebug(t *testing.T) {
	if _autoLevel.Level() == zap.DebugLevel {
		t.Fail()
	}
	RegisterDebug(true)
	if _autoLevel.Level() != zap.DebugLevel {
		t.Fail()
	}
	RegisterDebug(false)
	if _autoLevel.Level() == zap.DebugLevel {
		t.Fail()
	}
}

func TestRegisterWriter(t *testing.T) {
	RegisterWriter(new(slave))
	I("msg", "key", "val")

	RegisterWriter(new(slave2))
	I("msg", "key", "val")
}

func TestRegisterWriter2(t *testing.T) {
	RegisterWriter(new(slave), new(slave2))
	I("msg", "key", "val")
}

func TestRegisterErrorWriter(t *testing.T) {
	RegisterErrorWriter(new(slave), new(slave2))
	RegisterServiceName("svc name")
	E("msg", "key", "val")
}

func TestRegisterServiceName(t *testing.T) {
	RegisterServiceName("Monaco")
	RegisterWriter(new(slave))
	I("msg", "key", "val")
}

func TestT(t *testing.T) {
	var mockStruct = struct{ key string }{key: "value"}
	type args struct {
		args []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "T",
			args: args{
				[]interface{}{"key1", mockStruct, "key2", mockStruct},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			T(tt.args.args...)
		})
	}
}
