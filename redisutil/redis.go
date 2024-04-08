package redisutil

import (
	"github.com/gomodule/redigo/redis"
	"github.com/gotoeasy/glang/cmn"
	"go-ocr/pkg/settings"
	"time"
)

const (
	DEFAULT_REDIS_PRE_KEY    = "ocr:image:"
	DEFAULT_REDIS_PRE_NX_KEY = "ocr:image:nx:"
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
func (r *RedisDataStore) SetNxEx(k string, v interface{}, ex int64) (interface{}, error) {
	c := r.RedisPool.Get()
	defer c.Close()
	result, err := c.Do("SETNX", k, v)
	return result, err
}
func (r *RedisDataStore) Del(k string) (interface{}, error) {
	c := r.RedisPool.Get()
	defer c.Close()
	result, err := c.Do("DEL", k)
	return result, err
}

var RDS RedisDataStore

func Start() {
	cmn.Info("Redis Init")
	RDS = RedisDataStore{
		RedisHost:       settings.OcrConfig.Redis.RedisHost,
		RedisDB:         settings.OcrConfig.Redis.RedisDB,
		RedisPwd:        settings.OcrConfig.Redis.RedisPwd,
		Timeout:         settings.OcrConfig.Redis.Timeout,
		PoolMaxIdle:     settings.OcrConfig.Redis.PoolMaxIdle,
		PoolMaxActive:   settings.OcrConfig.Redis.PoolMaxActive,
		PoolIdleTimeout: settings.OcrConfig.Redis.PoolIdleTimeout,
		PoolWait:        settings.OcrConfig.Redis.PoolWait,
		RedisPool:       nil,
	}
	RDS.RedisPool = RDS.NewPool()
}
