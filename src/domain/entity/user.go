package domain

type User interface {
	All() *[]User
}

func Hey() string {
	return "Hey"
}
