package dto

type SendNotificationRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	Token string `json:"token"`
}

type SendNotificationResponse struct {
	MessageID string `json:"message_id"`
}
