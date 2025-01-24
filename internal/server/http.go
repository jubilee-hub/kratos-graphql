package server

import (
	"arkgo/api/graphql/generated"
	"arkgo/internal/conf"
	"arkgo/internal/data"
	"arkgo/internal/service"
	"context"
	"errors"
	"runtime/debug"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, logger log.Logger, data *data.Data, resolver *service.Resolver, greeter *service.GreeterService) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)

	// GraphQL setup
	schema := generated.NewExecutableSchema(generated.Config{
		Resolvers: resolver,
	})

	// 创建GraphQL服务器
	graphqlHandler := handler.New(schema)
	graphqlHandler.AddTransport(transport.POST{})
	graphqlHandler.AddTransport(transport.GET{})

	// 添加错误恢复处理
	graphqlHandler.SetRecoverFunc(func(ctx context.Context, err interface{}) error {
		logger.Log(log.LevelError, "msg", "graphql panic", "error", err)
		debug.PrintStack()
		return errors.New("internal server error")
	})

	srv.HandlePrefix("/graphql", graphqlHandler)

	// GraphQL playground
	playgroundHandler := playground.Handler("GraphQL playground", "/graphql")
	srv.HandlePrefix("/playground", playgroundHandler)

	return srv
}
