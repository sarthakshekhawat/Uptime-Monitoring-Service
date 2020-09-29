package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func response(c *gin.Context, data DataBase) {
	status := "inactive"
	if data.Status == true {
		status = "active"
	}
	c.JSON(http.StatusOK, gin.H{
		"id":                data.ID,
		"url":               data.URL,
		"crawl_timeout":     data.CrawlTimeout,
		"frequency":         data.Frequency,
		"failure_threshold": data.FailureThreshold,
		"status":            status,
		"failure_count":     data.FailureCount,
	})
}

// StartMonitoring starts making the requests to the url provided in the request body
func StartMonitoring(c *gin.Context) {
	var data requestData
	c.BindJSON(&data)

	dbData := addingDataInDatabase(data)
	go startRequest(dbData)
	response(c, dbData)
}

// FetchMonitoringStatus fetches the data from the database using the id
func FetchMonitoringStatus(c *gin.Context) {
	fetchData, err := checkUUID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": err.Error()})
		return
	}
	response(c, fetchData)
}

// UpdateMonitoring updates the value in the database using the values provided in the request body
func UpdateMonitoring(c *gin.Context) {
	fetchData, err := checkUUID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": err.Error()})
		return
	}
	var data requestData
	c.BindJSON(&data)

	fetchData = updatingDataInDatabase(data, fetchData)
	response(c, fetchData)
}

// ActivateMonitoring starts crawling the deactivated url
func ActivateMonitoring(c *gin.Context) {
	fetchData, err := checkUUID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": err.Error()})
		return
	}
	if fetchData.Status == true {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Already Active!",
		})
		return
	}
	fetchData = activateInDatabase(fetchData)
	go startRequest(fetchData)
	response(c, fetchData)
}

// DeactivateMonitoring stops crawling the activated url
func DeactivateMonitoring(c *gin.Context) {
	fetchData, err := checkUUID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": err.Error()})
		return
	}
	if fetchData.Status == false {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Already Inactive!",
		})
		return
	}
	fetchData = deactivateInDatabase(fetchData)
	response(c, fetchData)
}

// DeleteMonitoring deletes the data from the database using id
func DeleteMonitoring(c *gin.Context) {
	fetchData, err := checkUUID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": err.Error()})
		return
	}
	dbRepo.deleteData(fetchData)
	c.JSON(http.StatusNoContent, gin.H{})
}
