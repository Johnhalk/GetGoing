package main

import (
	"fmt"
	"strconv"
)

const myConst float64 = 1.67

func main() {
	var (
		actorName = "Maisie Williams"
		noOne = "noone"
		season = 7
	)
	fmt.Println(actorName, noOne, season)

	var (
		i = 42
		j float32
	)

	j = float32(i)

	fmt.Printf("%v, %T\n", j, j)

	var k string
	k = strconv.Itoa(i)

	fmt.Printf("%v, %T\n", k, k)

	var n bool = true
	m := 1 == 2
	fmt.Printf("%v, %T\n", n, n)
	fmt.Printf("%v, %T\n", m, m)

	//constants

	const myConst float64 = 1.57;
	fmt.Printf("%v, %T \n", myConst, myConst)

	const (
		integerA int = 12
		integerB int = 15
	)

	fmt.Printf("%v, %T \n", integerA + integerB, integerA + integerA)

	const (
		integerC = iota
		integerD
		integerE
	)

	fmt.Printf("%v, %T \n", integerC, integerC)
	fmt.Printf("%v, %T \n", integerD, integerD)
	fmt.Printf("%v, %T \n", integerE, integerE)
}
