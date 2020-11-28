package authenticator

type Email string
type Password string
type Token string
type VerificationCode string

type Base struct {
	Email    Email    `json:"email,omitempty"`
	Password Password `json:"pass,omitempty"`
	Token    Token    `json:"token,omitempty"`
}

type SignUpRequest struct{ Base }

type SignInRequest struct{ Base }

type VerificationRequest struct {
	Code VerificationCode `json:"code,omitempty"`
}

type ResetPassRequest struct {
	Email       Email            `json:"email,omitempty"`
	Code        VerificationCode `json:"code,omitempty"`
	NewPassword Password         `json:"newPassword,omitempty"`
}

type SignInResponse struct {
	Token        Token `json:"token,omitempty"`
	RefreshToken Token `json:"refreshToken,omitempty"`
	ExpiresIn    int64 `json:"expiresIn,omitempty"`
}

type SignOutRequest struct{ Base }
