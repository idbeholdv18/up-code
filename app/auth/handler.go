package auth

type RegisterRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	FullName string `json:"full_name"`
}

type RegisterResponsePayload struct {
	Email    string `json:"email"`
	FullName string `json:"full_name"`
}

type LoginRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponsePayload struct {
	AccessToken string `json:"access_token"`
}
