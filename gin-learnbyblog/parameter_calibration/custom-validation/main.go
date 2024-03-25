package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// import (
// 	"net/http"
// 	"reflect"

// 	"github.com/gin-gonic/gin"
// 	"github.com/gin-gonic/gin/binding"
// 	"github.com/go-playground/validator"
// )

/*
对绑定解析到结构体上的参数，自定义验证功能
比如我们要对name字段做校验，要不能为空并且不等于admin，类似这种需求无法binding现成的方法
需要我们自己验证方法才能实现 官网示例（https://godoc.org/gopkg.in/go-playground/validator.v8#hdr-Custom_Functions）
这里需要下载引入下 gopkg.in/go-playground/validator.v8
*/

/*
type Person struct {
	Age int `form:"age" binding:"required,gt=10"`
	// 2.在参数binding上使用自定义的校验方法函数注册时候的名称
	Name    string `form:"name" binding:"NotNullAndAdmin"`
	Address string `form:"address" binding:"required"`
}

// 1.自定义的校验方法(!!!应换成最新版,https://gin-gonic.com/zh-cn/docs/examples/custom-validators/)
func nameNotNullAndAdmin(
	v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value, field reflect.Value,
	fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	if value, ok := field.Interface().(string); ok {
		// 字段不能为空，且不等于admin
		return value != "" && !("51mh" == value)
	}

	return true
}

func main() {
	r := gin.Default()

	// 3.将我们自定义的校验方法注册到validation中
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 这里的key和fn可以不一样，最终在struct使用的是key
		v.RegisterValidation("NotNullAndAdmin", nameNotNullAndAdmin)
	}


	// curl -X GET "http://127.0.0.1:8080/testing?name=&age=12&address=beijing"
	// curl -X GET "http://127.0.0.1:8080/testing?name=lmh&age=12&address=beijing"
	// curl -X GET "http://127.0.0.1:8080/testing?name=adz&age=12&address=beijing"

	r.GET("/51mh", func(c *gin.Context) {
		var person Person
		if e := c.ShouldBind(&person); e == nil {
			c.String(http.StatusOK, "%v", person)
		} else {
			c.String(http.StatusOK, "person bind err:%v", e.Error())
		}
	})

	r.Run()
}

*/

type Booking struct {
	// 定义一个预约的事件大于今天的时间
	CheckIn time.Time `form:"check_in" binding:"required,bookabledate" time_format:"2006-01-02"`
	// gtfield=CheckIn退出的时间大于预约的时间
	CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn,bookabledate" time_format:"2006-01-02"`
}

// func bookableDate(
// 	v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
// 	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string,
// ) bool {
// 	// field.Interface().(time.Time)获取参数值并且转换为时间格式
// 	if date, ok := field.Interface().(time.Time); ok {
// 		today := time.Now()
// 		if today.Unix() > date.Unix() {
// 			return false
// 		}
// 	}
// 	return true
// }

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

func main() {
	route := gin.Default()
	// 注册验证
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 绑定第一个参数是验证的函数,第二个参数是自定义的验证函数
		v.RegisterValidation("bookabledate", bookableDate)
	}

	route.GET("/bookable", getBookable)
	route.Run()
}

func getBookable(c *gin.Context) {
	var b Booking
	if err := c.ShouldBindWith(&b, binding.Query); err == nil {
		c.JSON(http.StatusOK, gin.H{"message": "Booking dates are valid!"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
