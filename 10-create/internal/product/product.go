package product

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"time"
)

// List gets all Products from the database
func List(db *sqlx.DB) ([]Product, error) {
	products := make([]Product, 0, 10)

	const query = `SELECT * FROM products`

	if err := db.Select(&products, query); err != nil {
		return nil, errors.Wrap(err, "selecting products")
	}

	return products, nil
}

func Retrieve(db *sqlx.DB, id string) (*Product, error) {
	var p Product
	const q = `select * from products where product_id = $1`
	if err := db.Get(&p, q, id); err != nil {
		return nil, errors.Wrap(err, "selecting single products")
	}
	return &p, nil
}

func Create(db *sqlx.DB, np NewProduct, now time.Time) (*Product, error) {
	p := Product{
		ID: uuid.New().String(),
		Name: np.Name,
		Quantity: np.Quantity,
		Cost: np.Cost,
		DateCreated: now.UTC(),
		DateUpdated: now.UTC(),
	}
	const q = `
		INSERT INTO products
		(product_id, name, cost, quantity, date_created, date_updated)
		VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := db.Exec(q,
		p.ID, p.Name,
		p.Cost, p.Quantity,
		p.DateCreated, p.DateUpdated)
	if err != nil {
		return nil, errors.Wrap(err, "inserting product")
	}

	return &p, nil
}
