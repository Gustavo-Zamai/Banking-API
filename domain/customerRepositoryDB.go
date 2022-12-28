package domain

import (
	"database/sql"
	"log"
	"time"
	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepositoryDB struct {
	client *sql.DB
}

//open connection to database
func (d CustomerRepositoryDB) FindAll() ([]Customer, error) {
	//define sql to query into the database
	findAllSql := "select * from customers"
	rows, err := d.client.Query(findAllSql)
	if err != nil {
		log.Println("Error while querying customer table: " + err.Error())
		return nil, err
	}

	//scanning and append customer
	c := make([]Customer, 0)
	for rows.Next() {
		var customer Customer
		err := rows.Scan(&customer.Id, &customer.Name, &customer.City, &customer.DateOfBirth, &customer.Zipcode, &customer.Status)
		if err != nil {
			log.Println("Error while scanning customers table: " + err.Error())
		}
		c = append(c, customer)
	}
	return c, nil
}


func NewCustomerRepositoryDb() CustomerRepositoryDB {
	client, err := sql.Open("mysql", "root:Etecfg1c@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return CustomerRepositoryDB{client}
}