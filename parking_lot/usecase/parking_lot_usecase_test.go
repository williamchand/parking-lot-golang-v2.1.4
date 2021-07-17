package usecase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/williamchand/parking-lot-golang-v2.1.4/parking_lot/domain"
	"github.com/williamchand/parking-lot-golang-v2.1.4/parking_lot/domain/mocks"
	ucase "github.com/williamchand/parking-lot-golang-v2.1.4/parking_lot/usecase"
)

func TestFetchStatus(t *testing.T) {
	mockParkingLotRepo := new(mocks.ParkingLotRepository)
	regNum1 := "B-1234-RFS"
	regNum2 := "B-1234-RFK"
	colour1 := "Black"
	colour2 := "Black"
	mockParkingLot := domain.ParkingLot{
		ID:                 1,
		RegistrationNumber: &regNum1,
		Colour:             &colour1,
		IsOccupied:         true,
		UpdatedAt:          time.Now(),
		CreatedAt:          time.Now(),
	}

	mockParkingLot2 := domain.ParkingLot{
		ID:                 2,
		RegistrationNumber: &regNum2,
		Colour:             &colour2,
		IsOccupied:         true,
		UpdatedAt:          time.Now(),
		CreatedAt:          time.Now(),
	}
	mockListParkingLot := make([]domain.ParkingLot, 0)
	mockListParkingLot = append(mockListParkingLot, mockParkingLot, mockParkingLot2)

	t.Run("fetch status success", func(t *testing.T) {
		mockParkingLotRepo.On("Fetch", mock.Anything, mock.Anything).Return(mockListParkingLot, nil).Once()
		u := ucase.NewParkingLotUsecase(mockParkingLotRepo, time.Second*2)
		list, err := u.FetchStatus(context.TODO())
		assert.NoError(t, err)
		assert.Len(t, list, len(mockListParkingLot))
		mockParkingLotRepo.AssertExpectations(t)
	})

	t.Run("FetchRegistrationNumber success", func(t *testing.T) {
		mockParkingLotRepo.On("Fetch", mock.Anything, mock.Anything).Return(mockListParkingLot, nil).Once()
		u := ucase.NewParkingLotUsecase(mockParkingLotRepo, time.Second*2)
		list, err := u.FetchRegistrationNumber(context.TODO(), "Black")
		assert.NoError(t, err)
		assert.Len(t, list, len(mockListParkingLot))
		mockParkingLotRepo.AssertExpectations(t)
	})

	t.Run("FetchCarsSlot success", func(t *testing.T) {
		mockParkingLotRepo.On("Fetch", mock.Anything, mock.Anything).Return(mockListParkingLot, nil).Once()
		u := ucase.NewParkingLotUsecase(mockParkingLotRepo, time.Second*2)
		list, err := u.FetchCarsSlot(context.TODO(), "Black")
		assert.NoError(t, err)
		assert.Len(t, list, len(mockListParkingLot))
		mockParkingLotRepo.AssertExpectations(t)
	})
}

func TestGetIdByRegistrationNumber(t *testing.T) {
	mockParkingLotRepo := new(mocks.ParkingLotRepository)
	regNum1 := "B-1234-RFS"
	colour1 := "Black"
	mockParkingLot := domain.ParkingLot{
		ID:                 1,
		RegistrationNumber: &regNum1,
		Colour:             &colour1,
		IsOccupied:         true,
		UpdatedAt:          time.Now(),
		CreatedAt:          time.Now(),
	}

	t.Run("success", func(t *testing.T) {
		mockParkingLotRepo.On("GetIdByRegistrationNumber", mock.Anything, mock.AnythingOfType("string")).Return(mockParkingLot, nil).Once()
		u := ucase.NewParkingLotUsecase(mockParkingLotRepo, time.Second*2)

		a, err := u.GetIdByRegistrationNumber(context.TODO(), regNum1)

		assert.NoError(t, err)
		assert.Equal(t, a, mockParkingLot.ID)
		mockParkingLotRepo.AssertExpectations(t)
	})

	t.Run("error-failed", func(t *testing.T) {
		mockParkingLotRepo.On("GetIdByRegistrationNumber", mock.Anything, mock.AnythingOfType("string")).Return(domain.ParkingLot{}, errors.New("Unexpected")).Once()
		u := ucase.NewParkingLotUsecase(mockParkingLotRepo, time.Second*2)

		_, err := u.GetIdByRegistrationNumber(context.TODO(), regNum1)

		assert.Error(t, err)
		mockParkingLotRepo.AssertExpectations(t)
	})
}

func TestCreateParkingLot(t *testing.T) {
	mockParkingLotRepo := new(mocks.ParkingLotRepository)

	t.Run("success", func(t *testing.T) {
		mockParkingLotRepo.On("DeleteAllSlot", mock.Anything).Return(nil).Once()
		mockParkingLotRepo.On("Store", mock.Anything, mock.AnythingOfType("*domain.ParkingLot")).Return(nil).Times(6)

		u := ucase.NewParkingLotUsecase(mockParkingLotRepo, time.Second*2)

		err := u.CreateParkingLot(context.TODO(), int64(6))

		assert.NoError(t, err)
		mockParkingLotRepo.AssertExpectations(t)
	})
}

func TestOccupyParkingLot(t *testing.T) {
	mockParkingLotRepo := new(mocks.ParkingLotRepository)
	regNum1 := "B-1234-RFS"
	colour1 := "Black"

	t.Run("success", func(t *testing.T) {
		mockParkingLotRepo.On("UpdateOccupied", mock.Anything, mock.AnythingOfType("*domain.ParkingLot")).Return(int64(2), nil).Once()

		u := ucase.NewParkingLotUsecase(mockParkingLotRepo, time.Second*2)

		res, err := u.OccupyParkingLot(context.TODO(), regNum1, colour1)

		assert.NoError(t, err)
		assert.Equal(t, res, int64(2))
		mockParkingLotRepo.AssertExpectations(t)
	})
	t.Run("error-failed", func(t *testing.T) {
		mockParkingLotRepo.On("UpdateOccupied", mock.Anything, mock.AnythingOfType("*domain.ParkingLot")).Return(int64(0), errors.New("Unexpected")).Once()
		u := ucase.NewParkingLotUsecase(mockParkingLotRepo, time.Second*2)

		_, err := u.OccupyParkingLot(context.TODO(), regNum1, colour1)

		assert.Error(t, err)
		mockParkingLotRepo.AssertExpectations(t)
	})
}

func TestUnOccupyParkingLot(t *testing.T) {
	mockParkingLotRepo := new(mocks.ParkingLotRepository)

	t.Run("success", func(t *testing.T) {
		mockParkingLotRepo.On("UpdateUnOccupied", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()

		u := ucase.NewParkingLotUsecase(mockParkingLotRepo, time.Second*2)

		err := u.UnOccupyParkingLot(context.TODO(), int64(2))

		assert.NoError(t, err)
		mockParkingLotRepo.AssertExpectations(t)
	})
	t.Run("error-failed", func(t *testing.T) {
		mockParkingLotRepo.On("UpdateUnOccupied", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("Unexpected")).Once()
		u := ucase.NewParkingLotUsecase(mockParkingLotRepo, time.Second*2)

		err := u.UnOccupyParkingLot(context.TODO(), int64(2))

		assert.Error(t, err)
		mockParkingLotRepo.AssertExpectations(t)
	})
}
