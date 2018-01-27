package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	httpDelivery "github.com/dodohq/backdo/delivery/http"
	adminRepo "github.com/dodohq/backdo/repository/admin"
	adminUsecase "github.com/dodohq/backdo/usecase/admin"
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
	httpDeliveryHandler := httpDelivery.Handler{Router: router}
	httpDeliveryHandler.InitAdminHandler(au)

	whereToListen := ":8080"
	if isDevEnv {
		whereToListen = "localhost" + whereToListen
	}
	log.Fatal(http.ListenAndServe(whereToListen, router))
}
