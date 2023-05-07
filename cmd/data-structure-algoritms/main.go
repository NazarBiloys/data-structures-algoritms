package main

import (
	"fmt"
	"github.com/NazarBiloys/data-structures-algoritms/internal/service"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Start balanced binary search tree test..")
	service.TestBalancedBinary()
	fmt.Println("Start sorting by counting sort..")
	service.TestCountingSort()
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":90", nil))
}
