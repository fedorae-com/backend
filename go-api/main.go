package main

import (
	"fmt"
	"github.com/fedorae-com/go-api/db"
	"github.com/go-pg/pg"
)

func main()  {
	fmt.Println("Main Project")
	db.Connect()
}

func GetById(dbRef *pg.DB)  {
	newProduct := &db.Product {
		ID: 1
	}
	newProduct.GetById
}