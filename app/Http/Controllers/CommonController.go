/**
 * File: CommonController.go 通用控制方法
 * Author: QylinFly (18612116114@163.com)
 * Created: 星期 2, 2017-12-19 6:18:50 pm
 * -----
 * Modified By: QylinFly (18612116114@163.com>)
 * Modified: 星期 2, 2017-12-19 6:19:24 pm
 * -----
 * Copyright 2017 - 2027 乐编程, 乐编程
 */

package Controller

import (
	// "database/sql"
	// "fmt"
	"go-crazy/app"

	// "go-crazy/util/logger"

	// "go-crazy/app/Models"
	// "encoding/json"
	"go-crazy/app/Services/User"

	"go-crazy/app/Repositories"
	"go-crazy/util"

	"github.com/kataras/iris"
)

func Ping(ctx iris.Context) {
	// Ping test
	util.Api_response(ctx, "pong")
}

func GetConfig(ctx iris.Context) {
	// Ping test
	util.Api_response(ctx, App.Config())
}

func GetTokeString(ctx iris.Context) {
	// Ping test
	// token := Auth.GetTokenByUser()
	util.Api_response(ctx, UserService.GetUserByName("xoxo2"))
}

func Dbtest(ctx iris.Context) {
	// Ping test

	// var user Model.User

	// App.DB().First(&user)

	// Scan
	type Result struct {
		Name string
		Id   uint
	}

	// data := make(map[string]interface{})

	// var result []map[string]interface{}
	// var result []Result
	// App.DB().Raw("SELECT name, id FROM users WHERE id > ?", 0).Scan(&result)

	rows, _ := App.DB().Raw("SELECT id,name,email,created_at,updated_at FROM users WHERE id > ?", 0).Rows()

	// App.DB().ScanRows(rows, &result)

	result := Repository.SuperQueryRowsToArrMap(rows) //RowsToArrMap(rows)

	// c, _ := rows.Columns()

	util.Api_response(ctx, result)
}

// type ExampleStringSlice []string

// func (l *ExampleStringSlice) Scan(input interface{}) error {
// 	return json.Unmarshal(input.([]byte), l)
// }
