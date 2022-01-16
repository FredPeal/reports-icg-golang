package models
import (
	config "reportsicg/config"
)
 
func getStringConnection() string {
	var c = config.GetConf()
	return c.Connection
}