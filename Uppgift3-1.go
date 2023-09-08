// Namn
// Datum
/*Förklara varför programmet nedan får deadlock.
Skriv detta som kommentarer i din kod.
Fixa sedan så att all data passerar genom kanalerna och printas ut som avsett.
För C och A ska två möjliga lösningar ges.
*/

package main

import "fmt"

// I want this program to print "Hello world!", but it doesn't work.
func main() {
	ch := make(chan string)
	ch <- "Hello world!"
	fmt.Println(<-ch)
}
