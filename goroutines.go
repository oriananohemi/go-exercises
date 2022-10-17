package goroutines

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	fmt.Println("Hellos")
	wg.Add(1)

	go say("world", &wg)

	wg.Wait()
}
