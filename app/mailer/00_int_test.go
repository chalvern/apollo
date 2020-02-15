package mailer

import (
	"fmt"
	"os"
	"testing"

	"github.com/chalvern/apollo/configs/initializer"
)

func TestMain(m *testing.M) {
	fmt.Println("begin test of mailer")

	initializer.InitViperWithFile("../../configs/config.yml")
	Init()

	resultCode := m.Run()
	fmt.Println("end test of mailer")
	os.Exit(resultCode)
}
