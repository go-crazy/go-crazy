// Application service app 实例
package App

import (
	_config "go-crazy/Config"

	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
)

var (
	App *Application
)

type Application struct {
	IrisApp *iris.Application
	DB      *gorm.DB
}

// 自动初始化
func init() {
	App = new(Application)
}

// 初始化
func InitIrisApp(app *iris.Application) {
	App.IrisApp = app
}

func InitDB() *gorm.DB {
	App.DB = _config.InitDB()
	return App.DB
}
func InitLogger() {
	_config.InitLogger()
}
func InitPath() {
	_config.InitPath()
}

// 获取
func Config() *_config.CONFIG {
	return &_config.Config
}
func DB() *gorm.DB {
	return App.DB
}

func ReleaseResources() {
	_config.CloseDB()
}

func CloseDB() {
	_config.CloseDB()
}
