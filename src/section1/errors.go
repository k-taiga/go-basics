package section1

import (
	"errors"
	"fmt"
	"os"
)

// sentinel error = 定義済みのエラー
var ErrCustom = errors.New("not found")

func errorsFunc() {
	err01 := errors.New("something wrong")
	err02 := errors.New("something wrong")

	// 0xc00000c0e0 *errors.errorString something wrong
	fmt.Printf("%[1]p %[1]T %[1]v\n", err01)
	fmt.Println(err01.Error())
	// これはError()メソッドを実装しているため中身を自動で表示してくれる
	fmt.Println(err01)
	// Newでは別のポインタでエラーが生成されるのでfalse
	fmt.Println(err01 == err02)

	// fmt.Errorfで%wでerrorsに付加情報を付けることができる
	err0 := fmt.Errorf("add info: %w", errors.New("original error"))
	// 0xc00000c0e0 *fmt.wrapError add info: original error
	fmt.Printf("%[1]p %[1]T %[1]v\n", err0)
	// original errorがUnwrapされる
	fmt.Println(errors.Unwrap(err0))
	// %vでもerrorsにwrapできる
	err1 := fmt.Errorf("add info: %v", errors.New("original error"))
	fmt.Println(err1)
	// 0xc00000c0e0 *errors.errorString add info: original error %wと違い*fmt.wrapErrorではないのでUnwrapできない
	fmt.Printf("%T\n", err1)
	// nilが返る
	fmt.Println(errors.Unwrap(err1))

	err2 := fmt.Errorf("in repository layer: %w", ErrCustom)
	// in repository layer: not found
	fmt.Println(err2)

	err2 = fmt.Errorf("in service layer: %w", err2)
	// in service layer: in repository layer: not found wrapが累積する
	fmt.Println(err2)

	// errors.Isでエラーの比較ができる Unwrapせずに比較できる
	if errors.Is(err2, ErrCustom) {
		fmt.Println("matched")
	}

	file := "dummy.txt"
	err3 := FileChecker(file)
	if err3 != nil {
		// sentinel errorのos.ErrNotExistが含まれていればtrue
		if errors.Is(err3, os.ErrNotExist) {
			fmt.Println("%v file not found\n", file)
		} else {
			fmt.Println("unknown error")
		}
	}
}

func FileChecker(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return fmt.Errorf("in FileChecker: %w", err)
	}

	defer f.Close()
	return nil
}
