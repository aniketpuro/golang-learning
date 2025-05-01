Receiver are the way you create a method for a specific type/struct

```go
type rect struct {
    width, height int
}

func (r *rect) area() int {
    return r.width * r.height
}

//used as 
r := rect{2,3}
areaX := r.area()
fmt.Println(areaX)
```
