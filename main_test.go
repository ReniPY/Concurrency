package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRandomElements(t *testing.T) {
	t.Run("ValidSize", func(t *testing.T) {
		size := 10
		result := generateRandomElements(size)
		assert.Len(t, result, size, "должна быть создана последовательность нужной длины")
	})

	t.Run("ZeroSize", func(t *testing.T) {
		size := 0
		result := generateRandomElements(size)
		assert.Nil(t, result, "результат должен быть равен nil при размере 0")
	})

	t.Run("NegativeSize", func(t *testing.T) {
		size := -10
		result := generateRandomElements(size)
		assert.Nil(t, result, "результат должен быть равен nil при отрицательном размере")
	})
}

func TestMaximum(t *testing.T) {
	t.Run("EmptyArray", func(t *testing.T) {
		arr := []int{}
		result := maximum(arr)
		assert.Equal(t, 0, result, "для пустого массива должно возвращаться 0")
	})

	t.Run("OneElement", func(t *testing.T) {
		arr := []int{42}
		result := maximum(arr)
		assert.Equal(t, 42, result, "одиночное значение должно возвращаться как максимум")
	})

	t.Run("MultipleElements", func(t *testing.T) {
		arr := []int{1, 5, 3, 8, 2}
		result := maximum(arr)
		assert.Equal(t, 8, result, "функция должна возвращать верный максимум")
	})
}

func TestMaxChunks(t *testing.T) {
	t.Run("EmptyArray", func(t *testing.T) {
		arr := []int{}
		result := maxChunks(arr)
		assert.Equal(t, 0, result, "для пустого массива должно возвращаться 0")
	})

	t.Run("OneChunk", func(t *testing.T) {
		arr := []int{1, 5, 3, 8, 2}
		result := maxChunks(arr)
		assert.Equal(t, 8, result, "функция должна возвращать верный максимум даже при одной порции")
	})

	t.Run("MultipleChunks", func(t *testing.T) {
		arr := []int{1, 5, 3, 8, 2, 9, 7, 4, 6}
		result := maxChunks(arr)
		assert.Equal(t, 9, result, "функция должна возвращать верный максимум при разбиении на куски")
	})
}
