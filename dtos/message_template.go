package dtos

type ToastResponse struct {
	Code      string `json:"code"`
	Message   string `json:"message"`
	ToastType string `json:"toastType"`
}
