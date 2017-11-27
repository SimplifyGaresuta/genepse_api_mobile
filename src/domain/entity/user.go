package entity

type User interface {
	Update() bool
}

func Hey() string {
	return "Hey"
}
