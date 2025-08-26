package main

import (
	"fair-ticket-be-tutorial/db"

	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "model",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})
	db := db.DB()
	g.UseDB(db)

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
	g.WithDataTypeMap(dataMap)
	// generate all tables
	g.GenerateAllTable()
	// Generate the code
	g.Execute()
}
