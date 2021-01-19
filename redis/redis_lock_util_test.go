package redis

import (
	"context"
	"github.com/google/uuid"
	"testing"
	"time"
)

func TestLock(t *testing.T) {
	unique := uuid.New().String()
	lock := Lock(context.Background(), "lock_key", unique, 5*time.Second)
	if lock {
		t.Log("加锁成功，开始处理业务")

		freeLock, err := FreeLock(context.Background(), "lock_key", unique)
		if err != nil {
			t.Error("error", err)
		}
		if freeLock {
			t.Log("解锁成功")
		}
	}
}

func TestFreeLock(t *testing.T) {
}
