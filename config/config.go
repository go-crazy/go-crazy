/**
 * 配置
 * File: config.go
 * Author: QylinFly (18612116114@163.com)
 * Created: 星期 2, 2017-12-19 8:39:49 pm
 * -----
 * Modified By: QylinFly (18612116114@163.com>)
 * Modified: 星期 2, 2017-12-19 8:40:14 pm
 * -----
 * Copyright 2017 - 2027 乐编程, 乐编程
 */

package Config

type CONFIG struct {
	Env     string
	APPName string
	Port    string
	Debug   string

	DB_CONNECTION string
	DB            struct {
		Host     string
		Name     string
		User     string `default:"root"`
		Password string `required:"true" env:"DBPassword"`
		Port     string `default:"3306"`
	}
}

var Config CONFIG

func init() {
	config := CONFIG{}

	// 默认值
	config.Env = "production"
	config.Debug = "false"

	config.APPName = "app name"
	config.Port = "80"
	config.DB_CONNECTION = "mysql"

	config.DB.Host = "127.0.0.1"
	config.DB.Port = "3306"
	config.DB.User = "root"

	// 复制
	Config = config
}
