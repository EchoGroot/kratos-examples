package log

import (
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
)

// LogFatal 替代kratos的log.Fatal()，log.Fatal()内部调用os.Exit()，会造成defer函数无法执行，这里调用panic
// 如果已经用logrus初始化kratos的log，则会在kratosx/contrib/log/logrus/logrus.go:KratosLogLevel2LogrusLevel处panic
func LogFatal(args ...interface{}) {
	log.Log(log.LevelFatal, log.DefaultMessageKey, fmt.Sprint(args...))
	panic(nil)
}

// LogFatalf 替代kratos的log.Fatalf()，log.Fatalf()内部调用os.Exit()，会造成defer函数无法执行，这里调用panic
// 如果已经用logrus初始化kratos的log，则会在kratosx/contrib/log/logrus/logrus.go:KratosLogLevel2LogrusLevel处panic
func LogFatalf(format string, args ...interface{}) {
	log.Log(log.LevelFatal, log.DefaultMessageKey, fmt.Sprintf(format, args))
	panic(nil)
}

// LogFatalw 替代kratos的log.Fatalw()，log.Fatalw()内部调用os.Exit()，会造成defer函数无法执行，这里调用panic
// 如果已经用logrus初始化kratos的log，则会在kratosx/contrib/log/logrus/logrus.go:KratosLogLevel2LogrusLevel处panic
func LogFatalw(keyvals ...interface{}) {
	log.Log(log.LevelFatal, keyvals...)
	panic(nil)
}
