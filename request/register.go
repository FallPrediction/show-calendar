package request

type RegisterRequest struct {
	Name            string `json:"name" binding:"required,max=255"`
	Password        string `json:"password" binding:"required,max=255,min=8"`
	ConfirmPassword string `json:"confirmPassword" binding:"required,eqfield=Password"`
	Email           string `json:"email" binding:"required,email,max=255,unique=users email"`
}
