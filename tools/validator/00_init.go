package validator

import (
	"github.com/chalvern/sugar"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v9"
)

// InitValidatorEnhancement 增强 校验器
func InitValidatorEnhancement() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("lenlte", LenLte)
		v.RegisterValidation("lengte", LenGte)
	} else {
		sugar.Fatalf("InitValidatorEnhancement 运行失败")
	}
}
