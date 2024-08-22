package main

import (
	"context"
	"log"
	"sync"
	"time"
)

var (
	targetTimes []time.Time
	mutex       sync.Mutex
)

func fn() {
	log.Printf("Время %s достигнуто. Завершение программы.\n", time.Now().Format("00:00"))
}

func checkTargetTime() {
	mutex.Lock()
	defer mutex.Unlock()

	currentTime := time.Now()
	for _, targetTime := range targetTimes {
		if currentTime.After(targetTime) && currentTime.Before(targetTime.Add(1*time.Second)) {
			fn()
			return
		}
	}
}

func initWaitTask(ctx context.Context) {
	// Когда вызывается cancel(), канал ctx.Done() закрывается, и управление передается в соответствующий case в select.
	// она не будет продолжать выполнять checkTargetTime() или другие операции.
	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Println("Горутина завершена.")
				return
			default:
				checkTargetTime()
				log.Printf("ждем....")
				time.Sleep(1 * time.Second)
			}
		}
	}()
}

func addTargetTime(newTime time.Time) {
	mutex.Lock()
	defer mutex.Unlock()

	targetTimes = append(targetTimes, newTime)
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	now := time.Now()
	targetTimes = []time.Time{
		time.Date(now.Year(), now.Month(), now.Day(), 11, 44, 0, 0, now.Location()),
		time.Date(now.Year(), now.Month(), now.Day(), 11, 44, 10, 0, now.Location()),
	}

	initWaitTask(ctx)

	newTime := time.Date(now.Year(), now.Month(), now.Day(), 11, 44, 13, 0, now.Location())
	addTargetTime(newTime)

	time.Sleep(time.Hour * 24)
}
