package main

import "fmt"

func main() {
	q1()
	q2()
}

// 赤緑青を0, 1, 2とする
// 赤緑青の順になるように並べ替える関数を作成する
// 順番はバラバラなので [2, 1, 0, 0, 0, 1] -> [0, 0, 0, 1, 1, 2]
func q1() {
	fmt.Println("Q1")

	input := []int{2, 1, 0, 1, 2}

	sortColors := func(colors []int) {
		countRed, countGreen, countBlue := 0, 0, 0
		for _, v := range colors {
			if v == 0 {
				countRed++
			}
			if v == 1 {
				countGreen++
			}
			if v == 2 {
				countBlue++
			}
		}

		for i := range colors {
			if countRed > 0 {
				colors[i] = 0
				countRed--
			} else if countGreen > 0 {
				colors[i] = 1
				countGreen--
			} else {
				colors[i] = 2
				countBlue--
			}
		}
	}
	sortColors(input)
	fmt.Println(input)

	inputOneLook := []int{2, 1, 0, 1, 2}

	sortColorsOneloop := func(colors []int) {
		left, right := 0, len(colors)-1
		current := 0
		for current <= right {
			if colors[current] == 0 {
				colors[left], colors[current] = colors[current], colors[left]
				left++
				current++
			} else if colors[current] == 1 {
				current++
			} else if colors[current] == 2 {
				colors[right], colors[current] = colors[current], colors[right]
				right--
			}
		}
	}
	sortColorsOneloop(inputOneLook)
	fmt.Println(inputOneLook)
}

func q2() {
	type node struct {
		value int
		left  *node
		right *node
	}
	n := node{}
	fmt.Println(n)

	var isSameTree func(a, b *node) bool
	isSameTree = func(a, b *node) bool {
		if a == nil && b == nil {
			return true
		}
		if a == nil || b == nil {
			return false
		}
		return a.value == b.value && isSameTree(a.left, b.left) && isSameTree(a.right, b.right)
	}

	var isSubtree func(a, b *node) bool

	isSubtree = func(a, b *node) bool {
		if b == nil {
			return true
		}
		if a == nil {
			return false
		}
		return isSameTree(a, b) || isSubtree(a.left, b) || isSubtree(a.right, b)
	}
}
