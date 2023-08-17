// Namn
// Datum
/*
Kolla på nedanstående program. Förklara vad som händer och varför det händer om du gör följande ändringar.
Försök först att resonera om det, och testa sedan din hypotes genom att ändra och köra programmet.

Vad händer om du tar bort go-kommandot från Seek-anropet i huvudfunktionen?
Svar:
Vad händer om du växlar deklarationen wg := new(sync.WaitGroup) till var wg sync.WaitGroup och parametern wg *sync.WaitGroup till wg sync.WaitGroup?
Svar:
Vad händer om du tar bort bufferten på kanalen match?
Svar:
Vad händer om du tar bort default-fallet från case-satsen i huvudfunktionen?
Svar:
Tips: Tänk på ordningen på instruktionerna och vad som händer med arrayer av olika längder.
*/

package main

import (
	"fmt"
	"sync"
)

// This programs demonstrates how a channel can be used for sending and
// receiving by any number of goroutines. It also shows how  the select
// statement can be used to choose one out of several communications.
func main() {
	people := []string{"Anna", "Bob", "Cody", "Dave", "Eva"}
	match := make(chan string, 1) // Make room for one unmatched send.
	wg := new(sync.WaitGroup)
	wg.Add(len(people))
	for _, name := range people {
		go Seek(name, match, wg)
	}
	wg.Wait()
	select {
	case name := <-match:
		fmt.Printf("No one received %s’s message.\n", name)
	default:
		// There was no pending send operation.
	}
}

// Seek either sends or receives, whichever possible, a name on the match
// channel and notifies the wait group when done.
func Seek(name string, match chan string, wg *sync.WaitGroup) {
	select {
	case peer := <-match:
		fmt.Printf("%s sent a message to %s.\n", peer, name)
	case match <- name:
		// Wait for someone to receive my message.
	}
	wg.Done()
}
