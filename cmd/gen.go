// cmd/gen.go
// 注意包名是main而不是cmd
package main

import (
	"fair-ticket-be-tutorial/db"

	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	cfg := gen.Config{
		// 设置输出路径
		OutPath: "model",
	}
	g := gen.NewGenerator(cfg)
	db := db.DB()
	g.UseDB(db)

	// 设置自定义数据类型转换
	g.WithDataTypeMap(dataMap())
	// 设置自定义json标签命名
	g.WithJSONTagNameStrategy(jsonTagNameStrategy)
	// 指定根据所有表生成model
	g.GenerateAllTable()
	// 指定根据某个表生成model
	// g.GenerateModel("xxx")

	// 执行
	g.Execute()
}

// dateTime 类型转换
func dataMap() map[string]func(gorm.ColumnType) (dataType string) {
	// 设置数据库字段与go结构体中字段数据类型的映射
	// DeletedAt设置是为了让gorm能够处理软删除的逻辑
	var dataMap = map[string]func(gorm.ColumnType) (dataType string){
		"datetime": func(columnType gorm.ColumnType) (dataType string) {
			if n, ok := columnType.Nullable(); ok && n {
				if columnType.Name() == "deleted_at" {
					return "gorm.DeletedAt"
				}
				return "*time.Time"
			}
			return "time.Time"
		},
	}
	return dataMap
}

// 将db里面的蛇形命名转换成驼峰命名
func jsonTagNameStrategy(columnName string) string {
	// 更优雅的下划线转驼峰（首字母小写）
	result := make([]rune, 0, len(columnName))
	upperNext := false
	for i, r := range columnName {
		if r == '_' {
			upperNext = true
			continue
		}
		if upperNext {
			if r >= 'a' && r <= 'z' {
				r -= 'a' - 'A'
			}
			result = append(result, r)
			upperNext = false
		} else {
			if i == 0 && r >= 'A' && r <= 'Z' {
				r += 'a' - 'A'
			}
			result = append(result, r)
		}
	}
	return string(result)
}
