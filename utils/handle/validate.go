package handle

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"strings"
)
var Trans ut.Translator

func InitValidate()  {
	uni := ut.New(zh.New())
	Trans, _ = uni.GetTranslator("zh")

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		//注册翻译器
		_= zh_translations.RegisterDefaultTranslations(v, Trans)
	}
}
func TransTagName( langs *map[string]string,err error) interface{} {
	for _, e := range err.(validator.ValidationErrors) {
		v:=e.Translate(Trans)
		for key,value:=range *langs  {
			v=strings.Replace(v,key,value,-1)
		}
		return v
	}
	return err
}





