package request

type CreateShowRequest struct {
	Name      string `json:"name" binding:"required,max=255"`
	TicketUrl string `json:"ticket_url" binding:"required,http_url"`
	StartDate string `json:"start_date" binding:"datetime=2006-01-02"`
	EndDate   string `json:"end_date" binding:"datetime=2006-01-02"`
}
