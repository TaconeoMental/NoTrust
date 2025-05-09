package models

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type OTPRequest struct {
	OTP string `json:"otp"`
}

type SendOTPRequest struct {
	Email string `json:"email"`
}

