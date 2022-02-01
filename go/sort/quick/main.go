package main

import "fmt"

type iList []int

func (list *iList) Swap(i int, j int) {
	fmt.Printf("[%d, %d] Swap ", (*list)[i], (*list)[j])
	(*list)[i], (*list)[j] = (*list)[j], (*list)[i]
}

func (list *iList) Quick() {
	if len(*list) < 2 {
		return
	}

	pivot := list.Divide()

	left := (*list)[0:pivot]
	right := (*list)[pivot+1 : len(*list)]

	left.Quick()
	right.Quick()
}

func (list *iList) Divide() int {
	pivot := (*list)[0]

	low := 1
	high := len(*list) - 1

	for low < high {
		// 피벗보다 작은 수의 인덱스를 찾기 위해
		// 뒤에서부터 pivot 보다 큰 수면 high - 1
		for (*list)[high] > pivot {
			high -= 1
		}
		// 피벗보다 큰 수의 인덱스를 찾기 위해
		// 앞에서부터 pivot 보다 작은수면 low + 1
		// 뒤에 low < high 가 오는 이유는 high 마지막 연산에서 low 인덱스까지 넘어와서 low 가 high 인덱스로 넘어가지 않게 하기 위함
		for (*list)[low] <= pivot && low < high {
			low += 1
		}
		(*list).Swap(low, high)
	}

	fmt.Printf("\n")

	fmt.Println("divide:", *list)

	(*list).Swap(0, low)

	fmt.Printf("\n")

	fmt.Println("Swap pivot:", *list)

	fmt.Printf("\n")

	return low
}

func main() {
	arr := iList{4, 0, 3, 7, 2, 1, 9, 6, 5, 8}
	arr.Quick()
	fmt.Println(arr)
}
