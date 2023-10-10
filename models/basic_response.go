package models

type BasicResponse struct {
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}
