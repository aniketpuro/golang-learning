- Again Everything is `passed by value` except **arrays, slices, maps and channels** which some calls r**eference types**, these types are passed by reference.
- unlike in C, it's perfectly OK to return the address of a local variable; the storage associated with the variable survives after the function returns.

## Typical functions

```go
// return void
func add(x int, y int)  {
	fmt.Println("Hello, World!")
}

//-------arguments------return------
func add(x int, y int)  int {
	return x + y
}

//-----same type arguments-----------
func add(x, y int)  int {
	return x + y
}
```

- [[Go functions multiple return]]
- [[Go function named returns]]
- [[Go variadic function OR arguments list]]
- [[Go function type & Returning functions]]
- [[Go function callbacks]]
- [[Defer keyword in go]]
- [[Go receivers]]
- [[Go overriding recievers]]
