package bootstrap

import "flag"

// CommandFlags 命令行参数
type CommandFlags struct {
	Conf     string
	LogLevel string
}

func NewCommandFlags() *CommandFlags {
	return &CommandFlags{
		Conf:     "",
		LogLevel: "",
	}
}

func (f *CommandFlags) Init() {
	flag.StringVar(&f.Conf, "conf", "./configs/config.yaml", "config path, eg: -conf ./configs/config.yaml")
	flag.StringVar(&f.LogLevel, "logLevel", "debug", "log level, eg: -logLevel info")
}
