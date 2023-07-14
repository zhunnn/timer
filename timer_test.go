package timer

import (
	"log"
	"testing"
	"time"
)

func TestPause(t *testing.T) {
	tr := NewTimer()

	go func() {
		time.Sleep(time.Second * 3)
		log.Println("timer pause")
		tr.Pause()
		time.Sleep(time.Second * 3)
		log.Println("timer pause")
		tr.Pause()
	}()

	log.Println("timer start")
	tr.Block(time.Second * 5)
	log.Println("done")
}

func TestRepeatBlock(t *testing.T) {
	tr := NewTimer()
	tr.Pause()
	go tr.Block(time.Second * 5)
	// time.Sleep(time.Second * 1)
	go tr.Block(time.Second * 5)
	// time.Sleep(time.Second * 1)
	go tr.Block(time.Second * 5)
	// time.Sleep(time.Second * 1)
	go tr.Block(time.Second * 5)

	time.Sleep(time.Second * 1)
	log.Println(tr.GeDuration())
	time.Sleep(time.Second * 1)
	log.Println(tr.GeDuration())

	time.Sleep(time.Second * 6)
}

func TestCancel(t *testing.T) {
	tr := NewTimer()
	go tr.Block(time.Second * 5)
	time.Sleep(time.Second * 3)
	tr.Cancel()
	time.Sleep(time.Second * 3)
}

func TestGetDuration(t *testing.T) {
	tr := NewTimer()
	go tr.Block(time.Second * 5)
	time.Sleep(time.Millisecond * 500)
	log.Println(tr.GeDuration())
	time.Sleep(time.Millisecond * 700)
	log.Println(tr.GeDuration())
	time.Sleep(time.Millisecond * 900)
	log.Println(tr.GeDuration())
	go tr.Block(time.Second * 5)
	time.Sleep(time.Millisecond * 500)
	log.Println(tr.GeDuration())
	time.Sleep(time.Millisecond * 700)
	log.Println(tr.GeDuration())
	time.Sleep(time.Millisecond * 900)
	log.Println(tr.GeDuration())
}
