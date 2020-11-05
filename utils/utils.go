package utils

type StandardResponse struct {
	Success      bool        `json:"success"`
	ErrorMessage error       `json:"error_message"`
	Data         interface{} `json:"data"`
}
