package models
import (	
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"fmt"
)
type MinutasLin struct {
	Sala int
	Mesa int
	Codcliente int
	Numlinea int
	Codarticulo int
	Descripcion string
	Unidades float64
	Precioiva float64
	Preciodefecto float64
	Preciocoste float64
	Codtasa1 int
	Codtasa2 int
	Orden int
	Codformato int
	Codmacro int
	Tipo string
	Impreso bool
	Esmenu bool
	Porcargo string
	Hora string
	Referencia string
	Factporhora string
	Cerrado bool
	Consumadic int
	Codfamilia int
	Dividida int
	Servido int
	Sordertiquet string
	Nordertiquet int
	Numreserva string
	Idtalonario int
	Codtarifa int
	Idlineaminuta int
	Familiaaena int
	Idmotivodto int
	Identrada string
	Tipotarjentrada int
	Horamarchado string
	Udsabonadas float64
	Tasaespecial int
	Cambiomesa int
	Horacocina string	
	Guidlinea string
	Estipoabono int
	Entradagenerada int
	Factporfranja int
	Guidlineaabono string
	Posicionimpresion int
	Esregalomixmatch int
	Tipodelivery int
	Idmotivoabono int
	Numconsumiciones int
	Isprecio2 string
}


func Guest(date1, date2 string) float64 {
	dsn := getStringConnection()
	var query = `
	SELECT SUM(preciosventa.valor) as total
	FROM dbo.TIQUETSLIN as tiqueslin
	JOIN dbo.tiquetscab as tiquetscab ON tiquetscab.serie = tiqueslin.serie AND tiquetscab.numero = tiqueslin.numero
	JOIN dbo.preciosventa as preciosventa ON preciosventa.codarticulo = tiqueslin.codarticulo AND preciosventa.idtarifav = tiqueslin.codtarifa
	WHERE tiqueslin.tipo = 'I' AND convert(varchar, fecha,23) BETWEEN @date1 AND @date2
	
	UNION 
	
	SELECT SUM(preciosventa.valor) as total
	FROM dbo.minutaslin as minutas
	JOIN dbo.preciosventa as preciosventa ON preciosventa.codarticulo = minutas.codarticulo AND preciosventa.idtarifav = minutas.codtarifa
	WHERE minutas.tipo = 'I'`
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	fmt.Println(err)
	var total float64
	db.Raw(query, map[string]interface{}{"date1": date1, "date2": date2}).Scan(&total)
	return total
}

