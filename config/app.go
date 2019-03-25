/**
 * 基本的app配置 包含中间件 配置等
 * File: app.go
 * Author: QylinFly (18612116114@163.com)
 * Created: 星期 2, 2017-12-19 6:27:12 pm
 * -----
 * Modified By: QylinFly (18612116114@163.com>)
 * Modified: 星期 2, 2017-12-19 7:02:17 pm
 * -----
 * Copyright 2017 - 2027 乐编程, 乐编程
 */

package Config

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

var Path = struct {
	App      string `default:"./"`
	Base     string `default:"./"`
	Static   string `default:"./static/"`
	Resource string `default:"./static/"`
	Storage  string `default:"./storage/"`
}{}

func InitPath() {
	var base = getCurrentDirectory()

	Path.App = base
	Path.Base = base
	Path.Static = base + "/static/"
	Path.Resource = base + "/static/"
	Path.Storage = base + "/storage/"
}

func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}
