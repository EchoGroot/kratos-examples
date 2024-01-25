package bootstrap

import (
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/pkg/errors"
)

// NewConfigProvider 读取配置
func NewConfigProvider(confPath string, bootstrap interface{}) error {
	c := config.New(
		config.WithSource(
			file.NewSource(confPath),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		return errors.Wrap(err, "load config file error")
	}

	if err := c.Scan(bootstrap); err != nil {
		return errors.Wrap(err, "read config file error")
	}

	if err := validate(bootstrap); err != nil {
		return errors.Wrap(err, "validate config file error")
	}

	return nil
}

type validator interface {
	Validate() error
}

func validate(i interface{}) error {
	if v, ok := i.(validator); ok {
		if err := v.Validate(); err != nil {
			return err
		}
	}
	return nil
}
