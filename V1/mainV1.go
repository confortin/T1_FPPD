//V1

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
const rounds = 20

var forks [numPhilosophers]sync.Mutex
var statsMu sync.Mutex

func Philosopher(id int, wg *sync.WaitGroup) {
	defer wg.Done() // avisa "terminei" quando a função acabar
	time_to_think := 200
	time_to_eat := 300
	left := id
	right := (id + 1) % numPhilosophers

	for iteration := 1; iteration <= rounds; iteration++ {
		// think
		fmt.Printf("%*s T%d\n", id*6, "", iteration) // Thinking
		time.Sleep(time.Duration(time_to_think) * time.Millisecond)

		waitStart := time.Now()
		forks[left].Lock()
		fmt.Printf("P%d pegou o garfo %d\n", id+1, left+1)
		//time.Sleep(100 * time.Millisecond) // Com time sleep, toda vez ocorre deadlock
		fmt.Printf("P%d esperando o garfo %d\n", id+1, right+1)
		forks[right].Lock()

		wait := time.Since(waitStart)
		statsMu.Lock()
		totalWait = totalWait + wait
		if wait > maxWait {
			maxWait = wait
		}
		statsMu.Unlock()

		// eat
		statsMu.Lock()
		eating++
		if eating > maxEating {
			maxEating = eating
		}
		statsMu.Unlock()

		fmt.Printf("%*s E%d\n", id*6, "", iteration) // Eating
		time.Sleep(time.Duration(time_to_eat) * time.Millisecond)

		statsMu.Lock()
		eating--     // NOVO
		cycles[id]++ // NOVO: mais um ciclo concluido
		statsMu.Unlock()

		forks[left].Unlock()
		forks[right].Unlock()
	}
}

func main() {
	var wg sync.WaitGroup

	fmt.Println("\n[P1]  [P2]  [P3]  [P4]  [P5]\n")

	// run philosophers

	start := time.Now()

	for j := 0; j < numPhilosophers; j++ {
		wg.Add(1)              // "vou lançar mais 1 filósofo"
		go Philosopher(j, &wg) // o "go" faz ele rodar em paralelo
	}
	wg.Wait() // main espera os 5 terminarem

	elapsed := time.Since(start)

	fmt.Println("Tempo total:", elapsed)
	fmt.Println("Tempo medio de espera:", totalWait/time.Duration(rounds*numPhilosophers))
	fmt.Println("Maior tempo de espera:", maxWait)
	fmt.Println("Maximo de filosofos comendo juntos:", maxEating)
	for i := 0; i < numPhilosophers; i++ {
		fmt.Printf("Ciclos concluidos P%d: %d\n", i+1, cycles[i])
	}
}
