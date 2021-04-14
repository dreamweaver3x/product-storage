package models



type Product struct {
	ID int `db:"id"`
	Article string `json:"article" db:"article"`
	Name string `json:"name" db:"name"`
	Amount uint `json:"amount" db:"amount"`
	Address string `json:"address" db:"address"`
}
