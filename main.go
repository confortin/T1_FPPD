// Example code filo_seq.go - Sequential Dinning Philosophers
// Since is sequential, elapsed timne could be simply calculated by (roubnds*philosophers*(time_to_think+time_to_eat))
// By Cesar De Rose

package main

import (
     "fmt"
     "time"
)

var cycles [5]int           // ciclos concluidos por filosofo
var eating int              // quantos filosofos estao comendo agora
var maxEating int           // maximo de filosofos comendo ao mesmo tempo
var totalWait time.Duration // soma dos tempos de espera pelos garfos
var maxWait time.Duration   // maior tempo de espera


func Philosopher(id, iteration int) {
	time_to_think := 200 
	time_to_eat := 300
 
	// think
	for s := 1; s <= id*6; s++ {
		fmt.Printf(" ")
	}
	fmt.Printf(" T%d\n", iteration) // Thinking
	time.Sleep(time.Duration(time_to_think) * time.Millisecond)
 
	// NOVO: adquirir recursos (na V0 nao ha garfos, entao nao espera nada)
	waitStart := time.Now()
	// (nas outras versoes, o Lock dos garfos entra aqui)
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
	// NOVO: liberar recursos (nas outras versoes, o Unlock dos garfos entra aqui)
}

func main() {
     philosophers := 5
     rounds := 20 //// mudamos para de 5 para 20

     fmt.Println("\n[P1]  [P2]  [P3]  [P4]  [P5]\n")
 
     // run philosophers

     start := time.Now()

     for i := 1; i <= rounds; i++ {
          for j := 0; j < philosophers; j++ {
               Philosopher(j, i)
          }
     }
     elapsed := time.Since(start)

     fmt.Printf("\nSequential Dinner took %s\n\n", elapsed)
}
