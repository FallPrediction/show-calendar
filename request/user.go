package request

type UserLikeShowRequest struct {
	ShowId uint32 `json:"show_id" binding:"required,exists=shows id"`
}
