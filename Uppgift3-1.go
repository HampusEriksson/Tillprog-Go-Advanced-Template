// Namn
// Datum
/*Förklara vad som är fel i programmet nedan.
Skriv detta som kommentarer i din kod.
Fixa sedan så att all data passerar genom kanalerna och printas ut som avsett.
*/

package main

import "fmt"

// I want this program to print "Hello world!", but it doesn't work.
func main() {
	ch := make(chan string)
	ch <- "Hello world!"
	fmt.Println(<-ch)
}
