package trace

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"go.opentelemetry.io/otel/baggage"
)

func NewContext(ctx context.Context, traceId string) context.Context {
	ctx = logx.ContextWithFields(ctx, logx.Field(traceIdKey, traceId))

	bg := baggage.FromContext(ctx)
	member, err := baggage.NewMember(traceIdKey, traceId)
	if err != nil {
		logx.WithContext(ctx).Error(err)
		return ctx
	}

	bg, err = bg.SetMember(member)
	if err != nil {
		logx.WithContext(ctx).Error(err)
		return ctx
	}

	ctx = baggage.ContextWithBaggage(ctx, bg)

	return ctx
}

const traceIdKey = "biz-trace-id"

func FromTraceId(ctx context.Context) (string, bool) {
	bg := baggage.FromContext(ctx)
	member := bg.Member(traceIdKey)
	return member.Value(), member.Key() != ""
}

func WithContext(ctx context.Context) logx.Logger {
	traceId, ok := FromTraceId(ctx)
	if !ok {
		return logx.WithContext(ctx)
	}

	return logx.WithContext(ctx).WithFields(logx.Field(traceIdKey, traceId))
}
