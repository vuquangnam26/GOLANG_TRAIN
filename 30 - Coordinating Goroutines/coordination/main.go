package main

import (
	"math"
	"math/rand"
	"sync"
	"time"

	// "math"
	// "math/rand"
	"context"
)

const (
	countKey = iota
	sleepPeriodKey
)

func processRequest(ctx context.Context, wg *sync.WaitGroup) {
	total := 0
	count := ctx.Value(countKey).(int)
	sleepPeriod := ctx.Value(sleepPeriodKey).(time.Duration)
	for i := 0; i < count; i++ {
		select {
		case <-ctx.Done():
			if ctx.Err() == context.Canceled {
				Printfln("Stopping processing - request cancelled")
			} else {
				Printfln("Stopping processing - deadline reached")
			}
			goto end
		default:
			Printfln("Processing request: %v", total)
			total++
			time.Sleep(sleepPeriod)
		}
	}
	Printfln("Request processed...%v", total)
end:
	wg.Done()
}

var waitGroup = sync.WaitGroup{}
var rwmutex = sync.RWMutex{}
var readCond = sync.NewCond(rwmutex.RLocker())
var squares = map[int]int{}

func generateSquares(max int) {
	rwmutex.Lock()
	Printfln("Genarating data ....")
	for val := 0; val < max; val++ {
		squares[val] = int(math.Pow(float64(val), 2))
	}
	rwmutex.Unlock()
	Printfln("Broadcasting conditon")
	readCond.Broadcast()
	waitGroup.Done()
}
func readSquares(id, max, iterations int) {
	readCond.L.Lock()
	for len(squares) == 0 {
		readCond.Wait()
	}
	for i := 0; i < iterations; i++ {

		key := rand.Intn(max)
		Printfln("#%v Read value: %v = %v", id, key, squares[key])
		time.Sleep(time.Millisecond * 100)
	}
	readCond.L.Unlock()
	waitGroup.Done()
}
func main() {
	rand.Seed(time.Now().UnixNano())
	 numRoutines := 2 
	 waitGroup.Add(numRoutines) 
	 for i := 0; i < numRoutines; i++ {
		 go readSquares(i, 10, 5) }
		  waitGroup.Add(1)
		   go generateSquares(10)
		    waitGroup.Wait() 
			Printfln("Cached values: %v", len(squares))

	// waitGroup := sync.WaitGroup{}
	// waitGroup.Add(1)
	// Printfln("Request dispatched...")
	// ctx, _ := context.WithTimeout(context.Background(), time.Second*2)
	// ctx = context.WithValue(ctx, countKey, 4)
	// ctx = context.WithValue(ctx, sleepPeriodKey, time.Millisecond*250)
	// go processRequest(ctx, &waitGroup)

	// time.Sleep(time.Second)
	// Printfln("Canceling request")
	// cancel()

	waitGroup.Wait()
}
