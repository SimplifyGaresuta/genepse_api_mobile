package middleware

type Login struct {
	LoginURL string `json:"login_url"`
}

type Callback struct {
	UserID uint `json:"user_id"`
}
