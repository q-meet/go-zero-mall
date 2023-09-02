package middleware

import (
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"rpc-common/trace"
	"rpc-common/util"
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
		ctx = trace.NewContext(ctx, traceId)

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

		logx.WithContext(r.Context()).Info("global 前面执行")
		next(w, r)
		logx.WithContext(r.Context()).Info("global 后面面执行")
	}
}
