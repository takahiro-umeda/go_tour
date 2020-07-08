package main

import (
	"fmt"
	"math"
	"math/rand"
)

func packages(){
	fmt.Println("My favorite number is", rand.Intn(10))
}

func exportedNames(){
	fmt.Println(math.Pi)
}

func add(x int, y int) int {
	return x + y
}

func functions(){
	fmt.Println(add(42, 13))
}

func add2(x, y int) int {
	return x + y
}

func functionContinued(){
	fmt.Println(add2(33,4))
}

func swap(x, y string) (string, string){
	return y, x
}

func multipleResults(){
	a, b :=swap("hello", "world")
	fmt.Println(a,b)
}

func split(sum int)(x, y int){
	x = sum + 2 //* 4 / 9
	y = sum - x
	return
}

func namedReturnValues(){
	fmt.Println((split(17)))
}

//var c, python, java bool

func variables(){
	//var i int
	//fmt.Println(i, c, python, java)
}

var i, j int = 1, 2

func variablesWithInitializers(){
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}

// 10ページ目
func shortVariableDeclarations(){
	var i, j int = 1,2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}

func main() {
	//packages()
	// exportedNames()
	//functionContinued()
	//multipleResults()
	//namedReturnValues()
	//variables()
	//variablesWithInitializers()
	shortVariableDeclarations()
}
