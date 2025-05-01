The rightmost argument can be a list of variable size (_slice_) of data.

```go
// x value here has no use for average. just illustrates the idea of having arguments
// then a variable number of arguments
func average(x int, values …int) float64{
	//print values
	fmt.Println("Single argument value: ", x)
	fmt.Println("Variable argument values: ", values)

	//calculate average
	total := 0
	for _, value := range values {
		total += value
	}

	return float64(total) / float64(len(values))
}

func main() {
	avg := average(10,20,30,40,50)
	println("Average:", avg)
}
```
