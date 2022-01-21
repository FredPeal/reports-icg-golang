package models

import (	
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"fmt"
)
type Transaction struct {
	Serie string `json:"serie"`
	Numero string `json:"numero"`
	Mesa string `json:"mesa"`
	Fecha string `json:"fecha"`
	Vendedor string `json:"vendedor"`
	Totalneto float64 `json:"totalneto"`
	Totalbruto float64 `json:"totalbruto"`
	TotalIva float64 `json:"totaliva"`
}

type DetailTransaction struct {
	Serie string `json:"serie"`
	Numero string `json:"numero"`
	Mesa string `json:"mesa"`
	Fecha string `json:"fecha"`
	Vendedor string `json:"vendedor"`
	Totalneto float64 `json:"totalneto"`
	Totalbruto float64 `json:"totalbruto"`
	TotalIva float64 `json:"totaliva"`
	Seriefiscal string `json:"seriefiscal"`
	Numerofiscal string `json:"numerofiscal"`
	Seriefiscal2 string `json:"seriefiscal2"`
	Fechareal string `json:"fechareal"`
	Sala string `json:"sala"`
	Ncf string `json:"ncf"`
	Products []ProductsDetail `json:"products"`
	TicketsPaid []TicketsPaid `json:"ticketspaid"`
}

type ProductsDetail struct {
	CODARTICULO string `json:"codarticulo"`
	Descripcion string `json:"descripcion"`
	Vendedor string `json:"vendedor"`
	Hora string `json:"hora"`
	Cantidad float64 `json:"cantidad"`
	Monto float64 `json:"monto"`
}

type TicketsPaid struct {
	Formapago string `json:"formapago"`
	Importe float64 `json:"importe"`
}

func GetTransactions(date1 string,date2 string) []Transaction {
	dsn := getStringConnection()
	var query string = `
			SELECT tiquetscab.serie, tiquetscab.numero, tiquetscab.mesa,
			convert(varchar, tiquetscab.fecha,23) as fecha, vendedores.nombrecorto as vendedor, tiquetscab.totalneto as totalneto, 
			tiquetscab.totalbruto as totalbruto, tiquetscab.totalcostEiva as totaliva
			FROM dbo.tiquetscab 
			JOIN dbo.vendedores ON dbo.tiquetscab.codvendedor = dbo.vendedores.codvendedor
			WHERE convert(varchar, tiquetscab.fecha,23) BETWEEN @date1 AND @date2
			AND tiquetscab.N = 'B'
		`
		var res []Transaction
		db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
		db.Raw(query, map[string]interface{}{"date1": date1, "date2": date2}).Scan(&res)
		fmt.Println("Error getTransaction: ", err)
		return res
}

func GetTransaction(numero string, serie string) DetailTransaction {

	type Header struct {
		Mesa string `json:"mesa"`
		Totalbruto float64 `json:"totalbruto"`
		Totalneto float64 `json:"totalneto"`
		Fecha string `json:"fecha"`
		Seriefiscal string `json:"seriefiscal"`
		Numerofiscal string `json:"numerofiscal"`
		Seriefiscal2 string `json:"seriefiscal2"`
		Numero string `json:"numero"`
		Nombrecorto string `json:"nombrecorto"`
		Fechareal string `json:"fechareal"`
		Totaliva float64 `json:"totaliva"`
		Sala string `json:"sala"`
		Ncf string `json:"ncf"`
		Seccion string `json:"seccion"`
	}

	// Header Query
	var queryHeader string = `
		SELECT mesa, totalbruto, totalneto, fecha, seriefiscal, numerofiscal, seriefiscal2, numero, vendedores.nombrecorto,
		fechareal, totalcosteiva as totaliva, salas.nombre as sala
		FROM dbo.tiquetscab
		JOIN dbo.vendedores ON tiquetscab.codvendedor = vendedores.codvendedor
		LEFT JOIN dbo.salas on salas.sala = tiquetscab.sala
		WHERE tiquetscab.numero = @numero AND tiquetscab.serie = @serie
	`
	dsn := getStringConnection()
	var resHeader Header
	fmt.Println("numero: ", numero)
	fmt.Println("serie: ", serie)
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	db.Raw(queryHeader, map[string]interface{}{"numero": numero, "serie": serie}).Scan(&resHeader)
	fmt.Println("Error Header: ", err)

	// Detail Query
	var Details []ProductsDetail
	var queryDetail string = `
		SELECT codarticulo, descripcion, vendedores.nombrecorto as vendedor, hora, unidades as cantidad, precioiva as monto
		FROM dbo.tiquetslin
		JOIN dbo.vendedores ON tiquetslin.codvendedor = vendedores.codvendedor
		WHERE tiquetslin.numero = @numero AND tiquetslin.serie = @serie
	`

	db.Raw(queryDetail, map[string]interface{}{"numero": numero, "serie": serie}).Scan(&Details)

	// Tickets Paid Query

	var queryPaid string = `
		SELECT formaspago.descripcion as Formapago, tiquetspag.importe
		FROM dbo.tiquetspag
		JOIN dbo.formaspago ON tiquetspag.codformapago = formaspago.codformapago
		WHERE tiquetspag.numero = @numero AND tiquetspag.serie =  @serie`


	var Paid []TicketsPaid
	db, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	db.Raw(queryPaid, map[string]interface{}{"numero": numero, "serie": serie}).Scan(&Paid)
	fmt.Println("Error Header: ", err)

	var response DetailTransaction;
	response.Serie = serie
	response.Numero = numero
	response.Mesa = resHeader.Mesa
	response.Fecha = resHeader.Fecha
	response.Vendedor = resHeader.Nombrecorto
	response.Totalneto = resHeader.Totalneto
	response.Totalbruto = resHeader.Totalbruto
	response.TotalIva = resHeader.Totaliva
	response.Seriefiscal = resHeader.Seriefiscal
	response.Numerofiscal = resHeader.Numerofiscal
	response.Seriefiscal2 = resHeader.Seriefiscal2
	response.Fechareal = resHeader.Fechareal
	response.Sala = resHeader.Sala
	response.Ncf = resHeader.Seriefiscal + resHeader.Numerofiscal
	response.Products = Details
	response.TicketsPaid = Paid
	return response
}