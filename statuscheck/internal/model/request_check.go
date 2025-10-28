package model

type CheckRequest struct {
	URLs []string `json:"urls" binding:"required,dive,url"`
}