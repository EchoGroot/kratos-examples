package reflectx

import "reflect"

func Indirect(reflectValue reflect.Value) reflect.Value {
	for reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}
	return reflectValue
}

func IndirectType(reflectType reflect.Type) reflect.Type {
	for reflectType.Kind() == reflect.Ptr || reflectType.Kind() == reflect.Slice {
		reflectType = reflectType.Elem()
	}
	return reflectType
}

// Interface2Array 将 interface{} 转换为 []interface{}
// s 可以是一个数组或者单个元素
func Interface2Array(s interface{}) []interface{} {
	var arr []interface{}

	v := Indirect(reflect.ValueOf(s))
	if v.Kind() == reflect.Slice {
		for i := 0; i < v.Len(); i++ {
			arr = append(arr, v.Index(i).Interface())
		}
	} else {
		arr = append(arr, v.Interface())
	}
	return arr
}
