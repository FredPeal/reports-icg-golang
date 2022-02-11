package models

import (
	"fmt"
	"strings"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type ProductsResume struct {
	Cantidad    float64
	Descripcion string
	Categoria   string
	Precio      float64
}

type ProductsGuest struct {
	Cantidad    float64
	Descripcion string
	Precioiva      float64
	Precio      float64
	Categoria   string
	Tipo        string
	Valor       float64
	Vendedor    string
	Hora        string

}

func ResumeProducts(date1, date2, caja string) []ProductsResume {
	dsn := getStringConnection()
	var query string = `
			SELECT SUM(minutas.unidades) as cantidad, minutas.descripcion, secciones.descripcion as categoria,  SUM(minutas.precioiva) as precio
			FROM dbo.articulos as products
			INNER JOIN dbo.minutaslin AS minutas ON minutas.codarticulo = products.codarticulo
			INNER JOIN dbo.secciones as secciones ON products.seccion = secciones.seccion
			WHERE minutas.tipo = 'V'
			GROUP BY minutas.descripcion, secciones.descripcion

			UNION 

			SELECT SUM(tiqueslin.unidades) as cantidad, tiqueslin.descripcion, secciones.descripcion as categoria,  SUM(tiqueslin.precioiva) as precio
			FROM dbo.articulos as products
			JOIN dbo.TIQUETSLIN as tiqueslin ON tiqueslin.codarticulo = products.codarticulo
			JOIN dbo.secciones as secciones ON products.seccion = secciones.seccion
			JOIN dbo.tiquetscab as tiquetscab ON tiquetscab.serie = tiqueslin.serie AND tiquetscab.numero = tiqueslin.numero
			WHERE tiqueslin.tipo = 'V' AND convert(varchar, fecha,23) BETWEEN @date1 AND @date2 AND tiquetscab.caja IN @caja
			GROUP BY tiqueslin.descripcion, secciones.descripcion
	`
	var products []ProductsResume
	cajas := strings.Split(caja, ",")
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	db.Raw(query, map[string]interface{}{"date1": date1, "date2": date2, "caja": cajas}).Scan(&products)
	fmt.Println(err)
	return products
}

func GuestProducts(date1, date2, caja string) []ProductsGuest {
	dsn := getStringConnection()
	var query string = `
		SELECT SUM(minutas.unidades) as cantidad, minutas.descripcion as descripcion, SUM(minutas.precioiva) as precioiva, SUM(minutas.precio) as precio,
		secciones.descripcion as categoria, minutas.tipo as tipo, precios.valor as valor, vendedores.nombrecorto as vendedor, minutas.hora
		FROM dbo.articulos as products
		INNER JOIN dbo.minutaslin AS minutas ON minutas.codarticulo = products.codarticulo
		LEFT JOIN dbo.secciones AS secciones ON secciones.seccion = products.seccion
		JOIN  dbo.vendedores AS vendedores ON vendedores.codvendedor = minutas.codvendedor
		JOIN dbo.preciosventa AS precios ON precios.codarticulo = minutas.codarticulo AND precios.idtarifav = minutas.codtarifa
		WHERE minutas.tipo = 'I'
		GROUP BY minutas.DESCRIPCION, secciones.DESCRIPCION, minutas.TIPO,precios.valor, vendedores.NOMBRECORTO, minutas.hora
	
		UNION
	
		SELECT 
		SUM(tiqueslin.unidades) as cantidad, tiqueslin.descripcion as descripcion, SUM(tiqueslin.precioiva) as precioiva, SUM(tiqueslin.precio) as precio,
		secciones.descripcion as categoria, tiqueslin.tipo as tipo, precios.valor as valor, vendedores.nombrecorto as vendedor, tiqueslin.hora
		FROM dbo.articulos as products
		INNER JOIN dbo.TIQUETSLIN AS tiqueslin ON tiqueslin.codarticulo = products.codarticulo
		INNER JOIN dbo.vendedores as vendedores ON vendedores.codvendedor = tiqueslin.codvendedor
		INNER JOIN dbo.tiquetscab as tiquetscab ON tiquetscab.serie = tiqueslin.serie AND tiquetscab.numero = tiqueslin.numero
		INNER JOIN dbo.preciosventa AS precios ON precios.codarticulo = tiqueslin.codarticulo AND precios.idtarifav = tiqueslin.codtarifa
		INNER JOIN dbo.secciones as secciones ON secciones.seccion = products.seccion
		WHERE tiqueslin.tipo = 'I' AND convert(varchar, fecha,23) BETWEEN @date1 AND @date2  AND tiquetscab.caja IN @caja
		GROUP BY tiqueslin.DESCRIPCION, secciones.DESCRIPCION, tiqueslin.TIPO, precios.VALOR, vendedores.NOMBRECORTO, tiqueslin.HORA
	`
	var products []ProductsGuest
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	cajas := strings.Split(caja, ",")
	db.Raw(query, map[string]interface{}{"date1": date1, "date2": date2, "caja": cajas}).Scan(&products)
	fmt.Println(err)
	return products
}
