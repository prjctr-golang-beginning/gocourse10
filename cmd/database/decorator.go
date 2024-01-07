package database

// DatabaseConnectionDecorator декоратор для DatabaseConnection.
type DatabaseConnectionDecorator struct {
	wrapped Database
}

func NewObservableDecorator(connection Database) *ObservableDecorator {
	return &ObservableDecorator{
		Observable: &DefaultObserver{},
		wrapped:    connection,
	}
}

// ObservableDecorator є декоратором, який додає функціональність логування.
type ObservableDecorator struct {
	Observable
	wrapped Database
}

func (c *ObservableDecorator) ConnectToDatabase() string {
	c.NotifyObservers("Connection established")
	return c.wrapped.ConnectToDatabase()
}

func (c *ObservableDecorator) DisconnectFromDatabase() string {
	c.NotifyObservers("Connection cancelled")
	return c.wrapped.DisconnectFromDatabase()
}

func (c *ObservableDecorator) Execute(cmd string) error {
	c.NotifyObservers("Query executed")
	return c.wrapped.Execute(cmd)
}

func (c *ObservableDecorator) SetConnection(connection DatabaseConnection) {
	c.NotifyObservers("Connection set")
	c.wrapped.SetConnection(connection)
}
