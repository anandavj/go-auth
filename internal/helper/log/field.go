package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Field(pckg string, method string, msg string, err error, params ...map[string]string) (res []zapcore.Field) {
	if len(msg) == 0 {
		msg = "-"
	}
	// Core
	res = append(res,
		zap.String("Package", pckg),
		zap.String("Method", method),
		zap.String("Message", msg),
		zap.Error(err),
	)
	// Additional
	if len(params) > 0 {
		for key, el := range params[0] {
			res = append(res, zap.String(key, el))
		}
	}
	return res
}
