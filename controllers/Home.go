package controllers
import (
	"github.com/gin-gonic/gin"
	model "reportsicg/models"
	"fmt"
)

func Cajas( c *gin.Context) {
	var response struct {
		Cajas []model.Cajas `json:"cajas"`
	}
	response.Cajas = model.GetCajas()
	c.JSON(200, gin.H{
		"status": "success",
		"data": response,

	})
}

func Resume(c *gin.Context) {
	var date1 string = c.Query("fromDate")
	var date2 string = c.Query("toDate")
	var caja string = c.Query("caja")
	var response struct {
		Products []model.ProductsResume `json:"products"`
		OpenTables float64 `json:"openTables"`
		Guest float64 `json:"guest"`
		CloseTables float64 `json:"closeTables"`
		Cancellations float64 `json:"cancellations"`
		PaymentsMethods []model.PaymentsMethodsResume `json:"paymentsMethods"`
		GuestProducts []model.ProductsGuest `json:"guestProducts"`
	}
	fmt.Println(date1)
	productsData:= model.ResumeProducts(date1, date2, caja)
	response.Products = productsData
	response.OpenTables = model.OpenTables()
	response.Guest = model.Guest(date1, date2, caja)
	response.CloseTables = model.CloseTables(date1, date2, caja)
	response.Cancellations = model.Cancellations(date1, date2)
	response.PaymentsMethods = model.MethodPaymentsResume(date1, date2, caja)
	response.GuestProducts = model.GuestProducts(date1, date2, caja)
	c.JSON(200, gin.H{
		"status": "success",
		"data": response,

	})
}