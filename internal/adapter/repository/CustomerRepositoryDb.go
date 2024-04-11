package repository

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sugandhasaxena1911/MyPracticeBankingApp/helpers/error"
	"github.com/sugandhasaxena1911/MyPracticeBankingApp/helpers/logger"
	"github.com/sugandhasaxena1911/MyPracticeBankingApp/internal/core/domain"
)

type CustomerRepositoryDB struct {
	DBClient *sql.DB
}

func (custrepoDB CustomerRepositoryDB) GetAllCustomers() ([]domain.Customers, *error.AppError) {
	db := custrepoDB.DBClient
	sqlst := "select customer_id, name , date_of_birth, city , zipcode , status from customers"
	rows, err := db.Query(sqlst)
	if err != nil {
		log.Println("the error in DB while querying customer  ", err)
		return nil, error.NewInternalServerAppError("Internal server error")
	}
	customers := []domain.Customers{}
	for rows.Next() {
		var c domain.Customers
		err = rows.Scan(&c.Custid, &c.Custname, &c.Custbirthdate, &c.Custcity, &c.Custzipcode, &c.Custstatus)
		if err != nil {
			log.Println("Error in scanning rows ", err)
			return nil, error.NewInternalServerAppError("Internal server error ")
		}
		customers = append(customers, c)
	}

	return customers, nil

}

func (custrepoDB CustomerRepositoryDB) GetAllCustomersByStatus(status string) ([]domain.Customers, *error.AppError) {
	sqlst := "select customer_id, name , date_of_birth, city , zipcode , status from customers where status = ?"
	db := custrepoDB.DBClient
	rows, err := db.Query(sqlst, status)
	if err != nil {
		logger.Info("Error while fetching customers")
		return nil, error.NewInternalServerAppError("Database Internal server error ")
	}
	var c domain.Customers
	custs := []domain.Customers{}

	for rows.Next() {

		err := rows.Scan(&c.Custid, &c.Custname, &c.Custbirthdate, &c.Custcity, &c.Custzipcode, &c.Custstatus)
		if err != nil {
			return nil, error.NewInternalServerAppError("Internal db error")

		}
		custs = append(custs, c)
	}

	return custs, nil
}

func (custrepoDB CustomerRepositoryDB) GetCustomerByID(custid string) (domain.Customers, *error.AppError) {
	log.Println("INSSIDE")
	db := custrepoDB.DBClient

	sqlst := "select customer_id, name , date_of_birth, city , zipcode , status from customers where customer_id =?"
	row := db.QueryRow(sqlst, custid)

	var c domain.Customers
	err := row.Scan(&c.Custid, &c.Custname, &c.Custbirthdate, &c.Custcity, &c.Custzipcode, &c.Custstatus)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("Customer do not exists")
			return c, error.NewNotFoundAppError("Customer not found")
		}
		log.Println("Error while scanning customer ", err)
		return c, error.NewInternalServerAppError("Internal server error")
	}
	return c, nil
}

func (custrepoDB CustomerRepositoryDB) PostCustomer(c domain.Customers) (domain.Customers, *error.AppError) {

	sqlst := "insert into customers( name , date_of_birth, city , zipcode , status) values (?,?,?,?,? )"
	//var customer domain.Customers
	result, err := custrepoDB.DBClient.Exec(sqlst, &c.Custname, &c.Custbirthdate, &c.Custcity, &c.Custzipcode, &c.Custstatus)
	if err != nil {
		logger.Error(fmt.Sprintf("db error in inserting customer %s", err))
		return c, error.NewInternalServerAppError("Error while insertion")

	}
	id, err := result.LastInsertId()
	if err != nil {
		logger.Error(fmt.Sprintf("db error in inserting customer %s ", err))
		return c, error.NewInternalServerAppError("Error while fetching customer id")
	}
	c.Custid = strconv.Itoa(int(id))
	return c, nil

}

func NewCustomerReposioryDB(dbconn *sql.DB) CustomerRepositoryDB {
	//create DB connection

	return CustomerRepositoryDB{DBClient: dbconn}
}
