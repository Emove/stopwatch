package back

import (
	"fmt"
	"strconv"
	"time"
)

// 秒表，计时工具，
// 使用该工具时，特别是做点记录
// 使用者应自己注意记录点的名称不要重复，该工具就不做是否存在和重复判断了
// 不然使用时还需要处理是否有error是很烦人的一件事
type Stopwatch struct {
	master    *Point
	isRunning bool
	points    map[string]*Point
}

//记录点
type Point struct {
	name    string
	startAt time.Time
	dur     int64
}

type TimeUnit uint8

const (
	// 秒
	TimeUnit_Second TimeUnit = 0
	// 毫秒
	TimeUnit_Milliseconds TimeUnit = 1
	// 纳秒
	TimeUnit_Nanoseconds TimeUnit = 2
)

func NewStopwatch() *Stopwatch {
	return &Stopwatch{
		master: &Point{
			name: "master",
		},
		isRunning: false,
		points:    make(map[string]*Point, 5),
	}
}

// 开始计时
func (s *Stopwatch) Start() {
	s.master.startAt = time.Now()
	s.isRunning = true
}

func (s *Stopwatch) Stop() {
	if !s.isRunning {
		return
	}
	s.master.dur = time.Since(s.master.startAt).Nanoseconds()
	s.isRunning = false

	// 处理记录点
	if nil == s.points {
		return
	}
	for _, point := range s.points {
		if 0 == point.dur {
			point.dur = s.master.dur - int64(point.startAt.Nanosecond()-s.master.startAt.Nanosecond())
		}
	}
}

//新建记录点
func (s *Stopwatch) Begin(name string) {
	_, have := s.points[name]
	if have {
		return
	}
	point := &Point{
		name:    name,
		startAt: time.Now(),
	}
	s.points[name] = point
}

func (s *Stopwatch) End(name string) {
	point, have := s.points[name]
	if !have {
		return
	}
	point.dur = time.Since(point.startAt).Nanoseconds()
}

func (s *Stopwatch) Duration(name string, timeUnit TimeUnit) string {
	point, have := s.points[name]
	if !have {
		return ""
	}
	if point.dur == 0 {
		point.dur = time.Since(point.startAt).Nanoseconds()
	}
	return convert(point.dur, timeUnit)
}

func (s *Stopwatch) Total(timeUnit TimeUnit) string {
	if s.isRunning {
		s.Stop()
	}
	result := fmt.Sprintf("total: %s", convert(s.master.dur, timeUnit))
	for _, point := range s.points {
		result += fmt.Sprintf(", %s: %s", point.name, convert(point.dur, timeUnit))
	}
	return result
}

func convert(duration int64, timeUnit TimeUnit) string {
	switch timeUnit {
	case TimeUnit_Second:
		return strconv.FormatInt(duration/1e9, 10) + "s"
	case TimeUnit_Milliseconds:
		return strconv.FormatInt(duration/1e6, 10) + "ms"
	case TimeUnit_Nanoseconds:
		return strconv.FormatInt(duration, 10) + "ns"
	}
	return ""
}
