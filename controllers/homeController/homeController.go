package homeController

import (
	"net/http"
	"text/template"
)

// membuat function handler berisi parameter response dan request
func Welcome(w http.ResponseWriter, r *http.Request) {
	// memanggil file index.html 
	// akan menangkap template dan errornya
	temp, err := template.ParseFiles("views/home/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, nil)
}
