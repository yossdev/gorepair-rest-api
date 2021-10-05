// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	entities "gorepair-rest-api/src/w-services/entities"

	mock "github.com/stretchr/testify/mock"

	workshopsentities "gorepair-rest-api/src/workshops/entities"
)

// WServicesMysqlRepositoryInterface is an autogenerated mock type for the WServicesMysqlRepositoryInterface type
type WServicesMysqlRepositoryInterface struct {
	mock.Mock
}

// GetAll provides a mock function with given fields:
func (_m *WServicesMysqlRepositoryInterface) GetAll() ([]entities.WServices, error) {
	ret := _m.Called()

	var r0 []entities.WServices
	if rf, ok := ret.Get(0).(func() []entities.WServices); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entities.WServices)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllWorkshop provides a mock function with given fields: city
func (_m *WServicesMysqlRepositoryInterface) GetAllWorkshop(city string) ([]workshopsentities.WorkshopAddress, error) {
	ret := _m.Called(city)

	var r0 []workshopsentities.WorkshopAddress
	if rf, ok := ret.Get(0).(func(string) []workshopsentities.WorkshopAddress); ok {
		r0 = rf(city)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]workshopsentities.WorkshopAddress)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(city)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDetails provides a mock function with given fields: id
func (_m *WServicesMysqlRepositoryInterface) GetDetails(id uint64) (entities.WServices, error) {
	ret := _m.Called(id)

	var r0 entities.WServices
	if rf, ok := ret.Get(0).(func(uint64) entities.WServices); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(entities.WServices)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
