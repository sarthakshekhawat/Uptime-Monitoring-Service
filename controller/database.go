package controller

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

var db *gorm.DB

// AssignValue assigns value to the variable "db"
func AssignValue(database *gorm.DB) {
	db = database
}

//DataBase is a struct for the table in database
type DataBase struct {
	ID               uuid.UUID `gorm:"primaryKey"`
	URL              string
	CrawlTimeout     int
	Frequency        int
	FailureThreshold int
	Status           bool
	FailureCount     int
}

// DatabaseInterface is the interface for all the fuctions which access the database
type DatabaseInterface interface {
	fetchDataWithID(id uuid.UUID) (DataBase, error)
	updateData(data DataBase)
	addData(data DataBase)
	deleteData(data DataBase)
}

// DatabaseReceiver is the receiver type for the functions which access the database
type DatabaseReceiver struct{}

var dbRepo DatabaseInterface

// AssignDbRepo assigns the value to the dbRepo
func AssignDbRepo(di DatabaseInterface) {
	dbRepo = di
}

func (cr *DatabaseReceiver) fetchDataWithID(id uuid.UUID) (DataBase, error) {
	var data DataBase
	result := db.First(&data, id)
	return data, result.Error
}

func (cr *DatabaseReceiver) updateData(data DataBase) {
	db.Save(&data)
}

func (cr *DatabaseReceiver) addData(data DataBase) {
	db.Create(&data)
}

func (cr *DatabaseReceiver) deleteData(data DataBase) {
	db.Delete(&data)
}
