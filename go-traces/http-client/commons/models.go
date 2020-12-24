package commons

import "fmt"

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func (p *Product) String() string {
	return fmt.Sprintf("ID[%d], Name[%s], Price[%f]",
		p.ID, p.Name, p.Price)
}
