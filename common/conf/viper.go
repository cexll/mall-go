package conf

import (
	"github.com/spf13/viper"
)

type (
	// Option defines the method to customize the config options.
	Option func(opt *options)

	options struct {
		env bool
	}
)

func MustLoad(path string, v interface{}, opts ...Option) error {
	vi := viper.New()

	viper.SetConfigType("yaml")

	vi.SetConfigFile(path)

	if err := vi.ReadInConfig(); err != nil {
		return err
	}
	if err := vi.Unmarshal(&v); err != nil {
		return err
	}

	return nil
}
