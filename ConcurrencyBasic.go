/* Concurrency

-- What is concurrency -- 

Concurrency is about creating multiple processes that execute *independently*
 
 concurrency is the composition of independtly executing processes, while parallelism
 is the simultaneous execution of (possibly related) computations. Concurrency is
 about dealing with lots of things at once.

Go doesnt use threads for its concurrency ...

--- goroutines ---

The concurrency happens with the help of goroutines
they are scheduled by the golang itself

goroutines vs OSThreads

- Lighter weight
- Go manages goroutines(simpler for programmers)
- Less switching

Go's concurrency model

Actor Model
Communication Sequential processes
(CSP)

channel */

package main
import(
	"fmt"
	"time"
	"sync"
	"runtime"
)
func main(){
	runtime.GOMAXPROCS(2)
	var waitGrp sync.WaitGroup
	waitGrp.Add(2)
	go func(){
		defer waitGrp.Done()
		time.Sleep(5* time.Second)
		fmt.Println("Hello")
	}()

	go func(){
		defer waitGrp.Done()
		fmt.Println("Pluralsight")
	}()

	waitGrp.Wait()
}

// channels can be either buffered or unbuffered
// We dont specify size for the number of channels
// Goroutine needs another goroutine to carry on a process in unbuffered channel making things synchronized
// Goroutine can send any amount of data as long as buffer is free creating more async behaviour 