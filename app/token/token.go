package token

type AccessTokenClaims struct {
	UserID int64 `json:"user_id"`
}

type ConfirmationTokenClaims struct {
	UserID int64  `json:"user_id"`
	Email  string `json:"email"`
}
