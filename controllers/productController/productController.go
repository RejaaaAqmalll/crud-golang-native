package productcontroller

import (
	"CRUD-golang/entities"
	categorymodel "CRUD-golang/models/CategoryModel"
	productModel "CRUD-golang/models/ProductModel"
	"net/http"
	"strconv"
	"text/template"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	//  menangkap category dari model dan memanggil function menampilkan semua data
	products := productModel.GetAll()
	// buat variavel dengan tipe data MAP
	data := map[string]any{
		// value map adalah categories dari modelnya
		"products": products,
	}

	// menampilkan data pada views
	temp, err := template.ParseFiles("views/product/index.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(w, data)

}

func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/product/create.html")
		if err != nil {
			panic(err)
		}

		categories := categorymodel.GetAll()
		data := map[string]any {
			"categories" : categories,
		}

		temp.Execute(w, data)
	}
	// KIRIM DATA
	if r.Method == "POST" {
	// variabel product tipe data struct
		var product entities.Product

		// parsing category id 
		categoryId, err := strconv.Atoi(r.FormValue("category_id"))

		if err != nil {
			panic(err)
		}

		stock, err := strconv.Atoi(r.FormValue("stock"))

		if err != nil {
			panic(err)
		}
		// memamggil inputan user
		product.Name = r.FormValue("name")
		product.Category.Id = uint(categoryId) 
		product.Stock= int64(stock)
		product.Description  = r.FormValue("description")
		product.CreateAt = time.Now()
		product.UpdateAt = time.Now()

		// BERHASIL
		if ok := productModel.Create(product); !ok {
			// jika gagal
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusTemporaryRedirect)
		}
		http.Redirect(w, r, "/products", http.StatusSeeOther)
	}
}

func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/product/edit.html")
		if err != nil {
			panic(err)
		}

		idStr := r. URL.Query().Get("id")

		id, err := strconv.Atoi(idStr)

		if err != nil {
			panic(err)
		}

		product := productModel.Detail(id)

		categories := categorymodel.GetAll()
		data := map[string]any {
			"categories" : categories,
			"product" : product,
		}

		temp.Execute(w, data)
	}

	if r.Method == "POST" {
		// variabel product tipe data struct
		var product entities.Product

		idStr := r.FormValue("id")

		id, err := strconv.Atoi(idStr)

		if err != nil {
			panic(err)
		}

		// parsing category id 
		categoryId, err := strconv.Atoi(r.FormValue("category_id"))

		if err != nil {
			panic(err)
		}

		stock, err := strconv.Atoi(r.FormValue("stock"))

		if err != nil {
			panic(err)
		}
		// memamggil inputan user
		product.Name = r.FormValue("name")
		product.Category.Id = uint(categoryId) 
		product.Stock= int64(stock)
		product.Description  = r.FormValue("description")
		product.UpdateAt = time.Now()

		// BERHASIL
		if ok := productModel.Update(id, product); !ok {
			// jika gagal
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusTemporaryRedirect)
		}
		http.Redirect(w, r, "/products", http.StatusSeeOther)
	}
}

func Detail(w http.ResponseWriter, r *http.Request) {
	// menangkap id
	idString := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idString)

	if err != nil {
		panic(err)
	}

	product := productModel.Detail(id)
	data := map[string]any {
		"product" : product,
	}

	temp, err := template.ParseFiles("views/product/detail.html")

	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)


}

func Delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")

	// konversi int ke str
	id, err := strconv.Atoi(idStr)

	if err != nil {
		panic(err)
	}

	// tangkap variabel dari model
	if err := productModel.Delete(id); err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/products", http.StatusSeeOther)

}
