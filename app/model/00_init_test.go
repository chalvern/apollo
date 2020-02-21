package model

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/chalvern/apollo/configs/initializer"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	fmt.Println("begin test of model")

	// 加载测试环境的 yaml
	initializer.InitViperWithFile("../../configs/config_test.yml")
	initializer.InitMysql(context.Background())
	Init()

	resultCode := m.Run()
	fmt.Println("end test of model")
	os.Exit(resultCode)
}

func TestDbArgs(t *testing.T) {
	db := dbArgs(mydb, "id=?", "1")
	assert.NotNil(t, db)

	db2 := dbArgs(mydb, "id=1")
	assert.NotNil(t, db2)
}
