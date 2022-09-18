package dto

type CreateDTO struct {
	Title string `json:"title" binding:"required,gte=8,lte=30"`
	Notes string `json:"notes" binding:"required"`
}

type UpdateDTO struct {
	ID    string `json:"id" binding:"required"`
	Title string `json:"title" binding:"omitempty,gte=8,lte=30"`
	Done  *bool  `json:"done" binding:"omitempty"`
}
