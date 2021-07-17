package repository_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"

	"github.com/williamchand/parking-lot-golang-v2.1.4/parking_lot/domain"
	parkingLotMysqlRepo "github.com/williamchand/parking-lot-golang-v2.1.4/parking_lot/repository"
)

func TestFetchStatus(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	regNum1 := "B-1234-RFS"
	regNum2 := "B-1999-RFD"
	colour1 := "Black"
	colour2 := "Green"
	mockParkingLot := []domain.ParkingLot{
		{
			ID:                 1,
			RegistrationNumber: &regNum1,
			Colour:             &colour1,
			IsOccupied:         true,
			UpdatedAt:          time.Now(),
			CreatedAt:          time.Now(),
		},
		{
			ID:                 3,
			RegistrationNumber: &regNum2,
			Colour:             &colour2,
			IsOccupied:         true,
			UpdatedAt:          time.Now(),
			CreatedAt:          time.Now(),
		},
	}

	rows := sqlmock.NewRows([]string{"id", "registration_number", "colour", "is_occupied", "created_at", "updated_at"}).
		AddRow(mockParkingLot[0].ID, mockParkingLot[0].RegistrationNumber, mockParkingLot[0].Colour,
			mockParkingLot[0].IsOccupied, mockParkingLot[0].CreatedAt, mockParkingLot[0].UpdatedAt).
		AddRow(mockParkingLot[1].ID, mockParkingLot[1].RegistrationNumber, mockParkingLot[1].Colour,
			mockParkingLot[1].IsOccupied, mockParkingLot[1].CreatedAt, mockParkingLot[1].UpdatedAt)

	query := `SELECT id, registration_number, colour, is_occupied, created_at, updated_at
						FROM parking_lot WHERE is_occupied = true ORDER BY id`

	mock.ExpectQuery(query).WillReturnRows(rows)
	a := parkingLotMysqlRepo.NewMysqlParkingLotRepository(db)
	list, err := a.Fetch(context.TODO(), nil)
	assert.NoError(t, err)
	assert.Len(t, list, 2)
}

func TestFetchByColour(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	regNum1 := "B-1234-RFS"
	regNum2 := "B-1999-RFD"
	colour1 := "Black"
	mockParkingLot := []domain.ParkingLot{
		domain.ParkingLot{
			ID:                 1,
			RegistrationNumber: &regNum1,
			Colour:             &colour1,
			IsOccupied:         true,
			UpdatedAt:          time.Now(),
			CreatedAt:          time.Now(),
		},
		domain.ParkingLot{
			ID:                 3,
			RegistrationNumber: &regNum2,
			Colour:             &colour1,
			IsOccupied:         true,
			UpdatedAt:          time.Now(),
			CreatedAt:          time.Now(),
		},
	}

	rows := sqlmock.NewRows([]string{"id", "registration_number", "colour", "is_occupied", "created_at", "updated_at"}).
		AddRow(mockParkingLot[0].ID, mockParkingLot[0].RegistrationNumber, mockParkingLot[0].Colour,
			mockParkingLot[0].IsOccupied, mockParkingLot[0].CreatedAt, mockParkingLot[0].UpdatedAt).
		AddRow(mockParkingLot[1].ID, mockParkingLot[1].RegistrationNumber, mockParkingLot[1].Colour,
			mockParkingLot[1].IsOccupied, mockParkingLot[1].CreatedAt, mockParkingLot[1].UpdatedAt)

	query := `SELECT id, registration_number, colour, is_occupied, created_at, updated_at
						FROM parking_lot WHERE is_occupied = true and colour = \? ORDER BY id`

	mock.ExpectQuery(query).WillReturnRows(rows)
	a := parkingLotMysqlRepo.NewMysqlParkingLotRepository(db)
	list, err := a.Fetch(context.TODO(), &colour1)
	assert.NoError(t, err)
	assert.Len(t, list, 2)
}

func TestGetIdByRegistrationNumber(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	regNum1 := "B-1999-RFD"
	colour1 := "Black"
	rows := sqlmock.NewRows([]string{"id", "registration_number", "colour", "is_occupied", "updated_at", "created_at"}).
		AddRow(1, &regNum1, &colour1,
			true, time.Now(), time.Now())

	query := `SELECT id, registration_number, colour, is_occupied, created_at, updated_at
						FROM parking_lot WHERE is_occupied = true and registration_number = \?`
	mock.ExpectQuery(query).WillReturnRows(rows)
	a := parkingLotMysqlRepo.NewMysqlParkingLotRepository(db)
	list, err := a.GetIdByRegistrationNumber(context.TODO(), regNum1)
	assert.NoError(t, err)
	assert.NotNil(t, list)
}

func TestGetIdByRegistrationNumberError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	regNum1 := "B-1999-RFD"
	rows := sqlmock.NewRows([]string{"id", "registration_number", "colour", "is_occupied", "updated_at", "created_at"})

	query := `SELECT id, registration_number, colour, is_occupied, created_at, updated_at
						FROM parking_lot WHERE is_occupied = true and registration_number = \?`
	mock.ExpectQuery(query).WillReturnRows(rows)
	a := parkingLotMysqlRepo.NewMysqlParkingLotRepository(db)
	_, err = a.GetIdByRegistrationNumber(context.TODO(), regNum1)
	assert.Error(t, err)
}

func TestDeleteAllSlot(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	query := `DELETE FROM parking_lot`
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WillReturnResult(sqlmock.NewResult(6, 1))

	a := parkingLotMysqlRepo.NewMysqlParkingLotRepository(db)

	err = a.DeleteAllSlot(context.TODO())
	assert.NoError(t, err)
}

func TestStore(t *testing.T) {
	ar := &domain.ParkingLot{
		ID:                 1,
		RegistrationNumber: nil,
		Colour:             nil,
		IsOccupied:         false,
		UpdatedAt:          time.Now(),
		CreatedAt:          time.Now(),
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	query := "INSERT parking_lot SET id=\\? , registration_number=\\? , colour=\\?, is_occupied=\\?, updated_at=\\? , created_at=\\?"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(ar.ID, ar.RegistrationNumber, ar.Colour, ar.IsOccupied, ar.CreatedAt, ar.UpdatedAt).WillReturnResult(sqlmock.NewResult(1, 1))

	a := parkingLotMysqlRepo.NewMysqlParkingLotRepository(db)

	err = a.Store(context.TODO(), ar)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), ar.ID)
}

func TestUpdateOccupied(t *testing.T) {
	regNum1 := "B-1999-RFD"
	colour1 := "Black"
	ar := &domain.ParkingLot{
		RegistrationNumber: &regNum1,
		Colour:             &colour1,
		UpdatedAt:          time.Now(),
		CreatedAt:          time.Now(),
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	rows := sqlmock.NewRows([]string{"id", "registration_number", "colour", "is_occupied", "updated_at", "created_at"}).
		AddRow(1, nil, nil,
			false, time.Now(), time.Now())

	query := `SELECT id, registration_number, colour, is_occupied, created_at, updated_at
  						FROM parking_lot WHERE is_occupied=false ORDER BY id LIMIT 1`
	mock.ExpectQuery(query).WillReturnRows(rows)

	query = `UPDATE parking_lot set registration_number=\?, colour=\?, is_occupied=TRUE, updated_at=\?
						WHERE id = \?`

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(ar.RegistrationNumber, ar.Colour, ar.UpdatedAt, 1).WillReturnResult(sqlmock.NewResult(1, 1))

	a := parkingLotMysqlRepo.NewMysqlParkingLotRepository(db)

	res, err := a.UpdateOccupied(context.TODO(), ar)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), res)
}

func TestUpdateOccupiedFailed(t *testing.T) {
	regNum1 := "B-1999-RFD"
	colour1 := "Black"
	ar := &domain.ParkingLot{
		RegistrationNumber: &regNum1,
		Colour:             &colour1,
		UpdatedAt:          time.Now(),
		CreatedAt:          time.Now(),
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	rows := sqlmock.NewRows([]string{"id", "registration_number", "colour", "is_occupied", "updated_at", "created_at"})

	query := `SELECT id, registration_number, colour, is_occupied, created_at, updated_at
  						FROM parking_lot WHERE is_occupied=false ORDER BY id LIMIT 1`
	mock.ExpectQuery(query).WillReturnRows(rows)

	a := parkingLotMysqlRepo.NewMysqlParkingLotRepository(db)

	_, err = a.UpdateOccupied(context.TODO(), ar)
	assert.Error(t, err)
}

func TestUpdateUnOccupied(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	query := `UPDATE parking_lot set registration_number=NULL, colour=NULL, is_occupied=false, updated_at = \?
						WHERE id = \? and is_occupied=true`
	currentTime := time.Now()
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(currentTime, 2).WillReturnResult(sqlmock.NewResult(1, 1))

	a := parkingLotMysqlRepo.NewMysqlParkingLotRepository(db)

	err = a.UpdateUnOccupied(context.TODO(), 2, currentTime)
	assert.NoError(t, err)
}
