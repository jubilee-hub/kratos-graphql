package data

import (
	"arkgo/ent"
	"arkgo/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewGreeterRepo,
	NewUserRepo,
)

// Data .
type Data struct {
	Ent *ent.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(logger)

	client, err := NewClient(c)
	if err != nil {
		log.Errorf("failed opening mysql client: %v", err)
		return nil, nil, err
	}

	d := &Data{
		Ent: client,
	}

	return d, func() {
		log.Info("message", "closing the data resources")
		if err := d.Ent.Close(); err != nil {
			log.Error(err)
		}
	}, nil
}

// NewClient
func NewClient(c *conf.Data) (*ent.Client, error) {
	var entOptions []ent.Option
	entOptions = append(entOptions, ent.Debug())
	return ent.Open(c.Database.Driver, c.Database.Source, entOptions...)
}

// GetClient returns the ent client
func (d *Data) GetClient() *ent.Client {
	return d.Ent
}
