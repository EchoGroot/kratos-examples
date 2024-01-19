package structx

import (
	fieldMaskUtils "github.com/EchoGroot/kratos-examples/pkg/fieldmask-utils"
	"github.com/EchoGroot/kratos-examples/pkg/lang/stringx"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

// StructToMapByFieldMask 根据fieldMask，struct转map
// 注意fieldMask里的field名称是蛇形（前端传的字段是小驼峰，pb定义的字段为蛇形，最终接收到的是蛇形），
// 而struct的字段是大驼峰，所以这里做了转换
// 已经wrap过了，切勿重复wrap
func StructToMapByFieldMask(fieldMask *fieldmaskpb.FieldMask, src interface{}) (map[string]interface{}, error) {
	field := make(map[string]interface{})
	mask, err := fieldMaskUtils.MaskFromProtoFieldMask(fieldMask, stringx.Snake2BigCamel)
	if err != nil {
		return nil, errors.Wrap(err, "mask from proto field mask error")
	}
	if err := fieldMaskUtils.StructToMap(mask, src, field); err != nil {
		return nil, errors.Wrap(err, "struct to map error")
	}
	return field, nil
}
