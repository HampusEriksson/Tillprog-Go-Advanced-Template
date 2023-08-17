// Namn
// Datum
/*Förklara vad som är fel i programmet nedan.
Skriv detta som kommentarer i din kod.
Fixa sedan så att all data passerar genom kanalerna och printas ut som avsett.
*/

package main

import "fmt"

// This program should go to 11, but sometimes it only prints 1 to 10.
func main() {
	ch := make(chan int)
	go Print(ch)
	for i := 1; i <= 11; i++ {
		ch <- i
	}
	close(ch)
}

// Print prints all numbers sent on the channel.
// The function returns when the channel is closed.
func Print(ch <-chan int) {
	for n := range ch { // reads from channel until it's closed
		fmt.Println(n)
	}
}
