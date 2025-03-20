package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"net/http"
	"reflect"
)

// 自定义结构体
type User1 struct {
	Name string `json:"name" binding:"required"`     // 标准required，不允许空
	Age  int    `json:"age" binding:"required"`      // 标准required，不允许0
	Age2 int    `json:"age2" binding:"existsOrZero"` // 自定义校验，允许0但必须存在
}

// 自定义校验函数
func existsOrZero(fl validator.FieldLevel) bool {
	// 获取字段值
	value := fl.Field()

	// 检查字段是否存在（非nil）且类型为int
	if value.IsValid() && value.Kind() == reflect.Int {
		return true // 允许任何整数值，包括0
	}
	return false // 字段不存在或类型错误则失败
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	// 获取Gin默认的validator引擎并注册自定义校验
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("existsOrZero", existsOrZero)
	}

	r.POST("/user", handleUser)
	return r
}

func handleUser(c *gin.Context) {
	var user User1
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"name": user.Name,
		"age":  user.Age,
		"age2": user.Age2,
	})
}

func main() {
	r := setupRouter()
	r.Run(":8090")
}
