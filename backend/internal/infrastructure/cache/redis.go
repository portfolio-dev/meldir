package cache

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
	"meldir-backend/internal/config"
)

type RedisClient struct {
	Client *redis.Client
}

func NewRedisClient(cfg *config.Config) (*RedisClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rdb := redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort),
		Password:     "", // Kosong jika tanpa password lokal
		DB:           0,
		DialTimeout:  5 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	})

	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("gagal melakukan ping Redis (%s:%s): %w", cfg.RedisHost, cfg.RedisPort, err)
	}

	log.Printf("⚡ Terhubung sukses ke Redis Server (%s:%s)", cfg.RedisHost, cfg.RedisPort)
	return &RedisClient{Client: rdb}, nil
}

func (r *RedisClient) BlacklistToken(ctx context.Context, tokenString string, ttl time.Duration) error {
	key := fmt.Sprintf("blacklist:token:%s", tokenString)
	return r.Client.Set(ctx, key, "revoked", ttl).Err()
}

func (r *RedisClient) IsTokenBlacklisted(ctx context.Context, tokenString string) (bool, error) {
	key := fmt.Sprintf("blacklist:token:%s", tokenString)
	exists, err := r.Client.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}
	return exists > 0, nil
}

func (r *RedisClient) Close() {
	if r.Client != nil {
		_ = r.Client.Close()
	}
}
