package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_password_check(t *testing.T) {
	p := password{val:111111}
	assert.Equal(t, p.check(), 0)

	p = password{val:223450}
	assert.NotEqual(t, p.check(), 0)

	p = password{val:123789}
	assert.Equal(t, p.check(), 1)

	p = password{val:123789}
	assert.Equal(t, p.check(), 1)
}

func Test_password_check2(t *testing.T) {
	p := password{val:112233}
	assert.Equal(t, p.check2(), 0)

	p = password{val:123444}
	assert.NotEqual(t, p.check2(), 0)

	p = password{val:111122}
	assert.Equal(t, p.check2(), 0)

	p = password{val:788889}
	assert.NotEqual(t, p.check2(), 0)

}