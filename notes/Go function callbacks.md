Passing functions as argument

```go
func visit(numbers []int, callback func(int2 int)){
	for _, n := range numbers {
		callback(n*2)
	}
}
func main() {
	visit([]int{1,2,3,4}, func(n int){
		fmt.Println(n, "Printed withing the callback function.")
	})
}
```
