// Namn
// Datum
/*
Detta program innehåller en översikt för ett program som kommer att besvara 'frågor'.
Komplettera Oracle-funktionen. Du ska inte ändra huvudfunktionen eller andra funktioner.
Observera att svar inte ska visas omedelbart; istället bör det finnas en fördröjning eller paus för eftertanke.
Dessutom kommer Oracle fortfarande att skriva ut hjälpsamma förutsägelser även om det inte finns några frågor.
Du kan strukturera din lösning i flera funktioner.

Ditt program bör innehålla två kanaler: En kanal för frågor och en för svar och förutsägelser. I Oracle-funktionen bör du starta tre obestämda go-rutiner.

En go-rutin som tar emot alla frågor och för varje inkommande fråga skapar en separat go-rutin som svarar på den frågan
En go-rutin som genererar förutsägelser
En go-rutin som tar emot alla svar och förutsägelser och skriver ut dem
Även om Oracle-funktionen är den viktigaste delen av uppgiften, kanske du också vill förbättra svarsalgoritmen.

Extra: Programmet kan verka mer mänskligt om Oracle skriver ut sina svar en bokstav i taget
*/
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	star   = "Pythia"
	venue  = "Delphi"
	prompt = "> "
)

func main() {
	fmt.Printf("Welcome to %s, the oracle at %s.\n", star, venue)
	fmt.Println("Your questions will be answered in due time.")

	oracle := Oracle()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		fmt.Printf("%s heard: %s\n", star, line)
		oracle <- line // The channel doesn't block.
	}
}

// Oracle returns a channel on which you can send your questions to the oracle.
// You may send as many questions as you like on this channel, it never blocks.
// The answers arrive on stdout, but only when the oracle so decides.
// The oracle also prints sporadic prophecies to stdout even without being asked.
func Oracle() chan<- string {
	questions := make(chan string)
	// TODO: Answer questions.
	// TODO: Make prophecies.
	// TODO: Print answers.
	return questions
}

// This is the oracle's secret algorithm.
// It waits for a while and then sends a message on the answer channel.
// TODO: make it better.
func prophecy(question string, answer chan<- string) {
	// Keep them waiting. Pythia, the original oracle at Delphi,
	// only gave prophecies on the seventh day of each month.
	time.Sleep(time.Duration(20+rand.Intn(10)) * time.Second)

	// Find the longest word.
	longestWord := ""
	words := strings.Fields(question) // Fields extracts the words into a slice.
	for _, w := range words {
		if len(w) > len(longestWord) {
			longestWord = w
		}
	}

	// Cook up some pointless nonsense.
	nonsense := []string{
		"The moon is dark.",
		"The sun is bright.",
	}
	answer <- longestWord + "... " + nonsense[rand.Intn(len(nonsense))]
}

func init() { // Functions called "init" are executed before the main function.
	// Use new pseudo random numbers every time.
	rand.Seed(time.Now().Unix())
}
