package middleware

import (
	"context"
	"errors"
	"one-dock/app/config"
	e "one-dock/app/error"
	"one-dock/core/comm"
	"strings"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/cast"
)

// 定义 context key 类型，避免 key 冲突

// GetUserMiddleware 创建用户认证中间件
// 从 Authorization header 或 cookie 中提取 JWT token 并验证
func GetUserMiddleware(cfg *config.Cfg) fiber.Handler {
	return func(ctx fiber.Ctx) error {
		// 如果用户已经存在，直接跳过
		if ctx.Context().Value("user") != nil {
			return ctx.Next()
		}

		// 提取 token
		token := extractToken(ctx)
		if token == "" {
			return ctx.Next()
		}

		// 解析和验证 token
		claims, err := parseAndValidateToken(token, cfg)
		if err != nil {
			// token 无效时继续执行，不中断请求
			setContextError(ctx, err)
			return ctx.Next()
		}

		// 检查 token 是否过期
		if isTokenExpired(claims) {
			setContextError(ctx, e.New(403, "登录过期！"))
			return ctx.Next()
		}

		// 从 claims 中提取用户状态数据
		userData, ok := claims["data"]
		if !ok {
			setContextError(ctx, e.New(403, "登录凭证数据无效！"))
			return ctx.Next()
		}

		// 检查用户状态
		if err := validateUserStatus(userData); err != nil {
			setContextError(ctx, err)
			return ctx.Next()
		}

		// 设置用户信息到上下文
		setContextUser(ctx, userData)

		return ctx.Next()
	}
}

// extractToken 从请求中提取 JWT token
// 优先级：Authorization header > Cookie
func extractToken(ctx fiber.Ctx) string {
	// 从 Authorization header 获取
	authHeader := ctx.Get("Authorization")
	if authHeader != "" {
		// 支持 Bearer token 格式
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) == 2 && strings.EqualFold(parts[0], "Bearer") {
			return parts[1]
		}
		// 兼容直接传 token 的情况
		return authHeader
	}

	// 从 Cookie 获取
	return ctx.Cookies("Authorization")
}

// parseAndValidateToken 解析并验证 JWT token
func parseAndValidateToken(tokenStr string, cfg *config.Cfg) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// 验证签名方法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(cfg.JWT.Secret), nil
	})

	if err != nil {
		return nil, wrapTokenError(err)
	}

	if !token.Valid {
		return nil, e.New(401, "身份验证失败！")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, e.New(401, "token 无效！")
	}

	return claims, nil
}

// isTokenExpired 检查 token 是否过期
// exp 字段应为 Unix 时间戳（秒）
func isTokenExpired(claims jwt.MapClaims) bool {
	exp, ok := claims["exp"].(int64)
	if !ok {
		// 如果没有 exp 字段，认为 token 永不过期
		return false
	}

	// 转换为时间戳进行比较（exp 通常是 Unix 时间戳，单位：秒）
	expTime := time.UnixMilli(exp)
	return time.Now().After(expTime)
}

// validateUserStatus 验证用户状态是否正常
func validateUserStatus(userData interface{}) error {
	userStatus, ok := cast.ToStringMap(userData)["status"]
	if !ok {
		return e.New(400, "用户状态信息缺失！")
	}

	if comm.UserStatus(cast.ToInt(userStatus)) != comm.UserStatusActive {
		return e.New(403, "用户被封禁！")
	}

	return nil
}

// setContextError 设置错误信息到上下文
func setContextError(ctx fiber.Ctx, err error) {
	c := context.WithValue(ctx.Context(), "err", err)
	ctx.SetContext(c)
}

// setContextUser 设置用户信息到上下文
func setContextUser(ctx fiber.Ctx, userData interface{}) {
	c := context.WithValue(ctx.Context(), "user", userData)
	ctx.SetContext(c)
}

// wrapTokenError 包装 JWT 解析错误，提供更友好的错误信息
func wrapTokenError(err error) error {
	if errors.Is(err, jwt.ErrTokenExpired) {
		return e.New(403, "token 已过期！")
	}
	if errors.Is(err, jwt.ErrTokenSignatureInvalid) {
		return e.New(401, "token 签名无效！")
	}
	if errors.Is(err, jwt.ErrTokenMalformed) {
		return e.New(401, "token 格式错误！")
	}
	return e.New(401, "禁止非法操作！")
}
