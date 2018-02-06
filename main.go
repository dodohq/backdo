package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"os"

	httpDelivery "github.com/dodohq/backdo/delivery/http"
	adminRepo "github.com/dodohq/backdo/repository/admin"
	companyRepo "github.com/dodohq/backdo/repository/company"
	deliveryRepo "github.com/dodohq/backdo/repository/delivery"
	driverRepo "github.com/dodohq/backdo/repository/driver"
	userRepo "github.com/dodohq/backdo/repository/user"
	adminUsecase "github.com/dodohq/backdo/usecase/admin"
	companyUsecase "github.com/dodohq/backdo/usecase/company"
	deliveryUsecase "github.com/dodohq/backdo/usecase/delivery"
	driverUsecase "github.com/dodohq/backdo/usecase/driver"
	userUsecase "github.com/dodohq/backdo/usecase/user"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
)

func main() {
	isDevEnv := os.Getenv("GO_ENV") == "development"
	if isDevEnv {
		if err := godotenv.Load(); err != nil {
			log.Fatal(err)
		}
	}

	dbConn, err := sql.Open("postgres", os.Getenv("DB"))
	if err != nil {
		log.Fatal(err)
	}
	defer dbConn.Close()

	auth := smtp.PlainAuth("", os.Getenv("DODO_EMAIL"), os.Getenv("DODO_EMAIL_PSSWD"), "smtp.gmail.com")

	router := httprouter.New()

	ar := adminRepo.NewAdminRepository(dbConn)
	au := adminUsecase.NewAdminUsecase(ar)
	cr := companyRepo.NewCompanyRepository(dbConn)
	dr := deliveryRepo.NewDeliveryRepository(dbConn)
	drr := driverRepo.NewDriverRepository(dbConn)
	cu := companyUsecase.NewCompanyUsecase(cr)
	ur := userRepo.NewUserRepo(dbConn, &auth)
	uu := userUsecase.NewUserUsecase(ur, cr)
	du := deliveryUsecase.NewDeliveryUsecase(dr, cr)
	dru := driverUsecase.NewDriverUsecase(drr, cr)

	httpDeliveryHandler := httpDelivery.Handler{Router: router}
	httpDeliveryHandler.InitAdminHandler(au).InitCompanyHandler(cu).InitUserHandler(uu).InitDeliveryHandler(du).InitDriverHandler(dru)

	whereToListen := ":" + os.Getenv("PORT")
	if isDevEnv {
		whereToListen = "localhost" + whereToListen
	}

	fmt.Println("Listening on", whereToListen)
	log.Fatal(http.ListenAndServe(whereToListen, router))
}
