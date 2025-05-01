- channels in layman terms are like a synchronized bucket that contains data, a go-routine can add data to the channel, or extract data from the channel. There are unbuffered channels, and buffered channels. for unbuffered channels if you're adding data to the channel, adding another data will be blocking until another go-routine extract such data. on the other hand receiving is also blocking until data is put in the channel. GO Buffered channel add a buffer to the go channel to avoid stalls, however it is not recommended to use it as a beginner, uses it only when it makes sense.
- Channels can be `bidirectional (chan)`, `receive (<-chan)` only, or `send only(chan <-)` , send/receive only channels are useful when channels are passed as arguments, this indicates(and rather enforces) that the passed channel can only be received from (and you can send to), so this introduces some sort of control over how channels are used. think of pkgs where I don't want users to send anything to my channel.

## Examples

### 1

Note that I am using time.Sleep at the end to wait for the code to execute as the program will instantly close after running the two go routines.

```go
c := make(chan int)

go func() {
	for i := 0; i < 9; i++ {
		time.Sleep(time.Second)
		c <i
	}
}()

go func() {
	for{
		fmt.Println( <c )
	}
}()

time.Sleep(time.Second * 15)
```

### 2

Using Range on a channel, it will iterate over values added on the channel until the channel is closed. No need to use time.sleep as the for-range is a blocking function.

```go
c := make(chan int)

go func() {
	for i := 0; i <= 10; i++ {
		time.Sleep(time.Second)
		c <i
	}
	close(c)
}()

for n := range c{
	fmt.Println(n)
}
```

### 3

Using wait-group to use more than 1 function to write to the same channel. Using a `waitGroup` to close the channel once the two writing functions signal `wg.Done()`

```go
c := make(chan int)

var wg sync.WaitGroup

wg.Add(2)

go func() {
	for i := 0; i <= 10; i++ {
		time.Sleep(time.Millisecond * 350)
		c <i
	}
	wg.Done()
}()

go func() {
	for i := 1000; i <= 1010; i++ {
		time.Sleep(time.Millisecond * 350)
		c <i
	}
	wg.Done()
}()

go func() {
	wg.Wait()
	close(c)
}()

for n := range c{
	fmt.Println(n)
}
```

### 4

Using dynamic number of function calls.

Also notice passing i inside the go func, this is because the value outside is in a for-loop, hence it is changing, so using it inside the the go-routine will lead to unexpected results.

```go
c := make(chan string)

var wg sync.WaitGroup

n := 10

wg.Add(n)

for i := 0; i < n; i++  {
	go func(i int) {
		for t := i*10; t < i*10 + 10; t++ {
			c <- "From " + 	strconv.Itoa(i) + " : " + strconv.Itoa(t)
		}
		wg.Done()
	}(i)
}

go func() {
	wg.Wait()
	close(c)
}()

for x := range c{
	fmt.Println(x)
}
```

### 5 - Semaphores

Using only channels without waitGroup. This is done using a channel that store bool (or anything really), and use a function to receive n-done signals then close the main channel.

```go
c := make(chan string)
done := make(chan bool)

n := 2


for i := 0; i < n; i++  {
	go func(i int) {
		for t := i*10; t < i*10 + 10; t++ {
			c <- "From " + 	strconv.Itoa(i) + " : " + strconv.Itoa(t)
		}
		done <true
	}(i)
}

go func() {
	//receive the n-dones from the go-routines
	for i := 0; i < n; i++{
		<done
	}
	close(c)
}()

for x := range c{
	fmt.Println(x)
}
```

### 6 - Using channels as arguments / returns

In this example we sum values from 0 to i. e.g( i = 3 → 0+1+2+3 = 6)

This code we create a go routine that feeds the **increment channels values** 1,2,3 another Channel called the **sum channel** will take the **increment channel and processes its values ( so the sum channel will run until the increment channel closes )**, then the sum channel will **put its sum value for the main to pick**. the point here that main can do other stuff while sum channel finish processing. also we can pass any type of channel to sum channel to sum not necessary an incrementing values in a decoupled way.

```go
func main() {

	i := 10

	//Return A Channel that produces 1, 2, 3…n
	c := incrementer(i)

	// Take a channel that produces 1,2,3…n and sum these numbers
	// returns a channel that have the data in it after summation (so it is not blocking the main itself)
	cSum := puller(c)

	/* DO STUFF WHILE Puller is working (that's why it is returning a channel */

	//Pull from the puller when we want the result (This is blocking now)
	//Result for i := 10 should be : 10 + 9 + 8 + 7 + 6 + 5 + 4 + 3 + 2 + 1 + 0 = 55
	fmt.Println("Final Sum", <cSum)

}

//returns an ACTIVE go routine that produces 1,2,3..n
func incrementer(n int) chan int  {
	out := make(chan int, 10)

	//no need to pass n as parameter as it is a non-changing variable in this context
	go func() {
		for i := 0; i <= n; i++ {
			fmt.Println("From incrementer: Produced i = ", i )
			out <i
			//just to illustrate it is blocking in main.
			time.Sleep(time.Millisecond * 100)
		}
		close(out)
	}()
	return out
}

//takes a channel that produces numbers that are to be summed together.
func puller(c chan int) chan int  {
	out := make (chan int)
	go func() {
		var sum int
		for n := range c{
			fmt.Println("From Puller go-routine: Sum + i ->", sum, "+", n, "=", sum + n)
			sum += n
		}
		fmt.Println("Summation Finished -> Outputing SUM")
		out <sum
		// also we can output each sum stage for whoever uses the channel and close when finish.
		//close(out)
	}()
	return out
}
```

```txt
Output:

    From incrementer: Produced i =  0
    From Puller go-routine: Sum + i -> 0 + 0 = 0
    From incrementer: Produced i =  1
    From Puller go-routine: Sum + i -> 0 + 1 = 1
    From incrementer: Produced i =  2
    From Puller go-routine: Sum + i -> 1 + 2 = 3
    From incrementer: Produced i =  3
    From Puller go-routine: Sum + i -> 3 + 3 = 6
    From incrementer: Produced i =  4
    From Puller go-routine: Sum + i -> 6 + 4 = 10
    From incrementer: Produced i =  5
    From Puller go-routine: Sum + i -> 10 + 5 = 15
    Summation Finished -> Outputing SUM
    Final Sum 15
```
