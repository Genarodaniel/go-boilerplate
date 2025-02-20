package response

type DefaultErrorResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}
