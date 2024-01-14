package section1

import "fmt"

func forFunc() {
	// forをloopとしてラベルをつける
loop:
	for i := 0; i < 10; i++ {
		switch {
		case i == 3:
			continue
		case i == 5:
			continue
		// loopを抜ける
		case i == 7:
			break loop
		default:
			fmt.Printf("%v ", i)
		}
	}

	items := []item{
		{price: 100},
		{price: 200},
		{price: 300},
	}

	// これはpriceの値が変わらない rangeは値渡しでコピーが渡されるため
	for _, i := range items {
		i.price *= 1.1
	}

	// itemsの中身を変更するためにはindexでアクセスする必要がある
	for i := range items {
		items[i].price *= 1.1
	}

}

type item struct {
	price float32
}
