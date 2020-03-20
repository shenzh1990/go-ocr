package redisutil

import (
	"BitCoin/pkg/settings"
	"github.com/gomodule/redigo/redis"
	"time"
)

type RedisDataStore struct {
	RedisHost string
	RedisDB   string
	RedisPwd  string
	Timeout   int64

	PoolMaxIdle     int
	PoolMaxActive   int
	PoolIdleTimeout int64
	PoolWait        bool
	RedisPool       *redis.Pool
}

func (r *RedisDataStore) NewPool() *redis.Pool {

	return &redis.Pool{
		Dial:        r.RedisConnect,
		MaxIdle:     r.PoolMaxIdle,
		MaxActive:   r.PoolMaxActive,
		IdleTimeout: time.Duration(r.PoolIdleTimeout) * time.Second,
		Wait:        r.PoolWait,
	}
}

func (r *RedisDataStore) RedisConnect() (redis.Conn, error) {
	c, err := redis.Dial("tcp", r.RedisHost)
	if err != nil {
		return nil, err
	}
	_, err = c.Do("AUTH", r.RedisPwd)

	if err != nil {
		return nil, err
	}

	_, err = c.Do("SELECT", r.RedisDB)
	if err != nil {
		return nil, err
	}

	redis.DialConnectTimeout(time.Duration(r.Timeout) * time.Second)
	redis.DialReadTimeout(time.Duration(r.Timeout) * time.Second)
	redis.DialWriteTimeout(time.Duration(r.Timeout) * time.Second)

	return c, nil
}

func (r *RedisDataStore) Get(k string) (interface{}, error) {
	c := r.RedisPool.Get()
	defer c.Close()
	v, err := c.Do("GET", k)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (r *RedisDataStore) Set(k, v string) error {
	c := r.RedisPool.Get()
	defer c.Close()
	_, err := c.Do("SET", k, v)
	return err
}

func (r *RedisDataStore) SetEx(k string, v interface{}, ex int64) error {
	c := r.RedisPool.Get()
	defer c.Close()
	_, err := c.Do("SET", k, v, "EX", ex)
	return err
}

var RDS RedisDataStore

func init() {
	RDS = RedisDataStore{
		RedisHost:       settings.BitConfig.Redis.RedisHost,
		RedisDB:         settings.BitConfig.Redis.RedisDB,
		RedisPwd:        settings.BitConfig.Redis.RedisPwd,
		Timeout:         settings.BitConfig.Redis.Timeout,
		PoolMaxIdle:     settings.BitConfig.Redis.PoolMaxIdle,
		PoolMaxActive:   settings.BitConfig.Redis.PoolMaxActive,
		PoolIdleTimeout: settings.BitConfig.Redis.PoolIdleTimeout,
		PoolWait:        settings.BitConfig.Redis.PoolWait,
		RedisPool:       nil,
	}
	RDS.RedisPool = RDS.NewPool()
}
