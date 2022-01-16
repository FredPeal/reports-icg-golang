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

type MinutasHeader struct {
	Mesa string `json:"mesa"`
	Sala string `json:"sala"`
	Codcliente string `json:"codcliente"`
	Fecha string `json:"fecha"`
	Totalbruto float64 `json:"totalbruto"`
	Fechareal string `json:"fechareal"`
	Seccion string `json:"seccion"`
	Totalneto float64 `json:"totalneto"`
}
type Minutas struct {
	Mesa string `json:"mesa"`
	Sala string `json:"sala"`
	Codcliente string `json:"codcliente"`
	Fecha string `json:"fecha"`
	Totalbruto float64 `json:"totalbruto"`
	Fechareal string `json:"fechareal"`
	Seccion string `json:"seccion"`
	Totalneto float64 `json:"totalneto"`
	Products []ProductsDetail `json:"products"`
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


func OpenTransactions() []MinutasHeader {
	dsn := getStringConnection()
	var query = `
		SELECT DISTINCT minutascab.mesa, minutascab.sala, minutaslin.precioiva as Totalbruto, minutascab.codcliente, minutascab.numero, minutascab.serie,
		minutascab.fechaini as fecha, minutaslin.precio as Totalneto
		FROM dbo.minutascab
		JOIN dbo.minutaslin ON minutaslin.sala = minutascab.sala AND minutaslin.mesa = minutascab.mesa AND minutaslin.codcliente = minutascab.codcliente
	`
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	fmt.Println(err)
	var minutas []MinutasHeader
	db.Raw(query).Scan(&minutas)
	return minutas
}

func DetailOpenTransaction(mesa, sala, codcliente string) Minutas {
	dsn := getStringConnection()
	var query = `
		SELECT DISTINCT minutascab.mesa, minutascab.sala, minutaslin.precioiva as Totalbruto, minutascab.codcliente, minutascab.numero, minutascab.serie,
		minutascab.fechaini as fecha, minutaslin.precio as Totalneto
		FROM dbo.minutascab
		JOIN dbo.minutaslin ON minutaslin.sala = minutascab.sala AND minutaslin.mesa = minutascab.mesa AND minutaslin.codcliente = minutascab.codcliente
		WHERE minutascab.mesa = @mesa AND minutascab.sala = @sala AND minutascab.codcliente = @codcliente
	`
	
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	fmt.Println(err)
	var minutas Minutas
	db.Raw(query, map[string]interface{}{"mesa": mesa, "sala": sala, "codcliente": codcliente}).Scan(&minutas)
	query = `
		SELECT  codarticulo, descripcion, vendedores.nombrecorto as vendedor, hora, unidades as cantidad,
		precioiva as monto
		FROM dbo.minutaslin
		JOIN dbo.vendedores ON vendedores.codvendedor = minutaslin.codvendedor
		WHERE minutaslin.sala = @sala AND minutaslin.mesa = @mesa AND minutaslin.codcliente = @codcliente
	`
	db.Raw(query, map[string]interface{}{"mesa": mesa, "sala": sala, "codcliente": codcliente}).Scan(&minutas.Products)
	return minutas
}
