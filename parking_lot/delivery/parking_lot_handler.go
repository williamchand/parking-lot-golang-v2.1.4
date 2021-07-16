package delivery

import (
	"net/http"
	"strconv"
	"fmt"
	"strings"
	"bytes"
	"io/ioutil"
	"github.com/spf13/viper"

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
	e.POST("/bulk", handler.Bulk)
	e.GET("/status", handler.FetchStatus)
	e.GET("/slot_number/car_registration_number/:registration_number", handler.GetIdByRegistrationNumber)
	e.GET("/cars_registration_numbers/colour/:colour", handler.FetchRegistrationNumber)
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
	response := fmt.Sprintf("Created a parking lot with %d slots\n", num)

	return c.String(http.StatusOK, response)
}

// OccupyParkingLot will occupy parking lot based on given params
func (a *ParkingLotHandler) OccupyParkingLot(c echo.Context) error {
	regNum := c.Param("registration_number")
	colour := c.Param("colour")
	ctx := c.Request().Context()

	slot_number, err := a.PUsecase.OccupyParkingLot(ctx, regNum, colour)
	if err == domain.ErrNotFound { 
		return c.String(getStatusCode(err),  "Sorry, parking lot is full\n")
	} else if err != nil {
		return c.JSON(getStatusCode(err),  err.Error())
	}
	response := fmt.Sprintf("Allocated slot number: %d\n", slot_number)

	return c.String(http.StatusOK, response)
}

// UnOccupyParkingLot will unoccupy parking lot based on given params
func (a *ParkingLotHandler) UnOccupyParkingLot(c echo.Context) error {
	numS := c.Param("id")
	num, _ := strconv.Atoi(numS)
	ctx := c.Request().Context()

	err := a.PUsecase.UnOccupyParkingLot(ctx, int64(num))
	if err == domain.ErrNotFound { 
		return c.String(getStatusCode(err),  "Not found\n")
	} else if err != nil {
		return c.JSON(getStatusCode(err),  err.Error())
	}
	response := fmt.Sprintf("Slot number %d is free\n", num)

	return c.String(http.StatusOK, response)
}

// FetchStatus will fetch the FetchStatus based on given params
func (a *ParkingLotHandler) FetchStatus(c echo.Context) error {
	ctx := c.Request().Context()

	listAr, err := a.PUsecase.FetchStatus(ctx)
	if err != nil {
		return c.JSON(getStatusCode(err),  err.Error())
	}

	response := fmt.Sprintf("Slot No. Registration No Colour\n")
	for _, element := range listAr {
		response = response + fmt.Sprintf("%d %s %s\n",element.ID,*element.RegistrationNumber,*element.Colour)
	}
	return c.String(http.StatusOK, response)
}

// FetchRegistrationNumber will fetch the FetchRegistrationNumber based on given params
func (a *ParkingLotHandler) FetchRegistrationNumber(c echo.Context) error {
	colour := c.Param("colour")
	ctx := c.Request().Context()

	listAr, err := a.PUsecase.FetchRegistrationNumber(ctx, colour)
	if err != nil {
		return c.JSON(getStatusCode(err),  err.Error())
	}

	return c.String(http.StatusOK, strings.Join(listAr, ", ") + "\n")
}

// FetchCarsSlot will fetch the FetchCarsSlot based on given params
func (a *ParkingLotHandler) FetchCarsSlot(c echo.Context) error {
	colour := c.Param("colour")
	ctx := c.Request().Context()

	listAr, err := a.PUsecase.FetchCarsSlot(ctx, colour)
	if err != nil {
		return c.JSON(getStatusCode(err),  err.Error())
	}

	return c.String(http.StatusOK, strings.Join(listAr, ", ") + "\n")
}
// GetByID will get article by given id
func (a *ParkingLotHandler) GetIdByRegistrationNumber(c echo.Context) error {
	regNum := c.Param("registration_number")
	ctx := c.Request().Context()

	registrationNumber, err := a.PUsecase.GetIdByRegistrationNumber(ctx, regNum)
	if err == domain.ErrNotFound { 
		return c.String(getStatusCode(err),  "Not found\n")
	} else if err != nil {
		return c.JSON(getStatusCode(err),  err.Error())
	}

	response := fmt.Sprintf("%d\n", registrationNumber)
	return c.String(http.StatusOK, response)
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

// Bulk will create bulk based on given params
func (a *ParkingLotHandler) Bulk(c echo.Context) error {
	request := c.Request().Body
	buf := new(bytes.Buffer)
	buf.ReadFrom(request)
	newStr := buf.String()
	temp := strings.Split(newStr,"\n")
	response := []string{}
	address :=  "http://localhost" + viper.GetString("server.address")
	for _, element := range temp {
		redirectString := strings.Split(element," ")
		resp := &http.Response{}
		switch redirectString[0] {
		case "create_parking_lot":
			resp, _ = http.Post(address+"/create_parking_lot/"+redirectString[1], "text/plain", nil)
		case "park":
			resp, _ = http.Post(address+"/park/"+redirectString[1]+ "/" + redirectString[2], "text/plain", nil)
		case "leave":
			resp, _ = http.Post(address+"/leave/"+redirectString[1], "text/plain", nil)
		case "status":
			resp, _ = http.Get(address+"/status")
		case "registration_numbers_for_cars_with_colour":
			resp, _ = http.Get(address+"/cars_registration_numbers/colour/"+redirectString[1])
		case "slot_numbers_for_cars_with_colour":
			resp, _ = http.Get(address+"/cars_slot/colour/"+redirectString[1])
		case "slot_number_for_registration_number":
			resp, _ = http.Get(address+"/slot_number/car_registration_number/"+redirectString[1])
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		response = append(response, string(body))
	}

	return c.String(http.StatusOK, strings.Join(response,""))
}
