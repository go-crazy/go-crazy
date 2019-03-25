/**
 * 日志纪录配置
 * File: logger.go
 * Author: QylinFly (18612116114@163.com)
 * Created: 星期 3, 2017-12-20 11:58:42 am
 * -----
 * Modified By: QylinFly (18612116114@163.com>)
 * Modified: 星期 3, 2017-12-20 11:58:47 am
 * -----
 * Copyright 2017 - 2027 乐编程, 乐编程
 */

package Config

import (
	"fmt"
	"log"
	"os"

	"go-crazy/util/logger"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var _logger *zap.Logger

func InitLogger() {
	var err error
	cfg := zap.NewProductionConfig()

	// config
	cfg.OutputPaths = []string{
		"stdout",
		Path.Storage + "logs/go-crazy.log",
	}
	cfg.ErrorOutputPaths = []string{
		"stderr",
		Path.Storage + "logs/go-crazy-error.log",
	}
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	// level
	cfg.Level.SetLevel(zap.DebugLevel)

	// 建立
	_logger, err = cfg.Build()
	if err != nil {
		log.Println(fmt.Sprintf("\n Init logger error, and got err=%+v\n", err))
	}

	// set logfile Stdout
	var logFileName = Path.Storage + "logs/go-crazy-sys.log"
	logFile, logErr := os.OpenFile(logFileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if logErr != nil {
		fmt.Println("Fail to find", *logFile, "cServer start Failed")
		os.Exit(1)
	}

	if os.Getenv("GO_ENV") == "production" {
		log.SetOutput(logFile)
	}

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	// start
	_logger.Info("--------------------------------------------------")
	_logger.Info("-------------------App start----------------------")
	_logger.Info("--------------------------------------------------")

	logger.SetLogger(_logger)
}
