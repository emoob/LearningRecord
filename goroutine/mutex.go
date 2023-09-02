package goroutine

import (
	"sync"
	"sync/atomic"
)

/*
	SourceGo :语言高级编程
在 worker 的循环中，为了保证 total.value += i 的原子性，
我们通过 sync.Mutex 加锁和解锁来保证该语句在同一时刻只被一个线程访问。
对于多线程模型的程序而言，进出临界区前后进行加锁和解锁都是必须的。
如果没有锁的保护，total 的最终值将由于多线程之间的竞争而可能会不正确。
*/

var total struct {
	sync.Mutex
	value int
}

func worker(w *sync.WaitGroup) {
	defer w.Done()
	for i := 0; i < 100; i++ {
		// 上锁
		total.Lock()
		total.value += 1
		// 解锁
		total.Unlock()
	}
}

/*
用互斥锁来保护一个数值型的共享资源，麻烦且效率低下。
标准库的 sync/atomic 包对原子操作提供了丰富的支持
- 原子操作是指一个操作要么全部执行完毕，要么完全不执行，
中间不会被其他线程或进程打断
*/

/*
atomic.AddUint64 函数调用保证了 total2 的读取、更新和保存是一个原子操作，因此在多线程中访问也是安全的。
原子操作配合互斥锁可以实现非常高效的单件模式。互斥锁的代价比普通整数的原子读写高很多，
在性能敏感的地方可以增加一个数字型的标志位，
通过原子检测标志位状态降低互斥锁的使用次数来提高性能。
*/
var total2 uint64

func worker2(wg *sync.WaitGroup) {
	defer wg.Done()
	var i uint64
	for i = 0; i <= 100; i++ {
		atomic.AddUint64(&total2, i)
	}
}
