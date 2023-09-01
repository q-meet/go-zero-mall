package util

import "github.com/google/uuid"

func Uuid() string {

	// 生成一个新的 UUID
	traceID := uuid.New()

	// 将 UUID 格式化为字符串
	return traceID.String()

}
