package domain

const (
	BUSINESS = 1 + iota
	ENGINEER
	DESIGNER
)

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
