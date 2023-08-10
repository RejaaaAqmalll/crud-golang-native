package main

import (
	"CRUD-golang/config"
	categoriesController "CRUD-golang/controllers/categoriesController"
	"CRUD-golang/controllers/homeController"
	productcontroller "CRUD-golang/controllers/productController"
	"log"
	"net/http"
)

func main() {
	// memanggil variabel database dari folder config
	config.KoneksiDB()

	// Memanggil homepage
	http.HandleFunc("/", homeController.Welcome)

	// Memanggil Category
	http.HandleFunc("/categories", categoriesController.Index)
	http.HandleFunc("/categories/add", categoriesController.Add)
	http.HandleFunc("/categories/edit", categoriesController.Edit)
	http.HandleFunc("/categories/delete", categoriesController.Delete)

	// memanggil Product
	http.HandleFunc("/products", productcontroller.Index)
	http.HandleFunc("/products/add", productcontroller.Add)
	http.HandleFunc("/products/detail", productcontroller.Detail)
	http.HandleFunc("/products/edit", productcontroller.Edit)
	http.HandleFunc("/products/delete", productcontroller.Delete)

	log.Println("Server berjalan di port 8000")
	// membuat server sendiri
	http.ListenAndServe(":8000", nil)
}
