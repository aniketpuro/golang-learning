Switch statements in GO doesn't require `break`; they will break by default, `fallthrough` keyword used to go to NEXT statement even if condition doesn't match, `fallthrough` is like a break so no code can be after it. however a workaround is to use `labels` and `goto`

```go
i := 2
fmt.Println("Switch for i = ", i, " goes to: ")
switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
		i = 4
		fallthrough //goes to NEXT case even if doesn't match.
	case 3:
		fmt.Println("three")
	case 4:
		fmt.Println("four")
	case 5,6:
		fmt.Println("five or six")
	default:
		fmt.Println("default")
}
```
