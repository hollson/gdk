package bloom

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/willf/bitset"
)

func TestBloomx(t *testing.T) {
	Foo()
	bar()
}

func Foo() {
	var b bitset.BitSet // 定义一个BitSet对象

	b.Set(1).Set(2).Set(3) // 添加3个元素
	if b.Test(2) {
		fmt.Println("2已经存在")
	}
	fmt.Println("总数：", b.Count())

	b.Clear(2)
	if !b.Test(2) {
		fmt.Println("2不存在")
	}
	fmt.Println("总数：", b.Count())
}

func bar() {
	fmt.Printf("Hello from BitSet!\n")
	var b bitset.BitSet
	// play some Go Fish
	for i := 0; i < 100; i++ {
		card1 := uint(rand.Intn(52))
		card2 := uint(rand.Intn(52))
		b.Set(card1)
		if b.Test(card2) {
			fmt.Println("Go Fish!")
		}
		b.Clear(card1)
	}

	// Chaining
	b.Set(10).Set(11)

	for i, e := b.NextSet(0); e; i, e = b.NextSet(i + 1) {
		fmt.Println("The following bit is set:", i)
	}
	// 交集
	if b.Intersection(bitset.New(100).Set(10)).Count() == 1 {
		fmt.Println("Intersection works.")
	} else {
		fmt.Println("Intersection doesn't work???")
	}
}
