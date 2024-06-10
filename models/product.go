package models

import (
	"gomarket/database"
)

type Product struct {
	Id, Quantity      int
	Name, Description string
	Price             float64
}

func GetAllProducts() []Product {
	db := database.ConnectDB()
	rows, err := db.Query("SELECT id, name, description, quantity, price FROM products")

	if err != nil {
		panic(err.Error())
	}

	product := Product{}
	var products []Product

	for rows.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = rows.Scan(&id, &name, &description, &quantity, &price)

		if err != nil {
			panic(err.Error())
		}

		product.Id = id
		product.Name = name
		product.Description = description
		product.Quantity = quantity
		product.Price = price

		products = append(products, product)
	}

	defer db.Close()
	return products
}

func CreateProduct(product Product) Product {
	db := database.ConnectDB()
	insert, err := db.Prepare("INSERT INTO products (name, description, quantity, price) VALUES ($1, $2, $3, $4)")

	if err != nil {
		panic(err.Error())
	}

	_, _ = insert.Exec(product.Name, product.Description, product.Quantity, product.Price)

	defer db.Close()

	return product
}

func DeleteProduct(id int) {
	db := database.ConnectDB()
	preparedProduct, err := db.Prepare("DELETE FROM products WHERE id = $1")

	if err != nil {
		panic(err.Error())
	}

	_, _ = preparedProduct.Exec(id)

	defer db.Close()
}
