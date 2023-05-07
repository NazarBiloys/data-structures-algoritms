package service

import (
	"fmt"
	"math/big"
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"time"
)

func TestBalancedBinary() {
	var trees []AVLTree

	// random with big integer
	for i := 0; i < 10; i++ {
		tree := AVLTree{}

		for j := 0; j < 10; j++ {
			source := rand.NewSource(time.Now().UnixNano())

			random := rand.New(source)

			bigInt := big.NewInt(0).Rand(random, big.NewInt(0).Exp(big.NewInt(10), big.NewInt(100), nil))

			tree.Insert(int(bigInt.Int64()))
		}

		trees = append(trees, tree)
	}

	// random with small integer
	for i := 0; i < 10; i++ {
		tree := AVLTree{}

		for j := 0; j < 10; j++ {
			tree.Insert(int(rand.Int63n(100)))
		}

		trees = append(trees, tree)
	}

	for i := 0; i < 10; i++ {
		tree := AVLTree{}

		startValue := rand.Int63n(100)

		for j := 0; j < 10; j++ {
			startValue = startValue + 1
			tree.Insert(int(startValue))
		}

		trees = append(trees, tree)
	}

	// random with big integer
	for i := 0; i < 10; i++ {
		tree := AVLTree{}

		source := rand.NewSource(time.Now().UnixNano())

		random := rand.New(source)

		bigInt := big.NewInt(0).Rand(random, big.NewInt(0).Exp(big.NewInt(10), big.NewInt(100), nil)).Int64()

		for j := 0; j < 10; j++ {
			bigInt = bigInt + 1
			tree.Insert(int(bigInt))
		}

		trees = append(trees, tree)
	}

	file := MakeFile("balanced-binary.txt")

	defer file.Close()

	for _, tree := range trees {
		WriteIntoFile(fmt.Sprintf("height in tree: %d", tree.root.height), file)
		tree.Output(file)

		WriteIntoFile("Test insert", file)

		DoAction(Insert{}, tree, file)

		WriteIntoFile("Test Search", file)

		DoAction(Search{}, tree, file)

		WriteIntoFile("Test delete", file)

		DoAction(Delete{}, tree, file)

		WriteIntoFile("----------------------", file)
	}
}

func TestCountingSort() {
	file := MakeFile("counting-sort.txt")

	defer file.Close()

	// small integer
	for i := 0; i < 10; i++ {
		WriteIntoFile("Make a new test with small integer..", file)
		var array []int

		for j := 0; j < 10; j++ {
			value := rand.Intn(10000)
			WriteIntoFile(strconv.Itoa(value), file)
			array = append(array, value)
		}

		testSorting(array, file)
	}

	for i := 0; i < 10; i++ {
		WriteIntoFile("Make a new test with ordered small integer..", file)
		var array []int

		value := rand.Intn(1000)

		for j := 0; j < 10; j++ {
			value = value + 1
			WriteIntoFile(strconv.Itoa(value), file)
			array = append(array, value)
		}

		testSorting(array, file)
	}

	// big integer
	for i := 0; i < 10; i++ {
		WriteIntoFile("Make a new test with big integer..", file)
		var array []int

		for j := 0; j < 10; j++ {
			source := rand.NewSource(time.Now().UnixNano())

			random := rand.New(source)

			value := random.Intn(90000) + 10000

			WriteIntoFile(strconv.Itoa(value), file)
			array = append(array, value)
		}

		testSorting(array, file)
	}

	for i := 0; i < 10; i++ {
		WriteIntoFile("Make a new test with ordered big integer..", file)
		var array []int

		source := rand.NewSource(time.Now().UnixNano())

		random := rand.New(source)

		value := random.Intn(90000) + 10000

		for j := 0; j < 10; j++ {
			value = value + 1
			WriteIntoFile(strconv.Itoa(value), file)
			array = append(array, value)
		}

		testSorting(array, file)
	}
}

func testSorting(array []int, file *os.File) {
	startTime := time.Now()

	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	startAlloc := memStats.Alloc

	value := rand.Intn(1000000)

	WriteIntoFile("Test with: "+strconv.Itoa(value), file)

	CountingSort(array, value)

	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)

	WriteIntoFile("Time taken:"+elapsedTime.String(), file)

	runtime.ReadMemStats(&memStats)
	endAlloc := memStats.Alloc
	memoryUsed := endAlloc - startAlloc

	WriteIntoFile("Memory used:"+strconv.Itoa(int(memoryUsed)), file)
	WriteIntoFile("--------------------", file)
}
