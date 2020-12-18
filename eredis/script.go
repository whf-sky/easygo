package eredis

import (
	"github.com/gomodule/redigo/redis"
)

type script struct {
	pool *redis.Pool
	s *redis.Script
}

func (st *script ) Script() (*redis.Pool, *redis.Script) {
	return st.pool, st.s
}

func (st *script ) Do(keysAndArgs ...interface{}) (interface{}, error) {
	c := st.pool.Get()
	defer c.Close()
	return st.s.Do(c, keysAndArgs...)
}

func (st *script ) Hash() string {
	return st.s.Hash()
}

func (st *script ) Load() error {
	c := st.pool.Get()
	defer c.Close()
	return st.s.Load(c)
}