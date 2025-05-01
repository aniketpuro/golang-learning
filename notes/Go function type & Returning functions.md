1. Functions can be assigned to variables `func0 := func() int {x++; return x}`
2. Functions that are returned from another functions has its own scope per returned function (yea ikr ? 🤷).

```go

package main

var x = 0

func main() {
	//local x
	x := 0

	func0 := func() int {x++; return x}
	func1 := incrementGlobalX //without ()
	func2 := wrapper()
	func3 := wrapper()

	println(func0(), " : func0 (local x)")
	println(func1(), " : func1 (global x)")
	println(func2(), " : func2 (per func scope x1)")
	println(func3(), " : func3 (per func scope x2)")
	println("Second Increment")
	println(func0(), " : func0 (local x)")
	println(func1(), " : func1 (global x)")
	println(func2(), " : func2 (per func scope x1)")
	println(func3(), " : func3 (per func scope x2)")
}

func incrementGlobalX() int  {
	x++
	return x
}

func wrapper() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
```
