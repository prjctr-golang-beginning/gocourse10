package database

import (
	"fmt"
	"sync"
)

// DatabaseConnection інтерфейс, який визначає методи для взаємодії з базою даних.
type DatabaseConnection interface {
	Open() string
	Close() string
}

type Database interface {
	SetConnection(connection DatabaseConnection)
	ConnectToDatabase() string
	DisconnectFromDatabase() string
	Execute(cmd string) error
}

type DatabaseExecutor interface {
	Execute(cmd string) error
}

// MedicalDatabase є синглтоном, який використовується для управління медичними даними.
type MedicalDatabase struct {
	connection DatabaseConnection
}

func (c *MedicalDatabase) SetConnection(connection DatabaseConnection) {
	c.connection = connection
}

func (c *MedicalDatabase) ConnectToDatabase() string {
	return c.connection.Open()
}

func (c *MedicalDatabase) DisconnectFromDatabase() string {
	return c.connection.Close()
}

func (c *MedicalDatabase) Execute(cmd string) error {
	fmt.Printf("%s executed\n", cmd)

	return nil
}

var (
	instance Database
	once     sync.Once
)

// GetInstance повертає єдиний екземпляр MedicalDatabase. Is singleton an antipattern???
func GetInstance() Database {
	once.Do(func() {
		fmt.Println(`instance created`)
		instance = &MedicalDatabase{}
		// Тут можна ініціалізувати екземпляр додатковими даними, якщо потрібно.
	})

	return instance
}

func SetInstance(db Database) {
	instance = db
	fmt.Println(`instance created`)
}
