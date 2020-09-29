package controller

import (
	"errors"
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/golang/mock/gomock"
)

func TestCheckUuid(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDb := NewMockDatabaseInterface(ctrl)
	AssignDbRepo(mockDb)

	id := uuid.New()
	data := DataBase{
		ID:               id,
		URL:              "mockURL",
		CrawlTimeout:     2,
		Frequency:        5,
		FailureThreshold: 10,
		Status:           true,
		FailureCount:     0,
	}
	// Case 1: ID found
	mockDb.EXPECT().fetchDataWithID(id).Return(data, nil).MaxTimes(1)

	returnedData, err := checkUUID(id.String())
	assert.Equal(t, returnedData, data)
	assert.Equal(t, err, nil)

	// Case 2: ID not found
	mockDb.EXPECT().fetchDataWithID(uuid.UUID{}).DoAndReturn(func(_ uuid.UUID) (DataBase, error) {
		return DataBase{}, errors.New("record not found")
	}).MaxTimes(1)

	returnedData, err = checkUUID(uuid.UUID{}.String())
	assert.Equal(t, returnedData, DataBase{})
	assert.Equal(t, err, errors.New("record not found"))

	// Case 3: Invalid ID
	mockDb.EXPECT().fetchDataWithID(gomock.Any()).DoAndReturn(func(_ uuid.UUID) (DataBase, error) {
		return DataBase{}, errors.New("record not found")
	}).MaxTimes(1)

	returnedData, err = checkUUID("random text")
	assert.Equal(t, returnedData, DataBase{})
	assert.Equal(t, err, errors.New("Invalid UUID"))
}

func TestSingleRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDb := NewMockDatabaseInterface(ctrl)

	ctrl = gomock.NewController(t)
	defer ctrl.Finish()
	mockRequest := NewMockRequestInterface(ctrl)
	AssignRequestRepo(mockRequest)
	AssignDbRepo(mockDb)

	id := uuid.New()
	data := DataBase{
		ID:               id,
		URL:              "mock url",
		CrawlTimeout:     2,
		Frequency:        5,
		FailureThreshold: 10,
		Status:           true,
		FailureCount:     0,
	}
	// Case 1: Correct
	mockRequest.EXPECT().httpRequest(data).Return(&http.Response{StatusCode: http.StatusOK}, nil)
	singleRequest(data)

	// Case 2: Bad Request
	mockRequest.EXPECT().httpRequest(data).Return(&http.Response{StatusCode: http.StatusBadRequest}, errors.New("Mock error"))
	mockDb.EXPECT().fetchDataWithID(data.ID).Return(data, nil)
	dataFailureCountInc := data
	dataFailureCountInc.FailureCount++
	mockDb.EXPECT().updateData(dataFailureCountInc)

	singleRequest(data)
}

func TestAddingDataInDatabase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDb := NewMockDatabaseInterface(ctrl)
	AssignDbRepo(mockDb)

	data := requestData{
		URL:              "mock url",
		CrawlTimeout:     1,
		Frequency:        2,
		FailureThreshold: 10,
	}

	mockDb.EXPECT().addData(gomock.Any())

	returnedData := addingDataInDatabase(data)
	assert.Equal(t, returnedData.URL, data.URL)
	assert.Equal(t, returnedData.CrawlTimeout, data.CrawlTimeout)
	assert.Equal(t, returnedData.Frequency, data.Frequency)
	assert.Equal(t, returnedData.FailureThreshold, data.FailureThreshold)
	assert.Equal(t, returnedData.Status, true)
}

func TestUpdatingDataInDatabase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDb := NewMockDatabaseInterface(ctrl)
	AssignDbRepo(mockDb)
	mockDb.EXPECT().updateData(gomock.Any()).MaxTimes(4)

	dbData := DataBase{
		ID:               uuid.UUID{},
		URL:              "mock url",
		CrawlTimeout:     2,
		Frequency:        5,
		FailureThreshold: 10,
		Status:           true,
		FailureCount:     0,
	}

	// changed CrawlTimeout, Frequency, FailureThreshold
	data := requestData{
		CrawlTimeout:     1,
		Frequency:        2,
		FailureThreshold: 5,
	}
	returnedData := updatingDataInDatabase(data, dbData)
	assert.Equal(t, returnedData.CrawlTimeout, data.CrawlTimeout)
	assert.Equal(t, returnedData.Frequency, data.Frequency)
	assert.Equal(t, returnedData.FailureThreshold, data.FailureThreshold)
	assert.Equal(t, returnedData.FailureCount, 0)

	// changed CrawlTimeout only
	data = requestData{
		CrawlTimeout: 1,
	}
	returnedData = updatingDataInDatabase(data, dbData)
	assert.Equal(t, returnedData.CrawlTimeout, data.CrawlTimeout)
	assert.Equal(t, returnedData.Frequency, dbData.Frequency)
	assert.Equal(t, returnedData.FailureThreshold, dbData.FailureThreshold)
	assert.Equal(t, returnedData.FailureCount, 0)

	// changed Frequency only
	data = requestData{
		Frequency: 2,
	}
	returnedData = updatingDataInDatabase(data, dbData)
	assert.Equal(t, returnedData.CrawlTimeout, dbData.CrawlTimeout)
	assert.Equal(t, returnedData.Frequency, data.Frequency)
	assert.Equal(t, returnedData.FailureThreshold, dbData.FailureThreshold)
	assert.Equal(t, returnedData.FailureCount, 0)

	// changed FailureThreshold only
	data = requestData{
		FailureThreshold: 5,
	}
	returnedData = updatingDataInDatabase(data, dbData)
	assert.Equal(t, returnedData.CrawlTimeout, dbData.CrawlTimeout)
	assert.Equal(t, returnedData.Frequency, dbData.Frequency)
	assert.Equal(t, returnedData.FailureThreshold, data.FailureThreshold)
	assert.Equal(t, returnedData.FailureCount, 0)
}

func TestActivateInDatabase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDb := NewMockDatabaseInterface(ctrl)
	AssignDbRepo(mockDb)

	data := DataBase{
		ID:               uuid.UUID{},
		URL:              "mock url",
		CrawlTimeout:     2,
		Frequency:        5,
		FailureThreshold: 10,
		Status:           false,
		FailureCount:     6,
	}

	mockDb.EXPECT().updateData(gomock.Any())

	returnedData := activateInDatabase(data)
	assert.Equal(t, returnedData.Status, true)
	assert.Equal(t, returnedData.FailureCount, 0)
	assert.Equal(t, returnedData.FailureThreshold, data.FailureThreshold)
	assert.Equal(t, returnedData.Frequency, data.Frequency)
	assert.Equal(t, returnedData.CrawlTimeout, data.CrawlTimeout)
}

func TestDeactivateInDatabase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDb := NewMockDatabaseInterface(ctrl)
	AssignDbRepo(mockDb)

	data := DataBase{
		ID:               uuid.UUID{},
		URL:              "mock url",
		CrawlTimeout:     2,
		Frequency:        5,
		FailureThreshold: 10,
		Status:           true,
		FailureCount:     6,
	}
	mockDb.EXPECT().updateData(gomock.Any())

	returnedData := deactivateInDatabase(data)
	assert.Equal(t, returnedData.Status, false)
	assert.Equal(t, returnedData.FailureCount, data.FailureCount)
	assert.Equal(t, returnedData.FailureThreshold, data.FailureThreshold)
	assert.Equal(t, returnedData.Frequency, data.Frequency)
	assert.Equal(t, returnedData.CrawlTimeout, data.CrawlTimeout)
}
