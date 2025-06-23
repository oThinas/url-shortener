package store

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type store struct {
	rdb *redis.Client
}

type Store interface {
	SaveShortenedURL(ctx context.Context, shortenedURL string) (string, error)
	GetFullURL(ctx context.Context, code string) (string, error)
}

func NewStore(rdb *redis.Client) Store {
	return store{rdb}
}

func (s store) SaveShortenedURL(ctx context.Context, shortenedURL string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var code string
loop:
	for {
		select {
		case <-ctx.Done():
			return "", fmt.Errorf("timeout while generating unique code: %w", ctx.Err())
		default:
			code = generateCode()

			if err := s.rdb.HGet(ctx, "shortener", code).Err(); err != nil {
				if errors.Is(err, redis.Nil) {
					break loop
				}
				return "", fmt.Errorf("failed to get code from shortener hash: %w", err)
			}
		}
	}

	if err := s.rdb.HSet(ctx, "shortener", code, shortenedURL).Err(); err != nil {
		return "", fmt.Errorf("failed to set code in shortener hash: %w", err)
	}

	return code, nil
}

func (s store) GetFullURL(ctx context.Context, code string) (string, error) {
	url, err := s.rdb.HGet(ctx, "shortener", code).Result()
	if err != nil {
		return "", fmt.Errorf("failed to get code from shortener hash: %w", err)
	}

	return url, nil
}
