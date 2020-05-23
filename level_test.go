package logger

import "testing"

func TestD(t *testing.T) {
	I("msg xxx", "key xxx", "val xxx", "key xxx2", map[string]string{"mkey":"mval"})
	D("msg xxx", "key xxx", "val xxx", "key xxx2", map[string]string{"mkey":"mval"})
	W("msg xxx", "key xxx", "val xxx", "key xxx2", map[string]string{"mkey":"mval"})
	E("msg xxx", "key xxx", "val xxx", "key xxx2", map[string]string{"mkey":"mval"})
}
