package models

import (
	"github.com/go-pg/pg"
	"log"
	"time"
)

type Products struct {
	RefPointer int `sql:"-"`
	tableName struct{} `sql:"products"`
	ID int `sql:"id,pk"`
	Name string `sql:"name"`
	Quantity int `sql:"quantity"`
	Price int `sql:"price"`
	Description string `sql:"description"`
	Images struct{
		Image1 string
		Image2 string
		Image3 string
	} `sql:"images,type:jsonb"`
	Status bool `sql:"status"`
	CreatedAt time.Time `sql:"created_at"`
	UpdatedAt time.Time `sql:"updated_at"`
	DeletedAt time.Time `sql:"deleted_at"`
}

func (product *Products) GetProducts(db *pg.DB) error  {
	getErr := db.Model(product).Where("id = ?0", product.ID).Select()
	if getErr != nil {
		log.Printf("Error while geting value; %v\n", getErr)
		return getErr
	}
	log.Printf("Result; %v\n", *product)
	return nil
}