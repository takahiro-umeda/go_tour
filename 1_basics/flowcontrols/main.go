package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func p(str string){
	fmt.Println(str + "\n")
}

func for1(){
	sum := 0
	for i := 0; i < 10; i++ {
		sum += 1
	}
	fmt.Println(sum)
}

func for2(){
	sum := 1
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
}

func forWhileVer(){
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum + 1)
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func if1(){
	fmt.Println(sqrt(2), sqrt(-4))
}

func pow1(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func if2(){
	fmt.Println(
		pow1(3,2,10),
		pow1(3,3,20),
		)
}

func pow2(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Println("%g >= %g\n", v, lim)
	}
	return lim
}

func ifElse(){
	fmt.Println(
		pow2(3, 2, 10),
		pow2(3, 3, 20),
	)
}

func switch_9(){
	fmt.Println("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.", os)
	}
}

func switch_10(){
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away")
	}
}

func switch_11(){
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

func defer12(){
	defer fmt.Println("world")

	fmt.Println("hello")
}

func stackingDefers(){
	fmt.Println("counting")
	for i:= 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}

func main(){
	//for1()
	//for2()
	//forWhileVer()
	//if1()
	//if2()
	//ifElse()
	//switch_9()
	//switch_10()
	//switch_11()
	//defer12()
	stackingDefers()
}