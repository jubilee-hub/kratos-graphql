package data

import (
	"arkgo/ent"
	"context"
	"fmt"
)

type TxManager interface {
	WithTx(context.Context, func(tx *ent.Tx) error) error
}

func (d *Data) WithTx(ctx context.Context, fn func(tx *ent.Tx) error) error {
	tx, err := d.Ent.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			panic(v)
		}
	}()
	if err := fn(tx); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("rolling back transaction: %w", rerr)
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}
	return nil
}
