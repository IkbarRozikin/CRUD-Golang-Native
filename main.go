package main

import (
	"crud-go-native/config"
	"crud-go-native/controllers/authcontroller"
	"crud-go-native/controllers/categorycontroller"
	"crud-go-native/controllers/homecontroller"
	"crud-go-native/controllers/productcontroller"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()

	// 1. Auth
	http.HandleFunc("/", homecontroller.Welcome)
	http.HandleFunc("/login", authcontroller.Login)
	http.HandleFunc("/register", authcontroller.Register)
	http.HandleFunc("/logout", authcontroller.Logout)
	// 2. Category
	http.HandleFunc("/categories", categorycontroller.Index)
	http.HandleFunc("/categories/add", categorycontroller.Add)
	http.HandleFunc("/categories/edit", categorycontroller.Edit)
	http.HandleFunc("/categories/delete", categorycontroller.Delete)
	// 3. Products
	http.HandleFunc("/products", productcontroller.Index)
	http.HandleFunc("/products/add", productcontroller.Add)
	http.HandleFunc("/products/detail", productcontroller.Detail)
	http.HandleFunc("/products/edit", productcontroller.Edit)
	http.HandleFunc("/products/delete", productcontroller.Delete)

	log.Println("Hello, sekarang berjalan di port :8080")
	http.ListenAndServe(":8080", nil)
}
