package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 定义结构体，默认使用 required
type User struct {
	Name string `json:"name" binding:"required"`  // 不允许空字符串
	Age  int    `json:"age" binding:"required"`   // 不允许0
	Age2 *int   `json:"age2"  binding:"required"` // 使用指针，允许零值
}

func main() {
	r := gin.Default()

	r.POST("/user", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"name": user.Name,
			"age":  user.Age,
			"age2": user.Age2,
		})
	})

	r.Run(":8090")
}
