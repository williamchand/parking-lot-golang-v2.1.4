package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/williamchand/parking-lot-golang-v2.1.4/parking_lot/domain"
)

type mysqlParkingLotRepository struct {
	Conn *sql.DB
}

// NewMysqlParkingLotRepository will create an object that represent the parkingLot.Repository interface
func NewMysqlParkingLotRepository(Conn *sql.DB) domain.ParkingLotRepository {
	return &mysqlParkingLotRepository{Conn}
}

func (m *mysqlParkingLotRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []domain.ParkingLot, err error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = rows.Close()
	}()

	result = make([]domain.ParkingLot, 0)
	for rows.Next() {
		t := domain.ParkingLot{}
		err = rows.Scan(
			&t.ID,
			&t.RegistrationNumber,
			&t.Colour,
			&t.CreatedAt,
			&t.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

func (m *mysqlParkingLotRepository) Fetch(ctx context.Context, colour *string) (res []domain.ParkingLot, err error) {
	args := ""
	if colour != nil {
		args = "and colour = ?"
	}
	query := fmt.Sprintf(`SELECT id,registration_number,colour, is_occupied, created_at, updated_at
											 FROM parking_lot WHERE is_occupied = TRUE %s ORDER BY id`, args)
	if colour != nil {
		res, err = m.fetch(ctx, query, colour)
	} else {
		res, err = m.fetch(ctx, query)
	}
	if err != nil {
		return nil, err
	}

	return
}

func (m *mysqlParkingLotRepository) GetIdByRegistrationNumber(ctx context.Context, registrationNumber string) (res domain.ParkingLot, err error) {
	query := `SELECT id,registration_number,colour, created_at, updated_at
  						FROM parking_lot WHERE is_occupied = TRUE and registration_number = ?`

	list, err := m.fetch(ctx, query, registrationNumber)
	if err != nil {
		return
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return res, domain.ErrNotFound
	}
	return
}

func (m *mysqlParkingLotRepository) DeleteAllSlot(ctx context.Context) (err error) {
	query := "DELETE FROM parking_lot"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx)
	if err != nil {
		return
	}

	_, err = res.RowsAffected()
	if err != nil {
		return
	}

	return
}

func (m *mysqlParkingLotRepository) Store(ctx context.Context, a *domain.ParkingLot) (err error) {
	query := `INSERT parking_lot SET id=?, registration_number=?, colour=?, is_occupied=?, created_at=?, updated_at=?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	_, err = stmt.ExecContext(ctx, a.ID, nil, nil, false, a.CreatedAt, a.UpdatedAt)
	if err != nil {
		return
	}
	return
}

func (m *mysqlParkingLotRepository) UpdateOccupied(ctx context.Context, ar *domain.ParkingLot) (res int64, err error) {
	query := `UPDATE parking_lot set registration_number=?, colour=?, is_occupied=?, updated_at=?
						WHERE id in (SELECT id FROM (SELECT id FROM parking_lot WHERE is_occupied=false ORDER BY id LIMIT 1) as t)`

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	result, err := stmt.ExecContext(ctx, ar.RegistrationNumber, ar.Colour, true, ar.UpdatedAt)
	if err != nil {
		return
	}
	affect, err := result.RowsAffected()
	if err != nil {
		return
	}
	if affect != 1 {
		err = fmt.Errorf("Weird  Behavior. Total Affected: %d", affect)
		return
	}

	res, err = result.LastInsertId()
	if err != nil {
		err = fmt.Errorf("Weird  can't find last insert id: %d", affect)
		return
	}
	return
}

func (m *mysqlParkingLotRepository) UpdateUnOccupied(ctx context.Context, id int64, updatedAt time.Time) (err error) {
	query := `UPDATE parking_lot set registration_number=NULL, colour=NULL, is_occupied=FALSE, updated_at = ?
						WHERE id = ? and is_occupied=TRUE`

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, id)
	if err != nil {
		return
	}
	affect, err := res.RowsAffected()
	if err != nil {
		return
	}
	if affect != 1 {
		err = fmt.Errorf("Weird  Behavior. Total Affected: %d", affect)
		return
	}

	return
}
