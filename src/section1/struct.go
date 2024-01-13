package main

import (
	"fmt"
	"unsafe"
)

type Task struct {
	Title    string
	Estimate int
}

func sampleStruct() {
	task1 := Task{
		Title:    "task1",
		Estimate: 60,
	}

	task1.Title = "Learning Go!!"
	// %+vでstructのフィールド名も表示
	fmt.Printf("%[1]T %+[1]v %v\n", task1, task1.Title)

	// structは値型なのでtask2にtask1を代入してもtask1の値は変わらない
	var task2 Task = task1
	task2.Title = "new"

	task1p := &Task{
		Title:    "Learn concurrency",
		Estimate: 2,
	}

	// task1pはポインタ型なので*task1pで値を取得
	fmt.Printf("task1p: %T %+v %v\n", task1p, *task1p, unsafe.Sizeof(task1p))

	(*task1p).Title = "Changed"
	// 構造体のdereferenceはtask1p.Titleと省略してかける
	task1p.Title = "Changed2"
	fmt.Printf("task1p: %+v\n", *task1p)

	var task2p *Task = task1p
	task2p.Title = "Changed3"
	fmt.Printf("task1p: %+v\n", *task1p)
	fmt.Printf("task2p: %+v\n", *task2p)

	// これは値型なのでtask1pの値は変わらない
	task1.extendEstimate()
	// これはポインタ型なのでtask1pの値も変わる
	(&task1).extendEstimatePointer()
	task1.extendEstimatePointer()
}

// 値レシーバーは呼び出されたときに値のコピーが作成される
func (task Task) extendEstimate() {
	task.Estimate += 10
}

// ポインタレシーバーは呼び出されたときそのポインタが渡される
func (taskp *Task) extendEstimatePointer() {
	taskp.Estimate += 10
}
