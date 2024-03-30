package redismock

import (
	"github.com/gomodule/redigo/redis"
	"github.com/rafaeljusto/redigomock"
)

type RedisClient struct {
	conn *redigomock.Conn
}

func NewRedisClient(conn *redigomock.Conn) *RedisClient {
	return &RedisClient{conn: conn}
}

func (r *RedisClient) SetValue(key, value string) error {
	_, err := r.conn.Do("SET", key, value)
	return err
}

func (r *RedisClient) GetValue(key string) (string, error) {
	value, err := redis.String(r.conn.Do("GET", key))
	return value, err
}
