package cache_loader

import (
	"context"
	"fmt"
	"sync"
	"tech-wb/internal/infrastructure/cache"
	"tech-wb/internal/model"
	"tech-wb/pkg/client/postgresql"
)

func LoadOrderCache(ctx context.Context, oCache cache.OrderCache, db postgresql.Client) {
	const op = "cache_loder Error"

	m := sync.Mutex{}
	m.Lock()
	defer m.Unlock()

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

	rows, err := db.Query(ctx, query, 100)
	if err != nil {
		fmt.Println(op)
		fmt.Println(err)

	}

	//ordersMap := make(map[string]*model.Order)
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

		orderInMap, ok := oCache.Get(o.OrderUid)
		if !ok {
			var items []model.Item
			items = append(items, i)

			o.Delivery = d
			o.Payment = p
			o.Items = items
			oCache.Set(o.OrderUid, &o)
			continue
		}

		if err != nil {
			fmt.Println(op)

			fmt.Println(err)
		}

		orderInMap.Items = append(orderInMap.Items, i)

	}
	return
}
