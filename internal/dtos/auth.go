package dtos

type AuthControllerPayloadSignUp struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
	IsAdmin  bool   `json:"is_admin" validate:"required"`
}

type AuthControllerPayloadGetAccessToken struct {
	Email    string `json:"email"  validate:"required,email"`
	Password string `json:"password"  validate:"required,min=8"`
}
