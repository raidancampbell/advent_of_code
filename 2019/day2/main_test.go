package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_loop(t *testing.T) {
	original := []int{1,1,1,4,99,5,6,0,99}
	testArr := []int{1,1,1,4,99,5,6,0,99}
	loop(testArr)

	assert.NotEqual(t, testArr, original)
	assert.Equal(t, 30, testArr[0])
}