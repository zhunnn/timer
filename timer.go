package timer

import (
	"sync"
	"time"
)

type Timer struct {
	isPause      bool
	isCancel     bool
	chCancel     chan struct{}
	mu           *sync.Mutex
	intervalTime time.Duration
}

func NewTimer() *Timer {
	t := &Timer{
		isPause:      false,
		isCancel:     false,
		chCancel:     make(chan struct{}, 1),
		mu:           &sync.Mutex{},
		intervalTime: time.Millisecond,
	}
	return t
}

// 開啟倒計時器，輸入計時時間
func (t *Timer) Block(duration time.Duration) {
	// 強制關閉同計時器的其他線程
	t.mu.Lock()
	// 建立自己的 channel
	chCancel := make(chan struct{}, 1)
	// 傳送刪除訊號給上一個 channel
	select {
	case t.chCancel <- struct{}{}:
	default:
	}
	// 替換自己的 channel
	t.chCancel = chCancel
	// 移除暫停
	t.isPause = false
	t.mu.Unlock()

	// 每毫秒倒計時
	// log.Println("start")
	tr := time.NewTicker(t.intervalTime)
	defer tr.Stop()

	// 毫秒迴圈
	for {
		select {
		case <-tr.C:
			// 取消
			if t.isCancel {
				// log.Println("tr cancel by command")
				return
			}
			// 暫停
			if t.isPause {
				// log.Println("tr pause")
				continue
			}
			// 結束
			if duration <= 0 {
				// log.Println("done")
				return
			}
			// log.Println("tr")
			duration -= t.intervalTime
		case <-chCancel:
			// log.Println("tr be canceled by another tr")
			return
		}
	}
}

// 設定間隔時間
func (t *Timer) SetIntervalTime(it time.Duration) {
	t.intervalTime = it
}

// 暫停計時器
func (t *Timer) Pause() {
	t.isPause = !t.isPause
}

// 刪除計時器
func (t *Timer) Cancel() {
	t.isCancel = true
}
