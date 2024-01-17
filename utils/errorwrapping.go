package utils

import (
	"errors"
	"fmt"
)

// What is error wrapping?
// Error wrapping is the process of encapsulating one error into another. Let’s say we have a web server which accesses a database and tries to fetch a record from the DB. If the database call returns an error, we can decide whether to wrap this error or send our own custom error from the webservice.

var noRows = errors.New("no rows found")

func getRecord() error {
	return noRows
}

func webService() error {
	if err := getRecord(); err != nil {
		return fmt.Errorf("Error %s when calling DB", err)
	}
	return nil
}

func ErrorWrappingExample() {
	if err := webService(); err != nil {
		if errors.Is(err, noRows) {
			fmt.Printf("The searched record cannot bef found. Error returned from DB is %s", err)
			return
		}
		fmt.Println("unknown error when searching record")
		return
	}
	fmt.Println("Web service called successfully")
}

type DBError struct {
	desc string
}

func (dbError DBError) Error() string {
	return dbError.desc
}

func getRecord1() error {
	return DBError{
		desc: "no rows found",
	}
}

func webService1() error {
	if err := getRecord1(); err != nil {
		return fmt.Errorf("Error %s when calling DB", err)
	}
	return nil
}

func AsFunction() {
	if err := webService1(); err != nil {
		var dbError DBError
		if errors.As(err, &dbError) {
			fmt.Printf("The searched record cannot be found. Error returned from DB is %s", dbError)
			return
		}
		fmt.Println("unknown error when searching record")
		return
	}
	fmt.Println("Web service called successfully")
}
