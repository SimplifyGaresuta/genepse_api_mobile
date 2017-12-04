package domain

const (
	BUSINESS = 1 + iota
	ENGINEER
	DESIGNER
)

func GetAttribute(id int) string {
	switch id {
	case BUSINESS:
		return "Business"
	case ENGINEER:
		return "Engineer"
	case DESIGNER:
		return "Designer"
	default:
		return "Unknown"
	}
}

const (
	MAN = 1 + iota
	WOMAN
)

func GetGender(id int) string {
	switch id {
	case MAN:
		return "男性"
	case WOMAN:
		return "女性"
	default:
		return "その他"
	}
}

/*
type AttributeID int

func (c AttributeID) String() string {
	switch c {
	case BUSINESS:
		return "Business"
	case ENGINEER:
		return "ENgineer"
	case DESIGNER:
		return "Designer"
	default:
		return "Unknown"
	}
}
*/
