package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return nil
	}
	data := make([]int, size)
	for i := range data {
		data[i] = rand.Int()
	}
	return data
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}
	maxVal := data[0]
	for _, val := range data {
		if val > maxVal {
			maxVal = val
		}
	}
	return maxVal
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}

	var wg sync.WaitGroup
	results := make([]int, CHUNKS)
	chunkSize := len(data) / CHUNKS

	wg.Add(CHUNKS)
	for i := 0; i < CHUNKS; i++ {
		startIndex := i * chunkSize
		endIndex := startIndex + chunkSize
		if i == CHUNKS-1 {
			endIndex = len(data)
		}

		go func(chunkData []int, index int) {
			defer wg.Done()
			results[i] = maximum(chunkData)
		}(data[startIndex:endIndex], i)
	}
	wg.Wait()

	return maximum(results)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	slice := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	maxSequential := maximum(slice)
	elapsedSequential := time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", maxSequential, elapsedSequential)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	start = time.Now()
	maxParallel := maxChunks(slice)
	elapsedParallel := time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", maxParallel, elapsedParallel)
}
