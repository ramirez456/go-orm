package main

import (
	"fmt"

	"github.com/ramirez456/go-orm/model"
	"github.com/ramirez456/go-orm/storage"
)

func main() {
	fmt.Println("hi, friend")
	drive := storage.MySQL
	storage.New(drive)
	//storage.DB().AutoMigrate(&model.Product{}, &model.InvoiceHeader{}, &model.InvoiceItem{})
	//storage.DB().Model(&model.InvoiceItem{})
	//storage.DB().Model(&model.Product{}).DisableForeignKeyConstraintWhenMigrating()
	// m := &model.Product{
	// 	Name:  "curso GO Avanzado",
	// 	Price: 90,
	// }
	//create
	//storage.DB().Create(m)
	//Get All
	// products := make([]model.Product, 0)
	// storage.DB().Find(&products)
	// for _, product := range products {
	// 	fmt.Printf("%d - %s\n", product.ID, product.Name)
	// }
	//GetByID
	/*
		product := model.Product{}
		storage.DB().First(&product, 1)
		fmt.Printf("%d - %s\n", product.ID, product.Name)
	*/
	//update
	// product := model.Product{}
	// product.ID = 2
	//storage.DB().First(&product, 1)
	// storage.DB().Model(&product).Updates(
	// 	model.Product{
	// 		Name:  "Como conseguir novia",
	// 		Price: 200,
	// 	},
	// )
	//delete
	//storage.DB().Delete(&product)
	//storage.DB().Unscoped().Delete(&product)

	//transaction
	invoice := model.InvoiceHeader{
		Client: "Max Houston Ramirez Martel",
		InvoiceItems: []model.InvoiceItem{
			{ProductID: 1},
			{ProductID: 3},
			{ProductID: 1},
		},
	}
	storage.DB().Create(&invoice)
}
