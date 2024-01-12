package errorsx

import (
	"fmt"
	"runtime"
)

func PanicToError(r interface{}) (err error) {
	buf := make([]byte, 64<<10)
	buf = buf[:runtime.Stack(buf, false)]
	err = fmt.Errorf("panic recovered: %s\n%s", r, buf)
	return
}
