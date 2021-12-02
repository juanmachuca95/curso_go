package main

import (
	"fmt"
	"sync"
)

func main() {
	whoMe := []string{}

	wg := &sync.WaitGroup{}
	var mutex sync.Mutex
	wg.Add(4)
	go func(wg *sync.WaitGroup) {
		fmt.Println("One")
		mutex.Lock()
		whoMe = append(whoMe, "Juan")
		mutex.Unlock()
		wg.Done()
	}(wg)
	go func(wg *sync.WaitGroup) {
		fmt.Println("Two")
		mutex.Lock()
		whoMe = append(whoMe, "Martin")
		mutex.Unlock()
		wg.Done()
	}(wg)
	go func(wg *sync.WaitGroup) {
		fmt.Println("Three")
		mutex.Lock()
		whoMe = append(whoMe, "Gretel")
		mutex.Unlock()
		wg.Done()
	}(wg)
	go func(wg *sync.WaitGroup) {
		fmt.Println("Four")
		mutex.Lock()
		whoMe = append(whoMe, "Maxi")
		mutex.Unlock()
		wg.Done()
	}(wg)

	wg.Wait()
	// Secuencial orden Juan - Martin - Gretel - Maxi
	fmt.Println(whoMe)
}
