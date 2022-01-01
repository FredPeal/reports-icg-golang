package models

import (	
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"fmt"
)
type ProductsResume struct {
	Cantidad int
	Descripcion string
	Categoria string
	Precio float64
}

func ResumeProducts(date1, date2 string) []ProductsResume {
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
			WHERE tiqueslin.tipo = 'V' AND convert(varchar, fecha,23) BETWEEN @date1 AND @date2
			GROUP BY tiqueslin.descripcion, secciones.descripcion
			SORT BY cantidad DESC
	`
	var products []ProductsResume
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	db.Raw(query, map[string]interface{}{"date1": date1, "date2": date2}).Scan(&products)
	fmt.Println(err)
	return products
}