package validator

import (
	"one-dock/pkgs/utils"
	"regexp"

	"github.com/go-playground/validator/v10"
)

// 手机号校验
func validatePhone(fl validator.FieldLevel) bool {
	return utils.Is.Phone(fl.Field().String())
}

// 字母数字特殊字符组合校验
func validatePassword(fl validator.FieldLevel) bool {
	alphanumSpecialRegex := regexp.MustCompile(`^[a-zA-Z0-9!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]+$`)
	return alphanumSpecialRegex.MatchString(fl.Field().String())
}
