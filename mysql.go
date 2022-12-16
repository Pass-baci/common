package common

import "github.com/asim/go-micro/v3/config"

// MysqlConfig Mysql配置结构体
type MysqlConfig struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Pwd      string `json:"pwd"`
	Database string `json:"database"`
	Port     int64  `json:"port"`
}

func GetMysqlFromConsul(config config.Config, path ...string) (*MysqlConfig, error) {
	mysqlConfig := &MysqlConfig{}
	if err := config.Get(path...).Scan(mysqlConfig); err != nil {
		return nil, err
	}
	return mysqlConfig, nil
}
