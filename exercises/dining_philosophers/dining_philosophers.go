package main

import (
	"fmt"
	"sync"
)

type stick struct {
	number int
	mutex  sync.Mutex
}

type philosopher struct {
	sticks                 []*stick
	number, completedMeals int
}

func (p *philosopher) eat(h *host, wg *sync.WaitGroup) {
	for p.completedMeals < 3 {
		if !(h.startEating(p)) {
			continue
		}
		h.finishEating(p)
	}
	wg.Done()
}

type host struct {
	philosophersEating []*philosopher
	mutex              sync.Mutex
}

func (h *host) startEating(p *philosopher) bool {
	h.mutex.Lock()
	if len(h.philosophersEating) == 2 {
		h.mutex.Unlock()
		return false
	}
	h.philosophersEating = append(h.philosophersEating, p)
	h.mutex.Unlock()
	p.sticks[0].mutex.Lock()
	p.sticks[1].mutex.Lock()
	fmt.Println("Starting to eat ", p.number)
	return true
}

func (h *host) finishEating(p *philosopher) {
	p.completedMeals++
	fmt.Println("Finishing eating ", p.number)
	h.mutex.Lock()
	for i, p2 := range h.philosophersEating {
		if p.number == p2.number {
			h.philosophersEating = remove(h.philosophersEating, i)
		}
	}
	h.mutex.Unlock()
	p.sticks[0].mutex.Unlock()
	p.sticks[1].mutex.Unlock()
}

func main() {
	var wg sync.WaitGroup
	h := host{philosophersEating: []*philosopher{}}
	philosophers := generatePhilosophers()
	wg.Add(len(philosophers))
	for i := range philosophers {
		go philosophers[i].eat(&h, &wg)
	}
	wg.Wait()
	fmt.Println("All meals completed!")
}

func generatePhilosophers() (ps []philosopher) {
	var sticks []*stick
	for i := 0; i < 5; i++ {
		s := stick{number: i + 1}
		p := philosopher{number: i + 1, sticks: []*stick{&s}}
		ps = append(ps, p)
		sticks = append(sticks, &s)
		if i > 0 {
			ps[i-1].sticks = append(ps[i-1].sticks, &s)
		}
	}
	ps[4].sticks = append(ps[4].sticks, sticks[0])
	return
}

func remove(s []*philosopher, i int) []*philosopher {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
