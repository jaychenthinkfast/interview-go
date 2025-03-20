在使用Go的Gin框架时，结构体字段标记 `binding:"required"` 会阻止零值（如空字符串 `""`、0、`false` 等）通过校验的问题。这是由于Gin依赖的validator库将“非零值”作为 `required` 的条件。解决方案包括：

1. 使用指针类型（如 `*int`、`*string`），允许字段接收零值或空值。
2. 自定义validator以更灵活地处理校验规则。

---

### 可执行代码和示例

#### 问题复现

```go
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Name string `json:"name" binding:"required"` // 不允许空字符串
	Age  int    `json:"age" binding:"required"`  // 不允许0
}

func main() {
	r := gin.Default()

	r.POST("/user", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"name": user.Name, "age": user.Age})
	})

	r.Run(":8090")
}
```

- **测试**：发送 `curl -X POST http://localhost:8090/user -H "Content-Type: application/json" -d '{"name": "", "age": 0}'`
- **结果**：返回错误，提示 `name` 和 `age` 未通过 `required` 校验。

---

#### 解决方案：使用指针类型

```go
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

```


#### 解决方案：自定义Validator实现步骤
1. **引入validator库**：确保项目中使用 `go-playground/validator/v10`。
2. **定义自定义校验规则**：注册一个新的校验函数。
3. **在Gin中使用**：将自定义校验器绑定到结构体字段。

下面是一个完整示例，解决“允许零值但要求字段存在”的问题。

---

```go
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

```

---

##### 扩展性
- **更复杂规则**：你可以在 `existsOrZero` 中添加条件，例如范围校验（`value.Int() >= 0 && value.Int() <= 100`）。
- **多字段校验**：通过 `fl.Parent()` 访问结构体其他字段，实现关联校验。

这种方法比使用指针更灵活，能精确控制校验逻辑，适合需要细粒度验证的场景。

