package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var cache = map[int]Book{}
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	wg := &sync.WaitGroup{}
	m := &sync.Mutex{} // use pointers to pass copies and not values
	for i := 0; i < 10; i++ {
		id := rnd.Intn(10) + 1
		fmt.Println("Here is id: ", id)
		wg.Add(2) // we could put wg.Add(1) before each one
		go func (idArg int, wgArg *sync.WaitGroup, mArg *sync.Mutex) {
			if b, ok := queryCache(idArg, mArg); ok {
				fmt.Println("From cache: id: ", id)
				fmt.Println(b)
				// continue
			}
			wgArg.Done()
		}(id, wg, m)
		go func (idArg int, wgArg *sync.WaitGroup, mArg *sync.Mutex) {
			if b, ok := queryDatabase(idArg, mArg); ok {
				fmt.Println("From database: id: ", id)
				fmt.Println(b)
				// continue
			}
			wgArg.Done()
		}(id, wg, m)
		// fmt.Println("No Book found with ID: ", id)
		time.Sleep(150 * time.Millisecond)
		
	}
	// time.Sleep(2 * time.Second) // to make sure all go routines run
	wg.Wait()
}

func queryCache(id int, mArg *sync.Mutex) (Book, bool) {
	mArg.Lock()
	b, ok := cache[id]
	mArg.Unlock()
	return b, ok
}

func queryDatabase(id int, mArg *sync.Mutex) (Book, bool) {
	time.Sleep(100 * time.Millisecond)
	for _, b := range books {
		if b.ID == id {
			mArg.Lock()
			cache[id] = b
			mArg.Unlock()
			return b, true
		}
	}
	return Book{}, false
}


