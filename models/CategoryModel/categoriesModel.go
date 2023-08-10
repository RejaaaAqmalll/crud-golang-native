package categorymodel

import (
	"CRUD-golang/config"
	"CRUD-golang/entities"
)

// membuat function untuk menampilkan semua data kategori

func GetAll() []entities.Category {
	// memanggil config database
	// menangkap rows dan error
	rows, err := config.DB.Query("SELECT * FROM categories")

	if err != nil {
		panic(err)
	}

	// menutup rows
	defer rows.Close()

	// variabel untuk menampung semua data category
	var categories []entities.Category

	// looping rows
	for rows.Next() {
		// membuat variabel temporary untuk menyimpan single data
		var category entities.Category
		// Scan digunakan untuk memindahkan data ke variabel category dan mengembalikan error
		if err := rows.Scan(&category.Id, &category.Name, &category.CreateAt, &category.UpdateAt); err != nil {
			panic(err)
		}
		// menambahkan pada slice category
		categories = append(categories, category)
	}

	return categories
}

func Create(category entities.Category) bool {
	// melakukan insert data maka panggil func Exec
	result, err := config.DB.Exec(
		"INSERT INTO categories (name, created_at, update_at) VALUE (? , ? , ? )",
		category.Name, category.CreateAt, category.UpdateAt,
	)

	if err != nil {
		panic(err)
	}

	lastInsertId, err := result.LastInsertId()

	if err != nil {
		panic(err)
	}

	return lastInsertId > 0
}

func Detail(id int) entities.Category {
	// melakukan edit data maka panggil function QueryRow
	rows := config.DB.QueryRow("SELECT id, name FROM categories WHERE id=?", id)

	// variabel penampung category
	var category entities.Category

	// scan rownya
	if err := rows.Scan(&category.Id, &category.Name); err != nil {
		panic(err.Error())
	}

	return category
}

func Update(id int, category entities.Category) bool {
	query, err := config.DB.Exec(
		"UPDATE categories SET name = ?, update_at = ? WHERE id = ?",
		category.Name, category.UpdateAt, id)
	if err != nil {
		panic(err)
	}

	// cek apakah data sudah terupdate atau belum
	result, err := query.RowsAffected()

	if err != nil {
		panic(err)
	}

	return result > 0
}

func Delete(id int) error {
	_, err := config.DB.Exec("DELETE FROM categories WHERE id = ?", id)

	return err

}
