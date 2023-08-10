package categoriesController

import (
	"CRUD-golang/entities"
	categoryModel "CRUD-golang/models/CategoryModel"
	"net/http"
	"strconv"
	"text/template"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	//  menangkap category dari model dan memanggil function menampilkan semua data
	categories := categoryModel.GetAll()
	// buat variabel dengan tipe data MAP
	data := map[string]any{
		// value map adalah categories dari modelnya
		"categories": categories,
	}

	// menampilkan data pada views
	temp, err := template.ParseFiles("views/category/index.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(w, data)
}

func Add(w http.ResponseWriter, r *http.Request) {
	// jika request methodnya get maka akan menampilkan halaman cerate.html
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/category/create.html")

		if err != nil {
			panic(err)
		}

		temp.Execute(w, nil)
	}

	if r.Method == "POST" {
		var category entities.Category

		// memanggil inputan name dari User
		category.Name = r.FormValue("name")
		// created at dan update at di set sekarang
		category.CreateAt = time.Now()
		category.UpdateAt = time.Now()

		// memanggil model dan mengembalikan boolean
		if ok := categoryModel.Create(category); !ok {
			// jika insert gagal maka akan kembali ke halaman create
			temp, _ := template.ParseFiles("views/category/create")
			temp.Execute(w, nil)
		}
		http.Redirect(w, r, "/categories", http.StatusSeeOther)
	}
}

func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/category/edit.html")

		if err != nil {
			panic(err)
		}
		// menangkap id dari action edit pada index.html
		idStr := r.URL.Query().Get("id")
		// konversi string jadi int
		id, err := strconv.Atoi(idStr)

		if err != nil {
			panic(err)
		}

		// buat variabel untuk menampung hasil edit
		category := categoryModel.Detail(id)
		// buat varibael data untuk melempar data agar tampil di halaman edit
		data := map[string]any{
			"category": category,
		}

		temp.Execute(w, data)
	}

	if r.Method == "POST" {
		// variabel penampung id hasil edit
		var category entities.Category

		idStr := r.FormValue("id")

		id, err := strconv.Atoi(idStr)

		if err != nil {
			panic(err)
		}
		// memnyimpan category yang sudah di buat
		category.Name = r.FormValue("name")
		category.UpdateAt = time.Now()
		// jika gagal update maka
		if ok := categoryModel.Update(id, category); !ok {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
			return
		}
		// jika berhasil maka
		http.Redirect(w, r, "/categories", http.StatusSeeOther)
	}

}

func Delete(w http.ResponseWriter, r *http.Request) {
	// menangkap id dari action edit pada index.html
	idStr := r.URL.Query().Get("id")
	// konversi string jadi int
	id, err := strconv.Atoi(idStr)

	if err != nil {
		panic(err)
	}

	// tangkap variabel error dari model func delete
	if err := categoryModel.Delete(id); err != nil {
		panic(err)
	}

	// jika tidak error maka redirect ke categories
	http.Redirect(w, r, "/categories", http.StatusSeeOther)
}
