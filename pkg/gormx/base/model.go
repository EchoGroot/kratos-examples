package base

import (
	"github.com/EchoGroot/kratos-examples/pkg/lang/reflectx"
	"reflect"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Model 实体类的基类
// 泛型参数 T 是 Id 字段的类型
type Model[T string | int32] struct {
	Id         T         `gorm:"primaryKey"`
	CreateTime time.Time `gorm:"autoCreateTime"` // 创建时间
	UpdateTime time.Time `gorm:"autoUpdateTime"` // 更新时间
}

// BeforeCreate 利用gorm的钩子，在insert前给string类型的Id填充uuid
func (r *Model[T]) BeforeCreate(tx *gorm.DB) (err error) {
	idValue := reflectx.Indirect(reflect.ValueOf(&r.Id))
	idInterface := idValue.Interface()
	if idStr, ok := idInterface.(string); ok {
		if idStr == "" {
			uuidStr := uuid.NewString()
			idValue.SetString(uuidStr)
		}
	}
	return
}
