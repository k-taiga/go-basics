package main

import "fmt"

func sliceMap() {
	// 0,0,0
	var a1 [3]int
	var a2 = [3]int{10, 20, 30}
	// 代入された要素数から配列の要素数を特定
	a3 := [...]int{10, 20}
	// len: 2 cap: 2
	fmt.Printf("%v %v\n", len(a3), cap(a3))
	fmt.Printf("%T %T\n", a2, a3)

	// 要素数が空の場合はslice
	var s1 []int
	// 明示的に空で代入する場合は{}で囲む
	s2 := []int{}
	fmt.Printf("s1: %[1]T %[1]v %v %v\n", s1, len(s1), cap(s1))
	fmt.Printf("s2: %[1]T %[1]v %v %v\n", s2, len(s2), cap(s2))
	fmt.Println(s1 == nil)
	// s2は空のsliceのためfalse
	fmt.Println(s2 == nil)
	s1 = append(s1, 1, 2, 3)
	s3 := []int{4, 5, 6}
	// s3の要素を...で展開しs1に追加
	s1 = append(s1, s3...)

	// 要素数が0で2要素分の容量を持つslice
	s4 := make([]int, 0, 2)

	// sliceなので要素数は可変長で増える
	s4 = append(s4, 1, 2, 3, 4)
	// [0 0 0 0]
	s5 := make([]int, 4, 6)
	s6 := s5[1:3]
	s6[1] = 10
	// sliceは参照型なのでs5の要素も変更される
	s6 = append(s6, 2)
	// コピーして新しいsliceを作成
	sc6 := make([]int, len(s5[1:3]))
	// s5[1:3]の要素をsc6にコピー
	copy(sc6, s5[1:3])
	// コピーした要素を変更してもs5の要素は変更されない
	sc6[1] = 12

	s5 = make([]int, 4, 6)
	// 3番目の引数を指定するとその引数までのsliceのメモリを共有する
	fs6 := s5[1:3:3]
	fs6[0] = 6
	fs6[1] = 7
	// これは3つ目のメモリ以降の要素を追加するためもとのs5には影響しない
	fs6 = append(fs6, 8)
	// これもfs6の要素外のため影響しない
	s5[3] = 9

	// map
	var m1 map[string]int
	// 明示的に空で代入する場合は{}で囲む
	m2 := map[string]int{}
	fmt.Printf("%v %v \n", m1, m1 == nil)
	// m2は空のmapのためfalse
	fmt.Printf("%v %v \n", m2, m2 == nil)
	m2["A"] = 10
	m2["B"] = 20
	m2["C"] = 0
	delete(m2, "A")
	// 0, false 存在しない気ワードなので0とfalseが返る
	v, ok := m2["A"]
	// 0, true 存在するキーワードなので0とtrueが返る
	v, ok = m2["C"]
	fmt.Println(v, ok)

	// keyとvalueを取得
	// hashmapなので取り出す順番は不定
	for k, v := range m2 {
		fmt.Printf("%v %v\n", k, v)
	}
}
