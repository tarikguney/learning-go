## Run the application

```
go run main.go &
```

## Stop

```
# Finding which process ID using port 8080
lsof -i tcp:8080
kill process_id
```

## Notes on Channels

What's important about Go channels is that they block go routines when they is nothing to read or write from/into channels. This is very important aspect about channels. Read is known as recieve. Write is known as send. 

For instance,

```
func main(){
    score := make(chan int) // chan<- means send-only.

    go func(){
        fmt.Println("Inside the go routine")
        score <- 1 // 3
        fmt.Println("Exiting...") // 2
        score <- 1 // 4
    }()

    fmt.Println(<-score) // 1
}
```
At first, 1 will block the main go routine since there is nothing to read from it. However, 3 will not block anything since the score channel is empty and can receive data. In fact, if there is enough time, 2 will be executed. However, 4 will block until the score channel is cleared out.  

The idea is simple, if I can write to a channel, then I write and continue. If I cannot, that means the data needs to be received and the underlying data needs to be cleared in order for me to write into it. Second, if I can read a value from the channel, I will not block the go routine. However, if there is nothing in there for me to read, then I will block the current goroutine until I can read a value from it.

The relationship between channels and go routines are actually amazingly well and simply designed. At first, it takes a while to completely understand the idea, but once you get the basics, you can add up to it.