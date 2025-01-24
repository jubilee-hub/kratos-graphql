package data

import (
	"arkgo/ent"
	"arkgo/internal/conf"
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

// NewEntClient creates new ent client
func NewEntClient(c *conf.Data, logger log.Logger) *ent.Client {
	log := log.NewHelper(logger)
	client, err := ent.Open(
		c.Database.Driver,
		c.Database.Source,
	)
	if err != nil {
		log.Fatalf("failed opening connection to db: %v", err)
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return client
}
