package models
import (	
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"fmt"
)
 
type CancellationsProducts struct {
	Descripcion string `json:"descripcion"`
	Cantidad float64 `json:"cantidad"`
	Monto float64 `json:"monto"`
	Vendedor string `json:"vendedor"`
	Fecha string `json:"fecha"`
	Hora string `json:"hora"`
	Sala string `json:"sala"`
	Mesa string `json:"mesa"`
}

func TransactionsCancelled(date1 string, date2 string) []CancellationsProducts {
	  
	var res []CancellationsProducts
	dsn := getStringConnection()
	query := `
		SELECT registroauditoria.descripcion, registroauditoria.uds as cantidad, 
		registroauditoria.precioiva as monto, vendedores.nombrecorto as vendedor,
		registroauditoria.fecha, registroauditoria.hora, salas.nombre as sala, mesa
		FROM dbo.registroauditoria
		JOIN dbo.vendedores as vendedores ON vendedores.codvendedor = registroauditoria.codempleado
		JOIN dbo.salas as salas ON salas.sala = registroauditoria.sala
		WHERE convert(varchar, registroauditoria.fecha,23) BETWEEN @date1 AND @date2
		AND tipo = 0
	`
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	db.Raw(query, map[string]interface{}{"date1": date1, "date2": date2}).Scan(&res)
	fmt.Println("Error getTransaction: ", err)
	return res
}

