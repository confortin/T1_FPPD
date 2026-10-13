// Example code filo_seq.go - Sequential Dinning Philosophers
// Since is sequential, elapsed timne could be simply calculated by (roubnds*philosophers*(time_to_think+time_to_eat))
// By Cesar De Rose

package main

import (
     "fmt"
	 "sync"
     "time"
)

var cycles [5]int           // ciclos concluidos por filosofo
var eating int              // quantos filosofos estao comendo agora
var maxEating int           // maximo de filosofos comendo ao mesmo tempo
var totalWait time.Duration // soma dos tempos de espera pelos garfos
var maxWait time.Duration   // maior tempo de espera

const numPhilosophers = 5

var forks [numPhilosophers]sync.Mutex

var table sync.Mutex
var wg sync.WaitGroup

func Philosopher(id, rounds int) {
	defer wg.Done()

	left := id
	right := (id + 1) % numPhilosophers

	time_to_think := 200 
	time_to_eat := 300
 
	for iteration := 1; iteration <= rounds; iteration++ {
	// think
	for s := 1; s <= id*6; s++ {
		fmt.Printf(" ")
	}
	fmt.Printf(" T%d\n", iteration) // Thinking
	time.Sleep(time.Duration(time_to_think) * time.Millisecond)
 
	// Adquirir recursos
	waitStart := time.Now()

	table.Lock()

	forks[left].Lock()
	forks[right].Lock()

	wait := time.Since(waitStart)
	totalWait = totalWait + wait

	if wait > maxWait {
		maxWait = wait
	}
 
	// eat
	eating++                // NOVO
	if eating > maxEating { // NOVO
		maxEating = eating
	}
 
	for s := 1; s <= id*6; s++ {
		fmt.Printf(" ")
	}
	fmt.Printf(" E%d\n", iteration) // Eating
	time.Sleep(time.Duration(time_to_eat) * time.Millisecond)
 
	eating--     // NOVO
	cycles[id]++ // NOVO: mais um ciclo concluido

	// Liberar os garfos e depois o mutex global
	forks[right].Unlock()
	forks[left].Unlock()

	table.Unlock()
	}
}

func main() {
     philosophers := 5
     rounds := 20 //// mudamos para de 5 para 20

     fmt.Println("\n[P1]  [P2]  [P3]  [P4]  [P5]\n")
 
     // run philosophers

     start := time.Now()

	 wg.Add(philosophers)

	 for j := 0; j < philosophers; j++ {
		go Philosopher(j, rounds)
	}

	 wg.Wait()

	 elapsed := time.Since(start)

     fmt.Printf("\nV2 - Jantar com sincronizacao excessiva levou %s\n\n", elapsed)

	fmt.Println("Tempo total:", elapsed)
	fmt.Println("Tempo medio de espera:", totalWait/time.Duration(rounds*philosophers))
	fmt.Println("Maior tempo de espera:", maxWait)
	fmt.Println("Maximo de filosofos comendo juntos:", maxEating)
	for i := 0; i < philosophers; i++ {
		fmt.Printf("Ciclos concluidos P%d: %d\n", i+1, cycles[i])
	}
}