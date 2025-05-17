package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRandomElements(t *testing.T) {
	testCases := []struct {
		name     string
		size     int
		expected bool
	}{
		{"PositiveSize", 10, true},
		{"ZeroSize", 0, false},
		{"NegativeSize", -10, false},
	}

	for _, tc := range testCases {
		result := generateRandomElements(tc.size)
		if !tc.expected {
			assert.Nil(t, result, "результат должен быть равен nil при некорректном размере")
		} else {
			assert.NotNil(t, result, "результат не должен быть nil при положительном размере")
			assert.Len(t, result, tc.size, "размер результата должен соответствовать входному значению")
		}
	}
}

func TestMaximum(t *testing.T) {
	testCases := []struct {
		input    []int
		expected int
	}{
		{[]int{}, 0},
		{[]int{42}, 42},
		{[]int{1, 5, 3, 8, 2}, 8},
	}

	for _, tc := range testCases {
		result := maximum(tc.input)
		assert.Equal(t, tc.expected, result, "неверное вычисление максимального значения")
	}
}

func TestMaxChunks(t *testing.T) {
	testCases := []struct {
		input    []int
		expected int
	}{
		{[]int{}, 0},
		{[]int{1, 5, 3, 8, 2}, 8},
		{[]int{1, 5, 3, 8, 2, 9, 7, 4, 6}, 9},
	}

	for _, tc := range testCases {
		result := maxChunks(tc.input)
		assert.Equal(t, tc.expected, result, "неверное вычисление максимального значения")
	}
}
