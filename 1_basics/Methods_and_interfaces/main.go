package main

import (
	"fmt"
	"image"
	"io"
	"math"
	"strings"
	"time"
)

// 20200819 ~P8

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

// めっちゃjsのprototypeっぽい。＆:=がデフォっぽい。冗長な型指定はしないんだ。
func methods(){
	v := Vertex{3,4}
	fmt.Println(v.Abs())
}

type Vertex2 struct {
	X, Y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func functionInsteadOfMethod(){
	v := Vertex{3,4}
	fmt.Println(Abs(v))
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// ~~めっちゃオープンクラスやモンキーパッチっぽい~~ ←却下。値オブジェクトやん。
func methodOfNonStructType() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}

type Vertex4 struct {
	X, Y float64
}

func (v Vertex4) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex4) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func pointerReceiver(){
	v := Vertex4{3,4}
	v.Scale(10)
	fmt.Println(v.Abs())
}

type Vertex5 struct {
	X, Y float64
}

func Abs5(v Vertex5) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *Vertex5, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func pointersAndFunctions(){
	v := Vertex5{3, 4}
	Scale(&v, 10)
	fmt.Println(Abs5(v))
}

type Vertex6 struct {
	X, Y float64
}

func (v Vertex6) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex6, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func methodsAndPointerIndirection(){
	v := Vertex6{3,4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &Vertex6{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}

type Vertex7 struct {
	X, Y float64
}

func (v Vertex7) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc7(v Vertex7) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func methodsAndPointerIndirection2(){
	v := Vertex7{3,4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc7(v))

	p := Vertex7{4,3}
	fmt.Println(p.Abs())
	fmt.Println(AbsFunc7(p))
}

type Vertex8 struct {
	X, Y float64
}

func (v *Vertex8) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex8) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func choosingAValueOrPointerReceiver(){
	 v := &Vertex8{3,4}
	 fmt.Println("Before scaling: %+v, Abs: %+v\n", v, v.Abs())
	 v.Scale(5)
	fmt.Println("Before scaling: %+v, Abs: %+v\n", v, v.Abs())
}

// 20200826 ~P14

type Abser interface {
	Abs() float64
}

type MyFloat9 float64

func (f MyFloat9) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex9 struct {
	X, Y float64
}

func (v *Vertex9) Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func interfaces() {
	var a Abser
	f := MyFloat9(-math.Sqrt2)
	v := Vertex9{3,4}

	a = f
	a = &v

	//a = v

	fmt.Println(a.Abs())
}

type I interface {
	M()
}

type T struct {
	S string
}

func (t T) M() {
	fmt.Println(t.S)
}

func implementedImplicitlyInterface(){
	var i I = T{"hello"}
	i.M()
}

type I11 interface {
	M11()
}

type T11 struct {
	S11 string
}

func (t *T11) M11(){
	fmt.Println(t.S11)
}

type F11 float64

func (f F11) M11() {
	fmt.Println(f)
}

func describe(i I11) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func interfaceValues() {
	var i I11

	i = &T11{"Hello"}
	describe(i)
	i.M11()

	i = F11(math.Pi)
	describe(i)
	i.M11()
}

type I12 interface {
	M()
}

type T12 struct {
	S string
}

func (t *T12) M() {
	// レシーバーがなくてもメソッドを実行できる
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func describe12(i I12) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func interfaceValuesWithNilUnderlyingValues() {
	var i I12

	var t *T12
	i = t
	describe12(i)
	i.M()

	i = &T12{"hello"}
	describe12(i)
	i.M()
}

type I13 interface {
	M()
}

func describe13(i I13) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func nilInterfaceValues() {
	var i I13
	describe13(i)
	i.M()
	// 実行時エラー
	// panic: runtime error: invalid memory address or nil pointer dereference
}

func emptyInterface() {
	// 他で言うところのobject?
	var i interface{}
	describe14(i)

	i = 42
	describe14(i)

	i = "Hello"
	describe14(i)
}

func describe14(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

// 20200909 ~P26(Exercise除く)
func typeAssertion(){
	var i interface{} = "hello"
	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	//f := i.(float64)
	//fmt.Println(f, ok)
}

func typePrint(i interface{}){
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T", v)
	}
}

func typeSwitches(){
	typePrint(21)
	typePrint("hello")
	typePrint(true)
}

type Person struct {
	Name string
	Age int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

// toStringなやつ
func stringers(){
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	 return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func errors(){
	if err := run(); err != nil {
		fmt.Println(err) // Error()が実行されている理由はrunの返り値がerrorインタフェースだから。Stringerっぽく文字列化を求められたらError()を実行する
	}
}

func readers(){
	r := strings.NewReader("Hello, Reader!")
	b := make([]byte, 4) // 8) 何バイトずつ取得するか定義。スライスがいっぱいになるまで取得する

	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF { // ← Errorに末尾判定が入ってくるの嫌い。statusってものだったらわかりやすいのに。errにstatusとerrorの2つの意味をもたせてる
			break
		}
	}
}

func images(){
	m := image.NewRGBA(image.Rect(0,0,100,100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0,0).RGBA())
}

func main() {
	//methods()
	//functionInsteadOfMethod()
	//methodOfNonStructType()
	//pointerReceiver()
	//methodsAndPointerIndirection()
	//choosingAValueOrPointerReceiver()
	//interfaces()
	//implementedImplicitlyInterface()
	//interfaceValues()
	//interfaceValuesWithNilUnderlyingValues()
	//nilInterfaceValues()
	//emptyInterface()
	//typeAssertion()
	//typeAssertion()
	//typeSwitches()
	//stringers()
	//errors()
	//readers()
	images()
}