// Namn
// Datum
/*
I denna uppgift ska du implementera funktionen Remind. Funktionen ska printa följande:
Klockan är XX.XX: + <text>

Funktionen ska fortsätta printa ut efter en given delay. "XX.XX" ska bytas ut mot nuvarande tiden och <text> ska bytas ut mot indatan text.

Skriv nu main-funktionen så att den producerar följande utskrifter:
- Var 3e timme: Klockan är XX.XX: Dags att äta
- Var 8e timme: Klockan är XX.XX: Dags att arbeta
- Var 24e timme: Klockan är XX.XX: Dags att sova

Tips: Testa med sekund istället för timme.

För att få nuvarande tid kan du kolla in: https://golang.org/pkg/time/
*/

package main

import "time"

func Remind(text string, delay time.Duration) {

}

func main() {

	// För att förhindra programmet att stängas av kan du göra en tom select

	select {}
}
