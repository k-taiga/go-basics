package section1

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func funcDefer() {
	defer fmt.Println("final defer")
	defer fmt.Println("semi defer")
	fmt.Println("hello world")
}

// ...で可変長引数を受け取る(1つでも複数でも可)
func trimExtension(files ...string) []string {
	out := make([]string, 0, len(files))
	// trimしてoutにappend
	for _, file := range files {
		out = append(out, strings.TrimSuffix(file, ".csv"))
	}
	return out
}

func fileChecker(name string) (string, error) {
	f, err := os.Open(name)
	if err != nil {
		return "", errors.New("file not found")
	}

	// 開いていたリソースを閉じる
	defer f.Close()
	return name, nil
}

// 関数を引数に取る関数 引数の関数の型は同じにしないといけない
func addExt(f func(file string) string, name string) {
	// 引数の関数を実行
	fmt.Println(f(name))
}

// 関数を返す関数
func multiPly() func(int) int {
	return func(n int) int {
		return n * 1000
	}
}

func countUp() func(int) int {
	count := 0

	// この関数は常に同じcountを参照する countは返却されていないため
	// クロージャー 簡単に言うと「関数の中の関数」で、外側の関数から変数を「覚えて」使える特別な関数
	return func(n int) int {
		count += n
		return count
	}
}

// var count int

func main() {
	funcDefer()
	files := []string{"file1.csf", "file2.csv", "file3.csv"}
	// 可変長引数を渡すときは...をつける
	fmt.Println(trimExtension(files...))
	name, err := fileChecker("file.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(name)

	i := 1

	// 無名関数 ()をつけると即座に実行される ()には渡す引数を書く
	func(i int) {
		fmt.Println(i)
	}(i)

	// f1に無名関数を代入 変数に入れておくと好きなタイミングで実行できる
	f1 := func(i int) int {
		return i + 1
	}

	fmt.Println(f1(i))

	f2 := func(file string) string {
		return file + ".csv"
	}

	// 関数を引数で渡して実行
	addExt(f2, "file")

	f3 := multiPly()
	fmt.Println(f3(2))

	f4 := countUp()
	for i := 0; i < 5; i++ {
		v := f4(2)
		// 2 4 6 8 10
		fmt.Println("%v\n", v)
	}
}
