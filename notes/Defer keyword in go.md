`Defer` used before a functions executes the function at the end of the scope of it, think of it as a destructor for the scope. usually used to close opened files/buffers so u open the file and closes it using defer in the next line to keep things clean. they're executed as a **stack**.

```go
fmt.Println("One")
defer fmt.Println("Four")
defer fmt.Println("Three")
fmt.Println("Two")

//Prints One Two Three Four
```
