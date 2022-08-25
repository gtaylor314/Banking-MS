package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gtaylor314/Banking-MS/domain"
	"github.com/gtaylor314/Banking-MS/service"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

// environmentVariableCheck() checks that all environment variables have been injected and, if not, exits with an error
func environmentVariableCheck() {
	switch {
	case os.Getenv("SERVER_ADDRESS") == "":
		log.Fatal("environment variable SERVER_ADDRESS not defined")
	case os.Getenv("SERVER_PORT") == "":
		log.Fatal("environment variable SERVER_PORT not defined")
	case os.Getenv("DB_USER") == "":
		log.Fatal("environment variable DB_USER not defined")
	case os.Getenv("DB_PASSWD") == "":
		log.Fatal("environment variable DB_PASSWD not defined")
	case os.Getenv("DB_ADDRESS") == "":
		log.Fatal("environment variable DB_ADDRESS not defined")
	case os.Getenv("DB_PORT") == "":
		log.Fatal("environment variable DB_PORT not defined")
	case os.Getenv("DB_NAME") == "":
		log.Fatal("environment variable DB_NAME not defined")
	}
}

func Start() {

	environmentVariableCheck()

	// using gorilla/mux package - creating a router instance
	// gorilla/mux is idiomatic - uses expressions native to standard HTTP library (like HandleFunc())
	router := mux.NewRouter()

	// create db connection
	db_conn := getDbConnection()
	// create customer repository
	customerRepositoryDb := domain.NewCustomerRepositoryDb(db_conn)
	// create account repository
	accountRepositoryDb := domain.NewAccountRepositoryDb(db_conn)
	// create transaction repository
	transactionRepositoryDb := domain.NewTransactionRepositoryDb(db_conn)

	// this custHandler uses NewCustomerRepositoryStub() which creates a predefined set of customers to test with
	// custHandler := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}

	// this custHandler uses customerRepositoryDb (defined above) which connects to the MySQL database
	custHandler := CustomerHandlers{service: service.NewCustomerService(customerRepositoryDb)}
	// this acctHandler uses accountRepositoryDb (defined above) which connects to the MySQL database
	acctHandler := AccountHandler{service: service.NewAccountService(accountRepositoryDb)}
	// this tranHandler uses transactionRepositoryDb (defined above) which connects to the MySQL database
	tranHandler := TransactionHandler{service: service.NewTransactionService(transactionRepositoryDb)}

	// registering the handler functions for the given patterns (routes)
	router.HandleFunc("/customers", custHandler.getAllCustomers).Methods(http.MethodGet)
	// here we use a variable (customer_id) in the path and use a regular expression (:[0-9]+) to indicate that the customer_id
	// will only be comprised of numerical digits 0 - 9 which can repeat (+)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", custHandler.getCustomer).Methods(http.MethodGet)
	// handler for creating an account - customer_id is required as accounts can only be created by existing customers
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account", acctHandler.newAccount).Methods(http.MethodPost)
	// handler for creating a transaction - customer_id is required as transactions can only be created by existing customers
	router.HandleFunc("/customers/{customer_id:[0-9]+}/transaction", tranHandler.newTransaction).Methods(http.MethodPost)

	// we use environment variables to inject the server address and port (see Makefile)
	serverAdd := os.Getenv("SERVER_ADDRESS")
	serverPort := os.Getenv("SERVER_PORT")
	// ListenAndServe listens on the TCP network address specified and then calls Serve with the handler to handle
	// incoming requests - if we were using the DefaultServeMux, we would pass nil for the handler
	// ListenAndServe returns an error, so we wrap it with log.Fatal so we will know if an error occurs during server startup
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", serverAdd, serverPort), router))
}

func getDbConnection() *sqlx.DB {
	// we use environment variables to inject the database connection information (see Makefile)
	dbUser := os.Getenv("DB_USER")
	dbPasswd := os.Getenv("DB_PASSWD")
	dbAddress := os.Getenv("DB_ADDRESS")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPasswd, dbAddress, dbPort, dbName)

	db, err := sqlx.Open("mysql", dbSource)
	if err != nil {
		panic(err)
	}

	// SetConnMaxLifetime sets the maximum amount of time that a connection may be reused
	db.SetConnMaxLifetime(time.Minute * 3)
	// SetMaxOpenConns sets the maximum number of open connections to the database
	db.SetMaxOpenConns(10)
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool
	db.SetMaxIdleConns(10)

	return db
}
