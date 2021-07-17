package delivery_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/bxcodec/faker"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	parkingLotHttp "github.com/williamchand/parking-lot-golang-v2.1.4/parking_lot/delivery"
	"github.com/williamchand/parking-lot-golang-v2.1.4/parking_lot/domain"
	"github.com/williamchand/parking-lot-golang-v2.1.4/parking_lot/domain/mocks"
)

// Custom type that allows setting the func that our Mock Do func will run instead
type MockDoType func(req *http.Request) (*http.Response, error)

// MockClient is the mock client
type MockClient struct {
	MockDo MockDoType
}

// Overriding what the Do function should "do" in our MockClient
func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	return m.MockDo(req)
}

func TestFetchStatus(t *testing.T) {
	var mockParkingLot domain.ParkingLot
	err := faker.FakeData(&mockParkingLot)
	assert.NoError(t, err)
	mockPCase := new(mocks.ParkingLotUsecase)
	mockListParkingLot := make([]domain.ParkingLot, 0)
	mockListParkingLot = append(mockListParkingLot, mockParkingLot)
	mockPCase.On("FetchStatus", mock.Anything).Return(mockListParkingLot, nil)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/status", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := parkingLotHttp.ParkingLotHandler{
		PUsecase: mockPCase,
	}
	err = handler.FetchStatus(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockPCase.AssertExpectations(t)
}

func TestFetchRegistrationNumber(t *testing.T) {
	var mockParkingLot string
	err := faker.FakeData(&mockParkingLot)
	assert.NoError(t, err)
	mockPCase := new(mocks.ParkingLotUsecase)
	mockListParkingLot := make([]string, 0)
	mockListParkingLot = append(mockListParkingLot, mockParkingLot)
	mockPCase.On("FetchRegistrationNumber", mock.Anything, mock.Anything).Return(mockListParkingLot, nil)
	colour := "Black"
	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/cars_registration_numbers/colour/"+colour, strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("cars_registration_numbers/colour/:colour")
	c.SetParamNames("colour")
	c.SetParamValues(colour)
	handler := parkingLotHttp.ParkingLotHandler{
		PUsecase: mockPCase,
	}
	err = handler.FetchRegistrationNumber(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockPCase.AssertExpectations(t)
}

func TestFetchCarsSlot(t *testing.T) {
	var mockParkingLot string
	err := faker.FakeData(&mockParkingLot)
	assert.NoError(t, err)
	mockPCase := new(mocks.ParkingLotUsecase)
	mockListParkingLot := make([]string, 0)
	mockListParkingLot = append(mockListParkingLot, mockParkingLot)
	mockPCase.On("FetchCarsSlot", mock.Anything, mock.Anything).Return(mockListParkingLot, nil)
	colour := "Black"
	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/cars_slot/colour/"+colour, strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("cars_slot/colour/:colour")
	c.SetParamNames("colour")
	c.SetParamValues(colour)
	handler := parkingLotHttp.ParkingLotHandler{
		PUsecase: mockPCase,
	}
	err = handler.FetchCarsSlot(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockPCase.AssertExpectations(t)
}

func TestGetIdByRegistrationNumberSuccess(t *testing.T) {
	mockPCase := new(mocks.ParkingLotUsecase)
	num := int64(2)
	registrationNumber := "B-1234-RFS"
	mockPCase.On("GetIdByRegistrationNumber", mock.Anything, mock.AnythingOfType("string")).Return(num, nil)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/slot_number/car_registration_number/"+registrationNumber, strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("slot_number/car_registration_number/:registration_number")
	c.SetParamNames("registration_number")
	c.SetParamValues(registrationNumber)
	handler := parkingLotHttp.ParkingLotHandler{
		PUsecase: mockPCase,
	}
	err = handler.GetIdByRegistrationNumber(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockPCase.AssertExpectations(t)
}

func TestGetIdByRegistrationNumberFailed(t *testing.T) {
	mockPCase := new(mocks.ParkingLotUsecase)
	num := int64(0)
	registrationNumber := "B-1234-RFS"
	mockPCase.On("GetIdByRegistrationNumber", mock.Anything, mock.AnythingOfType("string")).Return(num, domain.ErrNotFound)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/slot_number/car_registration_number/"+registrationNumber, strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("slot_number/car_registration_number/:registration_number")
	c.SetParamNames("registration_number")
	c.SetParamValues(registrationNumber)
	handler := parkingLotHttp.ParkingLotHandler{
		PUsecase: mockPCase,
	}
	err = handler.GetIdByRegistrationNumber(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusNotFound, rec.Code)
	mockPCase.AssertExpectations(t)
}

func TestCreateParkingLot(t *testing.T) {
	mockPCase := new(mocks.ParkingLotUsecase)
	mockPCase.On("CreateParkingLot", mock.Anything, mock.AnythingOfType("int64")).Return(nil)

	e := echo.New()
	num := int64(6)
	req, err := http.NewRequest(echo.POST, "/create_parking_lot/"+fmt.Sprintf("%d", num), strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("create_parking_lot/:parking_lot")
	c.SetParamNames("parking_lot")
	c.SetParamValues(fmt.Sprintf("%d", num))

	handler := parkingLotHttp.ParkingLotHandler{
		PUsecase: mockPCase,
	}
	err = handler.CreateParkingLot(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockPCase.AssertExpectations(t)
}

func TestOccupyParkingLotSuccess(t *testing.T) {
	mockPCase := new(mocks.ParkingLotUsecase)
	mockPCase.On("OccupyParkingLot", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(int64(2), nil)
	registrationNumber := "B-1234-RFS"
	colour := "Black"

	e := echo.New()
	req, err := http.NewRequest(echo.POST, "/park/"+registrationNumber+"/"+colour, strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("park/:registration_number/:colour")
	c.SetParamNames("registration_number", "colour")
	c.SetParamValues(registrationNumber, colour)

	handler := parkingLotHttp.ParkingLotHandler{
		PUsecase: mockPCase,
	}
	err = handler.OccupyParkingLot(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockPCase.AssertExpectations(t)
}

func TestOccupyParkingLotFailed(t *testing.T) {
	mockPCase := new(mocks.ParkingLotUsecase)
	mockPCase.On("OccupyParkingLot", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(int64(0), domain.ErrNotFound)
	registrationNumber := "B-1234-RFS"
	colour := "Black"

	e := echo.New()
	req, err := http.NewRequest(echo.POST, "/park/"+registrationNumber+"/"+colour, strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("park/:registration_number/:colour")
	c.SetParamNames("registration_number", "colour")
	c.SetParamValues(registrationNumber, colour)

	handler := parkingLotHttp.ParkingLotHandler{
		PUsecase: mockPCase,
	}
	err = handler.OccupyParkingLot(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusNotFound, rec.Code)
	mockPCase.AssertExpectations(t)
}

func TestUnOccupyParkingLotSuccess(t *testing.T) {
	mockPCase := new(mocks.ParkingLotUsecase)
	mockPCase.On("UnOccupyParkingLot", mock.Anything, mock.AnythingOfType("int64")).Return(nil)
	num := int64(2)

	e := echo.New()
	req, err := http.NewRequest(echo.POST, "/leave/"+fmt.Sprintf("%d", num), strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("leave/:id")
	c.SetParamNames("id")
	c.SetParamValues(fmt.Sprintf("%d", num))

	handler := parkingLotHttp.ParkingLotHandler{
		PUsecase: mockPCase,
	}
	err = handler.UnOccupyParkingLot(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockPCase.AssertExpectations(t)
}

func TestUnOccupyParkingLotFailed(t *testing.T) {
	mockPCase := new(mocks.ParkingLotUsecase)
	mockPCase.On("UnOccupyParkingLot", mock.Anything, mock.AnythingOfType("int64")).Return(domain.ErrNotFound)
	num := int64(2)

	e := echo.New()
	req, err := http.NewRequest(echo.POST, "/leave/"+fmt.Sprintf("%d", num), strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("leave/:id")
	c.SetParamNames("id")
	c.SetParamValues(fmt.Sprintf("%d", num))

	handler := parkingLotHttp.ParkingLotHandler{
		PUsecase: mockPCase,
	}
	err = handler.UnOccupyParkingLot(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusNotFound, rec.Code)
	mockPCase.AssertExpectations(t)
}

func TestBulk(t *testing.T) {
	// build our response JSON
	jsonResponse := `Test Success`

	// create a new reader with that JSON
	r := ioutil.NopCloser(bytes.NewReader([]byte(jsonResponse)))

	Client := &MockClient{
		MockDo: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       r,
			}, nil
		},
	}
	mockPCase := new(mocks.ParkingLotUsecase)
	e := echo.New()
	req, err := http.NewRequest(echo.POST, "/bulk", strings.NewReader("create_parking_lot 6"))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("bulk")

	handler := parkingLotHttp.ParkingLotHandler{
		PUsecase: mockPCase,
		Client:   Client,
	}
	err = handler.Bulk(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockPCase.AssertExpectations(t)
}
