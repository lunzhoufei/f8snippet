package cat

import (
	"sync"
	"sync/atomic"
	"time"
)

// Limiter 限频器接口
type Limiter interface {
	Acquire() bool
}

// ============================================================================
// LeakyBucketLimiter 漏桶算法限频器 秒级控制速率
// ============================================================================

type LeakyBucketLimiter struct {
	lastAccessTime time.Time
	capacity       int64
	avail          int64
	mu             sync.Mutex
}

// Acquire 漏桶算法限频器实现，以漏桶原理比喻算法，想象一下，有一个源源不断的水桶正在以capacity per second（capacity/e9 per nanasecond）的速度往下滴水，下面有一个容量为capacity的小桶1在接水，满了就直接溢出抛弃。另外还有一个正在使用中的水桶2，每次请求来的时候，先从水桶2取出水滴，如果水桶2是空的，直接把水桶1全部倒入水桶2，水桶1继续接水。
func (l *LeakyBucketLimiter) Acquire() bool {
	if l == nil {
		return true
	}
	l.mu.Lock()
	if l.avail > 0 {
		l.avail-- //水桶2取出水滴
		l.mu.Unlock()
		return true
	}
	//水桶2空了
	now := time.Now()
	add := l.capacity * now.Sub(l.lastAccessTime).Nanoseconds() / 1e9 //在使用水桶2的过程中，水桶1已经接到的水量
	if add > 0 {
		l.lastAccessTime = now
		l.avail += add //水桶1全部倒入水桶2
		if l.avail > l.capacity {
			l.avail = l.capacity //溢出抛弃
		}
		l.avail--
		l.mu.Unlock()
		return true
	}
	l.mu.Unlock()
	return false
}

// NewLeakyBucketLimiter 新建一个容量为capacity的漏桶算法限频器
func NewLeakyBucketLimiter(capacity int64) Limiter {
	if capacity <= 0 {
		return nil
	}
	return &LeakyBucketLimiter{capacity: capacity, lastAccessTime: time.Now(), avail: capacity}
}

// ============================================================================
// TokenBucketLimiter 令牌桶算法限频器 秒级控制速率
// ============================================================================
type TokenBucketLimiter struct {
	capacity int64
	avail    int64
}

// Acquire 令牌桶算法限频器实现
func (l *TokenBucketLimiter) Acquire() bool {
	if l == nil {
		return true
	}
	if atomic.AddInt64(&l.avail, 1) > l.capacity {
		return false
	}

	return true
}

// NewTokenBucketLimiter 新建一个容量为capacity的令牌桶算法限频器
func NewTokenBucketLimiter(capacity int64) Limiter {
	if capacity <= 0 {
		return nil
	}
	l := &TokenBucketLimiter{capacity: capacity, avail: 0}
	go func() {
		for {
			time.Sleep(time.Second)
			atomic.StoreInt64(&l.avail, 0)
		}
	}()

	return l
}
