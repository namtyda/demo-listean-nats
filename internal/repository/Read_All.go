package repository

import (
	"context"

	"github.com/namtyda/demo-listean-nats/internal/models"
)

func (r *repository) ReadAll() (rowSlice []models.Cache, err error) {
	query := `
	SELECT order_uuid, data FROM orders;
	`
	rows, err := r.pool.Query(context.Background(), query)

	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var r models.Cache

		if err = rows.Scan(
			&r.Order_uuid,
			&r.Data,
		); err != nil {
			return
		}
		rowSlice = append(rowSlice, r)
	}
	return
}
