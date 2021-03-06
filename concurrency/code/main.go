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
	m := &sync.RWMutex{} // use pointers to pass copies and not values
	cacheCh := make(chan Book)
	dbCh := make(chan Book)
	for i := 0; i < 10; i++ {
		id := rnd.Intn(10) + 1
		fmt.Println("Here is id: ", id)
		wg.Add(2) // we could put wg.Add(1) before each one
		go func (idArg int, wgArg *sync.WaitGroup, mArg *sync.RWMutex, ch chan<- Book) {
			if b, ok := queryCache(idArg, mArg); ok {
				ch <- b
				fmt.Println("From cache: id: ", id)
				fmt.Println(b)
				// continue
			}
			wgArg.Done()
		}(id, wg, m, cacheCh)
		go func (idArg int, wgArg *sync.WaitGroup, mArg *sync.RWMutex, ch chan<- Book) {
			if b, ok := queryDatabase(idArg, mArg); ok {
				fmt.Println("From database: id: ", id)
				fmt.Println(b)
				ch <- b
				// continue
			}
			wgArg.Done()
		}(id, wg, m, dbCh)
		// fmt.Println("No Book found with ID: ", id)
		time.Sleep(150 * time.Millisecond)
		
		go func(cacheCh, dbCh <-chan Book) {
			select {
				case b := <-cacheCh:
				fmt.Println("from cache: ", b)
				<-dbCh
				case b := <-dbCh:
				fmt.Println("from DB: ", b)
			}
		}(cacheCh, dbCh)
		
	} // for loop
	time.Sleep(2 * time.Second) // to make sure all go routines run
	wg.Wait()
	close(cacheCh)
	close(dbCh)

}

func queryCache(id int, mArg *sync.RWMutex) (Book, bool) {
	mArg.RLock()
	b, ok := cache[id]
	mArg.RUnlock()
	return b, ok
}

func queryDatabase(id int, mArg *sync.RWMutex) (Book, bool) {
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


