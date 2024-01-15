package section1

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type customConstraints interface {
	~int | int16 | float64 | float32 | string
}

// 独自の型もgenericsに含める事ができる ~で指定する
type NewInt int

func add[T customConstraints](x, y T) T {
	return x + y
}

// これがorderedの定義
// type Ordered interface {
// Integer | Float | ~string
// }
func min[T constraints.Ordered](x, y T) T {
	if x < y {
		return x
	}
	return y
}

// mapのkはKのgenerics, vはVのgenerics
func sumValues[K int | string, V constraints.Float | constraints.Integer](m map[K]V) V {
	var sum V
	for _, v := range m {
		sum += v
	}
	return sum
}

func generics() {
	// genericsを使うと型を指定しなくてもよくなる
	fmt.Printf("%v\n", add(1, 2))
	fmt.Printf("%v\n", add(1.1, 2.2))
	fmt.Printf("%v\n", add("hello", "world"))

	var i1, i2 NewInt = 1, 2
	// NewIntはintの独自の型だが~intでgenericsに含まれているのでエラーにならない
	fmt.Printf("%v\n", add(i1, i2))
	fmt.Printf("%v\n", min(1, 2))

	m1 := map[string]uint{"a": 1, "b": 2, "c": 3}
	m2 := map[int]float32{1: 1.1, 2: 2.2, 3: 3.3}

	fmt.Println(sumValues(m1))
	fmt.Println(sumValues(m2))
}
