package syncs

import "sync"

type safeResource struct {
	// 一般用法是将Mutex，或者RWMutex和需要被保住的资源封装在一个结构体内
	// 使用锁时优先使用RWMutex
	resource interface{}
	lock     sync.Mutex
	// 使用RWMuTex 实现Double-check
	// 1.  加锁先检查一遍
	// 2. 释放读锁
	// 3. 加写锁
	// 4. 再检查一遍
}

func (s *safeResource) DoSomethingToResource() {
	s.lock.Lock()
	defer s.lock.Unlock()
}
