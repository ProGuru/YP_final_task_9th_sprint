package main

import (
	"fmt"
	"math/rand"
	"slices"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	// ваш код здесь
	if size == 0 {
		return nil
	}
	src := rand.NewSource(time.Now().Unix())

	randomElements := make([]int, size)
	for i := 0; i < size; i++ {
		randomElements[i] = int(src.Int63())
	}

	return randomElements
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	// ваш код здесь
	if len(data) == 0 {
		return 0
	}
	return slices.Max(data)
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	// ваш код здесь
	if len(data) == 0 {
		return 0
	}
	var wg sync.WaitGroup

	// делим слайс на 8 частей
	chunkSize := len(data) / CHUNKS
	chunk := CHUNKS
	if chunkSize == 0 {
		chunkSize = 1
		chunk = len(data)
	}

	// сохраняем все 8 максимальных значений в другом новом слайсе
	maxValuesFromChunk := make([]int, chunk)

	wg.Add(chunk)
	for i := 1; i <= chunk; i++ {
		// создайте переменную типа sync.WaitGroup и используйте её при запуске и ожидании горутин

		var clip []int
		switch i {
		case chunk:
			clip = data[chunkSize*(i-1):]
		default:
			clip = data[chunkSize*(i-1) : chunkSize*i]
		}

		go func(idx int, newClip []int) {
			defer wg.Done()
			maxFromChunk := slices.Max(newClip)
			maxValuesFromChunk[idx-1] = maxFromChunk
		}(i, clip)
	}

	wg.Wait()
	// находим среди 8-ми максимальных значений максимум
	return slices.Max(maxValuesFromChunk)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел", SIZE)
	// ваш код здесь
	randomElements := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	// ваш код здесь
	start := time.Now()
	max := maximum(randomElements)
	elapsed := time.Since(start).Milliseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков", CHUNKS)
	// ваш код здесь
	start = time.Now()
	max = maxChunks(randomElements)
	elapsed = time.Since(start).Milliseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
