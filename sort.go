package main

import (
	"fmt"
	"math/rand"
	"time" 
)

func swap_int(a *int, b *int) {
	*a = *a + *b
	*b = *a - *b
	*a = *a - *b
}

func rand_int(rnd *rand.Rand,src []int, size int) {
	var t int
	for i := 0; i < size; i++ {
		t = (*rnd).Intn(50)
		//fmt.Println(t)
		src[i] = t
	}
}

func iter(src []int, size int) {
	for i := 0; i < size; i++ {
		fmt.Print(src[i])
		fmt.Print(" ")
	}
	fmt.Println()
}

func sort(src []int, size int) {
	for i := 0; i < size - 1; i++ {
		for j := i + 1; j < size; j++ {
			if(src[i] < src[j]) {
				swap_int(&src[i], &src[j])
			}	
		}
	}
}

func quick_sort(src []int, l int, h int) {

	if (h > l) {
		var point int = src[l]
		var low int = l
		var high int = h

		for;(low < high); {
			for;(src[high] >= point &&  low < high); {
				high--;
			}
			src[low] = src[high]

			for;(src[low] <= point && low < high);{
				low++;
			}
			src[high] = src[low]
		}
		src[low] = point;

		//iter(src, len(src))

		quick_sort(src, l, low - 1)
		quick_sort(src, low + 1, h)
	}
}

func quick(src []int, size int) {
	 quick_sort(src , 0, size - 1)
}


func main() {

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))  
	var src = make([]int, 7)
	rand_int(rnd, src, 7)

	iter(src, len(src))
	//sort(src, len(src))

	quick(src, len(src))

	iter(src, len(src))
}