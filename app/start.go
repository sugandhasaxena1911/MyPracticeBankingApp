package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
<<<<<<< HEAD
	"github.com/sugandhasaxena1911/MyPracticeBankingApp/helpers/logger"
	handlers "github.com/sugandhasaxena1911/MyPracticeBankingApp/internal/adapter/handlers"
	repository "github.com/sugandhasaxena1911/MyPracticeBankingApp/internal/adapter/repository"
	"github.com/sugandhasaxena1911/MyPracticeBankingApp/internal/core/service"
=======
	"github.com/sugandhasaxena19/MyPracticeBankingApp/helpers/logger"
	handlers "github.com/sugandhasaxena19/MyPracticeBankingApp/internal/adapter/handlers"
	repository "github.com/sugandhasaxena19/MyPracticeBankingApp/internal/adapter/repository"
	"github.com/sugandhasaxena19/MyPracticeBankingApp/internal/core/service"
>>>>>>> 27c2bab4ae9973b95478eedee3812fd6c4b17ae0
)

var DBClient *sql.DB

func init() {
	// Load env variables
	if err := godotenv.Load(); err != nil {
		log.Fatalln("Cannot load env variables ")
	}
	// sanity check

	if os.Getenv("DBNAME") == "" || os.Getenv("DBHOST") == "" || os.Getenv("DBPASSWORD") == "" || os.Getenv("DBUSER") == "" || os.Getenv("DBPORT") == "" {
		logger.Error("Env variables are not defined")
	}

}

func init() {
	// Take Db connection
	fmt.Println("I AM HERE")
	dbuser := os.Getenv("DBUSER")
	dbpassword := os.Getenv("DBPASSWORD")
	dbhost := os.Getenv("DBHOST")
	dbport := os.Getenv("DBPORT")
	dbname := os.Getenv("DBNAME")
	dbconn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbuser, dbpassword, dbhost, dbport, dbname)
	logger.Info(fmt.Sprintln("Connection string ", dbconn))
	var err error
	DBClient, err = sql.Open("mysql", dbconn)
	if err != nil {
		logger.Error("Cannot get DB connection")
		log.Panicln("Error in DB connection ", err)
	}
	err = DBClient.Ping()
	if err != nil {
		log.Panicln("Cannot ping DB", err)
	}
	logger.Info("Db ping successful ")
	log.Println(DBClient)
	DBClient.SetConnMaxLifetime(time.Minute * 3)
	DBClient.SetMaxOpenConns(10)
	DBClient.SetMaxIdleConns(10)
}

func Start() {

	log.Println("Defining Routers")
	log.Println(DBClient)
	router := mux.NewRouter()
	// create handler
	//cservice := service.NewCustomerServiceDefault(repository.NewCustomerReposioryStub())
	cservice := service.NewCustomerServiceDefault(repository.NewCustomerReposioryDB(DBClient))
	cushandler := handlers.Custhandlers{Custservice: cservice}

	router.HandleFunc("/customers", cushandler.GetAllCustomers).Methods(http.MethodGet)             // get all customers
	router.HandleFunc("/customers/{id:[0-9]+}", cushandler.GetCustomerByID).Methods(http.MethodGet) // get customer by id
	router.HandleFunc("/customers", cushandler.PostCustomer).Methods(http.MethodPost)               // post customer
	// get all customers by status
	serverhost := os.Getenv("SERVERHOST")
	serverport := os.Getenv("SERVERPORT")

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", serverhost, serverport), router))

}
