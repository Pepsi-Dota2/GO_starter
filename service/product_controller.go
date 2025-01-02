package service

import (
	"database/sql"
	"log"

	"github.com/pepsi/go-fiber/model"
)

var DB *sql.DB

func CreateProduct(product *model.Product) error {
	_, err := DB.Exec(
		"INSERT INTO public.products(name, price)VALUES ($1, $2);",
		product.Name,
		product.Price,
	)
	if err != nil {
		log.Printf("Failed to create product: %v\n", err)
		return err
	}
	return nil
}

func GetProductById(id int) (model.Product, error) {

	var product model.Product
	err := DB.QueryRow(
		"SELECT id, name, price FROM products WHERE id = $1;",
		id,
	).Scan(
		&product.ID,
		&product.Name,
		&product.Price,
	)
	if err != nil {
		log.Printf("Failed to get product: %v\n", err)
		return model.Product{}, err
	}
	return product, nil
}

func GetAllProducts() ([]model.Product, error) {
	rows, err := DB.Query("SELECT id, name, price FROM products;")
	if err != nil {
		log.Printf("Failed to get products: %v\n", err)
		return nil, err
	}
	var products []model.Product
	for rows.Next() {
		var p model.Product
		err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.Price,
		)
		if err != nil {
			log.Printf("Failed to scan product: %v\n", err)
			return nil, err
		}
		products = append(products, p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func UpdateProduct(id int, product *model.Product) (model.Product, error) {
	var p model.Product
	row := DB.QueryRow(
		"UPDATE products SET name=$1, price=$2 WHERE id=$3 RETURNING id, name , price;",
		product.Name,
		product.Price,
		id,
	)

	err := row.Scan(
		&p.ID,
		&p.Name,
		&p.Price,
	)
	if err != nil {
		log.Printf("Invalid request body: %v\n", err)
		return model.Product{}, err
	}
	return p, nil
}

func DeleteProduct(id int) error {
	_, err := DB.Exec(
		"DELETE FROM products WHERE id=$1;",
		id,
	)
	if err != nil {
		log.Printf("Invalid request body: %v\n", err)
		return err
	}
	return nil
}
