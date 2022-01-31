package models
import (	
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"fmt"
)
type Cajas struct {
	Name string `json:"name"`
}

func GetCajas() []Cajas {
	var cajas []Cajas
	dsn := getStringConnection()
	var query = `
		SELECT caja as name
		FROM dbo.tiquetscab
		GROUP BY caja
	`
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	fmt.Println(err)
	db.Raw(query).Scan(&cajas)
	return cajas
}
