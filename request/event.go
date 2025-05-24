package request

type CreateEventRequest struct {
	Url       string `json:"url" binding:"required,http_url"`
	Name      string `json:"name" binding:"omitempty,max=50"`
	ShowId    uint32 `json:"show_id" binding:"required,exists=shows id"`
	StartDate string `json:"start_date" binding:"omitempty,datetime=2006-01-02"`
	EndDate   string `json:"end_date" binding:"omitempty,datetime=2006-01-02"`
}
