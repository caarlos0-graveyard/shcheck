package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIgnore(t *testing.T) {
	var assert = assert.New(t)
	assert.True(ignore([]string{"vendor/**/*"}, "vendor/blah/adad/asdasd"))
}
