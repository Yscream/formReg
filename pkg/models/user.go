package models

type User struct {
	Name     string
	LastName string
	Email    string
	Password string
}

type LoginUser struct {
	Email    string
	Password string
}

type TypeOfErrors struct {
	FieldName  string
	MessageErr string
}

type Config struct {
	Username string
	Password string
	Host     string
	Port     int
	Dbname   string
	Sslmode  string
}
