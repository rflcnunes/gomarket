package models

import "database/sql"

type Product struct {
	Id, Quantity      int
	Name, Description string
	Price             float64
}

func GetAllProducts(db *sql.DB) []Product {
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

	return products
}
