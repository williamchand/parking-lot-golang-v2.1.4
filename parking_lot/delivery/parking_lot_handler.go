package delivery

import (
	"net/http"
	"strconv"
	"fmt"

	"github.com/labstack/echo"

	"github.com/williamchand/parking-lot-golang-v2.1.4/parking_lot/domain"
)


// ParkingLotHandler  represent the httphandler for parking lot
type ParkingLotHandler struct {
	PUsecase domain.ParkingLotUsecase
}

// NewParkingLotHandler will initialize the Parking lot/ resources endpoint
func NewParkingLotHandler(e *echo.Echo, us domain.ParkingLotUsecase) {
	handler := &ParkingLotHandler{
		PUsecase: us,
	}
	e.POST("/create_parking_lot/:parking_lot", handler.CreateParkingLot)
	e.POST("/park/:registration_number/:colour", handler.OccupyParkingLot)
	e.POST("/leave/:id", handler.UnOccupyParkingLot)
	e.GET("/status", handler.FetchStatus)
	e.GET("/slot_number/car_registration_number/:registration_number", handler.GetIdByRegistrationNumber)
	e.GET("/cars_registration_number/colour/:colour", handler.FetchRegistrationNumber)
	e.GET("/cars_slot/colour/:colour", handler.FetchCarsSlot)
}

// CreateParkingLot will create parking slot based on given params
func (a *ParkingLotHandler) CreateParkingLot(c echo.Context) error {
	numS := c.Param("parking_lot")
	num, _ := strconv.Atoi(numS)
	ctx := c.Request().Context()

	err := a.PUsecase.CreateParkingLot(ctx, int64(num))
	if err != nil {
		return c.JSON(getStatusCode(err),  err.Error())
	}
	response := fmt.Sprintf("Created a parking lot with %d slots", num)

	return c.JSON(http.StatusOK, response)
}

// OccupyParkingLot will occupy parking lot based on given params
func (a *ParkingLotHandler) OccupyParkingLot(c echo.Context) error {
	regNum := c.Param("registration_number")
	colour := c.Param("colour")
	ctx := c.Request().Context()

	slot_number, err := a.PUsecase.OccupyParkingLot(ctx, regNum, colour)
	if err == domain.ErrNotFound { 
		return c.JSON(getStatusCode(err),  "Sorry, parking lot is full")
	} else if err != nil {
		return c.JSON(getStatusCode(err),  err.Error())
	}
	response := fmt.Sprintf("Allocated slot number: %d", slot_number)

	return c.JSON(http.StatusOK, response)
}

// UnOccupyParkingLot will unoccupy parking lot based on given params
func (a *ParkingLotHandler) UnOccupyParkingLot(c echo.Context) error {
	numS := c.Param("id")
	num, _ := strconv.Atoi(numS)
	ctx := c.Request().Context()

	err := a.PUsecase.UnOccupyParkingLot(ctx, int64(num))
	if err != nil {
		return c.JSON(getStatusCode(err),  err.Error())
	}
	response := fmt.Sprintf("Slot number %d is free", num)

	return c.JSON(http.StatusOK, response)
}

// FetchStatus will fetch the FetchStatus based on given params
func (a *ParkingLotHandler) FetchStatus(c echo.Context) error {
	ctx := c.Request().Context()

	listAr, err := a.PUsecase.FetchStatus(ctx)
	if err != nil {
		return c.JSON(getStatusCode(err),  err.Error())
	}

	return c.JSON(http.StatusOK, listAr)
}

// FetchRegistrationNumber will fetch the FetchRegistrationNumber based on given params
func (a *ParkingLotHandler) FetchRegistrationNumber(c echo.Context) error {
	colour := c.Param("colour")
	ctx := c.Request().Context()

	listAr, err := a.PUsecase.FetchRegistrationNumber(ctx, colour)
	if err != nil {
		return c.JSON(getStatusCode(err),  err.Error())
	}

	return c.JSON(http.StatusOK, listAr)
}

// FetchCarsSlot will fetch the FetchCarsSlot based on given params
func (a *ParkingLotHandler) FetchCarsSlot(c echo.Context) error {
	colour := c.Param("colour")
	ctx := c.Request().Context()

	listAr, err := a.PUsecase.FetchCarsSlot(ctx, colour)
	if err != nil {
		return c.JSON(getStatusCode(err),  err.Error())
	}

	return c.JSON(http.StatusOK, listAr)
}
// GetByID will get article by given id
func (a *ParkingLotHandler) GetIdByRegistrationNumber(c echo.Context) error {
	regNum := c.Param("registration_number")
	ctx := c.Request().Context()

	registrationNumber, err := a.PUsecase.GetIdByRegistrationNumber(ctx, regNum)
	if err == domain.ErrNotFound { 
		return c.JSON(getStatusCode(err),  "Not found")
	} else if err != nil {
		return c.JSON(getStatusCode(err),  err.Error())
	}

	return c.JSON(http.StatusOK, registrationNumber)
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	switch err {
	case domain.ErrInternalServerError:
		return http.StatusInternalServerError
	case domain.ErrNotFound:
		return http.StatusNotFound
	case domain.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
