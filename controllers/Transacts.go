package controllers


import (
	"github.com/gin-gonic/gin"
	model "reportsicg/models"
)

func GetAllTransacts(c *gin.Context) {
	var date1 string = c.Query("fromDate")
	var date2 string = c.Query("toDate")
	var caja string = c.Query("caja")

	var response = model.GetTransactions(date1, date2, caja)
	c.JSON(200, gin.H{
		"status": "success",
		"data": response,
	})
}

func GetTransaction(c *gin.Context) {
	var numero string = c.Param("numero")
	var serie string = c.Param("serie")
	var response = model.GetTransaction(numero, serie)
	c.JSON(200, gin.H{
		"status": "success",
		"data": response,
	})
}

func GetTransactionsCancelled(c *gin.Context) {
	var date1 string = c.Query("fromDate")
	var date2 string = c.Query("toDate")
	var response = model.TransactionsCancelled(date1, date2)
	c.JSON(200, gin.H{
		"status": "success",
		"data": response,
	})
}

func GetOpenTransactions(c *gin.Context) {
	var response = model.OpenTransactions()
	c.JSON(200, gin.H{
		"status": "success",
		"data": response,
	})
}

func GetDetailsOpenTransactions(c *gin.Context) {
	var mesa string = c.Query("mesa")
	var sala string = c.Query("sala")
	var cliente string = c.Query("cliente")
	var numero string = c.Query("numero")
	var response = model.DetailOpenTransaction(mesa, sala, cliente, numero)
	c.JSON(200, gin.H{
		"status": "success",
		"data": response,
	})
}