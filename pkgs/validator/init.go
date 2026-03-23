package validator

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

var ValidCls *validCls

type validCls struct {
	validate *validator.Validate
}

func init() {
	ValidCls = &validCls{validate: validator.New()}

	// 注册自定义验证器
	_ = ValidCls.validate.RegisterValidation("phone", validatePhone)       // 手机号验证器
	_ = ValidCls.validate.RegisterValidation("password", validatePassword) // 密码验证器
}

func Valid(data any) error {
	if err := ValidCls.validate.Struct(data); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			// 获取第一个错误
			firstErr := validationErrors[0]

			// 处理嵌套字段名
			fieldPath := strings.SplitN(firstErr.StructNamespace(), ".", 2)
			if len(fieldPath) > 1 {
				// 去掉结构体名，只保留字段路径
				fieldPath = strings.Split(fieldPath[1], ".")
			}

			// 获取字段别名路径
			fieldAliasPath := getFieldAliasPath(data, fieldPath)

			// 构造完整的字段别名路径
			fullFieldAlias := strings.Join(fieldAliasPath, "")

			// 获取验证错误信息
			errVal := firstErr.Param()
			msg := fmt.Sprintf(validMsg[firstErr.Tag()], errVal)
			if msg == "" {
				msg = "验证失败！"
			}

			return errors.New(fullFieldAlias + msg)
		}
		return err
	}
	return nil
}

// 递归获取字段的别名路径
func getFieldAliasPath(data interface{}, fieldPath []string) []string {
	var aliases []string
	current := data

	// 处理指针
	val := reflect.ValueOf(current)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
		current = val.Interface()
	}

	for i, fieldName := range fieldPath {
		t := reflect.TypeOf(current)
		if t.Kind() == reflect.Ptr {
			t = t.Elem()
		}

		// 如果是结构体数组/slice，取第一个元素
		if t.Kind() == reflect.Slice || t.Kind() == reflect.Array {
			if val.Len() > 0 {
				current = val.Index(0).Interface()
				t = reflect.TypeOf(current)
				if t.Kind() == reflect.Ptr {
					t = t.Elem()
				}
			}
		}

		field, found := t.FieldByName(fieldName)
		if !found {
			aliases = append(aliases, fieldName)
			continue
		}

		// 获取字段别名
		alias := field.Tag.Get("alias")
		if alias == "" {
			alias = fieldName
		}
		aliases = append(aliases, alias)

		// 如果是嵌套结构体，继续处理下一级
		if i < len(fieldPath)-1 {
			// 获取字段值
			fieldVal := val.FieldByName(fieldName)
			if fieldVal.Kind() == reflect.Ptr {
				if fieldVal.IsNil() {
					break
				}
				fieldVal = fieldVal.Elem()
			}

			current = fieldVal.Interface()
			val = fieldVal
		}
	}

	return aliases
}
