package controllers

import (
	"School_api/models"
	"School_api/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetSchools(service services.SchoolService) gin.HandlerFunc {
	return func(c *gin.Context) {
		schools, err := service.GetSchools()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, schools)
	}
}

func GetSchoolByID(service services.SchoolService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		school, err := service.GetSchoolByID(uint(id))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "School not found"})
			return
		}
		c.JSON(http.StatusOK, school)
	}
}

func CreateSchool(service services.SchoolService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var school models.School
		if err := c.ShouldBindJSON(&school); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		createdSchool, err := service.CreateSchool(school)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, createdSchool)
	}
}

func UpdateSchool(service services.SchoolService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		var school models.School
		if err := c.ShouldBindJSON(&school); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		updatedSchool, err := service.UpdateSchool(uint(id), school)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, updatedSchool)
	}
}

func DeleteSchool(service services.SchoolService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		if err := service.DeleteSchool(uint(id)); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "School deleted"})
	}
}
