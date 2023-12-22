package order

import (
	"context"
	"fmt"
	"sync"
	"tech-wb/internal/infrastructure/cache"
	def "tech-wb/internal/infrastructure/repository"
	"tech-wb/internal/model"
	"tech-wb/pkg/client/postgresql"
	utilsCacheLoader "tech-wb/pkg/utils/cache-loader"
)

var _ def.OrderRepository = (*repository)(nil)

type repository struct {
	db    postgresql.Client
	cache cache.OrderCache
	m     sync.RWMutex
}

func (r *repository) GetByUUId(ctx context.Context, uuid string) (*model.Order, error) {

	order, ok := r.cache.Get(uuid)

	if ok {
		fmt.Println("from cache")
		return order, nil
	}

	r.m.RLock()
	defer r.m.RUnlock()

	query := `SELECT
		o.order_uid, o.track_number, o.entry, o.locale, o.internal_signature,
		o.customer_id, o.delivery_service, o.shardkey, o.sm_id,
		o.date_created,

		d.id, d.name, d.phone, d.zip, d.city, d.address, d.region,
		d.email,

		p.id, p.transaction, p.request_id, p.currency, p.provider, p.amount,
	  	p.bank, p.delivery_cost, p.goods_total,
		p.custom_fee, p.payment_dt,
		
		i.id, i.chrt_id, i.track_number, i.price, i.rid, i.name, i.sale, i.size,
    	i.total_price, i.nm_id, i.brand,
    	i.status

	FROM orders as o
	LEFT JOIN delivery d
        ON o.order_uid = d.order_uid
	LEFT JOIN payment p
        ON o.order_uid = p.order_uid
	 LEFT JOIN  orders_item oi
        ON o.order_uid = oi.order_uid
	LEFT JOIN item i
         ON oi.item_id = i.id
	WHERE o.order_uid = $1
`

	var o model.Order
	var d model.Delivery
	var p model.Payment
	var items []model.Item

	rows, err := r.db.Query(ctx, query, uuid)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	for rows.Next() {
		var i model.Item

		err := rows.Scan(
			&o.OrderUid, &o.TrackNumber, &o.Entry, &o.Locale, &o.InternalSignature, &o.CustomerId, &o.DeliveryService, &o.Shardkey, &o.SmId, &o.DateCreated,
			&d.Id, &d.Name, &d.Phone, &d.Zip, &d.City, &d.Address, &d.Region, &d.Email,
			&p.Id, &p.Transaction, &p.RequestId, &p.Currency, &p.Provider, &p.Amount, &p.Bank, &p.DeliveryCost, &p.GoodsTotal, &p.CustomFee, &p.PaymentDt,
			&i.Id, &i.ChrtId, &i.TrackNumber, &i.Price, &i.Rid, &i.Name, &i.Sale, &i.Size, &i.TotalPrice, &i.NmId, &i.Brand, &i.Status,
		)

		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		items = append(items, i)

	}

	o.Delivery = d
	o.Payment = p
	o.Items = items

	return &o, nil
}

func (r *repository) GetListDESCCreated(ctx context.Context, limit int) ([]*model.Order, error) {
	r.m.RLock()
	defer r.m.RUnlock()

	query := `SELECT
		o.order_uid, o.track_number, o.entry, o.locale, o.internal_signature,
		o.customer_id, o.delivery_service, o.shardkey, o.sm_id,
		o.date_created,

		d.id, d.name, d.phone, d.zip, d.city, d.address, d.region,
		d.email,

		p.id, p.transaction, p.request_id, p.currency, p.provider, p.amount,
	  	p.bank, p.delivery_cost, p.goods_total,
		p.custom_fee, p.payment_dt,
		
		i.id, i.chrt_id, i.track_number, i.price, i.rid, i.name, i.sale, i.size,
    	i.total_price, i.nm_id, i.brand,
    	i.status

	FROM orders as o
	LEFT JOIN delivery d
        ON o.order_uid = d.order_uid
	LEFT JOIN payment p
        ON o.order_uid = p.order_uid
	 LEFT JOIN  orders_item oi
        ON o.order_uid = oi.order_uid
	LEFT JOIN item i
         ON oi.item_id = i.id
	
	ORDER BY date_created DESC
	
	LIMIT $1
`

	var orders []*model.Order

	rows, err := r.db.Query(ctx, query, limit)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	ordersMap := make(map[string]*model.Order)
	for rows.Next() {
		var o model.Order
		var i model.Item
		var d model.Delivery
		var p model.Payment

		err := rows.Scan(
			&o.OrderUid, &o.TrackNumber, &o.Entry, &o.Locale, &o.InternalSignature, &o.CustomerId, &o.DeliveryService, &o.Shardkey, &o.SmId, &o.DateCreated,
			&d.Id, &d.Name, &d.Phone, &d.Zip, &d.City, &d.Address, &d.Region, &d.Email,
			&p.Id, &p.Transaction, &p.RequestId, &p.Currency, &p.Provider, &p.Amount, &p.Bank, &p.DeliveryCost, &p.GoodsTotal, &p.CustomFee, &p.PaymentDt,
			&i.Id, &i.ChrtId, &i.TrackNumber, &i.Price, &i.Rid, &i.Name, &i.Sale, &i.Size, &i.TotalPrice, &i.NmId, &i.Brand, &i.Status,
		)

		orderInMap, ok := ordersMap[o.OrderUid]
		if !ok {
			var items []model.Item
			items = append(items, i)

			o.Delivery = d
			o.Payment = p
			o.Items = items
			ordersMap[o.OrderUid] = &o
			continue
		}

		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		orderInMap.Items = append(orderInMap.Items, i)

	}

	for _, v := range ordersMap {
		orders = append(orders, v)
	}

	return orders, nil
}

func NewRepository(db postgresql.Client, c cache.OrderCache) *repository {
	utilsCacheLoader.LoadOrderCache(context.TODO(), c, db)
	return &repository{
		db:    db,
		cache: c,
	}

}
