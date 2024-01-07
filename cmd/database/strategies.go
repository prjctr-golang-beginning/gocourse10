package database

var _ DatabaseConnection = &SqlDatabaseConnection{}
var _ DatabaseConnection = &NoSqlDatabaseConnection{}

// SqlDatabaseConnection конкретна реалізація для SQL бази даних.
type SqlDatabaseConnection struct{}

func (s *SqlDatabaseConnection) Open() string {
	return "SQL Database Connection Opened"
}

func (s *SqlDatabaseConnection) Close() string {
	return "SQL Database Connection Closed"
}

// NoSqlDatabaseConnection конкретна реалізація для NoSQL бази даних.
type NoSqlDatabaseConnection struct{}

func (n *NoSqlDatabaseConnection) Open() string {
	return "NoSQL Database Connection Opened"
}

func (n *NoSqlDatabaseConnection) Close() string {
	return "NoSQL Database Connection Closed"
}
