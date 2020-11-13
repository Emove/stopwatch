package back

import (
	"fmt"
	"testing"
	"time"
)

func TestNewStopwatch(t *testing.T) {
	stopwatch := NewStopwatch()
	stopwatch.Start()
	fmt.Println(stopwatch.isRunning)
	time.Sleep(1 * time.Second)
	stopwatch.Stop()
	fmt.Println(stopwatch.Total(TimeUnit_Second))
	fmt.Println(stopwatch.Total(TimeUnit_Milliseconds))
	fmt.Println(stopwatch.Total(TimeUnit_Nanoseconds))
}

func TestStopwatch_Point(t *testing.T) {
	stopwatch := NewStopwatch()
	stopwatch.Start()
	fmt.Println(stopwatch.isRunning)
	// 模拟业务检测耗时
	//time.Sleep(15 * time.Millisecond)

	// 数据查询耗时埋点
	stopwatch.Begin("queryAccount")
	time.Sleep(100 * time.Millisecond)
	stopwatch.End("queryAccount")

	// 模拟业务处理
	//time.Sleep(60 * time.Millisecond)

	// 数据更新耗时埋点
	stopwatch.Begin("updateAccount")
	time.Sleep(25 * time.Millisecond)
	stopwatch.End("updateAccount")

	stopwatch.Stop()
	//fmt.Println(stopwatch.duration("queryAccount", TimeUnit_Nanoseconds))
	//fmt.Println(stopwatch.duration("updateAccount", TimeUnit_Nanoseconds))
	//fmt.Println(15 + 100 + 60 + 25)
	fmt.Println(stopwatch.Duration("updateAccount", TimeUnit_Milliseconds))
	fmt.Println(stopwatch.Total(TimeUnit_Milliseconds))
}

func BenchmarkNewStopwatch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stopwatch := NewStopwatch()
		stopwatch.Start()
		stopwatch.Begin("test")
		stopwatch.End("test")
		stopwatch.Stop()
		//stopwatch.Total(TimeUnit_Milliseconds)
	}
}
