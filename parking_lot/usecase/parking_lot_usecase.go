package usecase

import (
	"context"
	"time"
	"fmt"

	"github.com/williamchand/parking-lot-golang-v2.1.4/parking_lot/domain"
)

type parkingLotUsecase struct {
	parkingLotRepo     domain.ParkingLotRepository
	contextTimeout time.Duration
}

// NewParkingLotUsecase will create new an NewParkingLotUsecase object representation of domain.NewParkingLotUsecase interface
func NewParkingLotUsecase(a domain.ParkingLotRepository, timeout time.Duration) domain.ParkingLotUsecase {
	return &parkingLotUsecase{
		parkingLotRepo:    a,
		contextTimeout: timeout,
	}
}

func (a *parkingLotUsecase) FetchStatus(c context.Context) (res []domain.ParkingLot, err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	res, err = a.parkingLotRepo.Fetch(ctx, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (a *parkingLotUsecase) FetchRegistrationNumber(c context.Context, colour string) (res []string, err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	values, err := a.parkingLotRepo.Fetch(ctx, &colour)
	if err != nil {
		return nil, err
	}

	res = []string{}
	for _, element := range values {
		res = append(res, *element.RegistrationNumber)
	}

	return
}

func (a *parkingLotUsecase) FetchCarsSlot(c context.Context, colour string) (res []string, err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	values, err := a.parkingLotRepo.Fetch(ctx, &colour)
	if err != nil {
		return nil, err
	}

	res = []string{}
	for _, element := range values {
		res = append(res, fmt.Sprintf("%d", element.ID))
	}

	return
}

func (a *parkingLotUsecase) GetIdByRegistrationNumber(c context.Context, registrationNumber string) (res int64, err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	result, err := a.parkingLotRepo.GetIdByRegistrationNumber(ctx, registrationNumber)
	if err != nil {
		return
	}
	res = result.ID
	return
}

func (a *parkingLotUsecase) CreateParkingLot(c context.Context, slots int64) (err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()
	err = a.parkingLotRepo.DeleteAllSlot(ctx)
	if err != nil {
		return
	}
	for i := int64(1); i <= slots; i++ {
		err = a.parkingLotRepo.Store(ctx, &domain.ParkingLot{
			ID: i,
			UpdatedAt: time.Now(),
			CreatedAt: time.Now(),
		})
		if err != nil {
			return
		}
	}

	return
}

func (a *parkingLotUsecase) OccupyParkingLot(c context.Context, registrationNumber string, colour string) (res int64, err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()
	ar := domain.ParkingLot{
		RegistrationNumber: &registrationNumber,
		Colour: &colour,
		UpdatedAt: time.Now(),
	}
	return a.parkingLotRepo.UpdateOccupied(ctx, &ar)
}

func (a *parkingLotUsecase) UnOccupyParkingLot(c context.Context, id int64) (err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	UpdatedAt:= time.Now()
	return a.parkingLotRepo.UpdateUnOccupied(ctx, id,UpdatedAt)
}
