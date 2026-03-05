package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var randomElements []int

// Пишите тесты в этом файле
func TestGenerateRandomElements(t *testing.T) {
	randomElements = generateRandomElements(0)
	require.Nil(t, randomElements)

	numElements := 12345
	randomElements = generateRandomElements(numElements)
	require.NotNil(t, randomElements)
	require.Len(t, randomElements, numElements)
}

// для нашего случая неплохо подошли бы параметризованные тесты, чтобы не дублировать один и тот же код
func TestMaxFunctions(t *testing.T) {
	t.Run("maximum", func(t *testing.T) {
		t.Parallel()
		testMaximum(t, maximum)
	})
	t.Run("maxChunks", func(t *testing.T) {
		t.Parallel()
		testMaximum(t, maxChunks)
	})
}

// в параметры добавим общую для maximum и для maxChunks сигнатуру функции: f func([]int) int
func testMaximum(t *testing.T, f func([]int) int) {
	randomElements = make([]int, 0)
	maxValue := f(randomElements)
	require.Zero(t, maxValue)

	randomElements = make([]int, 1)
	maxNum := 50
	randomElements[0] = maxNum
	maxValue = f(randomElements)
	require.Equal(t, maxNum, maxValue)

	randomElements = make([]int, 5)
	randomElements = []int{10, -50, 0, 803040, 100}
	maxValue = f(randomElements)
	require.Equal(t, 803040, maxValue)

	randomElements = []int{0, 0, 0, 0, 0}
	maxValue = f(randomElements)
	require.Equal(t, 0, maxValue)

	randomElements = []int{40, 40, 40, 40, 40}
	maxValue = f(randomElements)
	require.Equal(t, 40, maxValue)

	randomElements = []int{-200, -200, -200, -200, -200}
	maxValue = f(randomElements)
	require.Equal(t, -200, maxValue)

	randomElements = make([]int, 8)
	randomElements = []int{10, -50, 0, 505236412, 100, -1250, 10100, 9}
	maxValue = f(randomElements)
	require.Equal(t, 505236412, maxValue)

	randomElements = make([]int, 15)
	randomElements = []int{12, -7, 45, 0, -100, 234, 89, -3, 678, 1, -56, 999, 2, -456, 3001}
	maxValue = f(randomElements)
	require.Equal(t, 3001, maxValue)
}
