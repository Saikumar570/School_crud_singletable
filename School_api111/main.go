package main

import (
	"School_api/controllers"
	"School_api/services"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	schoolService := services.NewSchoolService()

	// School routes
	r.GET("/schools", controllers.GetSchools(schoolService))
	r.GET("/schools/:id", controllers.GetSchoolByID(schoolService))
	r.POST("/schools", controllers.CreateSchool(schoolService))
	r.PUT("/schools/:id", controllers.UpdateSchool(schoolService))
	r.DELETE("/schools/:id", controllers.DeleteSchool(schoolService))

	r.Run(":8080")
}
