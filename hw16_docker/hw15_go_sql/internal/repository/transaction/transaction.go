package transaction

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/srg77global/home_work_basic/hw15_go_sql/internal/repository/online_shop"
)

func CreateOrderByUsernameProductnameAndQuantity(
	ctx context.Context,
	db *pgxpool.Pool,
	dbConn *online_shop.Queries,
	username, productname string,
	quantity pgtype.Numeric,
) error {
	tx, err := db.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.RepeatableRead})
	if err != nil {
		return fmt.Errorf("error creating tx: %w", err)
	}
	defer func() {
		if err = tx.Rollback(ctx); err != nil {
			log.Printf("error rollback tx: %v", err)
		}
	}()
	dbConn = dbConn.WithTx(tx)

	pProductname, err := dbConn.CreateOrderByUsernameProductnameAndQuantityFirst(
		ctx,
		online_shop.CreateOrderByUsernameProductnameAndQuantityFirstParams{
			Column1: &username,
			Column2: &productname,
			Column3: quantity,
		},
	)
	if err != nil {
		return fmt.Errorf("error function CreateOrderByUsernameProductnameAndQuantityFirst: %w", err)
	}
	if *pProductname != productname {
		return errors.New("invalid productname")
	}

	orderID, err := dbConn.CreateOrderByUsernameProductnameAndQuantitySecond(ctx, productname)
	if err != nil {
		return fmt.Errorf("error function CreateOrderByUsernameProductnameAndQuantitySecond: %w", err)
	}
	log.Println("Order created. Order ID:", orderID)

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("error tx commit: %w", err)
	}
	return nil
}
