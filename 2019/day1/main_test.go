package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMain_calcFuel2(t *testing.T) {
	assert.Equal(t, 50346, calcFuel2(100756, 0))
}