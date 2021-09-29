package core

import (
	zh2 "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhtranslations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"regexp"
	"strings"
)

func mobileValidate(fl validator.FieldLevel) bool {
	mobile := fl.Field().String()
	re := `^1[3456789]\d{9}$`
	r := regexp.MustCompile(re)
	return r.MatchString(mobile)

}
func mobileRegister(ut ut.Translator) error {
	return ut.Add("mobile", "{0}填写不正确哦", true)
}
func mobileTranslation(ut ut.Translator, fe validator.FieldError) string {
	t, _ := ut.T("mobile", fe.Field())
	return t
}

var (
	validate *validator.Validate
	trans    ut.Translator
)

func InitValidate() {
	validate = validator.New()
	// 中文翻译
	zh := zh2.New()
	uni := ut.New(zh, zh)
	trans, _ = uni.GetTranslator("zh")
	e := zhtranslations.RegisterDefaultTranslations(validate, trans)
	if e != nil {
		panic(e)
	}
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		return fld.Tag.Get("comment")
	})

	return
}

func Translate(errs validator.ValidationErrors) string {
	var errList []string

	for _, e := range errs {
		// can translate each error one at a time.
		errList = append(errList, e.Translate(trans))
	}
	return strings.Join(errList, "|")
}
