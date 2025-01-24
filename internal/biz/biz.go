package biz

import (
	"arkgo/ent"

	"github.com/google/wire"
)

type Cursor = ent.Cursor

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(
	NewGreeterUsecase,
	NewUserUsecase,
)
