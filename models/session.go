package models

type SessionInput struct {
	Key   string `json:"key" binding:"required"`
	Value string `json:"value" binding:"required"`
}
