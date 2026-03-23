package utils

var Map *mapCls

type mapCls struct{}

// Assign 合并两个 map[string]any 类型的对象，将 obj2 中的键值对赋值给 obj1
func (m *mapCls) Assign(obj1 map[string]any, obj2 map[string]any) map[string]any {
	for k, v := range obj2 {
		obj1[k] = v
	}
	return obj1
}

// Keys 返回 obj 中所有的键
func (m *mapCls) Keys(obj map[string]any) []string {
	keys := make([]string, 0, len(obj))
	for k := range obj {
		keys = append(keys, k)
	}
	return keys
}

// Values 返回 obj 中所有的值
func (m *mapCls) Values(obj map[string]any) []any {
	values := make([]any, 0, len(obj))
	for _, v := range obj {
		values = append(values, v)
	}
	return values
}
