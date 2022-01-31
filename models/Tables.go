package models
import (	
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"fmt"
	"strings"
)

func OpenTables() float64 {
	dsn := getStringConnection()
	var query = "SELECT SUM(precioiva) FROM dbo.minutaslin WHERE tipo = 'V'"
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	fmt.Println(err)
	var total float64
	db.Raw(query).Scan(&total)
	return total
}

func CloseTables(date1,date2, caja string) float64 {
	dsn := getStringConnection()
	var query = `
		SELECT SUM(totalneto)
		FROM dbo.tiquetscab 
		WHERE convert(varchar, fecha,23) BETWEEN @date1 AND @date2
		AND N = 'B'
		AND caja IN @caja
	`
	var total float64
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	fmt.Println(err)
	cajas := strings.Split(caja, ",")
	db.Raw(query, map[string]interface{}{"date1": date1, "date2": date2, "caja": cajas}).Scan(&total)
	return total
}


func Cancellations(date1,date2 string) float64 {

	dsn := getStringConnection()
	var query = `
		SELECT SUM(precioiva)
		FROM dbo.registroauditoria
		WHERE convert(varchar, fecha,23) BETWEEN @date1 AND @date2
		AND tipo = 0
	`
	var total float64
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	fmt.Println(err)
	db.Raw(query, map[string]interface{}{"date1": date1, "date2": date2}).Scan(&total)
	return total
}