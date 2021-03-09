package main

import (
	"fmt"
	"sort"
	"sync"
	"time"
)
import "golang.org/x/tour/tree"

// 20200909 ~6章 この日は5までしか進まなかった。今後のエクササイズをやるか、EffectiveGoを先に読むかは迷う
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func goroutines(){
	// 何をやってるのか、何がメリットなのか全くわからない
	// →並行処理をしているから、後続のsayをやってるときに前にやってるやつが行われることをあらわしたい例
	go say("world")
	say("hello")
}

func sum(s []int, c chan int, wait time.Duration){
	sum := 0
	for _, v := range s {
		sum += v
	}
	time.Sleep(wait * time.Millisecond)
	c <- sum
}

func channels()  {
	s := []int{7, 2, 8, -9, 4, 0}

	// 何をやってるのか、何がメリットなのか全くわからない
	// →cはネットワークプログラミングでプロセス間通信をするときに使ってたような共有キュー。
	// `c<-`でキューイングして<-cでデキュー
	// waitを設定しなかったら前半が早く終わるから順番通りだけど、waitを設定したら後者が先に終わる可能性がある
	// 並行処理だから当たり前だけど、順番が関係あることをやっちゃいけない
	// 真価が発揮されるのは今回みたいにgoルーチンを複数回しているとき
	c := make(chan int)
	go sum(s[:len(s)/2], c, 1000)
	go sum(s[len(s)/2:], c, 2000)
	x, y := <- c, <- c
	fmt.Println(x, y, x+y)
}

func bufferedChannel(){
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	//fmt.Println(<-ch)
	fmt.Println(<-ch)
	ch <- 3
	fmt.Println(<-ch)
}

func fibonacci(n int, c chan int){
	x, y := 0,1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func rangeAndClose(){
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}



func fibonacci5(c, quit chan int){
	x, y := 0,1
	for {
		fmt.Println("wait") // ③for処理開始
		// UNIXネットワークプログラミングのwait的なやつだと推測
		// ゴルーチン間のチャンネルの送受信を監視してトリガーにする？
		// 確かネットワークプログラミングにもselectあった気がする
		select {/*
		推測:
		全てのcaseのできるやつを実行？
		caseでパターンマッチしていると思ったら、c<-xという挿入処理しているところがここしかないため推測。
		というかcaseに副作用起こして成功しているだけか？
		↑一応selectの説明にも書いてあるけど、どれかを実行するまで待ち続ける。defaultを書くと無限ループはdefaultで周りまくる。
		*/
		case c <- x: // ④ cに値を入れる
			x, y = y, x+y
		case <-quit:
			// ↑シグナル的な使い方
			fmt.Println("quit")
			// selectによる監視を解除？
			return
		}
	}
}

func select5(){
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 1; i < 10; i++ {
			// 多分`<-c`は非同期の値待ち？
			fmt.Println(i, ":", <-c) // ①cが入るの待ち  // ⑤cが入ったからcの値を描画（なぜかこの時点でprintされないけど）
			//if i == 5{
			//	quit <- 0
			//	//close(c) // ←だけを入れるとバグる
			//}
		}
		quit <- 0
	}()
	// JSでいうイベント登録？。onClickみたいにチャンネルの送受信時にcaseで実行される？。returnが行われてイベントが解除される
	fibonacci5(c, quit) // ②fibonacci開始
}

type Tree struct {
	Left *Tree
	Value int
	Right *Tree
}

func Walk(t *tree.Tree, ch chan int){
	ch <- t.Value
	if t.Left != nil {
		go Walk(t.Left, ch)
	}
	if t.Right != nil {
		go Walk(t.Right, ch)
	}
}

func extractTreeValues(t *tree.Tree) []int {
	ch := make(chan int, 10)
	var result []int

	go Walk(t, ch)

	for i := 0; i < 10 ; i++ {
		result = append(result, <-ch)
	}
	fmt.Println(result)
	return result
}

func isSameSlice(a, b []int) bool {
	sort.Ints(a)
	sort.Ints(b)

	existDifferentValueFlg := false
	for i := 0; i < 10 ; i++ {
		if a[i] != b[i] {
			existDifferentValueFlg = true
		}
	}
	return !existDifferentValueFlg
}

func Same(t1, t2 *tree.Tree) bool {
	t1Values := extractTreeValues(t1)
	t2Values := extractTreeValues(t2)

	return isSameSlice(t1Values, t2Values)
}

func exerciseBinalyTree(){
	fmt.Println(Same(tree.New(2), tree.New(1)))
}

// 20200930 P9

type SafeCounter struct {
	// muはどこまで作用するの？muを持つstruct?
	mu sync.Mutex
	v map[string]int
}

func (c *SafeCounter) Inc(key string){
	c.mu.Lock()
	c.v[key]++
	c.mu.Unlock()
}

func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.v[key]
}

func mutex(){
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++  {
		go c.Inc("somekey")
		time.Sleep(1)
	}

	time.Sleep(2)
	// 実行結果から見て1000並行処理が全部実行されているから、
	// ロックされているからいじれないというよりも、
	// ロックされていたら待つみたいな感じっぽそう
	fmt.Println(c.Value("somekey"))
}

func main()  {
	//goroutines()
	//channels()
	//bufferedChannel()
	//bufferedChannel()
	//rangeAndClose()
	//select5()
	//exerciseBinalyTree()
	mutex()
}