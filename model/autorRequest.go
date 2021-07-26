package model

type AutorRequest struct {
	Nome      string `json:"nome" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Descricao string `json:"descricao" binding:"required"`
}
