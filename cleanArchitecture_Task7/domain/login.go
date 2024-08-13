package domain


type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Message string `json:"message"`
	Token string `json:"token"`
}

type LoginUsecase interface {
	CreateToken(user *User , secret string , expiry int) (string, error)	
}