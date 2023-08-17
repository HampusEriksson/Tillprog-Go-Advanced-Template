// Namn
// Datum
/*
Detta program innehåller fyra produceras som tillsammans skickar 32 strängar över en kanal.
I andra änden finns två consumers som tar emot strängarna.
Beskriv vad som händer, och förklara varför det händer, om du gör följande ändringar i programmet.
Försök först resonera dig fram, och testa sedan din hypotes genom att ändra och köra programmet.
Lämna in svaren på frågorna som kommentarer i denna fil.

Vad händer om du växlar ordningen på uttalandena wgp.Wait() och close(ch) i slutet av huvudfunktionen?
Svar:
Vad händer om du flyttar close(ch) från huvudfunktionen och istället stänger kanalen i slutet av funktionen Produce?
Svar:
Vad händer om du tar bort uttalandet close(ch) helt och hållet?
Svar:
Vad händer om du ökar antalet konsumenter från 2 till 4?
Svar:
Kan du vara säker på att alla strängar skrivs ut innan programmet avslutas?
Svar:
Slutligen, ändra koden genom att lägga till en ny WaitGroup som väntar på att alla konsumenter ska avslutas.
*/
package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

func main() {
	// Use different random numbers each time this program is executed.
	rand.Seed(time.Now().Unix())

	const strings = 32
	const producers = 4
	const consumers = 2

	before := time.Now()
	ch := make(chan string)
	wgp := new(sync.WaitGroup)
	wgp.Add(producers)
	for i := 0; i < producers; i++ {
		go Produce("p"+strconv.Itoa(i), strings/producers, ch, wgp)
	}
	for i := 0; i < consumers; i++ {
		go Consume("c"+strconv.Itoa(i), ch)
	}
	wgp.Wait() // Wait for all producers to finish.
	close(ch)
	fmt.Println("time:", time.Now().Sub(before))
}

// Produce sends n different strings on the channel and notifies wg when done.
func Produce(id string, n int, ch chan<- string, wg *sync.WaitGroup) {
	for i := 0; i < n; i++ {
		RandomSleep(100) // Simulate time to produce data.
		ch <- id + ":" + strconv.Itoa(i)
	}
	wg.Done()
}

// Consume prints strings received from the channel until the channel is closed.
func Consume(id string, ch <-chan string) {
	for s := range ch {
		fmt.Println(id, "received", s)
		RandomSleep(100) // Simulate time to consume data.
	}
}

// RandomSleep waits for x ms, where x is a random number, 0 ≤ x < n,
// and then returns.
func RandomSleep(n int) {
	time.Sleep(time.Duration(rand.Intn(n)) * time.Millisecond)
}
