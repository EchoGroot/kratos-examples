package copier

import (
	"github.com/EchoGroot/kratos-examples/pkg/lang/reflectx"
	"reflect"
	"time"

	"github.com/pkg/errors"

	"google.golang.org/protobuf/proto"

	"github.com/mennanov/fmutils"

	jinzhuCopier "github.com/jinzhu/copier"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type copier struct {
	to   interface{}
	from interface{}
	// 已经wrap过了，切勿重复wrap
	Error error
}

// Copy 封装了jinzhu copier.Copy，用于对象拷贝、切片拷贝
// 注：这是浅拷贝，如果需要深拷贝，自己设置Option
func Copy(to interface{}, from interface{}) *copier {
	err := jinzhuCopier.Copy(to, from)
	if err != nil {
		err = errors.Wrap(err, "copy error")
	}
	return &copier{
		to:    to,
		from:  from,
		Error: err,
	}
}

func CopyWithOption(to interface{}, from interface{}, op jinzhuCopier.Option) *copier {
	err := jinzhuCopier.CopyWithOption(to, from, op)
	if err != nil {
		err = errors.Wrap(err, "copy error")
	}
	return &copier{
		to:    to,
		from:  from,
		Error: err,
	}
}

// Time2Timestamppb golang的time转Protocol Buffer的timestamp
func (c *copier) Time2Timestamppb() *copier {
	if c.Error != nil {
		return c
	}
	to := reflectx.Indirect(reflect.ValueOf(c.to))
	from := reflectx.Indirect(reflect.ValueOf(c.from))
	if to.Kind() == reflect.Slice && from.Kind() == reflect.Slice {
		for i := 0; i < from.Len(); i++ {
			timeMap := getTime(from.Index(i))
			setTimestamppb(to.Index(i), timeMap)
		}
	} else if to.Kind() == reflect.Struct && from.Kind() == reflect.Struct {
		timeMap := getTime(from)
		setTimestamppb(to, timeMap)
	}
	return c
}

func getTime(fromValue reflect.Value) map[string]time.Time {
	timeMap := make(map[string]time.Time)
	fromValue = reflectx.Indirect(fromValue)
	recursiveGetExtendTimeField(fromValue, timeMap)
	return timeMap
}

// recursiveGetExtendTimeField 递归获取内嵌结构体的time.Time字段
func recursiveGetExtendTimeField(fromValue reflect.Value, timeMap map[string]time.Time) {
	fromType := reflectx.IndirectType(fromValue.Type())
	for i := 0; i < fromType.NumField(); i++ {
		fromField := fromType.Field(i)
		fromFieldType := reflectx.IndirectType(fromField.Type)
		if fromField.Anonymous {
			recursiveGetExtendTimeField(fromValue.Field(i), timeMap)
		} else {
			if fromFieldType.Name() == "Time" {
				timeMap[fromField.Name] = fromValue.Field(i).Interface().(time.Time)
			}
		}
	}
}

func setTimestamppb(toValue reflect.Value, timeMap map[string]time.Time) {
	toValue = reflectx.Indirect(toValue)
	toType := reflectx.IndirectType(toValue.Type())
	for i := 0; i < toType.NumField(); i++ {
		toField := toType.Field(i)
		toFieldType := reflectx.IndirectType(toField.Type)
		if toFieldType.Name() == "Timestamp" {
			if value, ok := timeMap[toField.Name]; ok {
				toValue.Field(i).Set(reflect.ValueOf(timestamppb.New(value)))
			}
		}
	}
}

// Filter 根据字段名进行过滤，只能用于Protocol Buffer定义的message
func (c *copier) Filter(fieldMask []string) *copier {
	if c.Error != nil {
		return c
	}
	toValue := reflectx.Indirect(reflect.ValueOf(c.to))
	if toValue.Kind() == reflect.Slice {
		for i := 0; i < toValue.Len(); i++ {
			toItemValue := toValue.Index(i)
			// 不能使用 toItemValue := reflectx.Indirect(toValue.Index(i))
			// 这里必须是指针的value，否则通过反射无法获取方法，但能获取字段
			// 会导致转接口时转换失败，因为没有获取到实现的方法，被坑了两小时
			fmutils.Filter(toItemValue.Interface().(proto.Message), fieldMask)
		}
	} else if toValue.Kind() == reflect.Struct {
		fmutils.Filter(c.to.(proto.Message), fieldMask)
	}
	return c
}
