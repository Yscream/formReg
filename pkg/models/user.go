package models

type User struct {
	ID       int
	Name     string
	LastName string
	Email    string
	Password string
}

type Credentials struct {
	ID   int    `db:"users_id"`
	Salt string `db:"salt"`
	Hash string `db:"hash"`
}

type AccessToken struct {
	ID    int
	Token string
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
	Token string `db:"tokens"`
	ID    int    `db:"id"`
	Name  string `db:"fname"`
	Lname string `db:"lname"`
	Email string `db:"email"`
}
