package ProductModel

import (
	"CRUD-golang/config"
	"CRUD-golang/entities"
)

// menampilkan semua data

func GetAll() []entities.Product {
	rows, err := config.DB.Query(
		"SELECT products.id, products.name,categories.name as category_name,products.stock,products.description,products.created_at,products.update_at FROM products JOIN categories ON products.category_id = categories.id ")

	if err != nil {
		panic(err)
	}

	// tutup rows
	defer rows.Close()

	// variabel untuk tampung semua data
	var products []entities.Product

	// loop rows
	for rows.Next() {
		// variabel untuk single data
		var product entities.Product
		// scan data
		err := rows.Scan(
			&product.Id,
			&product.Name,
			// karena ada relasi
			&product.Category.Name,
			&product.Stock,
			&product.Description,
			&product.CreateAt,
			&product.UpdateAt,
		)
		if err != nil {
			panic(err)
		}

		products = append(products, product)
	}
	return products
}

func Create(product entities.Product) bool  {
	result, err := config.DB.Exec(
		"INSERT INTO products (name,category_id,stock,description,created_at,update_at) VALUE (?, ?, ?, ?, ?, ?)",
		product.Name,product.Category.Id,product.Stock,product.Description,product.CreateAt,product.UpdateAt,
	)
	if err !=  nil {
		panic(err)
	}

	lastInsertId, err :=  result.LastInsertId()

	if err != nil {
		 panic(err)
	}

	return lastInsertId > 0
}


func Detail(id int) entities.Product {
	rows := config.DB.QueryRow(
		"SELECT products.id, products.name,categories.name as category_name,products.stock,products.description,products.created_at,products.update_at FROM products JOIN categories ON products.category_id = categories.id WHERE products.id = ?", id)
	// variabel penampung
	var product entities.Product

	// Scan rows
	err := rows.Scan(
		&product.Id,
		&product.Name,
		// karena ada relasi
		&product.Category.Name,
		&product.Stock,
		&product.Description,
		&product.CreateAt,
		&product.UpdateAt,
	)
	if err != nil {
		panic(err)
	}
	return product
}

func Update(id int, product entities.Product) bool {
	query, err := config.DB.Exec(
		"UPDATE products SET name = ?, category_id = ?, stock = ?, description = ?, update_at = ? WHERE id = ?", product.Name,product.Category.Id,product.Stock,product.Description,product.UpdateAt,id,)
		if err != nil {
			panic(err)
		}
	result, err := query.RowsAffected()

	if err != nil {
		panic(err)
	}
	return result > 0
}

func Delete(id int) error{
	_, err := config.DB.Exec("DELETE FROM products WHERE id = ?", id)

	return err
}
