package main

import (
	"github.com/go-vgo/robotgo"
	"log"
	"time"
)

func main() {
	// 获取当前鼠标所在的位置
	x1, y1 := robotgo.GetMousePos()

	postman := make(chan int)

	ticker := time.NewTicker(time.Second * 5)
	go func() {
		for range ticker.C {
			x2, y2 := robotgo.GetMousePos()
			if x1 == x2 && y1 == y2 {
				log.Println(time.Now().Format("2006-01-02 15:04:05"), ": timeout")
				postman <- 1
			}
			x1 = x2
			y1 = y2
		}
	}()

	w, _ := robotgo.GetScreenSize()

	go func(postman <-chan int) {
		for v := range postman {
			if v == 1 {
				robotgo.Move(w, 0)
				robotgo.MoveSmooth(w, 100)
				robotgo.MoveSmooth(w-100, 100)
				robotgo.MoveSmooth(w-100, 0)
			}
		}
	}(postman)

	for {
		mleft := robotgo.AddEvent("mleft")
		if mleft {
			log.Println(time.Now().Format("2006-01-02 15:04:05"), ": click")
			postman <- 0
		}
	}
}
