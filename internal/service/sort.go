package service

func CountingSort(arr []int, k int) []int {
	n := len(arr)

	count := make([]int, k+1)

	for i := 0; i < n; i++ {
		count[arr[i]]++
	}

	for i := 1; i <= k; i++ {
		count[i] += count[i-1]
	}

	output := make([]int, n)

	for i := n - 1; i >= 0; i-- {
		output[count[arr[i]]-1] = arr[i]
		count[arr[i]]--
	}

	return output
}
