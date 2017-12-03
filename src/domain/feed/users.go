package feed

type Response struct {
	HasNext bool   `json:"has_next"`
	Users   []User `json:"users"`
}

type User struct {
	ID        uint     `json:"id"`
	Name      string   `json:"name"`
	AvatarURL string   `json:"avatar_url"`
	Skills    []string `json:"skills"`
	Overview  string   `json:"overview"`
}

func GetResponse() *Response {
	return &Response{}
}
