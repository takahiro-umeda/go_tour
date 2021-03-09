package main

import (
	"fmt"
	"math"
	"strings"
)

func pointers(){
	i, j := 42, 2701

	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)
}
type Vertex struct {
	X int
	Y int
}

func structs(){

	fmt.Println(Vertex{1, 2})
}

func structsField(){
	v := Vertex{1,2}
	v.X = 4
	fmt.Println(v)
}

func pointersToStructs(){
	v := Vertex{1,2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}

func structLiterals(){
	var(
		v1 = Vertex{1,2}
		v2 = Vertex{X: 1}
		v3 = Vertex{}
		p = &Vertex{1,2}
	)
	fmt.Println(v1, p, v2, v3)
}

func arrays(){
	var a [2]string
	a[0] = "hello"
	a[1] = "world"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2,3,5,7,11,13}
	fmt.Println(primes)
}

type Person struct {
	name string
	age int
}

// 参照渡しの書き方勉強1
func caller(){
	person := Person{"Kevin", 32}
	called(&person)
	fmt.Println(person)
}

// 参照渡しの書き方勉強2
func called(p *Person){
	p.age = 31
}

func slices7()  {
	primes := [6]int{2,3,5,7,13}
	var s []int = primes[1:5]
	fmt.Println(s)
}

func slice8()  {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}

func sliceLiteral(){
	q := []int{2,3,5,7,11,13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)

}

func sliceDefaults()  {
	s := []int{2,3,5,7,11,13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func sliceLengthAndCapacity()  {
	s := []int{2,3,5,7,11,13}
	printSlice(s)

	s = s[:0]
	printSlice(s)

	s = s[:4]
	printSlice(s)

	s = s[2:]
	printSlice(s)
}
//$ gohr main
//Reloading... `main` was built!
//
//----- Results -----
//len=6 cap=6 [2 3 5 7 11 13]
//len=0 cap=6 []
//len=4 cap=6 [2 3 5 7]
//len=2 cap=4 [5 7]

func nilSlices() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}

func createSliceWithMake()  {
	a := make([]int, 5)
	printSlice2("a", a)

	b := make([]int, 0, 5)
	printSlice2("b", b)

	c := b[:2]
	printSlice2("c", c)

	d := c[2:5]
	printSlice2("d", d)

}

func printSlice2(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

func slicesOfSlices()  {
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	board[0][0] = "X"
	board[2][2] = "◯"
	board[1][2] = "X"
	board[1][0] = "◯"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func appendToSlice()  {
	var s []int
	printSlice(s)

	s = append(s, 0)
	printSlice(s)

	s = append(s, 1)
	printSlice(s)
	s = append(s, 2,3,4)
	printSlice(s)
}

// each_with_index的な？
func range16()  {
	var pow = []int{1,2,4,8,16,32,64,128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

func rangeContinued()  {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}

// 18のgo-tourのエディタで入力したやつ
//func Pic(dx, dy int) [][]uint8 {
//	var board = make([][]uint8, dx)
//	for dxIdx := 0; dxIdx < dx; dxIdx++ {
//		var row = make([]uint8, dy)
//		for dyIdx := 0; dyIdx < dy; dyIdx++ {
//			row[dyIdx] = uint8((dxIdx + dyIdx) / 2)
//		}
//		board[dxIdx] = row
//	}
//	return board
//}
//
//func main() {
//	pic.Show(Pic)
//}

type GeoVertex struct {
	Lat, Long float64
}

var m map[string]GeoVertex

func maps(){
	m = make(map[string]GeoVertex)
	m["BEll Labs"] = GeoVertex{}
	fmt.Println(m["BEll Labs"])
}

func mapLiterals(){
	var m = map[string]GeoVertex{
		"Bell Labs": GeoVertex{40.68433, -74.39967},
		"Google": GeoVertex{37.42202, -122.08408,},
	}
	fmt.Println(m)
}



func mapLiterals2(){
	var m = map[string]GeoVertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google": {37.42202, -122.08408,},
	}
	fmt.Println(m)
}

func mapMutating(){
	m := make(map[string]int)
	var key string = "Answer"

	m[key] = 42
	fmt.Println("The value", m[key])

	m[key] = 56
	fmt.Println("The value", m[key])

	delete(m, key)
	fmt.Println("The value", m[key])

	v, ok := m[key]
	fmt.Println("The value:", v, "Present?", ok)
}

//func WordCount(s string) (results map[string]int){
//	var words []string = strings.Fields(s)
//	for i := 0; i < len(words); i++ {
//		var _, ok =
//	}
//	return
//}
//
//func mapExercise(){
//	fmt.Println(WordCount("x")["x"] == 1)
//	fmt.Println(WordCount("xxx")["xxx"] == 3)
//}

func functionValues(){
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(5, 12))
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func functionClosures(){
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2 * i))
	}
}

func fibonacci() func() int {
	prev := 1
	prevBeforePrev := -1
	current := 1
	return func () int {
		current = prev + prevBeforePrev
		prevBeforePrev = prev
		prev = current
		return current
	}
}

func fibonacci2() func() int {
	prev := 0
	current := 1
	return func () int {
		prevBeforePrev := prev
		prev = current
		current = prevBeforePrev + prev
		return prevBeforePrev
	}
}

func fibonacciClosure(){
	f := fibonacci2()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

func main(){
	//pointers()
	//structs()
	//structsField()
	//pointersToStructs()
	//structLiterals()
	//arrays()
	//caller()
	//slices7()
	//slice8()
	//sliceLiteral()
	//sliceDefaults()
	//sliceLengthAndCapacity()
	//nilSlices()
	//createSliceWithMake()
	//slicesOfSlices()
	//appendToSlice()
	//range16()
	//rangeContinued()
	//maps()
	//mapLiterals()
	//mapLiterals2()
	//mapMutating()
	//mapExercise()
	//hoge()
	//mapMutating()
	//functionValues()
	//functionClosures()
	fibonacciClosure()
}





















