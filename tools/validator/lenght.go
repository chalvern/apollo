package validator

import (
	"fmt"
	"reflect"
	"strconv"
	"unicode/utf8"

	v "gopkg.in/go-playground/validator.v9"
)

// LenLte 判定长度不大于等于某个值
// 比如限定字符串不长于 50 个字符
func LenLte(fl v.FieldLevel) bool {
	field := fl.Field()
	param := fl.Param()

	switch field.Kind() {

	case reflect.String:
		p := asInt(param)
		return int64(utf8.RuneCountInString(field.String())) <= p

	case reflect.Slice, reflect.Map, reflect.Array:
		p := asInt(param)
		return int64(field.Len()) <= p
	}

	panic(fmt.Sprintf("Bad field type %T", field.Interface()))
}

// LenGte 判定长度不小于等于某个值
// 比如限定字符串不小于 8 个字符
func LenGte(fl v.FieldLevel) bool {
	field := fl.Field()
	param := fl.Param()

	switch field.Kind() {

	case reflect.String:
		p := asInt(param)
		return int64(utf8.RuneCountInString(field.String())) >= p

	case reflect.Slice, reflect.Map, reflect.Array:
		p := asInt(param)
		return int64(field.Len()) >= p
	}

	panic(fmt.Sprintf("Bad field type %T", field.Interface()))
}

// asInt returns the parameter as a int64
// or panics if it can't convert
func asInt(param string) int64 {

	i, err := strconv.ParseInt(param, 0, 64)
	panicIf(err)

	return i
}

func panicIf(err error) {
	if err != nil {
		panic(err.Error())
	}
}
