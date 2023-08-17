# Uppgifter GO i Tillämpad Programmering

## Dessa uppgifter är tagna ifrån https://github.com/korthaj/pallinda21

[Step-by-step guide to concurrency](http://yourbasic.org/golang/concurrent-programming/)

### Förkunskaper Uppgift 1-2

- [Why Go? – Key advantages you may have overlooked](https://yourbasic.org/golang/advantages-over-java-python/)
- [Goroutines](http://yourbasic.org/golang/goroutines-explained/)
- [Channels](https://yourbasic.org/golang/channels-explained/)


### Uppgift 1 - Alarm Clock

In this task you will explore time functions using Go. Write a function `Remind(text string, delay time.Duration)` that will print the following output:

    Klockan är XX.XX: + <text>

Skeleton-code can be found [here](https://github.com/HampusEriksson/Tillprog-Go/blob/main/Skeletons/Upp2Skeleton.go)
The output will repeatedly print the output after the given delay, and `XX.XX` should be replaced with the current time, and `<text>` should be replaced by the contents of `text`.

Now, write a complete program that runs indefinitely and prints the following reminders:

- every 3rd hour: `Klockan är XX.XX: Dags att äta`
- every 8th hour: `Klockan är XX.XX: Dags att arbeta`
- every 24th hour: `Klockan är XX.XX: Dags att sova`

Tips: Try with second instead of hour.
To prevent the main program from exiting early, the following statement can be used:

```Go
select { }
```

In order to access time related functions, you should investigate the [time package](https://golang.org/pkg/time/), and discover how to get the current time in Go and also how you can format it neatly for human users to understand. Remember to test and format your code.

### Uppgift 2

In this task you will complete the following partial program. It adds all of the numbers in an array by splitting the array in half, then having two Go routines take care of each half. Partial results are then sent over a channel. Remember to test and format your code.

```Go
package main

// Add adds the numbers in a and sends the result on res.
func Add(a []int, res chan <- int) {
    // TODO
}

func main() {
    a := []int{1, 2, 3, 4, 5, 6, 7}
    n := len(a)
    ch := make(chan int)
    go Add(a[:n/2], ch)
    go Add(a[n/2:], ch)

    // TODO: Get the subtotals from the channel and print their sum.
}
```

### Förkunskaper Uppgift 3-5

- [Channels](http://yourbasic.org/golang/channels-explained/)
- [Select](http://yourbasic.org/golang/select-explained/)
- [Data races](http://yourbasic.org/golang/data-races-explained/)
- [Detecting data races](http://yourbasic.org/golang/detect-data-races/)
- [Deadlock](http://yourbasic.org/golang/detect-deadlock/)
- [Waitgroups](https://yourbasic.org/golang/wait-for-goroutines-waitgroup/)

### Uppgift 3 - Debugging Concurrent Programs

Förklara vad som är fel i de två programmen nedan. Skriv detta som kommentarer i din kod.
Fixa sedan så att all data passerar genom kanalerna och printas ut som avsett.

#### Bug 1

```Go
package main

import "fmt"

// I want this program to print "Hello world!", but it doesn't work.
func main() {
    ch := make(chan string)
    ch <- "Hello world!"
    fmt.Println(<-ch)
}
```

See: [bug01.go](code/bug01.go) for source code to modify.

#### Bug 2

```Go
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
```

See: [bug02.go](code/bug02.go) for source code to modify.

### Uppgift 4 - Many Senders; Many Receivers

The program [many2many.go](code/many2many.go) contains four producers that together send 32 strings over a channel.
At the other end there are two consumers that receive the strings.
Describe what happens, and explain why it happens, if you make the following changes in the program.
Try first to reason your way through, and then test your hypothesis by changing and running the program.
Lämna in svaren på frågorna i ett google-doc eller som kommentarer i din kod.

- What happens if you switch the order of the statements `wgp.Wait()` and `close(ch)` in the end of the `main` function?
- What happens if you move the `close(ch)` from the `main` function and instead close the channel in the end of the function `Produce`?
- What happens if you remove the statement `close(ch)` completely?
- What happens if you increase the number of consumers from 2 to 4?
- Can you be sure that all strings are printed before the program stops?

Finally, modify the code by adding a new WaitGroup that waits for all consumers to finish.

### Uppgift 5 - Pythia, the Oracle of Delphi

The code in [oracle.go](code/oracle.go) contains the outline for a program that will answer 'questions'.
Complete the `Oracle` function. **You should not modify the `main` function or other function signatures.**
Note that answers should not appear immediately; instead there should be a delay or **pause for thought**.
Also, the Oracle will still print **helpful predictions** even if there are not any questions.
You may structure your solution into multiple functions.

Your program should contain two channels: One channel for questions, and one for answers and predictions.
In the `Oracle` function you should start three indefinite go-routines.

- A go-routine that receives all questions, and for each incoming question, creates a separate go-routine that answers that question
- A go-routine that generates predictions
- A go-routine that receives all answers and predictions, and prints them

Whilst the `Oracle` function is the most important of the assignment, you may also want to improve the answer-algorithm.

- The [strings](https://golang.org/pkg/strings/) and [regex](https://golang.org/pkg/regexp/) packages may be of some help
- The program can seem more human if the Oracle prints it answers one character at a time
- Take a look at the story of [ELIZA](https://en.wikipedia.org/wiki/ELIZA)

### Förkunskaper Uppgift 6-8

- [Mutual exclusion](http://yourbasic.org/golang/mutex-explained/)
- [Efficient parallel computation](http://yourbasic.org/golang/efficient-parallel-computation/)
- [Create a new image](https://yourbasic.org/golang/create-image/)
- [HTTP server example](https://yourbasic.org/golang/http-server-example/)

### Uppgift 6 - Matching Behaviour

Take a look at the program [matching.go](code/matching.go). Explain what happens and why it happens if you make the following changes. Try first to reason about it, and then test your hypothesis by changing and running the program.

- What happens if you remove the `go-command` from the `Seek` call in the `main` function?
- What happens if you switch the declaration `wg := new(sync.WaitGroup`) to `var wg sync.WaitGroup` and the parameter `wg *sync.WaitGroup` to `wg sync.WaitGroup`?
- What happens if you remove the buffer on the channel match?
- What happens if you remove the default-case from the case-statement in the `main` function?

> **Hint:** Think about the order of the instructions and what happens with arrays of different lengths.

### Uppgift 7 - Fractal Images

The file [julia.go](code/julia.go) contains a program that creates images and writes them to file. The program is pretty slow. Your assignment is to divide the computations so that they run in parallel on all available CPUs. Use the ideas from the example in the [efficient parallel computation](http://yourbasic.org/golang/efficient-parallel-computation/).

You can also make changes to the program, such as using different functions and other colourings.

How many CPUs does you program use? How much faster is your parallel version?

> **Hint:** In more recent versions of Golang (since 1.5), the runtime will default to use as many operating system threads as it is allowed. To see differences in behaviour, refer to the [GOMAXPROCS](https://golang.org/pkg/runtime/#GOMAXPROCS) function and vary the value.

### Uppgift 8 - Weather station

The file [server.go](code/server.go) contains a program that simulates three independent weather stations that show the temperature at KTH. The results are published at the following addresses once the serves are operational:

- `http://localhost:8080`
- `http://localhost:8081`
- `http://localhost:8082`

Start the program and try to visit the three addresses in your browser. You'll soon find that the three services aren't very reliable; they're pretty slow and sometimes you get no response at all. You might also get the error message "Service unavailable".

Your assignment is to write a client that simultaneously asks all servers and terminates the search as soon as one has responded with a correct temperature. The request should also terminate if no-one has answered within a given time. The file [client.go](code/client.go) contains a template from which you should build on.

- Read through the code and start the client whilst the weather stations are operational
- Implement the function `MultiGet`
