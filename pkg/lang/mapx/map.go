package mapx

import (
	"github.com/EchoGroot/kratos-examples/pkg/lang/reflectx"
	"reflect"
)

// DeleteNil 删除 map 中 value 为 nil 的元素
func DeleteNil(m map[string]interface{}) {
	for k, v := range m {
		// 如果 v 是 map
		indirectType := reflectx.Indirect(reflect.ValueOf(v))
		if indirectType.Kind() == reflect.Map {
			DeleteNil(v.(map[string]interface{}))
			continue
		}
		// 如果 v 是 数组
		indirectValue := reflectx.Indirect(reflect.ValueOf(v))
		if indirectValue.Kind() == reflect.Slice {
			for _, v := range v.([]interface{}) {
				indirectType = reflectx.Indirect(reflect.ValueOf(v))
				if indirectType.Kind() == reflect.Map {
					DeleteNil(v.(map[string]interface{}))
				}
			}
			continue
		}
		if v == nil {
			delete(m, k)
		}
	}
}
