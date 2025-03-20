### GORM 零值问题背景

在 GORM 中，当使用struct更新时（如 `Update` 或 `Updates`）会忽略零值字段（例如 `0`、`""`、`false` 等），因为这些值被视为“未设置”。只有当字段值发生变化时，才会被包含在更新语句中。而 `Save` 方法会更新所有字段（包括零值），但有时开发者需要更精细的控制。

### 解决办法

文档中提到或隐含的解决 GORM 零值问题的方法包括以下几种：

1. **使用 `Save` 方法**

   - `Save` 会更新所有字段，即使字段是零值。
   - 示例：
     ```go
     user.Name = "jinzhu 2"
     user.Age = 0
     db.Save(&user) // 更新所有字段，包括 Age=0
     ```
   - 适用场景：需要保存所有字段时，不用担心零值被忽略。
2. **使用 `Updates` 并指定字段（Map 方式）**

   - 通过传递 `map[string]interface{}`，可以强制更新特定字段，即使是零值。
   - 示例：
     ```go
     db.Model(&user).Updates(map[string]interface{}{"name": "jinzhu", "age": 0})
     ```
   - 优点：明确指定字段，零值也会被更新。
3. **使用 `Update` 指定单个字段**

   - 使用 `Update` 方法更新单个字段时，零值会被正常处理。
   - 示例：
     ```go
     db.Model(&user).Update("age", 0) // 强制将 age 更新为 0
     ```
   - 适用场景：只更新特定字段时。
4. **使用指针或 `sql.NullXXX` 类型**

   - 文档未直接提到，但这是 GORM 常见的解决零值的方式。通过将字段定义为指针（如 `*int`）或 `sql.NullInt64` 等类型，可以区分“未设置”和“零值”。
   - 示例：
     ```go
     type User struct {
         Age *int
     }
     age := 0
     user := User{Age: &age}
     db.Save(&user) // Age 更新为 0
     ```
   - 优点：更灵活，能明确表示字段状态。
5. **选择性更新（`Select`）**

   - 使用 `Select` 指定需要更新的字段，结合 `Updates`，可以确保零值字段被更新。
   - 示例：
     ```go
     db.Model(&user).Select("age").Updates(User{Age: 0}) // 只更新 age，即使是 0
     ```
   - 适用场景：需要控制更新哪些字段时。

### 注意事项

- **默认行为**：`Updates` 只更新非零值的字段，除非使用上述方法强制更新。
- **性能考虑**：频繁使用 `Save` 更新所有字段可能影响性能，建议根据需求选择合适方法。

### 总结

GORM 零值问题的解决办法主要包括：

- 用 `Save` 更新所有字段（包括零值）。
- 用 `Updates` 配合 `map` 或 `Select` 强制更新零值字段。
- 用 `Update` 单独更新特定零值字段。
- 通过指针或 `sql.NullXXX` 类型区分零值和未设置状态。

### 参考
* https://gorm.io/zh_CN/docs/update.html