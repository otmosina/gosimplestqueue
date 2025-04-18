package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/otmosina/gosimplestqueue/queue"
)

type Point struct {
	X, Y int
}

type LazyPrinter struct {
	ID int
}

func NewLazyPrinter(id int) *LazyPrinter {
	return &LazyPrinter{
		ID: id,
	}
}

func (l *LazyPrinter) Exec() error {
	time.Sleep(time.Duration(700+rand.Intn(500)) * time.Millisecond)
	fmt.Printf("LazyPrinter ID %d finished\n", l.ID)
	return nil
}

func main() {
	var lazyPrinters map[int]*LazyPrinter = map[int]*LazyPrinter{
		0: NewLazyPrinter(0),
		1: NewLazyPrinter(1),
		2: NewLazyPrinter(2),
	}

	q := queue.New()

	for i := 0; i < 9; i++ {
		go func() {
			q.Add(lazyPrinters[i%3], time.Now().Add(time.Duration((i%3)+i)*time.Second))
		}()
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	for {
		sig := <-sigs
		fmt.Println("получен сигнал выхода:", sig)
		break
	}
}
