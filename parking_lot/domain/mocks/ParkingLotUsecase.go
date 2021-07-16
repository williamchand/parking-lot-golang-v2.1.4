// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	domain "github.com/williamchand/parking-lot-golang-v2.1.4/parking_lot/domain"
)

// ParkingLotUsecase is an autogenerated mock type for the ParkingLotUsecase type
type ParkingLotUsecase struct {
	mock.Mock
}

// Delete provides a mock function with given fields: ctx, id
func (_m *ParkingLotUsecase) Delete(ctx context.Context, id int64) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Fetch provides a mock function with given fields: ctx, cursor, num
func (_m *ParkingLotUsecase) Fetch(ctx context.Context, cursor string, num int64) ([]domain.ParkingLot, string, error) {
	ret := _m.Called(ctx, cursor, num)

	var r0 []domain.ParkingLot
	if rf, ok := ret.Get(0).(func(context.Context, string, int64) []domain.ParkingLot); ok {
		r0 = rf(ctx, cursor, num)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.ParkingLot)
		}
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(context.Context, string, int64) string); ok {
		r1 = rf(ctx, cursor, num)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string, int64) error); ok {
		r2 = rf(ctx, cursor, num)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *ParkingLotUsecase) GetByID(ctx context.Context, id int64) (domain.ParkingLot, error) {
	ret := _m.Called(ctx, id)

	var r0 domain.ParkingLot
	if rf, ok := ret.Get(0).(func(context.Context, int64) domain.ParkingLot); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(domain.ParkingLot)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByTitle provides a mock function with given fields: ctx, title
func (_m *ParkingLotUsecase) GetByTitle(ctx context.Context, title string) (domain.ParkingLot, error) {
	ret := _m.Called(ctx, title)

	var r0 domain.ParkingLot
	if rf, ok := ret.Get(0).(func(context.Context, string) domain.ParkingLot); ok {
		r0 = rf(ctx, title)
	} else {
		r0 = ret.Get(0).(domain.ParkingLot)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, title)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: _a0, _a1
func (_m *ParkingLotUsecase) Store(_a0 context.Context, _a1 *domain.ParkingLot) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.ParkingLot) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: ctx, ar
func (_m *ParkingLotUsecase) Update(ctx context.Context, ar *domain.ParkingLot) error {
	ret := _m.Called(ctx, ar)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.ParkingLot) error); ok {
		r0 = rf(ctx, ar)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
