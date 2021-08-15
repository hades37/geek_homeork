package main

import (
	"sync"
	"time"
)

type bucketValue struct {
	Value int64
}

// 计数结构
type Counter struct {
	// 存放时间窗口内的访问信息,key为对应时间戳
	Buckets map[int64]*bucketValue
	Mutex   *sync.RWMutex
	// 窗口大小
	Window int64
}

func NewCounter() *Counter {
	c := &Counter{
		Buckets: make(map[int64]*bucketValue),
		Mutex:   &sync.RWMutex{},
	}
	return c
}

// 获取当前时间bucket
func (c *Counter) getCurrentBucket() *bucketValue {
	now := time.Now().Unix()
	var bucket *bucketValue
	var ok bool
	// 当前时间的bucket不存在则插入一个
	if bucket, ok = c.Buckets[now]; !ok {
		c.Buckets[now] = &bucketValue{}
	}

	return bucket
}

// 移动窗口淘汰删除旧数据
func (c *Counter) removeOldBuckets() {
	// 淘汰当前时间-window length之前的所有buket
	oldTime := time.Now().Unix() - c.Window
	for timeStamp := range c.Buckets {
		if timeStamp < oldTime {
			delete(c.Buckets, timeStamp)
		}
	}
}

// 增加bucket计数
func (c *Counter) AddCount(num int64) {
	if num == 0 {
		return
	}
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	c.getCurrentBucket().Value += 1
	c.removeOldBuckets()
}

// 遍历Buckets即可得到当前计数器的均值,极值以及某个时间点的占比等信息
