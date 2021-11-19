package rabbitConnection

import "gorm.io/gorm"

type Pessoa struct {
	gorm.Model
	Nome             string `json:"nome"`
	Cpf              int    `json:"cpf"`
	DataDeNascimento string `json:"dataDeNascimento"`
}
