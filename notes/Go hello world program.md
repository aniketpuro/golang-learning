```go
package main        //Excuetables must be of package main
import "fmt"
        
func main() {
	var g string = "Hello golang"
	println(g)
}
        
func function() string {
	return "Five of Diamonds"
}
```

## Notes

1. Program Excuetable is at `package main/func main()`
2. **Blank Identifier**: Use `_` to replace unused variables.
3. there are two primitive ways to allocate to a pointer, `new()` and `make()`, they differ though, we will discuss that later, briefly new returns pointer, make return value. read [Doc](https://golang.org/doc/effective_go.html#allocation_new).
4. Every thing is `passed by value` except **arrays, slices, maps and channels** which some calls **reference types**, these types are passed by reference ( they internally have pointers, so no copying of the actual data happens when passing them) .
5. Unlike in C, it's perfectly OK to return the address of a local variable; the storage associated with the variable survives after the function returns.
