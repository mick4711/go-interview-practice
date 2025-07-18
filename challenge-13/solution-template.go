package main

import (
	"database/sql"
	"errors"

	_ "github.com/mattn/go-sqlite3"
)

// Product represents a product in the inventory system
type Product struct {
	ID       int64
	Name     string
	Price    float64
	Quantity int
	Category string
}

// ProductStore manages product operations
type ProductStore struct {
	db *sql.DB
}

// NewProductStore creates a new ProductStore with the given database connection
func NewProductStore(db *sql.DB) *ProductStore {
	return &ProductStore{db: db}
}

// InitDB sets up a new SQLite database and creates the products table
func InitDB(dbPath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	// Check if the database is accessible
	if err = db.Ping(); err != nil {
		return nil, err
	}
	// Create the products table if it doesn't exist
	// The table should have columns: id, name, price, quantity, category
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS products (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		price REAL NOT NULL,
		quantity INTEGER NOT NULL,
		category TEXT NOT NULL
	);
	`
	_, err = db.Exec(createTableSQL)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// CreateProduct adds a new product to the database
func (ps *ProductStore) CreateProduct(product *Product) error {
	// Insert the product into the database
	result, err := ps.db.Exec(`
		INSERT INTO products (name, price, quantity, category)
		VALUES (?, ?, ?, ?)
	`, product.Name, product.Price, product.Quantity, product.Category)
	if err != nil {
		return err
	}

	// Update the product.ID with the database-generated ID
	product.ID, err = result.LastInsertId()
	return err
}

// GetProduct retrieves a product by ID
func (ps *ProductStore) GetProduct(id int64) (*Product, error) {
	// Query the database for a product with the given ID
	row := ps.db.QueryRow(`
		SELECT id, name, price, quantity, category
		FROM products
		WHERE id = ?
	`, id)

	// Scan the result into a Product struct
	product := &Product{}
	if err := row.Scan(&product.ID, &product.Name, &product.Price, &product.Quantity, &product.Category); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("product not found")
		}
		return nil, err
	}
	return product, nil
}

// UpdateProduct updates an existing product
func (ps *ProductStore) UpdateProduct(product *Product) error {
	// TODO: Update the product in the database
	// TODO: Return an error if the product doesn't exist
	return errors.New("not implemented")
}

// DeleteProduct removes a product by ID
func (ps *ProductStore) DeleteProduct(id int64) error {
	// TODO: Delete the product from the database
	// TODO: Return an error if the product doesn't exist
	return errors.New("not implemented")
}

// ListProducts returns all products with optional filtering by category
func (ps *ProductStore) ListProducts(category string) ([]*Product, error) {
	// TODO: Query the database for products
	// TODO: If category is not empty, filter by category
	// TODO: Return a slice of Product pointers
	return nil, errors.New("not implemented")
}

// BatchUpdateInventory updates the quantity of multiple products in a single transaction
func (ps *ProductStore) BatchUpdateInventory(updates map[int64]int) error {
	// TODO: Start a transaction
	// TODO: For each product ID in the updates map, update its quantity
	// TODO: If any update fails, roll back the transaction
	// TODO: Otherwise, commit the transaction
	return errors.New("not implemented")
}

func main() {
	// Optional: you can write code here to test your implementation
}
