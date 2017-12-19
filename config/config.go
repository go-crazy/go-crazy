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

 
var Config = struct {
	APPName string `default:"app name"`
	Port string `default:"80"`

	DB struct {
		Name     string
		User     string `default:"root"`
		Password string `required:"true" env:"DBPassword"`
		Port     uint   `default:"3306"`
	}

	Contacts []struct {
		Name  string
		Email string `required:"true"`
	}
}{}
