package timer

import (
	"sync"
	"time"
)

type Timer struct {
	isPause   bool
	isCancel  *bool
	mu        *sync.Mutex
	frequency time.Duration
	duration  time.Duration
}

func NewTimer() *Timer {
	t := &Timer{
		isPause:   false,
		mu:        &sync.Mutex{},
		isCancel:  new(bool),
		frequency: time.Millisecond,
	}
	return t
}

// 開啟倒計時器，輸入計時時間
func (t *Timer) Block(duration time.Duration) {
	// 強制關閉同計時器的其他線程
	t.mu.Lock()
	// 關閉上個 start
	*t.isCancel = true
	// 替換自己的 cancel
	myCancel := new(bool)
	t.isCancel = myCancel
	// 移除暫停
	t.isPause = false
	t.duration = duration
	t.mu.Unlock()

	// 每毫秒倒計時
	// log.Println("start")
	tr := time.NewTicker(t.frequency)
	defer tr.Stop()

	// 毫秒迴圈
	for range tr.C {
		// 取消
		if *myCancel {
			// log.Println("tr be cancel")
			return
		}
		// 暫停
		if t.isPause {
			// log.Println("tr pause")
			continue
		}
		// 結束
		if t.duration <= 0 {
			// log.Println("done")
			return
		}
		// log.Println("tr")
		t.duration -= t.frequency
	}
}

// 暫停計時器
func (t *Timer) Pause() {
	t.isPause = !t.isPause
}

// 刪除計時器
func (t *Timer) Cancel() {
	*t.isCancel = true
}

// 設定倒計時頻率
func (t *Timer) SetFrequency(fq time.Duration) {
	t.frequency = fq
}

// 取得剩餘時間
func (t *Timer) GeDuration() time.Duration {
	return t.duration
}
