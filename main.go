package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host         = "localhost"
	port         = 5432
	databaseName = "mydatabase"
	username     = "myuser"
	password     = "mypassword"
)

var db *sql.DB

type Product struct {
	ID    int
	Name  string
	Price int
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, username, password, databaseName)

	sdb, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	db = sdb

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	print("Connection Database Successful!")

	err = deleteProduct(3)
	fmt.Println("delete successful")

}

func createProduct(product *Product) error {
	_, err := db.Exec(
		"INSERT INTO public.products(name, price) VALUES ($1, $2);",
		product.Name,
		product.Price,
	)

	return err
}

func getProduct(id int) (Product, error) {
	var p Product
	row := db.QueryRow(
		"SELECT id, name, price FROM products WHERE id=$1;",
		id,
	)

	err := row.Scan(&p.ID, &p.Name, &p.Price)

	if err != nil {
		log.Fatal(err)
		return Product{}, err
	}

	return p, nil
}

func updateProduct(id int, product *Product) (Product, error) {
	var p Product
	row := db.QueryRow(
		"UPDATE products SET name=$2, price=$3 WHERE id=$1 RETURNING id, name, price;",
		id,
		product.Name,
		product.Price,
	)

	err := row.Scan(&p.ID, &p.Name, &p.Price)
	if err != nil {
		return Product{}, err
	}

	return p, nil
}

func deleteProduct(id int) error {
	_, err := db.Exec(
		"DELETE FROM products WHERE id=$1;",
		id,
	)

	return err
}
