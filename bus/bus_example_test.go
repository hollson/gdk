package bus

import (
	"fmt"
	"sync"
)

func Example() {
	queueSize := 100
	_bus := New(queueSize)

	var wg sync.WaitGroup
	wg.Add(2)

	_ = _bus.Subscribe("topic", func(v bool) {
		defer wg.Done()
		fmt.Println("s1", v)
	})

	_ = _bus.Subscribe("topic", func(v bool) {
		defer wg.Done()
		fmt.Println("s2", v)
	})

	// Publish block only when the buffer of one of the subscribers is full.
	// change the buffer size altering queueSize when creating new messagebus
	_bus.Publish("topic", true)
	wg.Wait()

	// Unordered output:
	// s1 true
	// s2 true
}

func Example_second() {
	queueSize := 2
	subscribersAmount := 3

	ch := make(chan int, queueSize)
	defer close(ch)

	bus := New(queueSize)

	for i := 0; i < subscribersAmount; i++ {
		_ = bus.Subscribe("topic", func(i int, out chan<- int) { out <- i })
	}

	go func() {
		for n := 0; n < queueSize; n++ {
			bus.Publish("topic", n, ch)
		}
	}()

	var sum = 0
	for sum < (subscribersAmount * queueSize) {
		select {
		case <-ch:
			sum++
		}
	}

	fmt.Println(sum)
	// Output:
	// 6
}
