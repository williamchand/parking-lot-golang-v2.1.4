package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/spf13/viper"

	_parkingLotHttpDelivery "github.com/williamchand/parking-lot-golang-v2.1.4/parking_lot/delivery/http"
	_parkingLotHttpDeliveryMiddleware "github.com/williamchand/parking-lot-golang-v2.1.4/parking_lot/delivery/http/middleware"
	_parkingLotRepo "github.com/williamchand/parking-lot-golang-v2.1.4/parking_lot/repository/mysql"
	_parkingLotUcase "github.com/williamchand/parking-lot-golang-v2.1.4/parking_lot/usecase"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	dbConn, err := sql.Open(`mysql`, dsn)

	if err != nil {
		log.Fatal(err)
	}
	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	e := echo.New()
	middL := _parkingLotHttpDeliveryMiddleware.InitMiddleware()
	e.Use(middL.CORS)
	ar := _parkingLotRepo.NewMysqlParkingLotRepository(dbConn)

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	au := _parkingLotUcase.NewParkingLotUsecase(ar, timeoutContext)
	_parkingLotHttpDelivery.NewParkingLotHandler(e, au)

	log.Fatal(e.Start(viper.GetString("server.address")))
}
