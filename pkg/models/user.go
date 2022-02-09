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

type Person struct {
	Tokenerr string
	Name     string
	Lname    string
	Email    string
}
