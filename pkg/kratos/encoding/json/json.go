package json

import (
	"github.com/go-kratos/kratos/v2/encoding/json"
	"google.golang.org/protobuf/encoding/protojson"
)

// ConfigMarshal 配置序列化，用于接口返回值序列化
func ConfigMarshal() {
	json.MarshalOptions = protojson.MarshalOptions{
		// 枚举值使用数字
		UseEnumNumbers: true,
		// 发送未填充字段
		EmitUnpopulated: true,
	}
}
