package Validate

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"reflect"
)

func nameNotNullAndAdmin(fl validator.FieldLevel) bool {
	if value, ok := fl.Field().Interface().(string); ok {
		return value != "" && (value != "admin")
	}
	return false
}
func ValidateRegister() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("NameNotNullAndAdmin", nameNotNullAndAdmin)
	}
}

//调取错误信息
func ProcessErr(u interface{}, err error) string {
	if err == nil { //如果为nil 说明校验通过
		return ""
	}

	invalid, ok := err.(*validator.InvalidValidationError) //如果是输入参数无效，则直接返回输入参数错误
	if ok {
		return "输入参数错误：" + invalid.Error()
	}
	validationErrs := err.(validator.ValidationErrors) //断言是ValidationErrors
	for _, validationErr := range validationErrs {
		fieldName := validationErr.Field()                    //获取是哪个字段不符合格式
		field, ok := reflect.TypeOf(u).FieldByName(fieldName) //通过反射获取filed
		if ok {
			errorInfo := field.Tag.Get("reg_error_info") //获取field对应的reg_error_info tag值
			return fieldName + ":" + errorInfo           //返回错误
		} else {
			return "缺失reg_error_info"
		}
	}
	return ""
}
