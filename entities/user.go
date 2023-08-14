package entities

type User struct {
	Id       uint
	Username string
	Email    string
	Password string
}

type UserInput struct {
	Username string
	Password string
}
