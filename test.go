package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func myfunc() {
	fmt.Printf("---- %v\n", time.Now())
}

func executeAt(ctx context.Context, targetTime time.Time) {
	ctx, cancel := context.WithDeadline(ctx, targetTime)
	defer cancel()

	go func() {
		<-ctx.Done()
		if ctx.Err() == context.Canceled {
			log.Println("Выполнение прервано. cancel")
			return
		}
		// if ctx.Err() == context.DeadlineExceeded {
		// 	log.Println("Срок действия истек, функция не будет выполнена.")
		// 	return
		// }
		myfunc()
	}()
	<-ctx.Done()
}

func main() {
	now := time.Now()

	targetTime := time.Date(now.Year(), now.Month(), now.Day(), 17, 10, 0, 0, now.Location())

	if targetTime.Before(now) {
		log.Println("Функция вызвана немедленно")
		myfunc()
		return
	}

	baseCtx := context.Background()
	ctx, cancel := context.WithTimeout(baseCtx, 40*time.Second)
	defer cancel()

	go executeAt(ctx, targetTime)

	// go func() {
	// 	time.Sleep(3 * time.Second) // Задержка на 3 секунды
	// 	cancel()
	// }()

	// <-ctx.Done()
	// if ctx.Err() == context.DeadlineExceeded {
	// 	log.Println("Функция не выполняется, так как срок действия истек.")
	// }

	log.Println("start server")
	time.Sleep(time.Hour * 24)
}

///-------------------------мне не нрав
// package main

// import (
// 	"context"
// 	"fmt"
// 	"time"
// )

// func myfunc() {
// 	fmt.Printf("+ %v\n", time.Now())
// }

// func executeAt(ctx context.Context, targetTime time.Time) {
// 	ctx, cancel := context.WithDeadline(ctx, targetTime)
// 	defer cancel()

// 	// Запускаем горутину для вызова функции в целевое время
// 	go func() {
// 		<-ctx.Done()
// 		if ctx.Err() == context.Canceled {
// 			fmt.Println("Выполнение прервано.")
// 			return
// 		}
// 		myfunc()
// 	}()

// 	// Ожидаем завершения горутины
// 	<-ctx.Done()
// }

// func main() {
// 	now := time.Now()

// 	targetTime := time.Date(now.Year(), now.Month(), now.Day(), 16, 14, 0, 0, now.Location())

// 	if targetTime.Before(now) {
// 		fmt.Println("Функция вызвана немедленно")
// 		myfunc()
// 		return
// 	}

// 	// Создаем базовый контекст
// 	baseCtx := context.Background()

// 	// Создаем контекст с таймаутом на 3 секунды
// 	ctx, cancel := context.WithTimeout(baseCtx, 3*time.Second)
// 	defer cancel()

// 	// Запускаем выполнение
// 	go executeAt(ctx, targetTime)

// 	// Ждем завершения выполнения
// 	<-ctx.Done()
// 	if ctx.Err() == context.DeadlineExceeded {
// 		fmt.Println("Время ожидания истекло.")
// 	}
// 	time.Sleep(time.Hour * 24)
// }

//////////////////////////// РАБОЧИЙ
// package main

// import (
// 	"context"
// 	"fmt"
// 	"time"
// )

// func myfunc(str string) {
// 	fmt.Printf("%w + %v\n", str, time.Now())
// }

// func executeAt(targetTime time.Time, str string) {
// 	ctx, cancel := context.WithDeadline(context.Background(), targetTime)
// 	defer cancel()

// 	go func() {
// 		<-ctx.Done()
// 		myfunc(str)
// 	}()

// 	<-ctx.Done()
// }

// func main() {
// 	// Получаем текущее время
// 	now := time.Now()

// 	targetTime1 := time.Date(now.Year(), now.Month(), now.Day(), 16, 7, 0, 0, now.Location())
// 	targetTime2 := time.Date(now.Year(), now.Month(), now.Day(), 16, 8, 0, 0, now.Location())

// 	// if targetTime.Before(now) {
// 	// 	fmt.Println("Функция вызвана немедленно")
// 	// 	myfunc()
// 	// 	return
// 	// }

// 	// Вызываем функцию для выполнения в целевое время
// 	executeAt(targetTime1, "1")
// 	executeAt(targetTime2, "2")
// 	time.Sleep(time.Hour * 24)
// }

//-----------------------------------------------
// package main

// import (
// 	"fmt"
// 	"time"
// )

// // Описание формата времени.
// const timeLayout = "Aug 21 2006 15:04:05 MST"

// // Вызов переданной функции в указанное время.
// func callAt(callTime string, f func()) error {
// 	// Разбираем время запуска.
// 	ctime, err := time.Parse(timeLayout, callTime)
// 	if err != nil {
// 		return err
// 	}

// 	// Вычисляем временной промежуток до запуска.
// 	duration := ctime.Sub(time.Now())

// 	go func() {
// 		time.Sleep(duration)
// 		f()
// 	}()

// 	return nil
// }

// // Ваша функция.
// func myfunc() {
// 	fmt.Printf("+ %v\n", time.Now())
// }

// // Пример использования.
// func main() {
// 	err := callAt("Aug 21 2024 15:45:00 MSK", myfunc)
// 	if err != nil {
// 		fmt.Printf("error: %v\n", err)
// 	}

// 	// Эмуляция дальнейшей работы программы.
// 	time.Sleep(time.Hour * 24)
// }
// -------------------------------------------------------------
// package main

// import (
// 	"context"
// 	"fmt"
// 	"time"
// )

// // Функция, которая принимает время и сообщение
// func scheduleFunction(at time.Time, message string) {
// 	// Вычисляем продолжительность до вызова
// 	duration := at.Sub(time.Now())

// 	fmt.Println(duration)

// 	if duration < 0 {
// 		fmt.Println("Указанное время уже прошло:", message)
// 		return
// 	}

// 	// Создаем контекст с таймаутом
// 	ctx, cancel := context.WithTimeout(context.Background(), duration)
// 	defer cancel()

// 	go func() {
// 		<-ctx.Done() // Ожидание завершения контекста
// 		if ctx.Err() == context.DeadlineExceeded {
// 			fmt.Println("Сообщение:", message, "вызвано в:", time.Now())
// 		}
// 	}()
// }

// func main() {
// 	// Пример использования функции
// 	scheduledTime := time.Date(2024, 8, 21, 15, 36, 0, 0, time.Local)
// 	message := "Это запланированное сообщение"
// 	fmt.Println(scheduledTime)
// 	scheduleFunction(scheduledTime, message)

// 	// Ожидание, чтобы программа не завершилась сразу
// 	time.Sleep(10 * time.Minute) // Убедитесь, что программа работает достаточно долго
// }
