package database

import (
	"database/sql"
)

func GetProducts(db *sql.DB, start, count int) ([]Product, error) {
	rows, err := db.Query(getProductsQuery, count, start)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	products := make([]Product,0)
	for rows.Next() {
		var prod Product
		if err := rows.Scan(&prod.ID, &prod.Name, &prod.Price); err != nil {
			return nil, err
		}
		products = append(products, prod)
	}

	return products, nil
}

func GetProduct(db *sql.DB, product *Product) error {
	return db.QueryRow(getProductQuery, product.ID).
		Scan(&product.Name, &product.Price)
}

func CreateProduct(db *sql.DB, product *Product) error {
	err := db.QueryRow(createProductQuery, product.Name, product.Price).Scan(&product.ID)
	if err != nil {
		return err
	}
	return nil
}

func UpdateProduct(db *sql.DB, product *Product) error {
	_, err := db.Exec(updateProductQuery, product.Name, product.Price, product.ID)
	return err
}

func DeleteProduct(db *sql.DB, product *Product) error {
	_, err := db.Exec(deleteProductQuery, product.ID)
	return err
}
