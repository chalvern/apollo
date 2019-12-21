package template

import (
	"fmt"
	"os"
	"path"
	"time"

	"github.com/chalvern/sugar"
)

func render(relativePath, hintName string, fileTemplate string) error {

	nowTimeString := time.Now().Format("200601021504")
	fileName := fmt.Sprintf("%s_%s.go", nowTimeString, hintName)
	longFileName := path.Join(relativePath, fileName)
	fd, err := os.OpenFile(longFileName, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		sugar.Fatalf("Open file %s failed: %v", fileName, err)
	}

	fileContent := fmt.Sprintf(fileTemplate, nowTimeString)
	_, err = fd.WriteString(fileContent)
	if err != nil {
		sugar.Fatalf("write file %s failed: %v", fileName, err)
	}
	return nil
}
