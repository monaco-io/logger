package logger

import (
	"fmt"
	"testing"

	"go.uber.org/zap"
)

type slave struct{}

func (slave) Write(p []byte) (n int, err error) {
	// do some with log data
	fmt.Println("slave: ", string(p))
	return
}

type slave2 struct{}

func (slave2) Write(p []byte) (n int, err error) {
	// do some with log data
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
