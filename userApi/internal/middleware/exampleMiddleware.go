package middleware

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
)

type ExampleMiddleware struct {
}

func NewExampleMiddleware() *ExampleMiddleware {
	return &ExampleMiddleware{}
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
		ctx := context.Background()
		r.WithContext(context.WithValue(ctx, "info", "123321"))
		logx.WithContext(r.Context()).Info("global 前面执行")
		next(w, r)
		logx.WithContext(r.Context()).Info("global 后面面执行")
	}
}
