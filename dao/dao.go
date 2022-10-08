package dao

import (
	"anohan/config"
	"github.com/gomodule/redigo"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/olivere/elastic"
)

type Dao struct {
	conf       *config.Config
	db         *sqlx.DB
	redis      *redis.Pool
	es         *elastic.Client
	expireTime int
}

func NewDao(c *config.Config) (d *Dao, err error) {
	d = &Dao{
		conf: c,
	}
	db, err := connectDB(c.Database.Master.Addr)
	d.db = db

	redis, err := connectRedis(c.Redis.Master.Addr, c.Redis.Master.Auth, c.Redis.Master.Db)
	d.redis = redis
	d.expireTime = 86400
	return
}
