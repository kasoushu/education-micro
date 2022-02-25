package data

import (
	"context"
	cgv1 "education/api/v1/course"
	userv1 "education/api/v1/user"
	"education/app/interface/internal/conf"
	"fmt"
	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	jwt2 "github.com/golang-jwt/jwt/v4"

	//_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	consulApi "github.com/hashicorp/consul/api"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDiscovery, NewCourseClient, NewUserClient, NewUserRepo, NewGradeRepo, NewSelectRepo, NewCourseRepo)

// Data .
type Data struct {
	userClient   userv1.UserClient
	courseClient cgv1.CourseClient
}

// NewData .
func NewData(c *conf.AppConfig, userClient userv1.UserClient, courseClient cgv1.CourseClient) (*Data, func(), error) {
	data := Data{
		userClient:   userClient,
		courseClient: courseClient,
	}
	cleanup := func() {
		fmt.Println("closing!")
	}
	return &data, cleanup, nil
}

func NewUserClient(discover registry.Discovery, conf *conf.AppConfig) userv1.UserClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///education.user"),
		grpc.WithDiscovery(discover),
		grpc.WithMiddleware(
			tracing.Server(),
			recovery.Recovery(),
			jwt.Client(func(token *jwt2.Token) (interface{}, error) {
				return []byte(conf.Auth.ServiceKey), nil
			}, jwt.WithSigningMethod(jwt2.SigningMethodHS256)),
		),
	)
	if err != nil {
		panic(err)
	}
	c := userv1.NewUserClient(conn)
	return c
}

func NewCourseClient(discover registry.Discovery, conf *conf.AppConfig) cgv1.CourseClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///education.select.course"),
		grpc.WithDiscovery(discover),
		grpc.WithMiddleware(
			tracing.Server(),
			recovery.Recovery(),
			jwt.Client(func(token *jwt2.Token) (interface{}, error) {
				return []byte(conf.Auth.ServiceKey), nil
			}, jwt.WithSigningMethod(jwt2.SigningMethodHS256)),
		),
	)
	if err != nil {
		panic(err)
	}
	c := cgv1.NewCourseClient(conn)
	return c
}
func NewDiscovery(conf *conf.AppConfig) registry.Discovery {
	apiConfig := consulApi.DefaultConfig()
	if conf.Consul.Address != "" {
		apiConfig.Address = conf.Consul.Address
	}

	if conf.Consul.Scheme != "" {
		apiConfig.Scheme = conf.Consul.Scheme
	}
	apiCli, err := consulApi.NewClient(apiConfig)
	if err != nil {
		panic(err)
	}
	rs := consul.New(apiCli, consul.WithHealthCheck(false))
	return rs
}
