```go
func swap(x, y string) (string, string) {
	return y, x
}

//in main
a, b := swap("hello", "world")
fmt.Println(a, b) //prints "world hello"
```
