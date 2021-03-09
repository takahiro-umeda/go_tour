package main

import (
	"fmt"
	"math"
	"math/cmplx"
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

func basicType(){
	var (
		ToBe bool = false
		MaxInt uint64 = 1<<64 - 1
		z complex128 = cmplx.Sqrt(-5 + 12i)
	)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}

func zeroValues(){
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

func typeConversions(){
	var x, y int = 3,4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}

func typeInference(){
	v := 42
	fmt.Printf("v is of type %T\n", v)
}

func constants(){
	const World = "世界"
	fmt.Println("hello", World)
	fmt.Println("happy", math.Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}

func numericConstrants(){
	const (
		// Create a huge number by shifting a 1 bit left 100 places.
		// In other words, the binary number that is 1 followed by 100 zeroes.
		Big = 1 << 100
		// Shift it right again 99 places, so we end up with 1<<1, or 2.
		Small = Big >> 99
	)
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	//packages()
	// exportedNames()
	//functionContinued()
	//multipleResults()
	//namedReturnValues()
	//variables()
	//variablesWithInitializers()
	//shortVariableDeclarations()
	//basicType()
	//zeroValues()
	//typeConversions()
	//typeInference()
	numericConstrants()
}

