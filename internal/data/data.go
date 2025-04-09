package data

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"

	// "gorm.io/driver/mysql"
	"gorm.io/driver/postgres"

	"gorm.io/gorm"

	"github.com/itsgn/sgn-layout/internal/conf"
	"github.com/itsgn/sgn-layout/pkg/middleware/requestInfo"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewRedisCmd,
	NewMysqlCmd,
	NewRedisClient,
	NewExampleRepo)

// Data .
type Data struct {
	cfg    *conf.Bootstrap
	logger *log.Helper
	db     *gorm.DB
	rdb    *redis.Client
}

func NewData(cfg *conf.Bootstrap, db *gorm.DB, redisCli *redis.Client, logger log.Logger) (*Data, func(), error) {
	logs := log.NewHelper(log.With(logger, "module", "sgn-layout/data"))
	cleanup := func() {
		logs.Info("closing the data resources")
	}

	return &Data{
		logger: logs,
		cfg:    cfg,
		db:     db,
		rdb:    redisCli,
	}, cleanup, nil
}

func NewRedisClient(conf *conf.Data) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:         conf.Redis.Addr,
		Password:     conf.Redis.Password,
		ReadTimeout:  conf.Redis.ReadTimeout.AsDuration(),
		WriteTimeout: conf.Redis.WriteTimeout.AsDuration(),
		DialTimeout:  time.Second * 2,
		PoolSize:     10,
		DB:           int(conf.Redis.Db),
	})
	err := client.Ping().Err()
	if err != nil {
		log.Fatalf("redis connect error: %v", err)
	}
	return client
}

func getDomain(ctx context.Context) string {
	domain := ctx.Value(requestInfo.DomainKey)
	return domain.(string)
}

func getDbWithDomain(ctx context.Context, db *gorm.DB) *gorm.DB {
	domain := ctx.Value(requestInfo.DomainKey)
	if domain != nil {
		db = db.Where("domain = ?", domain)
	}
	return db
}

func NewRedisCmd(conf *conf.Data, logger log.Logger) redis.Cmdable {
	logs := log.NewHelper(log.With(logger, "module", "serviceName/data/redis"))
	client := redis.NewClient(&redis.Options{
		Addr:         conf.Redis.Addr,
		Password:     conf.Redis.Password,
		ReadTimeout:  conf.Redis.ReadTimeout.AsDuration(),
		WriteTimeout: conf.Redis.WriteTimeout.AsDuration(),
		DialTimeout:  time.Second * 2,
		PoolSize:     10,
	})
	err := client.Ping().Err()
	if err != nil {
		logs.Fatalf("redis connect error: %v", err)
	}
	return client
}

func NewMysqlCmd(conf *conf.Bootstrap, logger log.Logger) *gorm.DB {
	logs := log.NewHelper(log.With(logger, "module", "serviceName/data/mysql"))
	// db, err := gorm.Open(mysql.Open(conf.Data.Database.Source), &gorm.Config{})
	db, err := gorm.Open(postgres.Open(conf.Data.Database.Source), &gorm.Config{})
	if err != nil {
		logs.Fatalf("postgres connect error: %v", err)
	}
	// If it is a development environment, print sql
	if conf.Env.Mode == "dev" {
		db = db.Debug()
	}
	// Data migration
	db.AutoMigrate(&ExampleEntity{})
	return db
}
