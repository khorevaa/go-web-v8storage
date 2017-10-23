package service

import (
	"time"

	"github.com/kataras/iris/sessions"
	"github.com/kataras/iris/sessions/sessiondb/redis"
	"github.com/kataras/iris/sessions/sessiondb/redis/service"
)

var Sessions *sessions.Sessions

func init() {
	Sessions = sessions.New(sessions.Config{
		Expires: time.Hour * 24 * 7,
	})

	db := redis.New(service.Config{
		Addr:     Config.Redis.Addr,
		Password: Config.Redis.Password,
		Database: Config.Redis.Database,
		Prefix:   Config.Redis.Prefix,
	})

	Sessions.UseDatabase(db)
}
