package syncx

import (
	"github.com/EchoGroot/kratos-examples/pkg/lang/errorsx"
	"github.com/go-kratos/kratos/v2/log"
)

// Go 具备recover能力的新建协程异步执行
func Go(fn func()) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				err := errorsx.PanicToError(r)
				log.Errorf("Go panic: %+v", err)
			}
		}()
		fn()
	}()
}
