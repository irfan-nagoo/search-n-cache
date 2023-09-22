package response

type BaseResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
