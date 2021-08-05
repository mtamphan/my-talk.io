package main

import (
	"log"
	"math/rand"
	"time"
)

const n = 2084

func main() {
	start := time.Now()
	A := [n][n]int{}
	B := [n][n]int{}
	C := [n][n]int{}
	// just like python init using 3 loop
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			A[i][j] = rand.Int()
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			B[i][j] = rand.Int()
		}
	}
	for i := 0; i < n; i++ {
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
