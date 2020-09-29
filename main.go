package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sarthakshekhawat/Uptime-Monitoring-Service/controller"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	//open a db connection
	// dsn := "root:rootroot@tcp(host.docker.internal:3306)/UptimeMonitoringService?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "root:rootroot@tcp(docker.for.mac.localhost:3306)/UptimeMonitoringService?charset=utf8mb4&parseTime=True&loc=Local"
	// dsn := "root:rootroot@tcp(127.0.0.1:3306)/UptimeMonitoringService?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
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
