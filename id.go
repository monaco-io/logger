package logger

import (
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc/metadata"
)

type (
	contextLogger struct {
		context.Context
	}
)

var (
	xRequestID = "x-request-id"
)

func WithContext(ctx *context.Context) *contextLogger {
	md, ok := metadata.FromOutgoingContext(*ctx)
	if !ok || len(md[xRequestID]) < 1 {
		*ctx = metadata.AppendToOutgoingContext(*ctx, xRequestID, uuid.New().String())
	}

	return &contextLogger{*ctx}
}
