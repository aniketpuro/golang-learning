Slices are of dynamic size.

```go
letters := []string{"a", "b", "c", "d"}

/*  using make -> make([]T, len, cap) */
var s []byte
s = make([]byte, 5, 5)
//OR
s := make([]byte, 5)

// both equiavlent to: s == []byte{0, 0, 0, 0, 0}
```

- A slice does not store any data, it just describes a section of an underlying array. Changing the elements of a slice modifies the corresponding elements of its underlying array. **Other slices that share the same underlying array will see those changes**.
- Slicing a slice changes pointers of the underlying array, so it is as efficient as manipulating array indices, size and capacity of the new slice are changed too, capacity is equal `old capacity - sliced part from the beginning only`

```go
names := [4]string{"John","Paul","George","Ringo",}
fmt.Println(names)   //[John Paul George Ringo]
a := names[0:2]
b := names[1:3]
fmt.Println(a, b)    //[John Paul] [Paul George]
b[0] = "XXX"
fmt.Println(a, b)    //[John XXX] [XXX George] 
fmt.Println(names)   //[John XXX George Ringo]

//ALSO

//This is an array literal:
[3]bool{true, true, false}

//And this creates the same array as above, then builds a slice that references it:
[]bool{true, true, false}
```

## Iterating over a slice

```go
for i, v := range arr {	//do stuff }
for _, v := range arr {	//do stuff }
for i, _ := range arr {	//do stuff }
```

## Appending to a slice

```go
var s []int
// append works on nil slices.
s = append(s, 0)
// The slice grows as needed.
s = append(s, 1)
```

Append add element at the end of the slice **if there is enough capacity** and r**eturn a reference type**!, if **not enough capacity** it allocate and copy to a new array and **return it as a new value**! and the **old array points to the old data**.

If Append had enough capacity (didn't allocate new array) then _**changing a value in the new returned array changes the value in the old**_! but if it allocated a new array to expand capacity, then **changing a value at an index of the newly returned array _DOESN'T_ change the old array!**

consider only using append where the left hand side is the same variable in the append first argument **(S = append(S, …..) )** to avoid any unexpected results

```go
//Allocate new capacity  
var s []int
s = make([]int, 5, 5)
x := append(s, 1, 2 )
x[0] = 1337
s[0] = 6800
fmt.Println(s,x) //[6800 0 0 0 0] [1337 0 0 0 0 1 2 3]

//Doesn't allocat new capacity and return reference
var s []int
s = make([]int, 5, 150)
x := append(s, 1, 2 )
x[0] = 1337
s[0] = 6800
fmt.Println(s,x) //[6800 0 0 0 0] [6800 0 0 0 0 1 2 3]
				 //notice that 1337 is overwritten
```

## Common slice functions

### Append another slice

```go
a = append(a, b…)
```

### Copy

Copy only copy elements of size = min(len(a), len(b)). so, the new slice to which a copy is to be made must have a size == size of the original array to have all elements copied.

```go
b = make([]T, len(a))
copy(b, a)
// or
b = append([]T(nil), a…)
```

### Cut

```go
a = append(a[:i], a[j:]…)
```

### Delete

```go
a = append(a[:i], a[i+1:]…)
// or
a = a[:i+copy(a[i:], a[i+1:])]
```

## Slices tricks

[golang/wiki/SliceTricks](https://github.com/golang/go/wiki/SliceTricks)
