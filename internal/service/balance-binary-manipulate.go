package service

import (
	"os"
	"runtime"
	"strconv"
	"time"
)

type Action interface {
	do(tree AVLTree)
}

type Search struct {
}

type Delete struct {
}

type Insert struct {
}

func (s Search) do(tree AVLTree) {
	tree.Search(1230)
}

func (s Insert) do(tree AVLTree) {
	tree.Insert(1230)
}

func (s Delete) do(tree AVLTree) {
	tree.Delete(1230)
}

func DoAction(action Action, tree AVLTree, file *os.File) {
	startTime := time.Now()

	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	startAlloc := memStats.Alloc

	action.do(tree)

	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)

	WriteIntoFile("Time taken:"+elapsedTime.String(), file)

	runtime.ReadMemStats(&memStats)
	endAlloc := memStats.Alloc
	memoryUsed := endAlloc - startAlloc

	WriteIntoFile("Memory used:"+strconv.Itoa(int(memoryUsed)), file)
}
