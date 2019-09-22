package main

import "fmt"

func main() {
	p := &Person{
		ID:   1,
		Name: "Krisna Putra",
		Age:  26,
	}
	fmt.Println(p.GetID())
	fmt.Println(p.GetName())
	fmt.Println(p.GetAge())
	fmt.Println("Change Name")
	p.ChangeName("Anak Agung Krisna Putra")
	fmt.Println(p.GetName())
}

type Person struct {
	ID   int
	Name string
	Age  int
}

func (p *Person) ChangeName(newName string) {
	p.Name = newName
}
func (p *Person) GetID() int {
	return p.ID
}
func (p *Person) GetName() string {
	return p.Name
}
func (p *Person) GetAge() int {
	return p.Age
}
