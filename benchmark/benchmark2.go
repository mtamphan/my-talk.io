package main

import (
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	n, _ := strconv.Atoi(os.Args[0])
	start := time.Now()
	A := make([][]int, n)
	B := make([][]int, n)
	C := make([][]int, n)
	// just like python init using 3 loop
	for i := 0; i < n; i++ {
		A[i] = make([]int, n)
		for j := 0; j < n; j++ {
			A[i][j] = rand.Int()
		}
	}
	for i := 0; i < n; i++ {
		B[i] = make([]int, n)
		for j := 0; j < n; j++ {
			B[i][j] = rand.Int()
		}
	}
	for i := 0; i < n; i++ {
		C[i] = make([]int, n)
		for j := 0; j < n; j++ {
			C[i][j] = rand.Int()
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				C[i][j] += A[i][k] * B[k][i]
			}
		}
	}
	elapsed := time.Since(start)
	log.Printf("program took %s", elapsed)
}
