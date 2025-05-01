You can declare return variables and name them at the beginning, they are returned in the end.

you can override the returns and return whatever you want at the return statement.

```go
//Returns x,y at the end.
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
	//return a,b  <u can override default return of x,y.
}
```
