package models
import (	
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"fmt"
	"strings"
)
type PaymentsMethodsResume struct {
	Name string
	Amount float64
}
func MethodPaymentsResume(date1 string, date2, caja string) []PaymentsMethodsResume {

	dsn:=getStringConnection()
	var query string = `
			SELECT formaspago.DESCRIPCION as name , SUM(tiquetspag.importe) as amount
			FROM dbo.tiquetspag 
			JOIN dbo.formaspago ON dbo.tiquetspag.CODFORMAPAGO = dbo.formaspago.CODFORMAPAGO
			JOIN dbo.tiquetscab ON dbo.tiquetspag.serie = dbo.tiquetscab.serie AND dbo.tiquetspag.numero = dbo.tiquetscab.numero
			WHERE convert(varchar, fecha,23) BETWEEN @date1 AND @date2
			AND
			dbo.tiquetscab.caja IN @caja
			GROUP BY formaspago.DESCRIPCION
		`
		var res []PaymentsMethodsResume
		db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
		cajas := strings.Split(caja, ",")

		db.Raw(query, map[string]interface{}{"date1": date1, "date2": date2, "caja": cajas}).Scan(&res)
		fmt.Println("Error Method Payments Resume: ", err)
		return res
}