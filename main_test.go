package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Пишите тесты в этом файле
func TestGenerateRandomElements(t *testing.T) {
	var randomElements []int = generateRandomElements(0)
	require.Nil(t, randomElements)

	var sumElements int = 12345
	randomElements = generateRandomElements(sumElements)
	require.NotNil(t, randomElements)
	require.Len(t, len(randomElements), sumElements)
}
