package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type chopStick struct {
	sync.Mutex
}

type philosopher struct {
	name   string
	rstick *chopStick
	lstick *chopStick
}

func (p philosopher) eat() {
	for j := 0; j < 3; j++ {
		p.lstick.Lock()
		p.rstick.Lock()
		fmt.Printf("%s: Starting to eat\n", p.name)
		fmt.Println("eating")
		fmt.Println("eating")
		fmt.Println("eating")
		fmt.Printf("%s: Finished eating\n", p.name)
		p.lstick.Unlock()
		p.rstick.Unlock()
	}
	defer wg.Done()
}

func main() {
	wg.Add(5)
	csticks := make([]*chopStick, 5)
	names := []string{"Aifread", "Yuri", "Flynn", "Regal", "Jon"}
	for i := 0; i < 5; i++ {
		csticks[i] = new(chopStick)
	}

	philos := make([]*philosopher, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &philosopher{
			lstick: csticks[i],
			rstick: csticks[(i+1)%5],
			name:   names[i],
		}
	}
	for _, val := range philos {
		fmt.Println(*val)
	}

	for i := 0; i < 5; i++ {
		go philos[i].eat()
	}
	wg.Wait()
}
