package delivery

type BasicResponse struct {
	Code    StatusCode `json:"code"`
	Message string     `json:"message"`
	Error   string     `json:"error"`
}
