package services

import (
	"errors"
	ipgeo "gorepair-rest-api/src/ip-geo"
	_ipgeomock "gorepair-rest-api/src/ip-geo/mocks"
	"gorepair-rest-api/src/w-services/entities"
	"gorepair-rest-api/src/w-services/entities/mocks"
	"testing"

	_ws "gorepair-rest-api/src/workshops/entities"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	wservicesRepo mocks.WServicesMysqlRepositoryInterface
	ipgeoRepo _ipgeomock.Repository
	wserviceUsecase entities.WServicesService
	wserviceDomain entities.WServices
	wsAddress _ws.WorkshopAddress
	locationIP ipgeo.Domain
)

func setup() {
	wserviceUsecase = NewWServicesService(&wservicesRepo, &ipgeoRepo)

	wserviceDomain = entities.WServices{
		ID: 1,
		WorkshopID: 1,
		Vehicle: "BMW 007",
		VehicleType: "Sports Car",
		Services: "Window Repair",
		Price: 123456789,
	}

	wsAddress = _ws.WorkshopAddress{
		WorkshopID: 1,
		City: "Jakarta",
	}

	locationIP = ipgeo.Domain{
		City: "Jakarta",
	}

}

func TestGetAll(t *testing.T) {
	setup()

	t.Run("Test GetAll Valid", func(t *testing.T) {
		wservicesRepo.On("GetAll").Return([]entities.WServices{wserviceDomain}, nil).Once()

		resp, err := wserviceUsecase.GetAll()

		assert.Nil(t, err)
		assert.Contains(t, resp, wserviceDomain)
	})

	t.Run("Test GetAll Error", func(t *testing.T) {
		wservicesRepo.On("GetAll").Return([]entities.WServices{wserviceDomain}, errors.New("")).Once()

		_, err := wserviceUsecase.GetAll()

		assert.NotNil(t, err)
	})
}

func TestGetDetails(t *testing.T) {
	setup()

	wservicesRepo.On("GetDetails",
	mock.Anything).Return(wserviceDomain, nil).Twice()
	
	t.Run("Test GetDetails 1", func(t *testing.T) {
		resp, err := wserviceUsecase.GetDetails("1")

		assert.Nil(t, err)
		assert.Equal(t, "BMW 007", resp.Vehicle)
	})

	t.Run("Test GetDetails 2", func(t *testing.T) {
		_, err := wserviceUsecase.GetDetails("1a")

		assert.NotNil(t, err)
	})

	t.Run("Test GetDetails 3", func(t *testing.T) {
		wservicesRepo.On("GetDetails",
	mock.Anything).Return(wserviceDomain, nil).Once()
		resp, _ := wserviceUsecase.GetDetails("2")

		assert.NotEqual(t, 2, resp.ID)
	})
}

func TestGetAllWorkshop(t *testing.T) {
	setup()

	t.Run("Test GetAllWorkshop", func(t *testing.T) {
		wservicesRepo.On("GetAllWorkshop",
			mock.AnythingOfType("string")).Return([]_ws.WorkshopAddress{wsAddress}, nil).Once()
		ipgeoRepo.On("GetLocationByIP",
			mock.AnythingOfType("string")).Return(locationIP, nil).Once()

		_, err := wserviceUsecase.GetAllWorkshop("8.8.8.8")

		assert.Nil(t, err)
	})

	t.Run("Test GetAllWorkshop", func(t *testing.T) {
		wservicesRepo.On("GetAllWorkshop",
			mock.AnythingOfType("string")).Return([]_ws.WorkshopAddress{wsAddress}, errors.New("")).Once()
		ipgeoRepo.On("GetLocationByIP",
			mock.AnythingOfType("string")).Return(locationIP, nil).Once()
		_, err := wserviceUsecase.GetAllWorkshop("8.8.8.8")

		assert.NotNil(t, err)
	})
}