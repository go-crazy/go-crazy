package test

import (
	"fmt"
	. "go-crazy/Config"
	"go-crazy/util/logger"

	"go.uber.org/zap"
)

func testLogger() {

	InitLogger()

	logger.Instance().Info("你好", zap.String("www", "shgadadagds"), zap.Int("sss", 1212))
}
