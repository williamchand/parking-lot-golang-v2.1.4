package domain

import (
	"context"
	"time"
)

type ParkingLot struct {
	ID        					int64     `json:"id"`
	RegistrationNumber  *string    `json:"registration_number"`
	Colour   						*string    `json:"colour"`
	IsOccupied 					bool    	`json:"is_occupied"`
	UpdatedAt 					time.Time `json:"updated_at"`
	CreatedAt 					time.Time `json:"created_at"`
}

type ParkingLotUsecase interface {
	FetchStatus(ctx context.Context) ([]ParkingLot, error)
	FetchRegistrationNumber(ctx context.Context, colour string) ([]string, error)
	FetchCarsSlot(ctx context.Context, colour string) ([]string, error)
	GetIdByRegistrationNumber(ctx context.Context, registrationNumber string) (int64, error)
	CreateParkingLot(ctx context.Context, slots int64) error
	OccupyParkingLot(ctx context.Context, registrationNumber string, colour string) (int64, error)
	UnOccupyParkingLot(ctx context.Context, id int64) error
}

type ParkingLotRepository interface {
	Fetch(ctx context.Context, colour *string) ([]ParkingLot, error)
	GetIdByRegistrationNumber(ctx context.Context, registrationNumber string) (ParkingLot, error)
	DeleteAllSlot(ctx context.Context) error
	Store(ctx context.Context, a *ParkingLot) error
	UpdateOccupied(ctx context.Context, ar *ParkingLot) (int64, error)
	UpdateUnOccupied(ctx context.Context, id int64, updatedAt time.Time) error
}
