package controller

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type requestData struct {
	URL              string `json:"url" binding:"required"`
	CrawlTimeout     int    `json:"crawl_timeout" binding:"required"`
	Frequency        int    `json:"frequency" binding:"required"`
	FailureThreshold int    `json:"failure_threshold" binding:"required"`
}

func singleRequest(data DataBase) {
	resp, err := requestRepo.httpRequest(data)
	if err != nil {
		data, err = dbRepo.fetchDataWithID(data.ID)
		if err != nil {
			return
		}
		data.FailureCount++ //increasing the failure count
		dbRepo.updateData(data)
		fmt.Printf("\"%v\" is down\n", data.URL) // Printing
	} else {
		fmt.Printf("\"%v\" is up, Status Code: %v\n", data.URL, resp.StatusCode) // Printing
	}
}

func startRequest(data DataBase) {
	for {
		var err error
		data, err = dbRepo.fetchDataWithID(data.ID)
		if err != nil || data.Status == false {
			return
		}
		if data.FailureCount >= data.FailureThreshold {
			data.Status = false
			dbRepo.updateData(data)
			return
		}
		go singleRequest(data)
		time.Sleep(time.Duration(data.Frequency) * time.Second)
	}
}

// check whether the UUID given is valid or not,
// if valid and present in DB the returns the data from DB
func checkUUID(id string) (DataBase, error) {
	uuid, err := uuid.Parse(id) // parsing the string into UUID
	if err != nil {
		err = errors.New("Invalid UUID")
		return DataBase{}, err
	}
	return dbRepo.fetchDataWithID(uuid) // return (DataBase, error)
}

func addingDataInDatabase(data requestData) DataBase {
	newUUID := uuid.New() //generating new uuid
	dataInsert := DataBase{
		ID:               newUUID,
		URL:              data.URL,
		CrawlTimeout:     data.CrawlTimeout,
		Frequency:        data.Frequency,
		FailureThreshold: data.FailureThreshold,
		Status:           true,
		FailureCount:     0,
	}
	dbRepo.addData(dataInsert)
	return dataInsert
}

func updatingDataInDatabase(data requestData, dbData DataBase) DataBase {
	if data.CrawlTimeout != 0 {
		dbData.CrawlTimeout = data.CrawlTimeout
	}
	if data.FailureThreshold != 0 {
		dbData.FailureThreshold = data.FailureThreshold
	}
	if data.Frequency != 0 {
		dbData.Frequency = data.Frequency
	}
	dbData.FailureCount = 0
	dbRepo.updateData(dbData)
	return dbData
}

func activateInDatabase(data DataBase) DataBase {
	data.Status = true
	data.FailureCount = 0
	dbRepo.updateData(data)
	return data
}

func deactivateInDatabase(data DataBase) DataBase {
	data.Status = false
	dbRepo.updateData(data)
	return data
}
