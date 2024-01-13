package main

import (
	"fmt"
	"unsafe"
)

type Os int

const (
	// iotaは連番を生成する
	Mac Os = iota + 1
	Windows
	Linux
)

var (
	// 下記の内容で初期値が設定される
	// 0
	i int
	// ""
	s string
	// false
	b bool
)

func variable() {
	i := 1
	ui := uint16(2)

	fmt.Println(i)
	// １つ目の引数は%、2つ目の引数は%[2]
	fmt.Printf("i: %[1]v %[1]T ui: %[2]v %[2]T\n", i, ui)

	f := 1.23456
	s := "string"
	b := true
	fmt.Printf("f: %v %T s: %v %T b: %v %T\n", f, f, s, s, b, b)

	pi, title := 3.14, "Go"
	fmt.Printf("pi: %v %T title: %v %T\n", pi, pi, title, title)

	x := 10
	y := 1.23
	z := float64(x) + y
	fmt.Printf("z: %v %T\n", z, z)

	const secret = "SECRET"
	fmt.Printf("secret: %v %T\n", secret, secret)

	fmt.Printf("Mac: %v Windows: %v Linux: %v\n", Mac, Windows, Linux)

	// 0で初期化されている uint16は2byte持つ
	var ui1 uint16
	// 先頭の1byteのアドレスを表示
	fmt.Printf("memory address of ui1: %p\n", &ui1)

	var ui2 uint16
	fmt.Printf("memory address of ui2: %p\n", &ui2)

	// uint16のポインタの変数 先頭の番地のアドレスを持つ 先頭からどれくらいの番地まであるかわからないためポインタ変数の型を定義をする
	var p1 *uint16
	fmt.Printf("memory address of p1: %p\n", &p1)
	p1 = &ui1
	fmt.Printf("memory address of p1: %p\n", &p1)
	// ポインタ変数のサイズは8byte
	fmt.Printf("size of p1: %d[byte]\n", unsafe.Sizeof(p1))
	// p1のポインタの先頭の番地のアドレスを表示 ポインタのポインタ ダブルポインタ
	fmt.Printf("memory address of p1: %p\n", &p1)
	// メモリの指し示す値を表示 * = dereference
	fmt.Printf("value of ui1: %v\n", *p1)
	// dereference先のui1の値を変更
	*p1 = 1

	// p1のポインタの先頭の番地のアドレス ダブルポインタ（ポインタのポインタ）の変数
	var pp1 **uint16 = &p1
	fmt.Printf("value of pp1: %p\n", pp1)
	fmt.Printf("memory address of pp1: %p\n", &pp1)
	fmt.Printf("size of pp1: %d[byte]\n", unsafe.Sizeof(pp1))
	fmt.Printf("value of p1: %v\n", *pp1)
	// ダブルポインタのためdereferenceが2回必要
	fmt.Printf("value of ui1: %v\n", **pp1)
	**pp1 = 10
	fmt.Printf("value of ui1: %v\n", ui1)

	// ok, result := true, "A"
	// // resultはifのスコープ内で別の変数として定義
	// if ok {
	// 	result := "B"
	// 	println(result)
	// } else {
	// 	result := "C"
	// 	println(result)
	// }
	// // B,Aが表示される
	// println(result)

	ok, result := true, "A"
	// resultはifのスコープ内
	if ok {
		result = "B"
		println(result)
	} else {
		result = "C"
		println(result)
	}
	// B,Bが表示される
	println(result)
}
