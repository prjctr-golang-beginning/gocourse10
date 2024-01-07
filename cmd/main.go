package main

import (
	"fmt"
	"mvc/cmd/database"
)

func main() {
	database.GetInstance().SetConnection(&database.SqlDatabaseConnection{})
	defer func() {
		fmt.Println(database.GetInstance().DisconnectFromDatabase())
	}()

	fmt.Println()
	createRecord()
	fmt.Println()
	updateRecord()
	fmt.Println()

	nObserver := &database.DatabaseActivityLogger{}

	ddb := database.NewObservableDecorator(database.GetInstance())
	ddb.RegisterObserver(nObserver)
	database.SetInstance(ddb)

	fmt.Println()
	createRecord()

	fmt.Println()
	deleteRecord()

	if odb, ok := database.GetInstance().(database.Observable); ok {
		odb.RemoveObserver(nObserver)
		odb.RegisterObserver(&database.DatabaseActivityMetric{})
	}

	fmt.Println()
	createRecord()

	// Тут можна використовувати db для роботи з медичними даними.
	fmt.Println("Singleton instance:", database.GetInstance())
}

func createRecord() {
	// Отримання екземпляра синглтона.
	_ = database.GetInstance().Execute(`Create`)
}

func updateRecord() {
	// Отримання екземпляра синглтона.
	_ = database.GetInstance().Execute(`Update`)
}

func deleteRecord() {
	// Отримання екземпляра синглтона.
	_ = database.GetInstance().Execute(`Delete`)
}
