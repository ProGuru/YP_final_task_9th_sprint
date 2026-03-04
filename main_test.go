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

	var sumElements int = 12345
	randomElements = generateRandomElements(sumElements)
	require.NotNil(t, randomElements)
	require.Len(t, len(randomElements), sumElements)
}

func TestMaximum(t *testing.T) {
	randomElements = make([]int, 0)
	maxValue := maximum(randomElements)
	require.Zero(t, maxValue)

	randomElements = make([]int, 1)
	maxNum := 50
	randomElements[0] = maxNum
	maxValue = maximum(randomElements)
	require.Equal(t, maxNum, maxValue)

	randomElements = make([]int, 5)
	randomElements = []int{10, -50, 0, 803040, 100}
	maxValue = maximum(randomElements)
	require.Equal(t, 803040, maxValue)

	randomElements = []int{0, 0, 0, 0, 0}
	maxValue = maximum(randomElements)
	require.Equal(t, 0, maxValue)

	randomElements = []int{40, 40, 40, 40, 40}
	maxValue = maximum(randomElements)
	require.Equal(t, 40, maxValue)

	randomElements = []int{-200, -200, -200, -200, -200}
	maxValue = maximum(randomElements)
	require.Equal(t, -200, maxValue)
}
