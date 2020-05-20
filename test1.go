package main

import (
	"fmt"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)
var trans ut.Translator
// Booking contains binded and validated data.
type Booking struct {
	CheckIn  time.Time `form:"check_in" json:"check_in" binding:"required,bookabledate" time_format:"2006-01-02" label:"输入时间"`
	CheckOut time.Time `form:"check_out" json:"check_out" binding:"required,gtfield=CheckIn" time_format:"2006-01-02" label:"输出时间"`
}

var bookableDate validator.Func = func(fl validator.FieldLevel) bool {
	date, ok := fl.Field().Interface().(time.Time)
	if ok {
		today := time.Now()
		if today.After(date) {
			return false
		}
	}
	return true
}

func checkMobile(fl validator.FieldLevel) bool {
	mobile := strconv.Itoa(int(fl.Field().Uint()))
	re := `^1\d{10}$`
	r := regexp.MustCompile(re)
	return r.MatchString(mobile)
}
func main() {
	route := gin.Default()
	uni := ut.New(zh.New())
	trans, _ = uni.GetTranslator("zh")

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		//注册翻译器
		_= zh_translations.RegisterDefaultTranslations(v, trans)
		//注册自定义函数
		_=v.RegisterValidation("bookabledate", bookableDate)

		//注册一个函数，获取struct tag里自定义的label作为字段名
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name:=fld.Tag.Get("label")
			return name
		})
		//根据提供的标记注册翻译
		v.RegisterTranslation("bookabledate", trans, func(ut ut.Translator) error {
			return ut.Add("bookabledate", "{0}不能早于当前时间或{1}格式错误!", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("bookabledate", fe.Field(), fe.Field())
			return t
		})


		

	}
	route.GET("/bookable", getBookable)
	route.Run(":8085")
}

func getBookable(c *gin.Context) {
	var b Booking
	if err := c.ShouldBindWith(&b, binding.Query); err == nil {
		c.JSON(http.StatusOK, gin.H{"message": "Booking dates are valid!"})
	} else {
		errs := err.(validator.ValidationErrors)

		fmt.Println(errs.Translate(trans))
		//for _, e := range errs {
		//	// can translate each error one at a time.
		//	fmt.Println(e.Translate(trans))
		//}
		c.JSON(http.StatusBadRequest, gin.H{"error": errs.Translate(trans)})
	}
}
