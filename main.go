package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sarthakshekhawat/Uptime-Monitoring-Service/controller"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func dsn() string {
	godotenv.Load(".env")
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	return fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)
}

func init() {
	//open a db connection
	db, err := gorm.Open(mysql.Open(dsn()), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&controller.DataBase{}) //Migrate the schema

	controller.AssignValue(db)
	db.Model(&controller.DataBase{}).Where("status = true").Update("status", false) //Making all the urls inactive while initilizing

	dbRepo := controller.DatabaseReceiver{}
	controller.AssignDbRepo(&dbRepo)

	reqRepo := controller.RequestReceiver{}
	controller.AssignRequestRepo(&reqRepo)
}

func main() {
	r := gin.Default()
	v1 := r.Group("/urls")
	{
		v1.POST("/", controller.StartMonitoring)
		v1.GET("/:id", controller.FetchMonitoringStatus)
		v1.PATCH("/:id", controller.UpdateMonitoring)
		v1.POST("/:id/activate", controller.ActivateMonitoring)
		v1.POST("/:id/deactivate", controller.DeactivateMonitoring)
		v1.DELETE("/:id", controller.DeleteMonitoring)
	}
	r.Run(":8080")
}
