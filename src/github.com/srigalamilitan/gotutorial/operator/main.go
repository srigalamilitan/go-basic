package main

import (
	"fmt"
)

func main() {
	x := 10
	y := 5
	//aritmatik operator
	//penambahan /addtiotional
	zAdd := x + y
	fmt.Println(zAdd)
	//pengurangan / subsubtract
	zSub := x - y
	fmt.Println(zSub)

	// perkalian / multplies
	zMulti := x * y
	fmt.Println(zMulti)

	// pembagian / divide
	zDiv := x / y
	fmt.Println(zDiv)

	// sisa hasil bagi/ modulus
	zMod := x % y
	fmt.Println(zMod)

	// ** Operator Relational
	// perbandingan 2 value
	i := 10
	j := 5
	fmt.Println(i == j) // false
	fmt.Println(i != j) //true
	fmt.Println(i > j)  //true
	fmt.Println(i < j)  //false
	fmt.Println(i >= j) //true
	fmt.Println(i <= j) //false

}
