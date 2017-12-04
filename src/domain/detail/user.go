package detail

// User is 詳細画面に表示するユーザー
type User struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	AvatarURL    string    `json:"avatar_url"`
	Attribute    string    `json:"attribute"`
	Skills       []string  `json:"skills"`
	Overview     string    `json:"overview"`
	Awards       []string  `json:"awards"`
	Products     []Product `json:"products"`
	Sns          []Sns     `json:"sns"`
	License      []string  `json:"license"`
	Gender       string    `json:"gender"`
	Age          int       `json:"age"`
	Address      string    `json:"address"`
	SchoolCareer string    `json:"school_career"`
}

type Product struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

type Sns struct {
	Provider string `json:"provider"`
	URL      string `json:"url"`
}

func GetUser(id uint) (user *User, err error) {
	return &User{}, nil
}
