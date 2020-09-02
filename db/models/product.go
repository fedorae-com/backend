package models

import (
	"encoding/json"
	"github.com/go-pg/pg"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/fedorae-com/backend/db"
	"github.com/gorilla/mux"
)

// type Product struct {
// 	tableName struct{} `sql:"products"`
// 	ID int `sql:"id,pk"`
// 	Name string `sql:"name"`
// 	Quantity int `sql:"quantity"`
// 	Price int `sql:"price"`
// 	Description string `sql:"description"`
// 	Images struct{
// 		Image1 string
// 		Image2 string
// 		Image3 string
// 	} `sql:"images,type:jsonb"`
// 	Status bool `sql:"status"`
// 	CreatedAt time.Time `sql:"created_at"`
// 	UpdatedAt time.Time `sql:"updated_at"`
// 	DeletedAt time.Time `sql:"deleted_at"`
// }

// Product struct (Model)
type Product struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Quantity  string  `json:"quantity"`
	Price  string  `json:"price"`
	Store *Store `json:"store"`
}

// Store struct
type Store struct {
	ID string `json:"id"`
	Name  string `json:"name"`
}

// Product Main func
func ProductInit()  {
	// DB Connection
	db.Connect()

	// Init router
	r := mux.NewRouter()

	// Hardcoded data - @todo: add database
	//products = append(products, Product{ID: "1", Name: "Product One", Quantity: "21", Price: "37.00", Store: &Store{ID: "1", Name: "Store One"}})
	//products = append(products, Product{ID: "2", Name: "Product Two", Quantity: "37", Price: "21.00", Store: &Store{ID: "2", Name: "Store Two"}})

	// Route handles & endpoints
	r.HandleFunc("/api/v1/products", getProducts).Methods("GET")
	r.HandleFunc("/api/v1/products/{id}", getProduct).Methods("GET")
	r.HandleFunc("/api/v1/products", createProduct).Methods("POST")
	r.HandleFunc("/api/v1/products/{id}", updateProduct).Methods("PUT")
	r.HandleFunc("/api/v1/products/{id}", deleteProduct).Methods("DELETE")

	// Start server
	log.Fatal(http.ListenAndServe(":8000", r))
}

// Init products var as a slice Product struct
var products []Product

// Get all products
func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//pg.Model(&products).Select()
	json.NewEncoder(w).Encode(pg.Model(&products).Select())
}

// Get single product
func getProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Gets params
	// Loop through products and find one with the id from the params
	for _, item := range products {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Product{})
}

// Add new product
func createProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var product Product
	_ = json.NewDecoder(r.Body).Decode(&product)
	product.ID = strconv.Itoa(rand.Intn(100000000)) // Mock ID - not safe
	products = append(products, product)
	json.NewEncoder(w).Encode(product)
}

// Update product
func updateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range products {
		if item.ID == params["id"] {
			products = append(products[:index], products[index+1:]...)
			var product Product
			_ = json.NewDecoder(r.Body).Decode(&product)
			product.ID = params["id"]
			products = append(products, product)
			json.NewEncoder(w).Encode(product)
			return
		}
	}
}

// Delete product
func deleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range products {
		if item.ID == params["id"] {
			products = append(products[:index], products[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(products)
}
