package repository

func (r *repository) AddOrder(OrderUID, rawJson string) (err error) {

	query := `
	INSERT INTO orders (order_uuid, data) VALUES ($1, $2);
	`
	_, errR := r.pool.Exec(r.ctx, query, OrderUID, rawJson)
	if err != nil {
		err = errR
	}
	return
}
