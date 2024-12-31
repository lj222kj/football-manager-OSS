package redis

import (
	"context"
	"time"

	"github.com/lj222kj/xpkg/config"
	"github.com/redis/go-redis/v9"
)

type service struct {
	cfg    *config.Redis
	client *redis.Client
}

func New(cfg *config.Redis) (*service, error) {
	svc := &service{
		cfg: cfg,
		client: redis.NewClient(&redis.Options{
			Addr:     cfg.Host + ":" + cfg.Port,
			Password: cfg.Password,
			DB:       0,
		})}

	if err := svc.client.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}

	return svc, nil
}

func (s *service) Get(ctx context.Context, key string) (string, error) {
	v, err := s.client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return v, nil
}

func (s *service) Put(ctx context.Context, key string, value interface{}, expires time.Duration) error {
	if err := s.client.Set(ctx, key, value, expires).Err(); err != nil {
		return err
	}
	return nil
}
