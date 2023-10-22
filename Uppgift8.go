// Namn
// Datum
/*
Din uppgift är att skriva en klient som samtidigt frågar alla servrar och avslutar sökningen så snart en har svarat med en korrekt temperatur.
Förfrågan ska också avslutas om ingen har svarat inom en given tid (t.ex. 15 sekunder).

Läs igenom koden och starta klienten medan väderstationerna är igång
Implementera funktionen MultiGet

För att uppgiften ska fungera måste du först starta filen server.go
*/
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	server := []string{
		"http://localhost:8080",
		"http://localhost:8081",
		"http://localhost:8082",
	}

	// Add a time limit for all requests made by this client.
	client := &http.Client{Timeout: 10 * time.Second}

	for {
		before := time.Now()
		// Raden nedan är endast för att du kan testa Get-funktionen. Kommentera bort den när du är klar
		res := Get(server[0], client)
		// Denna rad ska kommenteras bort när du är redo att köra funktionen MultiGet
		//res := MultiGet(server, client)
		after := time.Now()
		fmt.Println("Response:", res)
		fmt.Println("Time:", after.Sub(before))
		fmt.Println()
		time.Sleep(500 * time.Millisecond)
	}
}

type Response struct {
	Body       string
	StatusCode int
}

func (r *Response) String() string {
	return fmt.Sprintf("%q (%d)", r.Body, r.StatusCode)
}

// Get makes an HTTP Get request and returns an abbreviated response.
// The response is empty if the request fails.
func Get(url string, client *http.Client) *Response {
	res, err := client.Get(url)
	if err != nil {
		return &Response{}
	}
	// res.Body != nil when err == nil
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("ReadAll: %v", err)
	}
	return &Response{string(body), res.StatusCode}
}

// MultiGet makes an HTTP Get request to each url and returns
// the response from the first server to answer with status code 200.
// If none of the servers answer before timeout, the response is 503
// – Service unavailable.
func MultiGet(urls []string, client *http.Client) *Response {
	return nil // TODO
}
