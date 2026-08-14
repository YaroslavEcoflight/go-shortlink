package restapi

type UserRegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLogoutRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type UserRefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type UserValidateTokenRequest struct {
	AccessToken string `json:"access_token"`
}
