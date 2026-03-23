package utils

import (
	"reflect"
	"regexp"

	"github.com/spf13/cast"
)

var Is *IsClass

type IsClass struct{}

// Email - 是否为邮箱
func (I *IsClass) Email(email any) (ok bool) {
	if email == nil {
		return false
	}
	return regexp.MustCompile(`\w[-\w.+]*@([A-Za-z0-9][-A-Za-z0-9]+\.)+[A-Za-z]{2,14}`).MatchString(cast.ToString(email))
}

// Phone - 是否为手机号
func (I *IsClass) Phone(phone any) (ok bool) {
	if phone == nil {
		return false
	}
	return regexp.MustCompile(`^1[3456789]\d{9}$`).MatchString(cast.ToString(phone))
}

// Empty - 是否为空
func (I *IsClass) Empty(value any) (ok bool) {
	if value == nil {
		return true
	}

	// 使用 reflect 包获取值的类型和值
	v := reflect.ValueOf(value)

	switch v.Kind() {
	case reflect.Array, reflect.Slice, reflect.Map, reflect.Chan, reflect.String, reflect.Ptr:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	default:
		return false
	}
}

// Domain - 是否为域名
func (I *IsClass) Domain(domain any) (ok bool) {
	if domain == nil {
		return false
	}
	return regexp.MustCompile(`^((https|http|ftp|rtsp|mms)?://)\S+`).MatchString(cast.ToString(domain))
}

// True - 是否为真
func (I *IsClass) True(value any) (ok bool) {
	return cast.ToBool(value)
}

// False - 是否为假
func (I *IsClass) False(value any) (ok bool) {
	return !cast.ToBool(value)
}

// Number - 是否为数字
func (I *IsClass) Number(value any) (ok bool) {
	if value == nil {
		return false
	}
	return regexp.MustCompile(`^[0-9]+$`).MatchString(cast.ToString(value))
}

// Float - 是否为浮点数
func (I *IsClass) Float(value any) (ok bool) {
	if value == nil {
		return false
	}
	return regexp.MustCompile(`^[0-9]+(.[0-9]+)?$`).MatchString(cast.ToString(value))
}

// Bool - 是否为bool
func (I *IsClass) Bool(value any) (ok bool) {
	return cast.ToBool(value)
}

// Accepted - 验证某个字段是否为为 yes, on, 或是 1
func (I *IsClass) Accepted(value any) (ok bool) {
	if value == nil {
		return false
	}
	return regexp.MustCompile(`^(yes|on|1)$`).MatchString(cast.ToString(value))
}

// Date - 是否为日期类型
func (I *IsClass) Date(date any) (ok bool) {
	if date == nil {
		return false
	}
	return regexp.MustCompile(`^\d{4}-\d{1,2}-\d{1,2}$`).MatchString(cast.ToString(date))
}

// Alpha - 只能包含字母
func (I *IsClass) Alpha(value any) (ok bool) {
	if value == nil {
		return false
	}
	return regexp.MustCompile(`^[a-zA-Z]+$`).MatchString(cast.ToString(value))
}

// AlphaNum - 只能包含字母和数字
func (I *IsClass) AlphaNum(value any) (ok bool) {
	if value == nil {
		return false
	}
	return regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString(cast.ToString(value))
}

// AlphaDash - 只能包含字母、数字和下划线_及破折号-
func (I *IsClass) AlphaDash(value any) (ok bool) {
	if value == nil {
		return false
	}
	return regexp.MustCompile(`^[a-zA-Z0-9_-]+$`).MatchString(cast.ToString(value))
}

// Chs - 是否为汉字
func (I *IsClass) Chs(value any) (ok bool) {
	if value == nil {
		return false
	}
	return regexp.MustCompile(`^[\x{4e00}-\x{9fa5}]+$`).MatchString(cast.ToString(value))
}

// ChsAlpha - 只能是汉字、字母
func (I *IsClass) ChsAlpha(value any) (ok bool) {
	if value == nil {
		return false
	}
	return regexp.MustCompile(`^[\x{4e00}-\x{9fa5}a-zA-Z]+$`).MatchString(cast.ToString(value))
}

// ChsAlphaNum - 只能是汉字、字母和数字
func (I *IsClass) ChsAlphaNum(value any) (ok bool) {
	if value == nil {
		return false
	}
	return regexp.MustCompile(`^[\x{4e00}-\x{9fa5}a-zA-Z0-9]+$`).MatchString(cast.ToString(value))
}

// ChsDash - 只能是汉字、字母、数字和下划线_及破折号-
func (I *IsClass) ChsDash(value any) (ok bool) {
	if value == nil {
		return false
	}
	return regexp.MustCompile(`^[\x{4e00}-\x{9fa5}a-zA-Z0-9_-]+$`).MatchString(cast.ToString(value))
}

// Cntrl - 是否为控制字符 - （换行、缩进、空格）
func (I *IsClass) Cntrl(value any) (ok bool) {
	if value == nil {
		return false
	}
	return regexp.MustCompile(`^[\x00-\x1F\x7F]+$`).MatchString(cast.ToString(value))
}

// Graph - 是否为可见字符 - （除空格外的所有可打印字符）
func (I *IsClass) Graph(value any) (ok bool) {
	if value == nil {
		return false
	}
	return regexp.MustCompile(`^[\x21-\x7E]+$`).MatchString(cast.ToString(value))
}

// Lower - 是否为小写字母
func (I *IsClass) Lower(value any) (ok bool) {
	if value == nil {
		return false
	}
	return regexp.MustCompile(`^[a-z]+$`).MatchString(cast.ToString(value))
}

// Upper - 是否为大写字母
func (I *IsClass) Upper(value any) (ok bool) {
	if value == nil {
		return false
	}
	return regexp.MustCompile(`^[A-Z]+$`).MatchString(cast.ToString(value))
}

// Space - 是否为空白字符 - （空格、制表符、换页符等）
func (I *IsClass) Space(value any) (ok bool) {
	if value == nil {
		return false
	}
	return regexp.MustCompile(`^[\s]+$`).MatchString(cast.ToString(value))
}

// Xdigit - 是否为十六进制字符 - （0-9、a-f、A-F）
func (I *IsClass) Xdigit(value any) (ok bool) {
	if value == nil {
		return false
	}
	return regexp.MustCompile(`^[\da-fA-F]+$`).MatchString(cast.ToString(value))
}

// ActiveUrl - 是否为有效的域名或者IP
func (I *IsClass) ActiveUrl(value any) (ok bool) {
	if value == nil {
		return false
	}
	return value == "localhost" || regexp.MustCompile(`^([a-z0-9-]+\.)+[a-z]{2,6}$`).MatchString(cast.ToString(value)) || regexp.MustCompile(`^(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])\.(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])\.(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])\.(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])$`).MatchString(cast.ToString(value))
}

// Ip - 是否为IP
func (I *IsClass) Ip(ip any) (ok bool) {
	if ip == nil {
		return false
	}
	return regexp.MustCompile(`^(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])\.(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])\.(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])\.(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])$`).MatchString(cast.ToString(ip))
}

// Port - 是否为端口号
func (I *IsClass) Port(value any) (ok bool) {

	if value == nil {
		return false
	}
	return regexp.MustCompile(`^([0-9]|[1-9]\d|[1-9]\d{2}|[1-9]\d{3}|[1-5]\d{4}|6[0-4]\d{3}|65[0-4]\d{2}|655[0-2]\d|6553[0-5])$`).MatchString(cast.ToString(value))

}

// Url - 是否为URL
func (I *IsClass) Url(url any) (ok bool) {
	if url == nil {
		return false
	}
	return regexp.MustCompile(`^((https|http|ftp|rtsp|mms)?://)\S+`).MatchString(cast.ToString(url))
}

// IdCard - 是否为有效的身份证号码
func (I *IsClass) IdCard(value any) (ok bool) {
	if value == nil {
		return false
	}
	return regexp.MustCompile(`(^\d{15}$)|(^\d{18}$)|(^\d{17}(\d|X|x)$)`).MatchString(cast.ToString(value))
}

// MacAddr - 是否为有效的MAC地址
func (I *IsClass) MacAddr(value any) (ok bool) {
	if value == nil {
		return false
	}
	return regexp.MustCompile(`^([A-Fa-f0-9]{2}:){5}[A-Fa-f0-9]{2}$`).MatchString(cast.ToString(value))
}

// Zip - 是否为有效的邮政编码
func (I *IsClass) Zip(value any) (ok bool) {
	if value == nil {
		return false
	}
	return regexp.MustCompile(`^[1-9]\d{5}$`).MatchString(cast.ToString(value))
}

// String - 是否为字符串
func (I *IsClass) String(value any) (ok bool) {
	if value == nil {
		return false
	}
	return reflect.TypeOf(value).Kind() == reflect.String
}

// Slice - 是否为切片
func (I *IsClass) Slice(value any) (ok bool) {
	if value == nil {
		return false
	}
	return reflect.TypeOf(value).Kind() == reflect.Slice
}

// Array - 是否为数组
func (I *IsClass) Array(value any) (ok bool) {
	if value == nil {
		return false
	}
	return reflect.TypeOf(value).Kind() == reflect.Array
}

// JsonString - 是否为json字符串
func (I *IsClass) JsonString(value any) (ok bool) {
	if value == nil {
		return false
	}
	return regexp.MustCompile(`^[{\[].*[}\]]$`).MatchString(cast.ToString(value))
}

// Map - 是否为map
func (I *IsClass) Map(value any) (ok bool) {
	if value == nil {
		return false
	}
	return reflect.TypeOf(value).Kind() == reflect.Map
}

// SliceSlice - 是否为二维切片
func (I *IsClass) SliceSlice(value any) (ok bool) {
	if value == nil {
		return false
	}
	return reflect.TypeOf(value).Kind() == reflect.Slice && reflect.TypeOf(value).Elem().Kind() == reflect.Slice
}

// MapAny - 是否为[]map[string]any
func (I *IsClass) MapAny(value any) (ok bool) {
	if value == nil {
		return false
	}
	return reflect.TypeOf(value).Kind() == reflect.Map && reflect.TypeOf(value).Elem().Kind() == reflect.Interface
}

// IsString - 是否为字符串
func IsString(value any) (ok bool) {
	if value == nil {
		return false
	}
	return reflect.TypeOf(value).Kind() == reflect.String
}

// IsSlice - 是否为切片
func IsSlice(value any) (ok bool) {
	if value == nil {
		return false
	}
	return reflect.TypeOf(value).Kind() == reflect.Slice
}

// IsArray - 是否为数组
func IsArray(value any) (ok bool) {
	if value == nil {
		return false
	}
	return reflect.TypeOf(value).Kind() == reflect.Array
}
