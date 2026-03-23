package response

import (
	e "one-dock/app/error"
	"one-dock/app/logger"

	"github.com/gofiber/fiber/v3"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Success 统一成功返回
func Success(c fiber.Ctx, data interface{}, msg ...string) error {
	message := "success"
	if len(msg) > 0 {
		message = msg[0]
	}
	return c.JSON(Response{
		Code: 200,
		Msg:  message,
		Data: data,
	})
}

// Fail 统一失败返回
func Fail(c fiber.Ctx, code int, msg string) error {
	return c.JSON(Response{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}

// ErrorHandler Fiber 全局错误拦截器
func ErrorHandler(logFactor *logger.LogFactor) func(c fiber.Ctx, err error) error {

	return func(c fiber.Ctx, err error) error {

		logFactor.Error(map[string]interface{}{
			"ip":     c.IP(),
			"path":   c.Path(),
			"method": c.Method(),
			"error":  err.Error(),
		})

		// 判断是否为我们自定义的业务错误
		if customErr, ok := err.(*e.CustomError); ok {

			return c.Status(fiber.StatusOK).JSON(Response{
				Code: customErr.Code,
				Msg:  customErr.Msg,
				Data: nil,
			})
		}

		// 处理 Fiber 内置错误 (如 404, 405 等)
		if fiberErr, ok := err.(*fiber.Error); ok {
			return c.Status(fiberErr.Code).JSON(Response{
				Code: fiberErr.Code,
				Msg:  fiberErr.Message,
				Data: nil,
			})
		}

		// 未知内部错误兜底
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Code: 500,
			Msg:  "服务器内部错误: " + err.Error(),
			Data: nil,
		})
	}
}
