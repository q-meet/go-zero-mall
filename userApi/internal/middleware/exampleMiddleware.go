package middleware

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc/metadata"
	"math/rand"
	"net/http"
	commonTrace "rpc-common/trace"
	"rpc-common/util"
	"time"
)

type ExampleMiddleware struct {
}

func NewExampleMiddleware() *ExampleMiddleware {
	return &ExampleMiddleware{}
}

func (m *ExampleMiddleware) BizTraceHandler(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		traceId := request.Header.Get("trace-id")
		if traceId == "" {
			next.ServeHTTP(writer, request)
			return
		}

		ctx := request.Context()
		ctx = commonTrace.NewContext(ctx, traceId)

		request = request.WithContext(ctx)
		logx.WithContext(request.Context()).Info("trace-id:" + traceId)
		next.ServeHTTP(writer, request)
	}
}

func (m *ExampleMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation

		next(w, r)
		// Passthrough to next handler if need
	}
}

func (*ExampleMiddleware) RegAndLoginHandler(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logx.WithContext(r.Context()).Info("login 和 register 前面执行")
		next(w, r)
		logx.WithContext(r.Context()).Info("login 和 register 后面面执行")
	}
}

func (*ExampleMiddleware) GlobalHandler(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//r = r.WithContext(context.WithValue(ctx, "info", "123321"))
		// 从请求中获取 trace_id，假设它是通过 HTTP 头部传递的
		traceID := r.Header.Get("trace_id")
		if traceID == "" {
			traceID = util.Uuid()
		}
		if traceID != "" {
			// 将 trace_id 添加到日志字段

			//ctx := context.WithValue(r.Context(), "TraceID", traceID)
			//r = r.WithContext(ctx)
			//WithContext(r.Context()).Info("xxxx")
			//ctx := logx.ContextWithFields(r.Context(), logx.Field("path", r.RequestURI), logx.Field(traceIdKey, traceID))
			//r = r.WithContext(ctx)
		}

		//next = BizTraceHandler()
		ctx := r.Context()

		ctx = context.WithValue(ctx, "xx", "xx1123")
		TraceState := trace.TraceState{}
		TraceState, _ = TraceState.Insert("foo", "bar")
		traceIdStr := generateTraceID(32)
		spanIDStr := generateTraceID(16)
		traceId, err := trace.TraceIDFromHex(traceIdStr)
		spanID, err := trace.SpanIDFromHex(spanIDStr)
		fmt.Println("----------------")
		fmt.Println("----------------traceIdStr:", traceIdStr, err)
		fmt.Println("----------------spanIDStr:", spanIDStr, err)
		fmt.Println("----------------")
		span := trace.SpanContextConfig{
			TraceID:    traceId,
			SpanID:     spanID,
			TraceFlags: 0x1,
			TraceState: TraceState,
			Remote:     true,
		}
		spanCtx := trace.NewSpanContext(span)

		ctx = trace.ContextWithRemoteSpanContext(ctx, spanCtx)

		md := propagation.MapCarrier{
			"xxx": "222",
		}
		otel.GetTextMapPropagator().Inject(ctx, md)
		// 设置
		ctx = metadata.NewOutgoingContext(ctx, metadata.New(md))

		// 获取
		//md, _ := metadata.FromIncomingContext(ctx)
		//mp := propagation.MapCarrier{}
		//for key, val := range md {
		//	mp[key] = val[0]
		//}
		//fmt.Println("---------------")
		//fmt.Println(mp)
		//fmt.Println("---------------")
		//
		//ctx = otel.GetTextMapPropagator().Extract(ctx, mp)
		r = r.WithContext(ctx)

		logx.WithContext(r.Context()).Info("global 前面执行")
		next(w, r)
		logx.WithContext(r.Context()).Info("global 后面面执行")
	}
}

func generateTraceID(length int) string {
	rand.Seed(time.Now().UnixNano())
	const charset = "abcdef012345678901234567890123456789"
	//6261559c186f37f6f8e7018615569d1d
	traceID := make([]byte, length)
	for i := 0; i < length; i++ {
		traceID[i] = charset[rand.Intn(len(charset))]
	}

	return string(traceID)
}
