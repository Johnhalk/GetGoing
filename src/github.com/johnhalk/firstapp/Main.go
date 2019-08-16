package main

import (
	"fmt"
	"strconv"
)

func main() {
	var (
		actorName = "Maisie Williams"
		noOne = "noone"
		season = 7
	)
	fmt.Println(actorName, noOne, season)

	var (
		i int = 42
		j float32
	)

	j = float32(i)

	fmt.Printf("%v, %T\n", j, j)

	var k string
	k = strconv.Itoa(i)

	fmt.Printf("%v, %T\n", k, k)
}
