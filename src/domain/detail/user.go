package detail

// User is 詳細画面に表示するユーザー
type User struct {
	ID           int
	Name         string
	AvatarURL    string
	Attribute    string
	Skills       []string
	Overview     string
	Awards       []string
	Products     []map[string]string
	Sns          []map[string]string
	License      []string
	Gender       string
	Age          int
	Address      string
	SchoolCareer string
}
