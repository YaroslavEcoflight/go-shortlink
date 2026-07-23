package entity

type Token struct {
	AccessToken  string
	RefreshToken string
	ExpiresIn    int64
}
