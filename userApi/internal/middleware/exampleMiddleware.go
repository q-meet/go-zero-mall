package middleware

import (
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
		logx.Info("login 和 register 前面执行")
		next(w, r)
		logx.Info("login 和 register 后面面执行")
	}
}
