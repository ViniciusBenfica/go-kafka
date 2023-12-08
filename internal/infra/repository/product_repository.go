package repository

import (
	"database/sql"
	"mensageria/internal/entity"
)

type ProductRepository struct {
	DB *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

func (r *ProductRepository) Create(product *entity.Product) error {
	_, err := r.DB.Exec("insert into products (id, name, price) values(?,?,?)", product.ID, product.Name, product.Price)
	if err != nil {
		return err
	}
	return nil
}

func (r *ProductRepository) FindAll() ([]*entity.Product, error) {
	rows, err := r.DB.Query("select id, name, price from products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*entity.Product
	for rows.Next() {
		var product entity.Product
		err = rows.Scan(&product.ID, &product.Name, &product.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	return products, nil
}
