package utils

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bytedance/sonic"
)

var Json = new(JsonCls)

type JsonCls struct{}

// Encode 编码为 JSON 字符串
func (j *JsonCls) Encode(data any) string {
	if data == nil {
		return ""
	}
	text, err := sonic.MarshalString(data)
	if err != nil {
		return ""
	}
	return text
}

// Decode 解码 JSON 字符串到 map
func (j *JsonCls) Decode(data any) (any, error) {
	var result any

	switch v := data.(type) {
	case string:
		err := sonic.UnmarshalString(v, &result)
		return result, err
	case []byte:
		err := sonic.Unmarshal(v, &result)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported type for Decode: %T", data)
	}
}

// IsJson 判断是否为有效的 JSON 格式
func (j *JsonCls) IsJson(v any) bool {
	if v == nil {
		return false
	}

	// 使用 sonic.Valid 高性能验证 JSON 格式
	switch val := v.(type) {
	case string:
		return sonic.ValidString(val)
	case []byte:
		return sonic.Valid(val)
	}

	// 对于其他 Go 结构体/类型，检查其是否能被成功序列化
	// 注意：sonic.Marshal 如果不报错，产出的必定是合法的 JSON 字节
	_, err := sonic.Marshal(v)
	return err == nil
}

// Get 从 map 数据中按路径获取值
func (j *JsonCls) Get(JSONData map[string]any, path string) any {
	if path == "" || JSONData == nil {
		return JSONData
	}

	parts := strings.Split(path, ".")
	var current any = JSONData

	for _, part := range parts {
		if current == nil {
			return nil
		}

		switch v := current.(type) {
		case map[string]any:
			current = v[part]
		case []any:
			idx, err := strconv.Atoi(part)
			if err == nil && idx >= 0 && idx < len(v) {
				current = v[idx]
			} else {
				return nil
			}
		default:
			// 遇到既不是 map 也不是 slice 的中间节点，直接返回 nil
			return nil
		}
	}
	return current
}
