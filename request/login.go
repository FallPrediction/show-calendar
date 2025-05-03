package request

type LoginRequest struct {
	Password string `json:"password" binding:"required,max=255,min=8"`
	Email    string `json:"email" binding:"required,email,max=255"`
}
