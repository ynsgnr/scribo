package authenticator

type Email string
type Password string
type Token string

type Base struct {
	Email    Email    `json:"email,omitempty"`
	Password Password `json:"pass,omitempty"`
	Token    Token    `json:"token,omitempty"`
}

type SignUpRequest struct{ Base }

type SignInRequest struct{ Base }

type SignInResponse struct{ Base }

type SignOutRequest struct{ Base }
