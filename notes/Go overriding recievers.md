```go
type Person struct {
	First string
	Last string
	Age int
 }

type Employee struct {
	Person
	ID string
	Salary int
}

func (p Person) FullName() string{
	return p.First + " " + p.Last
}

 //Override
func (p Employee) FullName() string{
	return p.ID + " " + p.First + " " + p.Last
}


func main() {
	x := Employee{
		Person{
			"Sherif",
			"Abdel-Naby",
			12},
		"0ID12000ID",
		9999,
	}

fmt.Println(x)
fmt.Println(x.Person.FullName()) //Sherif Abdel-Naby
fmt.Println(x.FullName()) 		   //0ID12000ID Sherif Abdel-Naby
```
