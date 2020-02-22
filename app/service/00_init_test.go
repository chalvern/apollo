package service

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/chalvern/apollo/app/model"
	"github.com/chalvern/apollo/configs/initializer"
)

func TestMain(m *testing.M) {
	fmt.Println("begin test of service")

	// 加载测试环境的 yaml
	initializer.InitViperWithFile("../../configs/config_test.yml")
	initializer.InitMysql(context.Background())
	model.Init()

	resultCode := m.Run()
	fmt.Println("end test of service")
	os.Exit(resultCode)
}
