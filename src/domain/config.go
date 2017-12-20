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
		return ""
	}
}

func GetAttributeID(attribute string) int {
	switch attribute {
	case "Business":
		return BUSINESS
	case "Engineer":
		return ENGINEER
	case "Designer":
		return DESIGNER
	default:
		return 0
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
	case 0:
		return ""
	default:
		return "その他"
	}
}

func GetGenderID(gender string) int {
	switch gender {
	case "男性":
		return MAN
	case "女性":
		return WOMAN
	default:
		return 3
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
