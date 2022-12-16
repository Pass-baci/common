package common

import (
	"fmt"
	"github.com/asim/go-micro/v3/config"
	"github.com/go-micro/plugins/v3/config/source/consul"
)

func GetConsulConfig(host string, port int64, prefix string) (config.Config, error) {
	consulSource := consul.NewSource(
		// 设置配置中心的地址
		consul.WithAddress(fmt.Sprintf("%s:%d", host, port)),
		// 设置前缀，不设置 /micro/config
		consul.WithPrefix(prefix),
		consul.StripPrefix(true),
	)
	conf, err := config.NewConfig()
	if err != nil {
		return nil, err
	}
	if err = conf.Load(consulSource); err != nil {
		return nil, err
	}
	return conf, nil
}
