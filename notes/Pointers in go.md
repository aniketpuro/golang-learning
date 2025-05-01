Pointers syntax is essentially like C/C++

```go
var value int = 1000
var pointer *int = &value
println(value)                //1000
println(pointer)              //0xfffffffff
println(*pointer)             //1000
(*pointer)++		  			      //1001
*pointer = *pointer + 10	    //1011
println(*pointer)			        //1011
println(*pointer + *pointer)  //1011 + 1011 = 2022
```
### 🔍 Tumhara code:

```go
func modify(y int) int {
    y += 15
    return y
}

func main() {
    y := 20
    modify(y)
    fmt.Println(y)
}
```
### 🤔 Problem kya hai?

- Jab tum `modify(y)` call karte ho, Go language **value pass karta hai, reference nahi**.
    
- Matlab: `y` ka **copy** function ke andar jaata hai.
    
- Toh jo `y += 15` hai, wo sirf copy pe hota hai, asli `y` **main()** function mein waise ka waise rehta hai.
    

---
### 🖨️ Output:

`20`

---
## ✅ Sahi tareeka:

Tumhe `modify` ka return value **wapis assign** karna padega:

```go
package main

import "fmt"

func modify(y int) int {
    y += 15
    return y
}

func main() {
    y := 20
    y = modify(y) // y ko updated value assign kar rahe hain
    fmt.Println(y)
}
```
### 🖨️ Ab Output:

`35`

---

### 🔁 Bonus: Agar pointer use karna ho (advanced)

Chaho toh pointer se bhi value modify kar sakte ho, directly:

```go
package main

import "fmt"

func modify(y *int) {
    *y += 15
}

func main() {
    y := 20
    modify(&y) // address pass kar rahe hain
    fmt.Println(y)
}

```

### 🖨️ Output:

`35`


### 🔍 Tera code:
```go
package main

import "fmt"

func main() {
	y := [3]int{10, 20, 30}
	py := &y
	fmt.Printf("%T %v \n", py, *py)
}

```
### 🔎 Ismein ho kya raha hai?

- `y` ek **array** hai of 3 integers: `[10, 20, 30]`
    
- `py := &y` — iska matlab: `py` ab `y` ka **pointer** hai, yaani `py` holds the **address** of the array
    
- `*py` dereference karta hai — matlab actual array value print hoti hai

### 🖨️ Output kya aayega?


`*[3]int [10 20 30]`

- `%T` se milega type → `*[3]int` (pointer to array of 3 ints)
    
- `%v` se milegi value → `[10 20 30]`
### Agar tu confusion mein hai:

- `*py` matlab: "dereference the pointer" — yaani jis variable ka address `py` hold kar raha hai, uski actual value le lo
    
- `&y` matlab: "address of `y`"
    

---

Tu chaahe toh `py[0]` bhi likh sakta hai, Go pointer dereference kar leta hai automatically for arrays/slices.

---

``` go
package main

import "fmt"

func main() {
	var y int          // default value: 0
	var ptr *int = &y  // ptr y ka address store karta hai
	fmt.Println(y)     // 0
	fmt.Println(*ptr)  // dereference ptr → again 0
}
```
### 🖨️ Output:

`0
`0`

### 💡Explanation:

var y int → Go assigns default value 0 to y

&y → y ka memory address 

*ptr → us address pe jo value hai, wo leke aata hai
package main

---

``` go
import (
    "fmt"
    "strings"
)

func main() {
    s := "hello"
    var ptr *string = &s
    fmt.Println(s)
    *ptr = *ptr + strings.ToUpper(s)
    fmt.Println(s)
}
```
### 🖨️ Output:

``hello
`helloHELLO```
### 🧠 Explanation:
s := "hello": Yeh ek string variable hai.

var ptr *string = &s: Yeh s ka pointer hai.

*ptr = *ptr + strings.ToUpper(s): Yeh line s ko update karti hai by appending its uppercase version.

fmt.Println(s): Yeh updated s ko print karega.​