package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	httpDelivery "github.com/dodohq/backdo/delivery/http"
	adminRepo "github.com/dodohq/backdo/repository/admin"
	companyRepo "github.com/dodohq/backdo/repository/company"
	adminUsecase "github.com/dodohq/backdo/usecase/admin"
	companyUsecase "github.com/dodohq/backdo/usecase/company"
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

	router := httprouter.New()

	ar := adminRepo.NewAdminRepository(dbConn)
	au := adminUsecase.NewAdminUsecase(ar)
	cr := companyRepo.NewCompanyRepository(dbConn)
	cu := companyUsecase.NewCompanyUsecase(cr)

	httpDeliveryHandler := httpDelivery.Handler{Router: router}
	httpDeliveryHandler.InitAdminHandler(au).InitCompanyHandler(cu)

	whereToListen := ":" + os.Getenv("PORT")
	if isDevEnv {
		whereToListen = "localhost" + whereToListen
	}

	fmt.Println("Listening on", whereToListen)
	log.Fatal(http.ListenAndServe(whereToListen, router))
}
