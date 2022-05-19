package repository

import (
	"context"
)

func (p *repository) AddOrder(OrderUID, rawJson string) (err error) {

	query := `
	INSERT INTO orders (order_uuid, data) VALUES ($1, $2);
	`
	_, errR := p.pool.Exec(context.Background(), query, OrderUID, rawJson)
	if err != nil {
		err = errR
	}
	return
}
