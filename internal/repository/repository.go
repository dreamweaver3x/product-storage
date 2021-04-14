package repository

import (
	"adverto/internal/models"
	"github.com/jmoiron/sqlx"

)

type LinksRepository struct {
	db *sqlx.DB
}

func NewLinksRepository(db *sqlx.DB) *LinksRepository {
	return &LinksRepository{db: db}
}


func (l *LinksRepository) Create(p models.Product) error {
	if _, err := l.db.NamedExec(`INSERT INTO products (article, name, amount, address) VALUES (:article, :name, :amount, :address)`, &p); err != nil {
		return err
	}
	return nil
}

func (l *LinksRepository) GetAll() ([]models.Product, error) {
	p := make([]models.Product, 0)
	if err := l.db.Select(&p,`SELECT * FROM products`); err != nil {
		return nil, err
	}
	return p, nil
}


func (l *LinksRepository) FindOne(p *models.Product) (*models.Product, error) {
	product := models.Product{}
	err := l.db.Get(&product, `SELECT * FROM products WHERE article = ?`, p.Article)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

