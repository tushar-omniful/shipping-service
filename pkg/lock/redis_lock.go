package lock

import (
	"context"
	"fmt"
	"github.com/omniful/go_commons/env"
	"github.com/omniful/go_commons/log"
	oredis "github.com/omniful/go_commons/redis"
	"time"
)

type RedisLock struct {
	Client *oredis.Client
}

func NewRedisLock(client *oredis.Client) *RedisLock {
	return &RedisLock{
		Client: client,
	}
}

func (r *RedisLock) Lock(ctx context.Context, key string, d time.Duration) error {
	logTag := fmt.Sprintf("RequestID: %s Function: Redis-Lock-Acquire", env.GetRequestID(ctx))
	log.Infof(logTag + "Acquiring lock for " + key)

	val, err := r.Client.SetNX(ctx, key, "true", d)
	if err != nil {
		log.Errorf(logTag+"failed to set lock,err: %+s", err.Error())

		return err
	}

	if !val {
		err = fmt.Errorf("failed to acquire lock")

		log.Errorf(logTag+"received val %+v,err: %s", val, err)

		return err
	}

	return nil
}

func (r *RedisLock) Release(ctx context.Context, lockKey string) error {
	logTag := fmt.Sprintf("RequestID: %s Function: Redis-Lock-Release", env.GetRequestID(ctx))

	log.Infof(logTag + "Releasing lock for " + lockKey)

	count, err := r.Client.Del(ctx, lockKey)
	if err != nil {
		log.Errorf(logTag+"failed to release lock. err: %s", err)
	}

	if count <= 0 {
		err = fmt.Errorf("failed to release lock")

		log.Errorf(logTag+"received val %+v,err: %s", count, err)

		return err
	}

	return nil
}
