// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	domain "github.com/williamchand/parking-lot-golang-v2.1.4/parking_lot/domain"
)

// ParkingLotRepository is an autogenerated mock type for the ParkingLotRepository type
type ParkingLotRepository struct {
	mock.Mock
}

// DeleteAllSlot provides a mock function with given fields: ctx
func (_m *ParkingLotRepository) DeleteAllSlot(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Fetch provides a mock function with given fields: ctx, colour
func (_m *ParkingLotRepository) Fetch(ctx context.Context, colour *string) ([]domain.ParkingLot, error) {
	ret := _m.Called(ctx, colour)

	var r0 []domain.ParkingLot
	if rf, ok := ret.Get(0).(func(context.Context, *string) []domain.ParkingLot); ok {
		r0 = rf(ctx, colour)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.ParkingLot)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *string) error); ok {
		r1 = rf(ctx, colour)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetIdByRegistrationNumber provides a mock function with given fields: ctx, registrationNumber
func (_m *ParkingLotRepository) GetIdByRegistrationNumber(ctx context.Context, registrationNumber string) (domain.ParkingLot, error) {
	ret := _m.Called(ctx, registrationNumber)

	var r0 domain.ParkingLot
	if rf, ok := ret.Get(0).(func(context.Context, string) domain.ParkingLot); ok {
		r0 = rf(ctx, registrationNumber)
	} else {
		r0 = ret.Get(0).(domain.ParkingLot)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, registrationNumber)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: ctx, a
func (_m *ParkingLotRepository) Store(ctx context.Context, a *domain.ParkingLot) error {
	ret := _m.Called(ctx, a)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.ParkingLot) error); ok {
		r0 = rf(ctx, a)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateOccupied provides a mock function with given fields: ctx, ar
func (_m *ParkingLotRepository) UpdateOccupied(ctx context.Context, ar *domain.ParkingLot) (int64, error) {
	ret := _m.Called(ctx, ar)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, *domain.ParkingLot) int64); ok {
		r0 = rf(ctx, ar)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *domain.ParkingLot) error); ok {
		r1 = rf(ctx, ar)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUnOccupied provides a mock function with given fields: ctx, id
func (_m *ParkingLotRepository) UpdateUnOccupied(ctx context.Context, id int64) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
