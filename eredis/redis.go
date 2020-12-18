package eredis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/whf-sky/easygo"
)

func NewRedis(group string) *Redis {
	return (&Redis{}).dial(group)
}

type Redis struct {
	pool *redis.Pool
}

func (r *Redis) dial(group string)  *Redis {
	if _, ok := configs[easygo.EASYGO_ENV];!ok {
		panic("There is no config info for " + easygo.EASYGO_ENV + " environment.")
	}
	if cnf, ok := configs[easygo.EASYGO_ENV][group];ok {
		fmt.Println(cnf.MaxIdle, cnf.IdleTimeout, cnf.Network, cnf.Address, cnf.Options)
		r.pool = RedisDial(cnf.MaxIdle, cnf.IdleTimeout, cnf.Network, cnf.Address, cnf.Options...)
		return r
	}
	panic("There is no config info for "+ easygo.EASYGO_ENV + " group.")
}

func (r *Redis) Pool()  *redis.Pool {
	return r.Pool()
}

func (r *Redis) Do(commandName string, args ...interface{}) (reply interface{}, err error)   {
	c := r.pool.Get()
	defer c.Close()
	return c.Do(commandName, args...)
}

func (r *Redis) Subscribe(backfun func(psc redis.PubSubConn) error, channel ...interface{})  error  {
	c := r.pool.Get()
	defer c.Close()
	psc := redis.PubSubConn{Conn: c}
	psc.Subscribe(channel...)
	return backfun(psc)
}

func (r *Redis) NewScript(keyCount int, src string)  *script  {
	s := redis.NewScript(keyCount, src)
	return &script{
		pool: r.pool,
		s: s,
	}
}