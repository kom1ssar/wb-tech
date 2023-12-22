package order

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"log"
	"tech-wb/internal/infrastructure/cache"
	def "tech-wb/internal/infrastructure/transaction"
	"tech-wb/internal/model"
	"tech-wb/pkg/client/postgresql"
)

var _ def.OrderTransaction = (*transaction)(nil)

type transaction struct {
	db    postgresql.Client
	cache cache.OrderCache
}

func (t *transaction) Insert(ctx context.Context, order *model.Order) error {
	tx, err := t.db.Begin(ctx)
	if err != nil {
		log.Printf("Failed to start a order.insert transaction: %v", err)
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		} else {
			tx.Commit(ctx)
		}
	}()

	err = t.insertOrder(ctx, tx, order)
	if err != nil {
		return err
	}

	err = t.insertPayment(ctx, order.OrderUid, tx, &order.Payment)
	if err != nil {
		return err

	}

	err = t.insertDelivery(ctx, order.OrderUid, tx, &order.Delivery)
	if err != nil {
		return err
	}

	err = t.insertItems(ctx, order.OrderUid, tx, order.Items)

	if err != nil {
		return err
	}

	t.cache.Set(order.OrderUid, order)

	return nil
}

func (t *transaction) insertOrder(ctx context.Context, tx pgx.Tx, order *model.Order) error {

	query := `
		INSERT INTO orders (
		                    order_uid, track_number, entry, internal_signature,
		                    locale, customer_id, delivery_service, shardkey,
		                    sm_id, date_created, oof_shard
		                    )
		VALUES (
		        $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
				)
`

	_, err := tx.Exec(
		ctx, query,
		order.OrderUid, order.TrackNumber, order.Entry, order.InternalSignature,
		order.Locale, order.CustomerId, order.DeliveryService, order.Shardkey,
		order.SmId, order.DateCreated, order.OofShard,
	)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil

}

func (t *transaction) insertPayment(ctx context.Context, orderId string, tx pgx.Tx, payment *model.Payment) error {
	query := `INSERT INTO payment (
                    transaction, currency, provider, amount,
                    payment_dt, bank, delivery_cost, goods_total,
                    custom_fee, order_uid
                    )

			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
`

	_, err := tx.Exec(
		ctx, query,
		payment.Transaction, payment.Currency, payment.Provider, payment.Amount,
		payment.PaymentDt, payment.Bank, payment.DeliveryCost, payment.GoodsTotal,
		payment.CustomFee, orderId,
	)

	return err

}

func (t *transaction) insertDelivery(ctx context.Context, orderId string, tx pgx.Tx, delivery *model.Delivery) error {
	query := `INSERT INTO delivery (
                      name, phone, zip, city,
                      address, region, email, order_uid
                      )

			VALUES (
					$1, $2, $3, $4,
					$5, $6, $7, $8
				    )
`

	_, err := tx.Exec(
		ctx, query,
		delivery.Name, delivery.Phone, delivery.Zip, delivery.City,
		delivery.Address, delivery.Region, delivery.Email, orderId,
	)

	return err
}

func (t *transaction) insertItem(ctx context.Context, orderId string, tx pgx.Tx, item *model.Item) error {
	query := `	
	WITH inserted_item AS (
  		INSERT INTO item(
							chrt_id, track_number, price, rid,
							 name, sale, size, total_price,
							 nm_id, brand, status
							) 
		VALUES (
				$1, $2, $3, $4,
				$5, $6, $7, $8,
				$9, $10, $11
				) 
		RETURNING id
	)
	INSERT INTO orders_item (item_id, order_uid)
	SELECT id, $12 FROM inserted_item;
`

	_, err := tx.Exec(
		ctx, query,
		item.ChrtId, item.TrackNumber, item.Price, item.Rid,
		item.Name, item.Sale, item.Size, item.TotalPrice,
		item.NmId, item.Brand, item.Status, orderId,
	)

	return err
}

func (t *transaction) insertItems(ctx context.Context, orderId string, tx pgx.Tx, items []model.Item) error {

	for _, item := range items {
		err := t.insertItem(ctx, orderId, tx, &item)

		if err != nil {
			return err
		}
	}

	return nil
}

func NewOrderTransaction(db postgresql.Client, c cache.OrderCache) *transaction {
	return &transaction{
		db:    db,
		cache: c,
	}

}
