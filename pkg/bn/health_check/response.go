package healthcheck

type HealthCheckResponse struct {
	Message string `json:"message"`
}

func NewHealthCheckResponseWith(message string) *HealthCheckResponse {
	if message == "" {
		message = "OK"
	}
	response := &HealthCheckResponse{Message: message}
	return response
}
