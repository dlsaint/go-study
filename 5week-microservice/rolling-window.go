package _week_microservice

import (
	"sync"
	"time"
)

type Bucket struct {
	mu sync.RWMutex
	// 请求总数
	Total int
	// 失败总数
	Failed    int
	Timestamp time.Time
}

func NewBucket() *Bucket {
	return &Bucket{
		Timestamp: time.Now(),
	}
}

// 记录请求数量
func (b *Bucket) Record(result bool) {
	b.mu.Lock()
	defer b.mu.Unlock()
	if !result {
		b.Failed++
	}
	b.Total++
}

type RollingWindow struct {
	mu sync.RWMutex
	// 是否触发熔断
	broken bool
	// 滑动窗口大小
	size int
	// 桶队列
	buckets []*Bucket
	// 触发熔断的请求总数阈值
	reqThreshold int
	// 出发熔断的失败率阈值
	failedThreshold float64
	// 上次熔断发生时间
	lastBreakTime time.Time
	// 熔断恢复的时间间隔
	brokeTimeGap time.Duration
}

// 新建滑动窗口
func NewRollingWindow(
	size int,
	reqThreshold int,
	failedThreshold float64,
	brokeTimeGap time.Duration,
) *RollingWindow {
	return &RollingWindow{
		size:            size,
		buckets:         make([]*Bucket, 0, size),
		reqThreshold:    reqThreshold,
		failedThreshold: failedThreshold,
		brokeTimeGap:    brokeTimeGap,
	}
}

// 启动滑动窗口, 按时间间隔追加桶, 每100 Millisecond, 并且开启监控
func (r *RollingWindow) Launch() {
	go func() {
		for {
			r.appendBucket()
			time.Sleep(time.Millisecond * 100)
		}
	}()

	go func() {
		for {
			if r.broken {
				if r.isOverBrokenTimeGap() {
					r.mu.Lock()
					r.broken = false
					r.mu.Unlock()
				}
				continue
			}
			if r.breakJudgement() {
				r.mu.Lock()
				r.broken = true
				r.lastBreakTime = time.Now()
				r.mu.Unlock()
			}
		}
	}()
}

// 在桶中记录当次结果
func (r *RollingWindow) RecordReqResult(result bool) {
	if len(r.buckets) == 0 {
		r.appendBucket()
	}
	r.buckets[len(r.buckets)-1].Record(result)
}

// 当前熔断状态
func (r *RollingWindow) Broken() bool {
	return r.broken
}

// 追加一个新桶
func (r *RollingWindow) appendBucket() {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.buckets = append(r.buckets, NewBucket())
	//固定队列不大于窗口大小
	if !(len(r.buckets) < r.size+1) {
		r.buckets = r.buckets[1:]
	}
}

// 当前滑动窗口判断是否需要触发熔断
func (r *RollingWindow) breakJudgement() bool {
	r.mu.RLock()
	defer r.mu.RUnlock()
	total := 0
	failed := 0
	for _, v := range r.buckets {
		total += v.Total
		failed += v.Failed
	}
	if float64(failed)/float64(total) > r.failedThreshold && total > r.reqThreshold {
		return true
	}
	return false
}

// 是否超过熔断间隔期
func (r *RollingWindow) isOverBrokenTimeGap() bool {
	return time.Since(r.lastBreakTime) > r.brokeTimeGap
}
