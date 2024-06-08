package http

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/IbnAnjung/dealls/pkg/cache"
	"github.com/IbnAnjung/dealls/pkg/crypt"
	pkghttp "github.com/IbnAnjung/dealls/pkg/http"
	"github.com/IbnAnjung/dealls/pkg/orm"
	"github.com/IbnAnjung/dealls/pkg/redis"
	"github.com/IbnAnjung/dealls/pkg/sql"
	"github.com/IbnAnjung/dealls/pkg/structvalidator"
	pkgtime "github.com/IbnAnjung/dealls/pkg/time"

	"github.com/IbnAnjung/dealls/cmd/http/config"
	"github.com/IbnAnjung/dealls/cmd/http/router"
	"github.com/IbnAnjung/dealls/pkg/jwt"
	"github.com/IbnAnjung/dealls/pkg/string"
	repository "github.com/IbnAnjung/dealls/repository/gorm"
	"github.com/IbnAnjung/dealls/usecase/auth"
	"github.com/IbnAnjung/dealls/usecase/user"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type echoHttpServer struct {
	e     *echo.Echo
	mysql sql.MysqlConnection
	redis redis.Redis
}

func NewEchoHttpServer() *echoHttpServer {
	return &echoHttpServer{}
}

func (server *echoHttpServer) Start(ctx context.Context) {
	e := echo.New()
	e.Logger.SetLevel(log.INFO)

	config := config.LoadConfig()
	t := pkgtime.NewTimeService()

	validator := structvalidator.NewStructValidator()
	hasher := crypt.NewHasherString()
	randomString := string.NewRandomString()

	jwt := jwt.NewJwtServive(jwt.JwtConfig{
		SecretKey:            config.Jwt.SecretKey,
		AccessTokenLifeTime:  time.Duration(config.Jwt.AccessTokenLifeTime) * time.Hour,
		RefreshTokenLifeTime: time.Duration(config.Jwt.RefreshTokenLifeTime) * time.Hour,
	})

	// open mysql connection
	mysql, err := sql.NewMysqlConnection(ctx, sql.MysqlConfig{
		User:               config.Mysql.User,
		Password:           config.Mysql.Password,
		Host:               config.Mysql.Host,
		Port:               config.Mysql.Port,
		DbName:             config.Mysql.Schema,
		Loc:                t.GetDefaultLoc(),
		Timeout:            time.Duration(config.Mysql.Timeout) * time.Second,
		MaxIddleConnection: config.Mysql.MaxIddleConnection,
		MaxOpenConnection:  config.Mysql.MaxOpenConnection,
		MaxLifeTime:        config.Mysql.MaxLifeTime,
	})
	if err != nil {
		panic(fmt.Sprintf("server startup panic: %s", err))
	}

	redis, err := redis.NewRedis(ctx, redis.RedisConfig{
		Addr:           config.Redis.Addr,
		Username:       config.Redis.Username,
		Password:       config.Redis.Password,
		ClientName:     "auth_service",
		DB:             config.Redis.DB,
		MinIdleConns:   config.Redis.MinIdleConns,
		MaxIdleConns:   config.Redis.MaxIdleConns,
		MaxActiveConns: config.Redis.MaxActiveConns,
	})
	if err != nil {
		panic(fmt.Sprintf("fail start redis connection %s", err.Error()))
	}

	redisCache := cache.NewRedisCache(redis.Client)

	guow, err := orm.NewGormOrm(orm.GormConfig{
		Connection: mysql.Db,
		Dialect:    orm.MysqlDialect,
	})

	if err != nil {
		panic(fmt.Sprintf("server startup panic: %s", err))
	}

	// repository
	userRepository := repository.NewGormUserRepository(guow)

	// usecase
	authUc := auth.NewUsecase(hasher, jwt, validator, t, randomString, userRepository)
	userUC := user.NewUsecase(validator, t, redisCache, userRepository)

	// default http middleware
	pkghttp.LoadEchoRequiredMiddleware(e)

	router.SetupRouter(
		e, authUc, userUC, jwt,
	)

	server.e = e
	server.mysql = mysql
	server.redis = redis

	if err := e.Start(fmt.Sprintf(":%s", config.Http.Port)); err != nil && err != http.ErrServerClosed {
		panic(fmt.Sprintf("server startup panic: %s", err))
	}
}

func (server *echoHttpServer) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.e.Shutdown(ctx); err != nil {
		server.e.Logger.Fatal(err)
	}

	if err := server.mysql.Cleanup(); err != nil {
		server.e.Logger.Fatal(err)
	}

	if err := server.redis.Cleanup(); err != nil {
		server.e.Logger.Fatal(err)
	}
}
