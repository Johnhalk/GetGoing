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

	// arrays

	grade1 := 97
	grade2 := 89
	grade3 := 92

	fmt.Printf("Grades: %v, %v, %v\n", grade1, grade2, grade3)

	grades := [...]int{97, 89, 92}
	fmt.Printf("Grades: %v\n", grades)

	var students [3]string
	fmt.Printf("Students: %v\n", students)
	students[0] = "John"
	students[1] = "Eam"
	students[2] = "Amber"
	students[1] = "Eammon"
	fmt.Printf("Students: %v\n", students)
	fmt.Printf("Students: %v\n", students[2])

	var identityMatrix [3][3]int
	identityMatrix[0] = [3]int{1, 0, 0}
	identityMatrix[0] = [3]int{0, 1, 0}
	identityMatrix[0] = [3]int{0, 0, 1}
	fmt.Printf("identityMatrix: %v\n", identityMatrix)

	//array pointer => &a

	a :=[...]int{1, 2, 3, 4}
	b := &a
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)

	// SLICES

	c :=[]int{1, 2, 3, 4}
	d := c
	d[1] = 5
	fmt.Printf("Length: %v\n", len(c))
	fmt.Printf("Capacity: %v\n", cap(c))


	sliceA := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sliceB := sliceA[:] // slice of all elements
	sliceC := sliceA[3:] // slice from 4th element to end
	sliceD := sliceA[:6] // slice first 6 elements
	sliceE := sliceA[3:6] // slice 4th, 5th and 6th elements

	fmt.Println(sliceB)
	fmt.Println(sliceC)
	fmt.Println(sliceD)
	fmt.Println(sliceE)


}
