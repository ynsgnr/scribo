package authenticator

type Email string
type Password string
type Token string

type SignUpRequest struct {
	Email    Email    `json:"email"`
	Password Password `json:"pass"`
}

type SignInRequest struct {
	Email    Email    `json:"email"`
	Password Password `json:"pass"`
	Token    Token    `json:"token"`
}
