package main

import (
	"fmt"
	"strconv"
	"reflect"
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

	sliceF := make([]int, 10) // create slice with capacity and length == 10
	sliceG := make([]int, 10, 100) // create slice with length == 10 and capacity == 100

	fmt.Println(sliceF)
	fmt.Println(sliceG)

	// maps and structs
	statePopulations := make(map[string]int)
	statePopulations = map[string]int{
		"California": 39250017,
		"Texas": 23409029,
		"Florida": 34324521,
		"New York": 3242432234,
		"Ohio": 12938901,
	}
	// Adds in to map
	statePopulations["Georgia"] = 12314321

	_, ok := statePopulations["Oho"]
	fmt.Println(ok)
	fmt.Println(statePopulations)
	fmt.Println(len(statePopulations))


	type Doctor struct {
		number int
		actorName string
		companions []string
		episodes []string
	}

	aDoctor := Doctor{
		number: 3,
		actorName: "Jon Pertwee",
		episodes: []string {
			"Episode 1",
			"Episode 2",
		},
		companions: []string {
			"Liz Shaw",
			"Jo Grant",
			"Sarah Jane Smith",
		},
	}

	fmt.Println(aDoctor);
	fmt.Println(aDoctor.companions);
	fmt.Println(aDoctor.episodes[1]);

	anotherDoctor := struct{name string}{name: "John Pertwee"}
	anotherOtherDoctor := &anotherDoctor
	anotherOtherDoctor.name = "Tom Baker"
	fmt.Println(anotherDoctor)
	fmt.Println(anotherOtherDoctor)

	type Animal struct {
		Name string
		Origin string
	}

	type Bird struct {
		Animal
		SpeedKPH float32
		CanFly bool
	}

	bird := Bird{}
	bird.Name = "Emu"
	bird.Origin = "Australia"
	bird.SpeedKPH = 48
	bird.CanFly = false
	fmt.Println(bird.Name)

	// tags in structs

	type Element struct {
		Name string `required max:"100"`
		Hot bool
	}

	t := reflect.TypeOf(Element{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)

	// if statements

	if false {
		fmt.Println("I will never print")
	}

	// comparison operators
	number := 50
	guess := 30

	if guess < 1 {
		fmt.Println("The guess must be greater than 1!")
	} else if guess > 100 {
		fmt.Println("The guess must be less than 100!")
	} else {
		if guess < number {
		fmt.Println("Too low")
		}
		if guess > number {
			fmt.Println("Too high")
		}
		if guess == number {
			fmt.Println("You got it!")
		}
	}

	fmt.Println(number <= guess, number >= guess, number != guess)

	it := 10

	// example of fallthrough, will hit next switch case rather than terminate
	switch {
	case it <= 10:
		fmt.Println("less than or qual to ten")
		fallthrough
	case it <= 20:
		fmt.Println("less than or equal to twenty")
	default:
		fmt.Println("greater than twenty.")
	}

	var inter interface{} = [3]int{3, 4, 5}

	switch inter.(type) {
	case int:
		fmt.Println("inter is an int")
	case string:
		fmt.Println("inter is string")
	case float64:
		fmt.Println("inter is a float64")
	case float32:
		fmt.Println("inter is a float32")
	default:
		fmt.Println("other type.")
	}

	// for loops

	for value, value2 := 0, 0; 
	value < 5; 
	value, value2 = value +1, value +2 {
		fmt.Println("value ", value)
		fmt.Println("value2 ", value2)
	}

	numbers := 0
	value := 0
	for {
		fmt.Println(numbers)
		numbers++
		if numbers == 5 {
			break
		}
	}


	numbers = 0
	for ; numbers < 10; numbers++ {
		fmt.Println(numbers)
		if numbers == 5 {
			break
		}
	}

	Loop:
		for numbers = 1; numbers <= 3; numbers++ {
			for value = 0; value <= 3; value++ {
				fmt.Println(" loop ", numbers * value)
				if numbers * value >= 3 {
					break Loop
				}
			}
		}


	s := []int{1, 2, 3}

	for k, v := range s{
		fmt.Println(k,v)
	}

	for k, v := range statePopulations {
		fmt.Println(k, v)
	}

	for _, v := range statePopulations {
		fmt.Println(v)
	}

	for k := range statePopulations {
		fmt.Println(k)
	}
}
