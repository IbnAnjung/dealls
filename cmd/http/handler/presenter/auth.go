package presenter

type RegisterRequest struct {
	Username        string `json:"username"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
	Fullname        string `json:"fullname"`
	Gender          *uint8 `json:"gender"`
	Age             uint8  `json:"age"`
}

type RegisterResponse struct {
	ID           string `json:"id"`
	Fullname     string `json:"fullname"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	ID           string `json:"id"`
	Fullname     string `json:"fullname"`
	IsPremium    bool   `json:"is_premium"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
