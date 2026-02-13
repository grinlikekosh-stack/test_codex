package main

func findSmallest(arr []int) int {
	smallest := arr[0]
	smallestIndex := 0

	for i := range len(arr) {
		if arr[i] < smallest {
			smallest = arr[i]
			smallestIndex = i
		}
	}
	return smallestIndex
}

func selectionSort(arr []int) []int {
	newArr := make([]int, len(arr))
	for i := range len(arr) {
		smallestIndex := findSmallest(arr)
		newArr[i] = arr[smallestIndex]

		arr = append(arr[:smallestIndex], arr[smallestIndex+1:]...)
	}
	return newArr
}

func main() {
	arr := []int{5, 3, 6, 2, 10}
	sortedArr := selectionSort(arr)

	println("Sorted array using selection sort:")
	for _, v := range sortedArr {
		print(v, " ")
	}
	println()
}
