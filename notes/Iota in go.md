iota in Go, is a value used within the **const** block, its value starts at 0 per block, and increment each time it is used again

```go
const (
  c0 = iota  // c0 == 0
  c1 = iota  // c1 == 1
  c2 = iota  // c2 == 2
)
```
