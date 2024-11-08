package model

import "fmt"

type Product struct {
	ID    int     `json:"id_product"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func (p *Product) ToString() string {
	return fmt.Sprintf("%+v", p)
}
