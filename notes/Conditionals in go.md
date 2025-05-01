- Braces must open in the same line of the if/else ( Aghh 😕 )
- in Go's if-statements **parentheses( )** around conditions **are optional**. but the **braces { } are required** even for oneliners.
- Go doesn't have Ternary Operator ( x < 0 ? A : B ) 🤷

```go
value := 10
if value < 10 {
	println("Less Than 10")
} else if value > 10 {
	println("Greater Than 10")
} else {
	println("Equals 10")
}

//if conditions with statment
//note that value is inscope of all if/else's 
if 	value := 10; value < 10 {
	println(value, "Less Than 10")
} else if value > 10{
	println(value, "Greater Than 10")
} else {
	println(value, "Equals 10")
}
```
```go
value := 10
if value < 10 {
	println("Less Than 10")
} else if value > 10 {
	println("Greater Than 10")
} else {
	println("Equals 10")
}

//if conditions with statment
//note that value is inscope of all if/else's 
if 	value := 10; value < 10 {
	println(value, "Less Than 10")
} else if value > 10{
	println(value, "Greater Than 10")
} else {
	println(value, "Equals 10")
}

// Adding logs
log.Println("Value is:", value)
if value < 10 {
	log.Println(value, "is Less Than 10")
} else if value > 10 {
	log.Println(value, "is Greater Than 10")
} else {
	log.Println(value, "is Equals 10")
}
```